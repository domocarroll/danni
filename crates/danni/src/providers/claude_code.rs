use anyhow::Result;
use async_stream::try_stream;
use async_trait::async_trait;
use rmcp::model::Role;
use serde::{Deserialize, Serialize};
use serde_json::{json, Value};
use std::ffi::OsString;
use std::path::PathBuf;
use std::process::Stdio;
use tokio::io::{AsyncBufReadExt, AsyncReadExt, BufReader};
use tokio::process::Command;

use super::base::{ConfigKey, MessageStream, Provider, ProviderMetadata, ProviderUsage, Usage};
use super::errors::ProviderError;
use super::utils::{filter_extensions_from_system_prompt, RequestLog};
use crate::config::base::ClaudeCodeCommand;
use crate::config::search_path::SearchPaths;
use crate::config::{Config, DanniMode};
use crate::conversation::message::{Message, MessageContent};
use crate::model::ModelConfig;
use crate::subprocess::configure_command_no_window;
use rmcp::model::Tool;

pub const CLAUDE_CODE_DEFAULT_MODEL: &str = "claude-sonnet-4-20250514";
pub const CLAUDE_CODE_KNOWN_MODELS: &[&str] = &["sonnet", "opus"];
pub const CLAUDE_CODE_DOC_URL: &str = "https://code.claude.com/docs/en/setup";

/// Represents a streaming event from Claude Code CLI's `--output-format stream-json`
/// Each line of output is a complete JSON object representing one of these event types.
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(tag = "type", rename_all = "snake_case")]
pub enum ClaudeCodeStreamEvent {
    /// System information at stream start
    System {
        #[serde(default)]
        subtype: Option<String>,
        #[serde(default)]
        session_id: Option<String>,
    },
    /// Assistant message with content
    Assistant {
        message: AssistantMessage,
    },
    /// Wrapped streaming event (when using --include-partial-messages)
    StreamEvent {
        event: InnerStreamEvent,
    },
    /// Content block started (for incremental streaming)
    ContentBlockStart {
        index: usize,
        content_block: ContentBlock,
    },
    /// Incremental text delta
    ContentBlockDelta {
        index: usize,
        delta: ContentDelta,
    },
    /// Content block finished
    ContentBlockStop {
        index: usize,
    },
    /// Final result with usage statistics
    Result {
        #[serde(default)]
        subtype: Option<String>,
        #[serde(default)]
        cost_usd: Option<f64>,
        #[serde(default)]
        duration_ms: Option<u64>,
        #[serde(default)]
        duration_api_ms: Option<u64>,
        #[serde(default)]
        is_error: Option<bool>,
        #[serde(default)]
        num_turns: Option<u32>,
        #[serde(default)]
        result: Option<String>,
        #[serde(default)]
        session_id: Option<String>,
        #[serde(default)]
        total_cost_usd: Option<f64>,
        #[serde(default)]
        usage: Option<StreamUsage>,
    },
    /// User message (contains tool results)
    User {
        message: UserMessage,
        #[serde(default)]
        tool_use_result: Option<ToolUseResult>,
    },
}

/// User message structure (for tool results)
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct UserMessage {
    #[serde(default)]
    pub role: Option<String>,
    #[serde(default)]
    pub content: Vec<UserContent>,
}

/// User message content (tool results)
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(tag = "type", rename_all = "snake_case")]
pub enum UserContent {
    ToolResult {
        tool_use_id: String,
        #[serde(default)]
        content: Option<String>,
    },
    #[serde(other)]
    Unknown,
}

/// Tool use result metadata
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct ToolUseResult {
    #[serde(default, rename = "type")]
    pub result_type: Option<String>,
    #[serde(default)]
    pub file: Option<FileResult>,
    // Bash tool results
    #[serde(default)]
    pub stdout: Option<String>,
    #[serde(default)]
    pub stderr: Option<String>,
    #[serde(default)]
    pub interrupted: Option<bool>,
}

/// File result from tool execution
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct FileResult {
    #[serde(default, rename = "filePath")]
    pub file_path: Option<String>,
    #[serde(default)]
    pub content: Option<String>,
    #[serde(default, rename = "numLines")]
    pub num_lines: Option<i32>,
}

/// Inner streaming event (unwrapped from stream_event)
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(tag = "type", rename_all = "snake_case")]
pub enum InnerStreamEvent {
    /// Message started
    MessageStart {
        #[serde(default)]
        message: Option<AssistantMessage>,
    },
    /// Content block started
    ContentBlockStart {
        index: usize,
        content_block: ContentBlock,
    },
    /// Incremental content delta
    ContentBlockDelta {
        index: usize,
        delta: ContentDelta,
    },
    /// Content block stopped
    ContentBlockStop {
        index: usize,
    },
    /// Message delta (stop reason, usage)
    MessageDelta {
        #[serde(default)]
        delta: Option<MessageDeltaContent>,
        #[serde(default)]
        usage: Option<StreamUsage>,
    },
    /// Message stopped
    MessageStop,
    /// Catch-all for unknown events
    #[serde(other)]
    Unknown,
}

/// Content of a message delta event
#[derive(Debug, Clone, Default, Serialize, Deserialize)]
pub struct MessageDeltaContent {
    #[serde(default)]
    pub stop_reason: Option<String>,
    #[serde(default)]
    pub stop_sequence: Option<String>,
}

/// Assistant message structure from streaming events
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct AssistantMessage {
    #[serde(default)]
    pub id: Option<String>,
    #[serde(default)]
    pub role: Option<String>,
    #[serde(default)]
    pub content: Vec<ContentBlock>,
    #[serde(default)]
    pub model: Option<String>,
    #[serde(default)]
    pub stop_reason: Option<String>,
    #[serde(default)]
    pub usage: Option<StreamUsage>,
}

/// Content block within a message
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(tag = "type", rename_all = "snake_case")]
pub enum ContentBlock {
    Text {
        text: String,
    },
    ToolUse {
        id: String,
        name: String,
        #[serde(default)]
        input: Option<Value>,
    },
    #[serde(other)]
    Unknown,
}

/// Delta for incremental content updates
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(tag = "type", rename_all = "snake_case")]
pub enum ContentDelta {
    TextDelta {
        text: String,
    },
    InputJsonDelta {
        #[serde(default)]
        partial_json: Option<String>,
    },
    #[serde(other)]
    Unknown,
}

/// Usage statistics from streaming events
#[derive(Debug, Clone, Default, Serialize, Deserialize)]
pub struct StreamUsage {
    #[serde(default)]
    pub input_tokens: Option<i64>,
    #[serde(default)]
    pub output_tokens: Option<i64>,
    #[serde(default)]
    pub cache_creation_input_tokens: Option<i64>,
    #[serde(default)]
    pub cache_read_input_tokens: Option<i64>,
}

/// Result of parsing a streaming event
#[derive(Debug)]
pub enum ParsedStreamEvent {
    /// Text content to yield immediately
    Text(String),
    /// Usage statistics (typically from result event)
    Usage(Usage),
    /// Text with usage (final assistant message)
    TextWithUsage(String, Usage),
    /// Tool call started (name, id)
    ToolCallStart { name: String, id: String },
    /// Tool call input being built (partial JSON)
    ToolCallInput { id: String, partial_json: String },
    /// Tool call result received
    ToolResult { tool_name: String, file_path: Option<String>, summary: String },
    /// Event was parsed but should be skipped (system, unknown, etc.)
    Skip,
}

/// Parse a single line from Claude Code CLI stream-json output
pub fn parse_stream_event(line: &str) -> Result<ParsedStreamEvent, ProviderError> {
    let line = line.trim();

    // Skip empty lines
    if line.is_empty() {
        return Ok(ParsedStreamEvent::Skip);
    }

    // Skip non-JSON lines (debug output, warnings)
    if !line.starts_with('{') {
        tracing::debug!("Skipping non-JSON line: {}", &line[..line.len().min(100)]);
        return Ok(ParsedStreamEvent::Skip);
    }

    // Parse the JSON event - skip unparseable events instead of failing
    let event: ClaudeCodeStreamEvent = match serde_json::from_str(line) {
        Ok(e) => e,
        Err(e) => {
            tracing::warn!("Skipping unparseable stream event: {} - line: {}", e, &line[..line.len().min(200)]);
            return Ok(ParsedStreamEvent::Skip);
        }
    };

    match event {
        ClaudeCodeStreamEvent::System { .. } => {
            Ok(ParsedStreamEvent::Skip)
        }
        ClaudeCodeStreamEvent::StreamEvent { event: inner } => {
            // Handle wrapped streaming events (from --include-partial-messages)
            match inner {
                InnerStreamEvent::ContentBlockDelta { delta, .. } => {
                    match delta {
                        ContentDelta::TextDelta { text } => {
                            if !text.is_empty() {
                                return Ok(ParsedStreamEvent::Text(text));
                            }
                        }
                        ContentDelta::InputJsonDelta { partial_json } => {
                            // Tool input being built - emit for visibility
                            if let Some(json) = partial_json {
                                if !json.is_empty() {
                                    return Ok(ParsedStreamEvent::ToolCallInput {
                                        id: String::new(), // ID tracked separately
                                        partial_json: json,
                                    });
                                }
                            }
                        }
                        _ => {}
                    }
                    Ok(ParsedStreamEvent::Skip)
                }
                InnerStreamEvent::ContentBlockStart { content_block, .. } => {
                    match content_block {
                        ContentBlock::Text { text } => {
                            if !text.is_empty() {
                                return Ok(ParsedStreamEvent::Text(text));
                            }
                        }
                        ContentBlock::ToolUse { id, name, .. } => {
                            // Tool call starting - emit for visibility
                            return Ok(ParsedStreamEvent::ToolCallStart { name, id });
                        }
                        _ => {}
                    }
                    Ok(ParsedStreamEvent::Skip)
                }
                InnerStreamEvent::MessageDelta { usage, .. } => {
                    if let Some(stream_usage) = usage {
                        let usage = Usage {
                            input_tokens: stream_usage.input_tokens.map(|v| v as i32),
                            output_tokens: stream_usage.output_tokens.map(|v| v as i32),
                            total_tokens: match (stream_usage.input_tokens, stream_usage.output_tokens) {
                                (Some(i), Some(o)) => Some((i + o) as i32),
                                _ => None,
                            },
                            ..Default::default()
                        };
                        return Ok(ParsedStreamEvent::Usage(usage));
                    }
                    Ok(ParsedStreamEvent::Skip)
                }
                _ => Ok(ParsedStreamEvent::Skip),
            }
        }
        ClaudeCodeStreamEvent::Assistant { message } => {
            // Extract text content from assistant message (skip when using partial messages
            // as we already got the text from stream_event deltas)
            // Only emit if we have content that wasn't already streamed
            let mut text_parts = Vec::new();
            for content in &message.content {
                if let ContentBlock::Text { text } = content {
                    text_parts.push(text.clone());
                }
            }

            let combined_text = text_parts.join("");

            // Extract usage if present
            if let Some(stream_usage) = message.usage {
                let usage = Usage {
                    input_tokens: stream_usage.input_tokens.map(|v| v as i32),
                    output_tokens: stream_usage.output_tokens.map(|v| v as i32),
                    total_tokens: match (stream_usage.input_tokens, stream_usage.output_tokens) {
                        (Some(i), Some(o)) => Some((i + o) as i32),
                        _ => None,
                    },
                    ..Default::default()
                };

                // When using partial messages, text already streamed - just emit usage
                // When not using partial messages, emit text with usage
                if combined_text.is_empty() {
                    Ok(ParsedStreamEvent::Usage(usage))
                } else {
                    // Skip text since it was already streamed via content_block_delta
                    // Just emit usage
                    Ok(ParsedStreamEvent::Usage(usage))
                }
            } else {
                // No usage info, skip (text already streamed)
                Ok(ParsedStreamEvent::Skip)
            }
        }
        ClaudeCodeStreamEvent::ContentBlockStart { content_block, .. } => {
            // Emit initial text if present
            if let ContentBlock::Text { text } = content_block {
                if !text.is_empty() {
                    return Ok(ParsedStreamEvent::Text(text));
                }
            }
            Ok(ParsedStreamEvent::Skip)
        }
        ClaudeCodeStreamEvent::ContentBlockDelta { delta, .. } => {
            // Emit text delta
            if let ContentDelta::TextDelta { text } = delta {
                if !text.is_empty() {
                    return Ok(ParsedStreamEvent::Text(text));
                }
            }
            Ok(ParsedStreamEvent::Skip)
        }
        ClaudeCodeStreamEvent::ContentBlockStop { .. } => {
            Ok(ParsedStreamEvent::Skip)
        }
        ClaudeCodeStreamEvent::Result { usage, .. } => {
            // Extract final usage statistics
            if let Some(stream_usage) = usage {
                let usage = Usage {
                    input_tokens: stream_usage.input_tokens.map(|v| v as i32),
                    output_tokens: stream_usage.output_tokens.map(|v| v as i32),
                    total_tokens: match (stream_usage.input_tokens, stream_usage.output_tokens) {
                        (Some(i), Some(o)) => Some((i + o) as i32),
                        _ => None,
                    },
                    ..Default::default()
                };
                Ok(ParsedStreamEvent::Usage(usage))
            } else {
                Ok(ParsedStreamEvent::Skip)
            }
        }
        ClaudeCodeStreamEvent::User { tool_use_result, .. } => {
            // Extract tool result information for visibility
            if let Some(result) = tool_use_result {
                let file_path = result.file.as_ref().and_then(|f| f.file_path.clone());
                let tool_name = result.result_type.clone().unwrap_or_else(|| "tool".to_string());

                // Generate summary based on what data we have
                let summary = if let Some(ref file) = result.file {
                    // File read result - show just filename, not full path
                    if let Some(ref path) = file.file_path {
                        let filename = std::path::Path::new(path)
                            .file_name()
                            .and_then(|n| n.to_str())
                            .unwrap_or(path);
                        let lines = file.num_lines.unwrap_or(0);
                        format!("Read {} ({} lines)", filename, lines)
                    } else {
                        "File read".to_string()
                    }
                } else if result.stdout.is_some() || result.stderr.is_some() {
                    // Bash/command - skip these (too noisy, not informative)
                    "Command completed".to_string()
                } else {
                    "Tool completed".to_string()
                };

                return Ok(ParsedStreamEvent::ToolResult {
                    tool_name,
                    file_path,
                    summary,
                });
            }
            Ok(ParsedStreamEvent::Skip)
        }
    }
}

#[derive(Debug, serde::Serialize)]
pub struct ClaudeCodeProvider {
    command: PathBuf,
    model: ModelConfig,
    #[serde(skip)]
    name: String,
}

impl ClaudeCodeProvider {
    pub async fn from_env(model: ModelConfig) -> Result<Self> {
        let config = crate::config::Config::global();
        let command: OsString = config.get_claude_code_command().unwrap_or_default().into();
        let resolved_command = SearchPaths::builder().with_npm().resolve(command)?;

        Ok(Self {
            command: resolved_command,
            model,
            name: Self::metadata().name,
        })
    }

    /// Convert danni messages to the format expected by claude CLI
    fn messages_to_claude_format(&self, _system: &str, messages: &[Message]) -> Result<Value> {
        let mut claude_messages = Vec::new();

        for message in messages.iter().filter(|m| m.is_agent_visible()) {
            let role = match message.role {
                Role::User => "user",
                Role::Assistant => "assistant",
            };

            let mut content_parts = Vec::new();
            for content in &message.content {
                match content {
                    MessageContent::Text(text_content) => {
                        content_parts.push(json!({
                            "type": "text",
                            "text": text_content.text
                        }));
                    }
                    MessageContent::ToolRequest(tool_request) => {
                        if let Ok(tool_call) = &tool_request.tool_call {
                            content_parts.push(json!({
                                "type": "tool_use",
                                "id": tool_request.id,
                                "name": tool_call.name,
                                "input": tool_call.arguments
                            }));
                        }
                    }
                    MessageContent::ToolResponse(tool_response) => {
                        if let Ok(tool_contents) = &tool_response.tool_result {
                            // Convert tool result contents to text
                            let content_text = tool_contents
                                .iter()
                                .filter_map(|content| match &content.raw {
                                    rmcp::model::RawContent::Text(text_content) => {
                                        Some(text_content.text.as_str())
                                    }
                                    _ => None,
                                })
                                .collect::<Vec<&str>>()
                                .join("\n");

                            content_parts.push(json!({
                                "type": "tool_result",
                                "tool_use_id": tool_response.id,
                                "content": content_text
                            }));
                        }
                    }
                    _ => {
                        // Skip other content types for now
                    }
                }
            }

            claude_messages.push(json!({
                "role": role,
                "content": content_parts
            }));
        }

        Ok(json!(claude_messages))
    }

    /// Parse the JSON response from claude CLI
    fn apply_permission_flags(cmd: &mut Command) -> Result<(), ProviderError> {
        let config = Config::global();
        let danni_mode = config.get_danni_mode().unwrap_or(DanniMode::Auto);

        match danni_mode {
            DanniMode::Auto => {
                cmd.arg("--dangerously-skip-permissions");
            }
            DanniMode::SmartApprove => {
                cmd.arg("--permission-mode").arg("acceptEdits");
            }
            DanniMode::Approve => {
                return Err(ProviderError::RequestFailed(
                    "\n\n\n### NOTE\n\n\n \
                    Claude Code CLI provider does not support Approve mode.\n \
                    Please use Auto (which will run anything it needs to) or \
                    SmartApprove (most things will run or Chat Mode)\n\n\n"
                        .to_string(),
                ));
            }
            DanniMode::Chat => {
                // Chat mode doesn't need permission flags
            }
        }
        Ok(())
    }

    fn parse_claude_response(
        &self,
        json_lines: &[String],
    ) -> Result<(Message, Usage), ProviderError> {
        let mut all_text_content = Vec::new();
        let mut usage = Usage::default();

        // Join all content (should typically be a single element now with the new reading approach)
        let full_response = json_lines.join("");
        let response_len = full_response.len();

        // Parse with detailed error diagnostics
        let json_array: Vec<Value> = serde_json::from_str(&full_response).map_err(|e| {
            // Provide detailed diagnostics for debugging truncation issues
            let preview_start: String = full_response.chars().take(500).collect();
            let preview_end: String = full_response.chars().rev().take(500).collect::<String>().chars().rev().collect();

            ProviderError::RequestFailed(format!(
                "Failed to parse JSON response (length: {} bytes): {}.\n\
                 First 500 chars: {}\n\
                 Last 500 chars: {}",
                response_len,
                e,
                preview_start,
                preview_end
            ))
        })?;

        for parsed in json_array {
            if let Some(msg_type) = parsed.get("type").and_then(|t| t.as_str()) {
                match msg_type {
                    "assistant" => {
                        if let Some(message) = parsed.get("message") {
                            // Extract text content from this assistant message
                            if let Some(content) = message.get("content").and_then(|c| c.as_array())
                            {
                                for item in content {
                                    if let Some(content_type) =
                                        item.get("type").and_then(|t| t.as_str())
                                    {
                                        if content_type == "text" {
                                            if let Some(text) =
                                                item.get("text").and_then(|t| t.as_str())
                                            {
                                                all_text_content.push(text.to_string());
                                            }
                                        }
                                        // Skip tool_use - those are claude CLI's internal tools
                                    }
                                }
                            }

                            // Extract usage information
                            if let Some(usage_info) = message.get("usage") {
                                usage.input_tokens = usage_info
                                    .get("input_tokens")
                                    .and_then(|v| v.as_i64())
                                    .map(|v| v as i32);
                                usage.output_tokens = usage_info
                                    .get("output_tokens")
                                    .and_then(|v| v.as_i64())
                                    .map(|v| v as i32);

                                // Calculate total if not provided
                                if usage.total_tokens.is_none() {
                                    if let (Some(input), Some(output)) =
                                        (usage.input_tokens, usage.output_tokens)
                                    {
                                        usage.total_tokens = Some(input + output);
                                    }
                                }
                            }
                        }
                    }
                    "result" => {
                        // Extract additional usage info from result if available
                        if let Some(result_usage) = parsed.get("usage") {
                            if usage.input_tokens.is_none() {
                                usage.input_tokens = result_usage
                                    .get("input_tokens")
                                    .and_then(|v| v.as_i64())
                                    .map(|v| v as i32);
                            }
                            if usage.output_tokens.is_none() {
                                usage.output_tokens = result_usage
                                    .get("output_tokens")
                                    .and_then(|v| v.as_i64())
                                    .map(|v| v as i32);
                            }
                        }
                    }
                    _ => {} // Ignore other message types
                }
            }
        }

        // Combine all text content into a single message
        let combined_text = all_text_content.join("\n\n");
        if combined_text.is_empty() {
            return Err(ProviderError::RequestFailed(
                "No text content found in response".to_string(),
            ));
        }

        let message_content = vec![MessageContent::text(combined_text)];

        let response_message = Message::new(
            Role::Assistant,
            chrono::Utc::now().timestamp(),
            message_content,
        );

        Ok((response_message, usage))
    }

    async fn execute_command(
        &self,
        system: &str,
        messages: &[Message],
        _tools: &[Tool],
    ) -> Result<Vec<String>, ProviderError> {
        let messages_json = self
            .messages_to_claude_format(system, messages)
            .map_err(|e| {
                ProviderError::RequestFailed(format!("Failed to format messages: {}", e))
            })?;

        let filtered_system = filter_extensions_from_system_prompt(system);

        if std::env::var("DANNI_CLAUDE_CODE_DEBUG").is_ok() {
            println!("=== CLAUDE CODE PROVIDER DEBUG ===");
            println!("Command: {:?}", self.command);
            println!("Original system prompt length: {} chars", system.len());
            println!(
                "Filtered system prompt length: {} chars",
                filtered_system.len()
            );
            println!("Filtered system prompt: {}", filtered_system);
            println!(
                "Messages JSON: {}",
                serde_json::to_string_pretty(&messages_json)
                    .unwrap_or_else(|_| "Failed to serialize".to_string())
            );
            println!("================================");
        }

        let mut cmd = Command::new(&self.command);
        configure_command_no_window(&mut cmd);
        cmd.arg("-p")
            .arg(messages_json.to_string())
            .arg("--system-prompt")
            .arg(&filtered_system);

        // Only pass model parameter if it's in the known models list
        if CLAUDE_CODE_KNOWN_MODELS.contains(&self.model.model_name.as_str()) {
            cmd.arg("--model").arg(&self.model.model_name);
        }

        cmd.arg("--verbose").arg("--output-format").arg("json");

        // Add permission mode based on DANNI_MODE setting
        Self::apply_permission_flags(&mut cmd)?;

        // CRITICAL: Set stdin to null to prevent subprocess from inheriting parent's stdin
        // This prevents deadlock when parent is reading JSON from stdin (TUI mode)
        cmd.stdin(Stdio::null())
            .stdout(Stdio::piped())
            .stderr(Stdio::piped());

        let mut child = cmd.spawn().map_err(|e| {
            ProviderError::RequestFailed(format!(
                "Failed to spawn Claude CLI command '{:?}': {}.",
                self.command, e
            ))
        })?;

        let stdout = child
            .stdout
            .take()
            .ok_or_else(|| ProviderError::RequestFailed("Failed to capture stdout".to_string()))?;

        let stderr = child
            .stderr
            .take()
            .ok_or_else(|| ProviderError::RequestFailed("Failed to capture stderr".to_string()))?;

        // CRITICAL: Read stdout and stderr concurrently to prevent deadlock.
        // If stderr fills up (64KB pipe buffer), Claude CLI will block waiting for it to drain.
        // This can cause partial reads on stdout, leading to truncated JSON.
        // Use a 1MB buffer to handle large JSON responses from Claude CLI.
        const BUFFER_CAPACITY: usize = 1024 * 1024; // 1MB buffer

        let stdout_future = async {
            let mut reader = BufReader::with_capacity(BUFFER_CAPACITY, stdout);
            let mut output = Vec::new();
            reader.read_to_end(&mut output).await?;
            Ok::<_, std::io::Error>(output)
        };

        let stderr_future = async {
            let mut reader = BufReader::with_capacity(65536, stderr); // 64KB for stderr
            let mut output = Vec::new();
            reader.read_to_end(&mut output).await?;
            Ok::<_, std::io::Error>(output)
        };

        // Read both streams concurrently - this is critical to avoid deadlock
        let (stdout_result, stderr_result) = tokio::join!(stdout_future, stderr_future);

        let stdout_bytes = stdout_result.map_err(|e| {
            ProviderError::RequestFailed(format!("Failed to read stdout: {}", e))
        })?;

        let stderr_bytes = stderr_result.map_err(|e| {
            ProviderError::RequestFailed(format!("Failed to read stderr: {}", e))
        })?;

        // Log stderr for debugging (filter common noise)
        if !stderr_bytes.is_empty() {
            let stderr_str = String::from_utf8_lossy(&stderr_bytes);
            // Only log non-trivial stderr output
            if !stderr_str.trim().is_empty()
                && !stderr_str.contains("Debugger attached")
                && !stderr_str.contains("Waiting for the debugger")
            {
                tracing::debug!("Claude CLI stderr ({} bytes): {}", stderr_bytes.len(), stderr_str);
            }
        }

        let exit_status = child.wait().await.map_err(|e| {
            ProviderError::RequestFailed(format!("Failed to wait for command: {}", e))
        })?;

        if !exit_status.success() {
            let stderr_str = String::from_utf8_lossy(&stderr_bytes);
            return Err(ProviderError::RequestFailed(format!(
                "Command failed with exit code: {:?}. Stderr: {}",
                exit_status.code(),
                stderr_str.chars().take(2000).collect::<String>()
            )));
        }

        // Convert stdout bytes to string
        let stdout_str = String::from_utf8(stdout_bytes).map_err(|e| {
            ProviderError::RequestFailed(format!("Invalid UTF-8 in stdout: {}", e))
        })?;

        tracing::debug!(
            "Command executed successfully, got {} bytes of output",
            stdout_str.len()
        );

        // Debug output for diagnosing truncation issues
        if std::env::var("DANNI_CLAUDE_CODE_DEBUG").is_ok() {
            eprintln!("=== CLAUDE CODE RESPONSE DEBUG ===");
            eprintln!("Response size: {} bytes", stdout_str.len());
            eprintln!("Stderr size: {} bytes", stderr_bytes.len());
            if stdout_str.len() > 0 {
                eprintln!("First 200 chars: {}", stdout_str.chars().take(200).collect::<String>());
                eprintln!("Last 200 chars: {}", stdout_str.chars().rev().take(200).collect::<String>().chars().rev().collect::<String>());
            }
            eprintln!("================================");
        }

        // Return as a single string (the whole output)
        // The JSON from Claude CLI is typically a single array on one or more lines
        Ok(vec![stdout_str])
    }

    /// Generate a simple session description without calling subprocess
    fn generate_simple_session_description(
        &self,
        messages: &[Message],
    ) -> Result<(Message, ProviderUsage), ProviderError> {
        // Extract the first user message text
        let description = messages
            .iter()
            .find(|m| m.role == Role::User)
            .and_then(|m| {
                m.content.iter().find_map(|c| match c {
                    MessageContent::Text(text_content) => Some(&text_content.text),
                    _ => None,
                })
            })
            .map(|text| {
                // Take first few words, limit to 4 words
                text.split_whitespace()
                    .take(4)
                    .collect::<Vec<_>>()
                    .join(" ")
            })
            .unwrap_or_else(|| "Simple task".to_string());

        if std::env::var("DANNI_CLAUDE_CODE_DEBUG").is_ok() {
            println!("=== CLAUDE CODE PROVIDER DEBUG ===");
            println!("Generated simple session description: {}", description);
            println!("Skipped subprocess call for session description");
            println!("================================");
        }

        let message = Message::new(
            Role::Assistant,
            chrono::Utc::now().timestamp(),
            vec![MessageContent::text(description.clone())],
        );

        let usage = Usage::default();

        Ok((
            message,
            ProviderUsage::new(self.model.model_name.clone(), usage),
        ))
    }
}

#[async_trait]
impl Provider for ClaudeCodeProvider {
    fn metadata() -> ProviderMetadata {
        ProviderMetadata::new(
            "claude-code",
            "Claude Code CLI",
            "Requires claude CLI installed, no MCPs. Use Anthropic provider for full features.",
            CLAUDE_CODE_DEFAULT_MODEL,
            CLAUDE_CODE_KNOWN_MODELS.to_vec(),
            CLAUDE_CODE_DOC_URL,
            vec![ConfigKey::from_value_type::<ClaudeCodeCommand>(true, false)],
        )
    }

    fn get_name(&self) -> &str {
        &self.name
    }

    fn get_model_config(&self) -> ModelConfig {
        // Return the model config with appropriate context limit for Claude models
        self.model.clone()
    }

    #[tracing::instrument(
        skip(self, model_config, system, messages, tools),
        fields(model_config, input, output, input_tokens, output_tokens, total_tokens)
    )]
    async fn complete_with_model(
        &self,
        model_config: &ModelConfig,
        system: &str,
        messages: &[Message],
        tools: &[Tool],
    ) -> Result<(Message, ProviderUsage), ProviderError> {
        // Check if this is a session description request (short system prompt asking for 4 words or less)
        if system.contains("four words or less") || system.contains("4 words or less") {
            return self.generate_simple_session_description(messages);
        }

        let json_lines = self.execute_command(system, messages, tools).await?;

        let (message, usage) = self.parse_claude_response(&json_lines)?;

        // Create a dummy payload for debug tracing
        let payload = json!({
            "command": self.command,
            "model": model_config.model_name,
            "system": system,
            "messages": messages.len()
        });
        let mut log = RequestLog::start(model_config, &payload)?;

        let response = json!({
            "lines": json_lines.len(),
            "usage": usage
        });

        log.write(&response, Some(&usage))?;

        Ok((
            message,
            ProviderUsage::new(model_config.model_name.clone(), usage),
        ))
    }

    async fn stream(
        &self,
        system: &str,
        messages: &[Message],
        _tools: &[Tool],
    ) -> Result<MessageStream, ProviderError> {
        // Check if this is a session description request - use non-streaming for these
        if system.contains("four words or less") || system.contains("4 words or less") {
            let (message, usage) = self.generate_simple_session_description(messages)?;
            return Ok(super::base::stream_from_single_message(message, usage));
        }

        let messages_json = self
            .messages_to_claude_format(system, messages)
            .map_err(|e| {
                ProviderError::RequestFailed(format!("Failed to format messages: {}", e))
            })?;

        let filtered_system = filter_extensions_from_system_prompt(system);

        if std::env::var("DANNI_CLAUDE_CODE_DEBUG").is_ok() {
            eprintln!("=== CLAUDE CODE STREAMING DEBUG ===");
            eprintln!("Command: {:?}", self.command);
            eprintln!("Using stream-json with --include-partial-messages");
            eprintln!("================================");
        }

        let mut cmd = Command::new(&self.command);
        configure_command_no_window(&mut cmd);
        cmd.arg("-p")
            .arg(messages_json.to_string())
            .arg("--system-prompt")
            .arg(&filtered_system);

        // Only pass model parameter if it's in the known models list
        if CLAUDE_CODE_KNOWN_MODELS.contains(&self.model.model_name.as_str()) {
            cmd.arg("--model").arg(&self.model.model_name);
        }

        // Use stream-json for streaming output with partial messages for real-time streaming
        cmd.arg("--verbose")
            .arg("--output-format")
            .arg("stream-json")
            .arg("--include-partial-messages");

        // Add permission mode based on DANNI_MODE setting
        Self::apply_permission_flags(&mut cmd)?;

        // Set up stdio
        cmd.stdin(Stdio::null())
            .stdout(Stdio::piped())
            .stderr(Stdio::piped());

        let mut child = cmd.spawn().map_err(|e| {
            ProviderError::RequestFailed(format!(
                "Failed to spawn Claude CLI command '{:?}': {}.",
                self.command, e
            ))
        })?;

        let stdout = child
            .stdout
            .take()
            .ok_or_else(|| ProviderError::RequestFailed("Failed to capture stdout".to_string()))?;

        let stderr = child
            .stderr
            .take()
            .ok_or_else(|| ProviderError::RequestFailed("Failed to capture stderr".to_string()))?;

        let model_name = self.model.model_name.clone();

        // Spawn background task to drain stderr to prevent deadlock
        tokio::spawn(async move {
            let mut reader = BufReader::new(stderr);
            let mut line = String::new();
            while let Ok(n) = reader.read_line(&mut line).await {
                if n == 0 {
                    break;
                }
                let trimmed = line.trim();
                if !trimmed.is_empty()
                    && !trimmed.contains("Debugger attached")
                    && !trimmed.contains("Waiting for the debugger")
                {
                    tracing::debug!("Claude CLI stderr: {}", trimmed);
                }
                line.clear();
            }
        });

        // Create async stream that reads stdout line by line
        let reader = BufReader::new(stdout);
        let mut lines = reader.lines();

        Ok(Box::pin(try_stream! {
            let mut last_usage: Option<Usage> = None;

            while let Some(line) = lines.next_line().await.map_err(|e| {
                ProviderError::RequestFailed(format!("Failed to read stream line: {}", e))
            })? {
                match parse_stream_event(&line)? {
                    ParsedStreamEvent::Text(text) => {
                        // Yield partial message with text
                        let message = Message::new(
                            Role::Assistant,
                            chrono::Utc::now().timestamp(),
                            vec![MessageContent::text(text)],
                        );
                        yield (Some(message), None);
                    }
                    ParsedStreamEvent::Usage(usage) => {
                        // Store usage for final yield
                        last_usage = Some(usage);
                    }
                    ParsedStreamEvent::TextWithUsage(text, usage) => {
                        // Yield message with text and usage
                        let message = Message::new(
                            Role::Assistant,
                            chrono::Utc::now().timestamp(),
                            vec![MessageContent::text(text)],
                        );
                        let provider_usage = ProviderUsage::new(model_name.clone(), usage);
                        yield (Some(message), Some(provider_usage));
                    }
                    ParsedStreamEvent::ToolCallStart { .. } => {
                        // Skip tool starts - they don't show useful context (what file, etc.)
                        // The result will show the actual details
                    }
                    ParsedStreamEvent::ToolCallInput { .. } => {
                        // Skip partial JSON input (too noisy)
                    }
                    ParsedStreamEvent::ToolResult { summary, file_path, .. } => {
                        // Only emit if we have meaningful info (not just "Tool completed")
                        if file_path.is_some() || !summary.contains("completed") {
                            let activity = format!("\n`✓ {}`\n", summary);
                            let message = Message::new(
                                Role::Assistant,
                                chrono::Utc::now().timestamp(),
                                vec![MessageContent::text(activity)],
                            );
                            yield (Some(message), None);
                        }
                    }
                    ParsedStreamEvent::Skip => {
                        // Continue to next line
                    }
                }
            }

            // Wait for child process to exit
            let exit_status = child.wait().await.map_err(|e| {
                ProviderError::RequestFailed(format!("Failed to wait for command: {}", e))
            })?;

            if !exit_status.success() {
                Err(ProviderError::RequestFailed(format!(
                    "Command failed with exit code: {:?}",
                    exit_status.code()
                )))?;
            }

            // Yield final usage if we have it
            if let Some(usage) = last_usage {
                let provider_usage = ProviderUsage::new(model_name.clone(), usage);
                yield (None, Some(provider_usage));
            }
        }))
    }

    fn supports_streaming(&self) -> bool {
        true
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_parse_stream_event_empty_line() {
        let result = parse_stream_event("").unwrap();
        assert!(matches!(result, ParsedStreamEvent::Skip));
    }

    #[test]
    fn test_parse_stream_event_whitespace_line() {
        let result = parse_stream_event("   \n  ").unwrap();
        assert!(matches!(result, ParsedStreamEvent::Skip));
    }

    #[test]
    fn test_parse_stream_event_non_json_line() {
        let result = parse_stream_event("Some debug output").unwrap();
        assert!(matches!(result, ParsedStreamEvent::Skip));
    }

    #[test]
    fn test_parse_stream_event_system() {
        let json = r#"{"type":"system","subtype":"init","session_id":"abc123"}"#;
        let result = parse_stream_event(json).unwrap();
        assert!(matches!(result, ParsedStreamEvent::Skip));
    }

    #[test]
    fn test_parse_stream_event_assistant_with_text() {
        // With --include-partial-messages, text is streamed via content_block_delta
        // so assistant message skips text (already streamed) and returns Skip if no usage
        let json = r#"{"type":"assistant","message":{"id":"msg_1","role":"assistant","content":[{"type":"text","text":"Hello, world!"}],"model":"claude-3","stop_reason":"end_turn"}}"#;
        let result = parse_stream_event(json).unwrap();
        // No usage present, so Skip
        assert!(matches!(result, ParsedStreamEvent::Skip));
    }

    #[test]
    fn test_parse_stream_event_assistant_with_usage() {
        // With --include-partial-messages, text is streamed via content_block_delta
        // so assistant message just returns Usage (text already streamed)
        let json = r#"{"type":"assistant","message":{"id":"msg_1","role":"assistant","content":[{"type":"text","text":"Hi"}],"model":"claude-3","usage":{"input_tokens":100,"output_tokens":50}}}"#;
        let result = parse_stream_event(json).unwrap();
        match result {
            ParsedStreamEvent::Usage(usage) => {
                assert_eq!(usage.input_tokens, Some(100));
                assert_eq!(usage.output_tokens, Some(50));
                assert_eq!(usage.total_tokens, Some(150));
            }
            _ => panic!("Expected Usage, got {:?}", result),
        }
    }

    #[test]
    fn test_parse_stream_event_assistant_empty_content() {
        let json = r#"{"type":"assistant","message":{"id":"msg_1","role":"assistant","content":[],"model":"claude-3"}}"#;
        let result = parse_stream_event(json).unwrap();
        assert!(matches!(result, ParsedStreamEvent::Skip));
    }

    #[test]
    fn test_parse_stream_event_content_block_start_text() {
        let json = r#"{"type":"content_block_start","index":0,"content_block":{"type":"text","text":"Starting"}}"#;
        let result = parse_stream_event(json).unwrap();
        match result {
            ParsedStreamEvent::Text(text) => assert_eq!(text, "Starting"),
            _ => panic!("Expected Text, got {:?}", result),
        }
    }

    #[test]
    fn test_parse_stream_event_content_block_start_empty() {
        let json = r#"{"type":"content_block_start","index":0,"content_block":{"type":"text","text":""}}"#;
        let result = parse_stream_event(json).unwrap();
        assert!(matches!(result, ParsedStreamEvent::Skip));
    }

    #[test]
    fn test_parse_stream_event_content_block_delta() {
        let json = r#"{"type":"content_block_delta","index":0,"delta":{"type":"text_delta","text":" more text"}}"#;
        let result = parse_stream_event(json).unwrap();
        match result {
            ParsedStreamEvent::Text(text) => assert_eq!(text, " more text"),
            _ => panic!("Expected Text, got {:?}", result),
        }
    }

    #[test]
    fn test_parse_stream_event_content_block_delta_empty() {
        let json = r#"{"type":"content_block_delta","index":0,"delta":{"type":"text_delta","text":""}}"#;
        let result = parse_stream_event(json).unwrap();
        assert!(matches!(result, ParsedStreamEvent::Skip));
    }

    #[test]
    fn test_parse_stream_event_content_block_stop() {
        let json = r#"{"type":"content_block_stop","index":0}"#;
        let result = parse_stream_event(json).unwrap();
        assert!(matches!(result, ParsedStreamEvent::Skip));
    }

    #[test]
    fn test_parse_stream_event_result_with_usage() {
        let json = r#"{"type":"result","subtype":"success","cost_usd":0.01,"duration_ms":1000,"is_error":false,"num_turns":1,"result":"done","session_id":"abc","total_cost_usd":0.01,"usage":{"input_tokens":500,"output_tokens":200}}"#;
        let result = parse_stream_event(json).unwrap();
        match result {
            ParsedStreamEvent::Usage(usage) => {
                assert_eq!(usage.input_tokens, Some(500));
                assert_eq!(usage.output_tokens, Some(200));
                assert_eq!(usage.total_tokens, Some(700));
            }
            _ => panic!("Expected Usage, got {:?}", result),
        }
    }

    #[test]
    fn test_parse_stream_event_result_without_usage() {
        let json = r#"{"type":"result","subtype":"success","is_error":false}"#;
        let result = parse_stream_event(json).unwrap();
        assert!(matches!(result, ParsedStreamEvent::Skip));
    }

    #[test]
    fn test_parse_stream_event_tool_use_block() {
        let json = r#"{"type":"content_block_start","index":1,"content_block":{"type":"tool_use","id":"tool_1","name":"read_file","input":{"path":"/test"}}}"#;
        let result = parse_stream_event(json).unwrap();
        // Tool use blocks should be skipped (they're handled internally by Claude CLI)
        assert!(matches!(result, ParsedStreamEvent::Skip));
    }

    #[test]
    fn test_parse_stream_event_input_json_delta() {
        let json = r#"{"type":"content_block_delta","index":1,"delta":{"type":"input_json_delta","partial_json":"{\"path\":"}}"#;
        let result = parse_stream_event(json).unwrap();
        // Input JSON deltas should be skipped
        assert!(matches!(result, ParsedStreamEvent::Skip));
    }

    #[test]
    fn test_parse_stream_event_multiple_text_blocks() {
        // When using --include-partial-messages, assistant messages skip text (already streamed)
        // and just return usage if present
        let json = r#"{"type":"assistant","message":{"content":[{"type":"text","text":"First "},{"type":"text","text":"Second"}]}}"#;
        let result = parse_stream_event(json).unwrap();
        // No usage, so Skip (text already streamed via content_block_delta)
        assert!(matches!(result, ParsedStreamEvent::Skip));
    }

    #[test]
    fn test_parse_stream_event_malformed_json() {
        // Malformed JSON is now skipped instead of erroring (for stream robustness)
        let result = parse_stream_event("{invalid json}").unwrap();
        assert!(matches!(result, ParsedStreamEvent::Skip));
    }

    // Tests for StreamEvent wrapper (--include-partial-messages)

    #[test]
    fn test_parse_stream_event_wrapper_content_block_delta() {
        let json = r#"{"type":"stream_event","event":{"type":"content_block_delta","index":0,"delta":{"type":"text_delta","text":"Hello"}},"session_id":"abc"}"#;
        let result = parse_stream_event(json).unwrap();
        match result {
            ParsedStreamEvent::Text(text) => assert_eq!(text, "Hello"),
            _ => panic!("Expected Text, got {:?}", result),
        }
    }

    #[test]
    fn test_parse_stream_event_wrapper_content_block_start() {
        let json = r#"{"type":"stream_event","event":{"type":"content_block_start","index":0,"content_block":{"type":"text","text":"Start"}},"session_id":"abc"}"#;
        let result = parse_stream_event(json).unwrap();
        match result {
            ParsedStreamEvent::Text(text) => assert_eq!(text, "Start"),
            _ => panic!("Expected Text, got {:?}", result),
        }
    }

    #[test]
    fn test_parse_stream_event_wrapper_content_block_start_empty() {
        let json = r#"{"type":"stream_event","event":{"type":"content_block_start","index":0,"content_block":{"type":"text","text":""}},"session_id":"abc"}"#;
        let result = parse_stream_event(json).unwrap();
        assert!(matches!(result, ParsedStreamEvent::Skip));
    }

    #[test]
    fn test_parse_stream_event_wrapper_message_delta_with_usage() {
        let json = r#"{"type":"stream_event","event":{"type":"message_delta","delta":{"stop_reason":"end_turn"},"usage":{"input_tokens":100,"output_tokens":50}},"session_id":"abc"}"#;
        let result = parse_stream_event(json).unwrap();
        match result {
            ParsedStreamEvent::Usage(usage) => {
                assert_eq!(usage.input_tokens, Some(100));
                assert_eq!(usage.output_tokens, Some(50));
                assert_eq!(usage.total_tokens, Some(150));
            }
            _ => panic!("Expected Usage, got {:?}", result),
        }
    }

    #[test]
    fn test_parse_stream_event_wrapper_message_stop() {
        let json = r#"{"type":"stream_event","event":{"type":"message_stop"},"session_id":"abc"}"#;
        let result = parse_stream_event(json).unwrap();
        assert!(matches!(result, ParsedStreamEvent::Skip));
    }

    #[test]
    fn test_parse_stream_event_wrapper_message_start() {
        let json = r#"{"type":"stream_event","event":{"type":"message_start","message":{"id":"msg_1","role":"assistant","content":[]}},"session_id":"abc"}"#;
        let result = parse_stream_event(json).unwrap();
        assert!(matches!(result, ParsedStreamEvent::Skip));
    }

    #[test]
    fn test_parse_stream_event_wrapper_content_block_stop() {
        let json = r#"{"type":"stream_event","event":{"type":"content_block_stop","index":0},"session_id":"abc"}"#;
        let result = parse_stream_event(json).unwrap();
        assert!(matches!(result, ParsedStreamEvent::Skip));
    }

    #[test]
    fn test_parse_stream_event_wrapper_tool_use() {
        let json = r#"{"type":"stream_event","event":{"type":"content_block_start","index":1,"content_block":{"type":"tool_use","id":"tool_1","name":"read_file"}},"session_id":"abc"}"#;
        let result = parse_stream_event(json).unwrap();
        // Tool use blocks should emit ToolCallStart
        match result {
            ParsedStreamEvent::ToolCallStart { name, id } => {
                assert_eq!(name, "read_file");
                assert_eq!(id, "tool_1");
            }
            other => panic!("Expected ToolCallStart, got {:?}", other),
        }
    }
}
