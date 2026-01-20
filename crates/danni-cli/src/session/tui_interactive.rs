/// TUI Mode Interactive Session Handler
///
/// This module handles the interactive loop when running in TUI mode,
/// reading JSON requests from stdin and emitting JSON events to stdout.

use super::tui_events::{TuiEvent, TuiRequest};
use super::CliSession;
use anyhow::Result;
use danni::conversation::message::Message;
use std::io::Write;
use tokio::io::{AsyncBufReadExt, BufReader};
use tokio_util::sync::CancellationToken;

impl CliSession {
    /// Run interactive session in TUI mode
    ///
    /// Reads JSON requests from stdin, processes them, and emits JSON events to stdout
    pub async fn interactive_tui(&mut self, initial_prompt: Option<String>) -> Result<()> {
        // Emit session info
        self.tui_emitter.emit(TuiEvent::session_info(
            self.session_id.clone(),
            None,
        ))?;

        // Process initial message if provided
        if let Some(prompt) = initial_prompt {
            self.tui_emitter.emit(TuiEvent::status(
                "Processing initial message...".to_string(),
            ))?;

            // Use headless to process initial prompt
            self.headless(prompt).await?;

            // Emit the response
            if let Some(last_msg) = self.messages.last() {
                let response_text = last_msg.as_concat_text();
                if !response_text.is_empty() {
                    self.tui_emitter.emit(TuiEvent::message_complete(
                        response_text,
                        "assistant".to_string(),
                    ))?;
                }
            }
        }

        // Emit ready signal
        self.tui_emitter.emit(TuiEvent::ready())?;

        // Main event loop - read JSON requests from stdin asynchronously
        let stdin = tokio::io::stdin();
        let reader = BufReader::new(stdin);
        let mut lines = reader.lines();

        while let Some(line_result) = lines.next_line().await.transpose() {
            let line = match line_result {
                Ok(l) => l,
                Err(e) => {
                    self.tui_emitter.emit(TuiEvent::error(format!(
                        "Failed to read stdin: {}",
                        e
                    )))?;
                    continue;
                }
            };

            if line.trim().is_empty() {
                continue;
            }

            // Parse request
            let request = match TuiRequest::from_json_line(&line) {
                Ok(req) => req,
                Err(e) => {
                    self.tui_emitter.emit(TuiEvent::error(format!(
                        "Failed to parse request: {}",
                        e
                    )))?;
                    continue;
                }
            };

            // Handle request
            match request {
                TuiRequest::Message { content, module } => {
                    if let Err(e) = self.handle_tui_message(content, module).await {
                        self.tui_emitter
                            .emit(TuiEvent::error(format!("Error processing message: {}", e)))?;
                    }
                    self.tui_emitter.emit(TuiEvent::ready())?;
                }

                TuiRequest::Stop => {
                    self.tui_emitter
                        .emit(TuiEvent::status("Stopping...".to_string()))?;
                    // TODO: Implement cancellation
                    break;
                }

                TuiRequest::SetModule { module } => {
                    self.tui_emitter.emit(TuiEvent::status(format!(
                        "Module set to: {}",
                        module
                    )))?;
                    // TODO: Implement module switching
                }

                TuiRequest::Ping => {
                    self.tui_emitter
                        .emit(TuiEvent::status("pong".to_string()))?;
                }
            }
        }

        Ok(())
    }

    /// Handle a message request in TUI mode
    async fn handle_tui_message(
        &mut self,
        content: String,
        _module: Option<String>,
    ) -> Result<()> {
        // Emit thinking state
        self.tui_emitter.emit(TuiEvent::thinking(None))?;

        // Use headless() which processes the message and returns response
        match self.headless(content).await {
            Ok(_) => {
                // Get the last message (should be assistant's response)
                if let Some(last_msg) = self.messages.last() {
                    let response_text = last_msg.as_concat_text();
                    if !response_text.is_empty() {
                        self.tui_emitter.emit(TuiEvent::message_complete(
                            response_text,
                            "assistant".to_string(),
                        ))?;
                    }
                }
            }
            Err(e) => {
                self.tui_emitter.emit(TuiEvent::error(format!("Error: {}", e)))?;
                return Err(e);
            }
        }

        // Update token usage
        if let Ok(session) = self.get_session().await {
            if let Some(total_tokens) = session.total_tokens {
                self.tui_emitter.emit(TuiEvent::token_update(
                    total_tokens,
                    100000, // TODO: Get actual token limit
                    0.0,    // TODO: Calculate cost
                ))?;
            }
        }

        // Emit complete
        self.tui_emitter.emit(TuiEvent::complete())?;

        Ok(())
    }
}
