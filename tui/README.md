# DANNI TUI

A beautiful terminal user interface for DANNI, built with the Charmbracelet ecosystem.

## ✨ Features

- **Sophisticated Design**: Purple/midnight blue/rose gold color scheme reflecting Danni's personality
- **Scrollable Chat**: Smooth viewport navigation through conversation history
- **Multi-line Input**: Rich textarea for composing messages
- **Markdown Rendering**: Beautiful markdown rendering with Glamour
- **Modular Architecture**: Support for Danni's 9 specialized modules
- **Real-time Updates**: Streaming responses with thinking indicators
- **Keyboard Navigation**: Intuitive vim-style shortcuts

## 🎨 Design Philosophy

The Danni TUI embodies sophistication, warmth, and intelligence through:

- **Color Palette**:
  - Primary Purple (#9F7AEA) - Intelligence & sophistication
  - Midnight Blue (#1E3A5F) - Depth & trust
  - Rose Gold (#B76E79) - Warmth & elegance
  - Soft Cream (#F7F3E9) - Clarity & readability

- **Layout**: Clean separation of concerns with header, chat area, input, and footer
- **Typography**: Bold headers, subtle hints, clear message hierarchy
- **Animation**: Gentle spinner for thinking states

## 🚀 Quick Start

### Build

```bash
cd tui
go build -o bin/danni-tui cmd/danni-tui/main.go
```

### Run

```bash
./bin/danni-tui
```

## ⌨️ Keyboard Shortcuts

- **Ctrl+C** or **Esc** - Quit
- **Enter** - Send message
- **↑/↓** or **k/j** - Scroll chat
- **PgUp/PgDn** - Page through chat

## 🏗️ Architecture

### Project Structure

```
tui/
├── cmd/
│   └── danni-tui/
│       └── main.go           # Main application entry point
├── internal/
│   ├── app/                  # Application logic (future)
│   ├── bridge/              # Rust CLI communication
│   │   ├── types.go         # Message type definitions
│   │   └── cli.go           # CLI bridge implementation
│   ├── styles/              # Visual styling
│   │   ├── theme.go         # Color palette & component styles
│   │   └── glamour.go       # Markdown renderer theming
│   └── components/          # Reusable UI components
│       ├── header.go        # Top bar (logo, module, session)
│       ├── footer.go        # Bottom bar (tokens, cost, help)
│       └── spinner.go       # Thinking state indicator
├── bin/                     # Compiled binaries
└── README.md               # This file
```

### Communication Flow

```
User Input → TUI (Go) → Bridge → CLI (Rust) → Agent → LLM
               ↑                      ↓
               └────── Events ────────┘
                   (JSON stream)
```

### Bridge Protocol

The TUI communicates with the Rust CLI via JSON over stdin/stdout:

**Request Format:**
```json
{
  "type": "message",
  "content": "User's message here",
  "module": "/strategy"
}
```

**Response Format:**
```json
{
  "type": "message",
  "content": "DANNI's response here"
}
```

**Event Types:**
- `message` - User message or assistant response
- `thinking` - Assistant is processing
- `complete` - Response complete
- `error` - Error occurred
- `tool_use` - Tool execution event
- `token` - Token usage update

## 🎯 Current Status

### ✅ Implemented (Week 1)
- [x] Go project structure
- [x] Charmbracelet dependencies (Bubble Tea, Lip Gloss, Glamour, Bubbles)
- [x] Danni color theme
- [x] Layout structure (Header, Chat, Input, Footer)
- [x] Viewport for scrollable chat
- [x] Textarea for multi-line input
- [x] Bridge package for CLI communication
- [x] JSON protocol definitions
- [x] Glamour markdown rendering
- [x] Thinking state spinner

### 🚧 Next Steps (Week 2)
- [ ] Integrate with actual Rust CLI
- [ ] Handle streaming responses
- [ ] Module selector UI
- [ ] Token/cost tracking
- [ ] Error handling UI
- [ ] Help overlay
- [ ] Settings panel

### 🌟 Future Enhancements
- [ ] Extension manager UI
- [ ] Session history browser
- [ ] Rich tooltips
- [ ] Keyboard shortcuts help
- [ ] Theme customization
- [ ] Animation polish

## 🔧 Development

### Dependencies

All dependencies are managed via Go modules:

```bash
go mod tidy
```

### Building

Development build:
```bash
go build -o bin/danni-tui cmd/danni-tui/main.go
```

Release build:
```bash
go build -ldflags="-s -w" -o bin/danni-tui cmd/danni-tui/main.go
```

### Testing

```bash
go test ./...
```

## 📚 Resources

### Charmbracelet Ecosystem
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Styling
- [Glamour](https://github.com/charmbracelet/glamour) - Markdown rendering
- [Bubbles](https://github.com/charmbracelet/bubbles) - TUI components

### Documentation
- [DANNI_TUI_ARCHITECTURE.md](../DANNI_TUI_ARCHITECTURE.md) - Complete technical design
- [DANNI_TUI_MOCKUPS.md](../DANNI_TUI_MOCKUPS.md) - Visual interface designs
- [DANNI_TUI_QUICKSTART.md](../DANNI_TUI_QUICKSTART.md) - 2-week implementation guide

## 🎨 Design Credits

Color palette and visual design by SUBFRACTURE, reflecting Danni's sophisticated, warm, and intelligent personality.

## 📝 License

Same as parent Danni project.

---

✨ **Built with love using Charmbracelet** ✨
