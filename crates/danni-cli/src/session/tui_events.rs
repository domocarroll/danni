/// TUI Event System - JSON Protocol for Go TUI ↔ Rust CLI Communication
///
/// This module defines the event types and serialization for communication
/// between the Go TUI frontend and Rust CLI backend via JSON over stdin/stdout.

use serde::{Deserialize, Serialize};
use std::collections::HashMap;

// ============================================================================
// OUTGOING EVENTS (Rust CLI → Go TUI)
// ============================================================================

/// Main event envelope sent from CLI to TUI
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(tag = "type", rename_all = "snake_case")]
pub enum TuiEvent {
    /// Assistant is generating a response
    Thinking {
        #[serde(skip_serializing_if = "Option::is_none")]
        message: Option<String>,
    },

    /// Text content being streamed
    MessageChunk {
        content: String,
        #[serde(skip_serializing_if = "Option::is_none")]
        role: Option<String>,
    },

    /// Complete message received
    MessageComplete {
        content: String,
        role: String,
    },

    /// Tool is being executed
    ToolUse {
        tool_name: String,
        #[serde(skip_serializing_if = "Option::is_none")]
        tool_input: Option<serde_json::Value>,
    },

    /// Tool execution result
    ToolResult {
        tool_name: String,
        #[serde(skip_serializing_if = "Option::is_none")]
        success: Option<bool>,
        #[serde(skip_serializing_if = "Option::is_none")]
        output: Option<String>,
    },

    /// Token usage update
    TokenUpdate {
        tokens_used: i32,
        tokens_total: i32,
        estimated_cost: f64,
    },

    /// Error occurred
    Error {
        message: String,
        #[serde(skip_serializing_if = "Option::is_none")]
        details: Option<String>,
    },

    /// Session information
    SessionInfo {
        session_id: String,
        #[serde(skip_serializing_if = "Option::is_none")]
        module: Option<String>,
    },

    /// Response generation complete
    Complete {
        #[serde(skip_serializing_if = "Option::is_none")]
        message: Option<String>,
    },

    /// System status update
    Status {
        message: String,
    },

    /// Ready to accept input
    Ready,
}

// ============================================================================
// INCOMING REQUESTS (Go TUI → Rust CLI)
// ============================================================================

/// Request from TUI to CLI
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(tag = "type", rename_all = "snake_case")]
pub enum TuiRequest {
    /// Send a message to the agent
    Message {
        content: String,
        #[serde(skip_serializing_if = "Option::is_none")]
        module: Option<String>,
    },

    /// Stop current generation
    Stop,

    /// Change active module
    SetModule {
        module: String,
    },

    /// Ping for connectivity check
    Ping,
}

// ============================================================================
// HELPER FUNCTIONS
// ============================================================================

impl TuiEvent {
    /// Serialize event to JSON line for stdout
    pub fn to_json_line(&self) -> anyhow::Result<String> {
        let json = serde_json::to_string(self)?;
        Ok(format!("{}\n", json))
    }

    /// Create a thinking event
    pub fn thinking(message: Option<String>) -> Self {
        TuiEvent::Thinking { message }
    }

    /// Create a message chunk event
    pub fn message_chunk(content: impl Into<String>) -> Self {
        TuiEvent::MessageChunk {
            content: content.into(),
            role: None,
        }
    }

    /// Create a complete message event
    pub fn message_complete(content: impl Into<String>, role: impl Into<String>) -> Self {
        TuiEvent::MessageComplete {
            content: content.into(),
            role: role.into(),
        }
    }

    /// Create a tool use event
    pub fn tool_use(tool_name: impl Into<String>, tool_input: Option<serde_json::Value>) -> Self {
        TuiEvent::ToolUse {
            tool_name: tool_name.into(),
            tool_input,
        }
    }

    /// Create a tool result event
    pub fn tool_result(tool_name: impl Into<String>, success: bool, output: Option<String>) -> Self {
        TuiEvent::ToolResult {
            tool_name: tool_name.into(),
            success: Some(success),
            output,
        }
    }

    /// Create a token update event
    pub fn token_update(tokens_used: i32, tokens_total: i32, estimated_cost: f64) -> Self {
        TuiEvent::TokenUpdate {
            tokens_used,
            tokens_total,
            estimated_cost,
        }
    }

    /// Create an error event
    pub fn error(message: impl Into<String>) -> Self {
        TuiEvent::Error {
            message: message.into(),
            details: None,
        }
    }

    /// Create a session info event
    pub fn session_info(session_id: impl Into<String>, module: Option<String>) -> Self {
        TuiEvent::SessionInfo {
            session_id: session_id.into(),
            module,
        }
    }

    /// Create a complete event
    pub fn complete() -> Self {
        TuiEvent::Complete { message: None }
    }

    /// Create a status event
    pub fn status(message: impl Into<String>) -> Self {
        TuiEvent::Status {
            message: message.into(),
        }
    }

    /// Create a ready event
    pub fn ready() -> Self {
        TuiEvent::Ready
    }
}

impl TuiRequest {
    /// Parse a JSON line from stdin into a request
    pub fn from_json_line(line: &str) -> anyhow::Result<Self> {
        let request = serde_json::from_str(line)?;
        Ok(request)
    }
}

// ============================================================================
// EVENT EMITTER
// ============================================================================

/// Helper for emitting TUI events to stdout
pub struct TuiEventEmitter {
    enabled: bool,
}

impl TuiEventEmitter {
    pub fn new(enabled: bool) -> Self {
        Self { enabled }
    }

    /// Emit an event to stdout if TUI mode is enabled
    pub fn emit(&self, event: TuiEvent) -> anyhow::Result<()> {
        if !self.enabled {
            return Ok(());
        }

        let json_line = event.to_json_line()?;
        print!("{}", json_line);
        use std::io::Write;
        std::io::stdout().flush()?;
        Ok(())
    }

    /// Emit multiple events
    pub fn emit_all(&self, events: Vec<TuiEvent>) -> anyhow::Result<()> {
        for event in events {
            self.emit(event)?;
        }
        Ok(())
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_event_serialization() {
        let event = TuiEvent::thinking(Some("Processing...".to_string()));
        let json = event.to_json_line().unwrap();
        assert!(json.contains("thinking"));
        assert!(json.contains("Processing"));
    }

    #[test]
    fn test_request_deserialization() {
        let json = r#"{"type":"message","content":"Hello"}"#;
        let request = TuiRequest::from_json_line(json).unwrap();
        match request {
            TuiRequest::Message { content, .. } => {
                assert_eq!(content, "Hello");
            }
            _ => panic!("Wrong request type"),
        }
    }
}
