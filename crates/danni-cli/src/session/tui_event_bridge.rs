/// TUI Event Bridge - Maps AgentEvents to TUI JSON Events
///
/// This module provides helper methods to emit TUI events based on AgentEvents

use super::tui_events::TuiEvent;
use super::CliSession;
use danni::agents::AgentEvent;
use danni::conversation::message::{Message, MessageContent};

impl CliSession {
    /// Emit TUI event for an AgentEvent if in TUI mode
    pub fn emit_agent_event(&self, event: &AgentEvent) -> anyhow::Result<()> {
        if !self.tui_mode {
            return Ok(());
        }

        match event {
            AgentEvent::Message(message) => {
                self.emit_message_event(message)?;
            }
            AgentEvent::McpNotification(_) => {
                // MCP notifications are handled separately
            }
            AgentEvent::ModelChange { .. } => {
                // Model changes don't need TUI events
            }
            AgentEvent::HistoryReplaced(_) => {
                // History replacement doesn't need TUI events
            }
        }

        Ok(())
    }

    /// Emit TUI events for a message
    fn emit_message_event(&self, message: &Message) -> anyhow::Result<()> {
        for content in &message.content {
            match content {
                MessageContent::Text(text) => {
                    self.tui_emitter.emit(TuiEvent::message_chunk(text.text.to_string()))?;
                }

                MessageContent::ToolRequest(tool_req) => {
                    if let Ok(tool_call) = &tool_req.tool_call {
                        let tool_input = serde_json::to_value(&tool_call.arguments).ok();
                        self.tui_emitter.emit(TuiEvent::tool_use(
                            tool_call.name.as_ref(),
                            tool_input,
                        ))?;
                    }
                }

                MessageContent::ToolResponse(tool_resp) => {
                    // Find the tool name from the corresponding request
                    let tool_name = "tool"; // TODO: Look up actual tool name
                    let success = tool_resp.tool_result.is_ok();
                    let output = match &tool_resp.tool_result {
                        Ok(content) => Some(format!("{:?}", content)),
                        Err(err) => Some(format!("Error: {}", err.message)),
                    };

                    self.tui_emitter.emit(TuiEvent::tool_result(
                        tool_name,
                        success,
                        output,
                    ))?;
                }

                MessageContent::ToolConfirmationRequest(_) => {
                    // Tool confirmations are handled specially in interactive mode
                }

                _ => {
                    // Other content types
                }
            }
        }

        Ok(())
    }

    /// Emit thinking state
    pub fn emit_thinking(&self, message: Option<String>) -> anyhow::Result<()> {
        if self.tui_mode {
            self.tui_emitter.emit(TuiEvent::thinking(message))?;
        }
        Ok(())
    }

    /// Emit status message
    pub fn emit_status(&self, message: impl Into<String>) -> anyhow::Result<()> {
        if self.tui_mode {
            self.tui_emitter.emit(TuiEvent::status(message.into()))?;
        }
        Ok(())
    }

    /// Emit error
    pub fn emit_error(&self, message: impl Into<String>) -> anyhow::Result<()> {
        if self.tui_mode {
            self.tui_emitter.emit(TuiEvent::error(message.into()))?;
        }
        Ok(())
    }

    /// Emit complete
    pub fn emit_complete(&self) -> anyhow::Result<()> {
        if self.tui_mode {
            self.tui_emitter.emit(TuiEvent::complete())?;
        }
        Ok(())
    }
}
