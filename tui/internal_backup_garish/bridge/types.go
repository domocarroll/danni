package bridge

// Message types for communication between Go TUI and Rust CLI

// Request represents a message sent from the TUI to the CLI
type Request struct {
	Type    string                 `json:"type"`
	Content string                 `json:"content,omitempty"`
	Module  string                 `json:"module,omitempty"`
	Params  map[string]interface{} `json:"params,omitempty"`
}

// Response represents a message sent from the CLI to the TUI
type Response struct {
	Type    string                 `json:"type"`
	Content string                 `json:"content,omitempty"`
	Error   string                 `json:"error,omitempty"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

// Event types
const (
	// Request types
	RequestTypeMessage   = "message"
	RequestTypeStop      = "stop"
	RequestTypeModuleSet = "module_set"

	// Response types (matching Rust TuiEvent types)
	ResponseTypeMessage       = "message_chunk"
	ResponseTypeMessageComplete = "message_complete"
	ResponseTypeThinking      = "thinking"
	ResponseTypeComplete      = "complete"
	ResponseTypeError         = "error"
	ResponseTypeToolUse       = "tool_use"
	ResponseTypeToolResult    = "tool_result"
	ResponseTypeToken         = "token_update"
	ResponseTypeSessionInfo   = "session_info"
	ResponseTypeStatus        = "status"
	ResponseTypeReady         = "ready"
)
