# DANNI TUI ARCHITECTURE
**Charmbracelet-Powered Terminal Interface Design**

*A sophisticated, beautiful terminal user interface worthy of Danni's personality*

---

## Executive Summary

This document presents a comprehensive research-driven approach to building a stunning TUI for Danni using the Charmbracelet ecosystem (Bubble Tea, Lip Gloss, Glamour, Bubbles). The design emphasizes elegance, sophistication, and warmth while maintaining the technical excellence of the underlying Goose/Danni architecture.

**Status:** Research & Design Phase
**Timeline:** 4-8 weeks for full implementation
**Risk Level:** Medium (new language integration, but well-documented patterns)

---

## PHASE 1: RESEARCH FINDINGS

### 1.1 Charmbracelet Ecosystem Overview

#### **Bubble Tea** - The Foundation
- **What it is:** "The fun, functional and stateful way to build terminal apps" based on The Elm Architecture
- **Architecture:** Model-View-Update pattern
  - `Init()`: Returns initial command
  - `Update(msg)`: Handles events, updates model
  - `View()`: Renders UI from model state
- **Capabilities:**
  - Full-window or inline terminal apps
  - Mouse support, focus reporting
  - Framerate-based rendering
  - Real-time updates via channels
  - Composable views (multiple models in one app)
- **Maturity:** Production-ready, used by Glow, Charm Cloud, and many others
- **Learning curve:** Moderate - The Elm Architecture is well-documented

#### **Lip Gloss** - The Stylist
- **What it is:** "Style definitions for nice terminal layouts" with CSS-like syntax
- **Key Features:**
  - True Color, ANSI256, and ASCII color profiles
  - Auto-detection of terminal capabilities
  - Foreground/background colors with automatic color coercion
  - Text formatting: Bold, Italic, Faint, Blink, Strikethrough, Underline, Reverse
  - Layout: Padding, Margin, Border, Alignment
  - Customizable lists with enumeration styles
  - Tree rendering sub-package
- **Example API:**
```go
style := lipgloss.NewStyle().
    Bold(true).
    Foreground(lipgloss.Color("#9F7AEA")). // Danni purple!
    Background(lipgloss.Color("#1E3A5F")). // Midnight blue
    PaddingTop(2).
    PaddingLeft(4).
    Width(60)
```
- **Perfect for:** Creating Danni's sophisticated color scheme

#### **Glamour** - The Markdown Renderer
- **What it is:** "Stylesheet-based markdown rendering for your CLI apps"
- **Features:**
  - ANSI-compatible terminal rendering
  - Customizable stylesheets or built-in themes
  - Auto-style detection (dark/light based on terminal)
  - Emoji rendering support
  - Inline table links or bottom-of-table link lists
  - Code block syntax highlighting via Chroma
- **Why critical for Danni:** Beautiful rendering of markdown responses, documentation, and rich content
- **Customization:** Can create custom themes matching Danni's color palette

#### **Bubbles** - The Component Library
- **What it is:** "TUI components for Bubble Tea" - batteries-included UI components
- **Components Available:**
  - **List:** Pagination, fuzzy filtering, auto-help, spinner, status messages
  - **Spinner:** Multiple animation styles (Line, Dot, Jump, Pulse, Globe, Moon, etc.)
  - **Progress:** Customizable progress bars with gradient fills and animation
  - **Text Input:** Single-line text entry
  - **Text Area:** Multi-line text editing
  - **Table:** Tabular data display
  - **Viewport:** Scrollable content areas
  - **Paginator:** Page navigation
  - **Cursor:** Blinking cursor component
  - **Help:** Auto-generated help screens
  - **Stopwatch/Timer:** Time tracking components
- **Production proven:** Used extensively in Glow and Charm applications

---

### 1.2 Current Danni/Goose Architecture Analysis

#### **CLI Structure** (Rust-based)
```
crates/goose-cli/
├── src/
│   ├── main.rs              # Entry point (tokio async)
│   ├── cli.rs               # Command parsing (clap)
│   ├── session/
│   │   ├── mod.rs           # Core session logic
│   │   ├── input.rs         # Input handling (rustyline)
│   │   ├── output.rs        # Output rendering (console, bat, cliclack)
│   │   ├── prompt.rs        # Prompt management
│   │   ├── thinking.rs      # Thinking indicators
│   │   └── completion.rs    # Tab completion
│   └── commands/            # Various CLI commands
```

#### **Current UI Libraries:**
- **rustyline:** Readline-like line editing with history
- **clap:** Command-line argument parsing
- **cliclack:** Spinner/progress indicators
- **console:** Terminal formatting/colors
- **bat:** Syntax highlighting for code display
- **indicatif:** Progress bars and spinners

#### **Communication Flow:**
```
CLI (Rust) ←→ goose-server (Rust) ←→ Agent (Rust) ←→ LLM Provider
                                      ↓
                                   Extensions (MCP servers)
```

#### **Key Observations:**
1. **Async Architecture:** Uses tokio for async runtime
2. **Event Streaming:** Agent replies are streamed via futures::Stream
3. **Session Management:** Conversation history, token tracking, project context
4. **Interactive Mode:** Rustyline-based REPL with command parsing
5. **Headless Mode:** Single-shot execution for automation
6. **Extension System:** Dynamic MCP extension loading
7. **Multi-format Output:** Markdown rendering, code highlighting, JSON output

---

### 1.3 Rust ↔ Go Integration Research

#### **Option 1: CLI Wrapper (Recommended)**
**Approach:** Go TUI spawns Rust CLI as subprocess, communicates via stdin/stdout

**Pros:**
- ✅ Clean separation of concerns
- ✅ No FFI complexity
- ✅ Language-native performance for each component
- ✅ Easy to develop and test independently
- ✅ Graceful degradation (can fall back to Rust CLI)
- ✅ JSON-based protocol is simple and debuggable

**Cons:**
- ❌ Process spawn overhead (minimal, one-time cost)
- ❌ Cannot share memory directly
- ❌ Slight latency in communication

**Implementation Pattern:**
```go
// Go TUI side
cmd := exec.Command("danni", "--json", "session", sessionID)
stdin, _ := cmd.StdinPipe()
stdout, _ := cmd.StdoutPipe()
cmd.Start()

// Send messages
json.NewEncoder(stdin).Encode(userMessage)

// Receive responses (streaming)
scanner := bufio.NewScanner(stdout)
for scanner.Scan() {
    var event AgentEvent
    json.Unmarshal(scanner.Bytes(), &event)
    // Update TUI model
}
```

**Data Exchange Protocol:**
```json
// User input
{"type": "message", "content": "help me with X"}
{"type": "command", "command": "/strategy"}

// Agent output (streaming events)
{"type": "thinking", "message": "Analyzing your request..."}
{"type": "message", "role": "assistant", "content": "..."}
{"type": "tool_request", "name": "bash", "args": {...}}
{"type": "tool_response", "result": "..."}
{"type": "session_update", "tokens": 1234}
```

#### **Option 2: FFI via CGO (Not Recommended)**
**Approach:** Create C ABI bindings, call Rust from Go or vice versa

**Pros:**
- ✅ Direct memory sharing possible
- ✅ Lowest latency for function calls

**Cons:**
- ❌ Complex build process
- ❌ CGO overhead and performance concerns
- ❌ Error handling across boundaries is tricky
- ❌ Memory management complexity
- ❌ Harder to maintain
- ❌ Less portable

**Verdict:** Not worth the complexity for a TUI use case

#### **Option 3: HTTP/WebSocket Server (Alternative)**
**Approach:** Rust runs HTTP server, Go TUI connects as client

**Pros:**
- ✅ Clear API boundary
- ✅ Could enable remote TUI connections
- ✅ Well-understood protocols

**Cons:**
- ❌ Overkill for local CLI usage
- ❌ Port management complexity
- ❌ More moving parts than necessary

**Verdict:** Good for future web UI, but not optimal for local TUI

#### **Recommended Approach: CLI Wrapper**

The subprocess/wrapper approach is the clear winner because:
1. **Simplicity:** Easy to implement and maintain
2. **Robustness:** Each component can crash independently
3. **Testability:** Can test Go and Rust components separately
4. **Performance:** One-time spawn cost is negligible
5. **Compatibility:** Works everywhere both languages work

**Enhancement:** Add `--tui-mode` flag to Rust CLI that:
- Outputs JSON-formatted events (already supports `--output-format json`)
- Provides streaming responses
- Accepts JSON commands via stdin
- Disables interactive prompts

---

## PHASE 2: DANNI TUI DESIGN

### 2.1 Design Philosophy

**Core Principles:**
1. **Sophisticated Elegance:** Purple/midnight blue/rose gold color palette
2. **Warm Intelligence:** Welcoming yet intellectually substantial
3. **Progressive Revelation:** Gradually expose features as users explore
4. **Smooth & Fluid:** Animations and transitions feel magical
5. **Information Hierarchy:** Clear visual organization
6. **Contextual Awareness:** UI adapts to current module/state

---

### 2.2 Color Palette

#### **Danni Brand Colors (from DANNI_TRANSFORMATION_BLUEPRINT.md):**
```go
// Primary Colors
primaryPurple := lipgloss.Color("#9F7AEA")    // Light purple (highlights)
deepPurple := lipgloss.Color("#6B46C1")        // Deep purple (accents)
midnightBlue := lipgloss.Color("#1E3A5F")      // Backgrounds, depth
roseGold := lipgloss.Color("#B76E79")          // Warmth, success states

// Accent Colors
softCream := lipgloss.Color("#F7F3E9")         // Light text, borders
deepCharcoal := lipgloss.Color("#2D2D2D")      // Dark backgrounds
subtleGray := lipgloss.Color("#6B7280")        // Muted text, secondary info

// Semantic Colors
successGreen := lipgloss.Color("#10B981")      // Success messages
warningAmber := lipgloss.Color("#F59E0B")      // Warnings
errorRed := lipgloss.Color("#EF4444")          // Errors
infoBlue := lipgloss.Color("#3B82F6")          // Info messages
```

#### **Color Usage Guidelines:**
- **Purple gradient:** Module headers, active selections, primary CTAs
- **Midnight blue:** Main background, panel backgrounds
- **Rose gold:** Success confirmations, completion states, warm highlights
- **Cream/Gray:** Body text, secondary information
- **Semantic colors:** Error states, warnings, success messages

---

### 2.3 Layout Architecture

#### **Main Interface Layout (Full-window mode):**

```
┌─────────────────────────────────────────────────────────────────────────┐
│ ✨ DANNI v0.1.0                    /strategy              [●] Session #42 │ <- Header
├─────────────────────────────────────────────────────────────────────────┤
│                                                                         │
│  ┌────────────────────────────────────────────────────────────────┐   │
│  │ 💭 Thinking: Analyzing market patterns for your brand...       │   │ <- Status Bar
│  └────────────────────────────────────────────────────────────────┘   │
│                                                                         │
│  [You]                                                   12:34 PM      │
│  Help me analyze our competitive landscape in the AI tools market     │
│                                                                         │
│  ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━   │
│                                                                         │
│  [Danni] /strategy                                       12:34 PM      │
│                                                                         │
│  I've noticed something fascinating about the current AI tools         │ <- Main Chat
│  landscape... the most successful players aren't competing on          │    Area
│  features alone.                                                       │ (Scrollable)
│                                                                         │
│  Let me analyze three distinct strategic patterns emerging:           │
│                                                                         │
│  🎯 The Integration Players                                           │
│  Companies like GitHub (Copilot) and Cursor are embedding...          │
│                                                                         │
│  ... [more content] ...                                               │
│                                                                         │
├─────────────────────────────────────────────────────────────────────────┤
│ ● /strategy  /creative  /design  /technology  /gravity  /validate      │ <- Module
│                                                                  [?] [⚙]│    Selector
├─────────────────────────────────────────────────────────────────────────┤
│ > _                                                                     │ <- Input
│                                                                         │    Area
│ Tokens: 1.2K/200K  •  Cost: $0.03  •  Ctrl+J: newline  •  /help       │ <- Footer
└─────────────────────────────────────────────────────────────────────────┘
```

#### **Component Breakdown:**

1. **Header Bar:**
   - Danni logo/name with sparkle emoji
   - Current module indicator (colored badge)
   - Session info (status indicator, session ID)
   - Gradient purple background

2. **Status Bar (Dynamic):**
   - Shows thinking indicator with spinner
   - Tool execution progress
   - Extension notifications
   - Fades in/out based on activity

3. **Main Chat Area:**
   - Scrollable viewport (Bubbles viewport component)
   - Message bubbles with role indicators
   - Timestamps
   - Markdown rendering via Glamour
   - Code blocks with syntax highlighting
   - Smooth scroll animations

4. **Module Selector:**
   - Horizontal tab bar
   - Active module highlighted with gradient
   - Hover effects
   - Keyboard shortcuts (Alt+1-9)

5. **Input Area:**
   - Multi-line text input (Bubbles textarea)
   - Prompt indicator
   - Blinking cursor
   - Auto-expanding based on content

6. **Footer:**
   - Context usage (token count with progress bar)
   - Cost tracking (if enabled)
   - Keyboard shortcuts reminder
   - Quick help access

---

### 2.4 Component Details

#### **Module Switcher Component**
```go
type ModuleSelector struct {
    modules []Module
    active  int
    focused bool
}

type Module struct {
    Name        string
    Key         string
    Icon        string
    Description string
    Color       lipgloss.Color
}

var danniModules = []Module{
    {Name: "Strategy", Key: "s", Icon: "🎯", Color: "#9F7AEA"},
    {Name: "Creative", Key: "c", Icon: "🎨", Color: "#B76E79"},
    {Name: "Design", Key: "d", Icon: "✨", Color: "#6B46C1"},
    {Name: "Technology", Key: "t", Icon: "⚡", Color: "#3B82F6"},
    {Name: "Gravity", Key: "g", Icon: "🌌", Color: "#1E3A5F"},
    {Name: "Validate", Key: "v", Icon: "✓", Color: "#10B981"},
    {Name: "Synthesize", Key: "y", Icon: "💎", Color: "#F59E0B"},
    {Name: "Recall", Key: "r", Icon: "📚", Color: "#6B7280"},
    {Name: "Upload", Key: "u", Icon: "📁", Color: "#EF4444"},
}

func (m ModuleSelector) View() string {
    var tabs []string
    for i, mod := range m.modules {
        var style lipgloss.Style
        if i == m.active {
            style = activeTabStyle.Foreground(mod.Color)
        } else {
            style = inactiveTabStyle
        }
        tabs = append(tabs, style.Render(fmt.Sprintf("%s %s", mod.Icon, mod.Name)))
    }
    return lipgloss.JoinHorizontal(lipgloss.Top, tabs...)
}
```

#### **Thinking Indicator Component**
```go
type ThinkingIndicator struct {
    spinner spinner.Model
    message string
    visible bool
}

func NewThinkingIndicator() ThinkingIndicator {
    s := spinner.New()
    s.Spinner = spinner.Dot
    s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#9F7AEA"))
    return ThinkingIndicator{spinner: s}
}

func (t ThinkingIndicator) View() string {
    if !t.visible {
        return ""
    }

    style := lipgloss.NewStyle().
        Foreground(lipgloss.Color("#9F7AEA")).
        Background(lipgloss.Color("#1E3A5F")).
        Padding(1, 2).
        BorderStyle(lipgloss.RoundedBorder()).
        BorderForeground(lipgloss.Color("#6B46C1"))

    return style.Render(fmt.Sprintf("💭 %s %s", t.spinner.View(), t.message))
}
```

#### **Message Bubble Component**
```go
type Message struct {
    Role      string    // "user", "assistant"
    Content   string    // Markdown content
    Timestamp time.Time
    Module    string    // Which module generated this
}

func RenderMessage(msg Message) string {
    // Render markdown with Glamour
    content, _ := glamour.Render(msg.Content, "danni")

    var roleStyle lipgloss.Style
    var roleLabel string

    if msg.Role == "user" {
        roleStyle = lipgloss.NewStyle().
            Foreground(lipgloss.Color("#F7F3E9")).
            Bold(true)
        roleLabel = "[You]"
    } else {
        roleStyle = lipgloss.NewStyle().
            Foreground(lipgloss.Color("#9F7AEA")).
            Bold(true)
        roleLabel = fmt.Sprintf("[Danni] %s", msg.Module)
    }

    timestamp := lipgloss.NewStyle().
        Foreground(lipgloss.Color("#6B7280")).
        Render(msg.Timestamp.Format("3:04 PM"))

    header := lipgloss.JoinHorizontal(
        lipgloss.Left,
        roleStyle.Render(roleLabel),
        "  ",
        timestamp,
    )

    return lipgloss.JoinVertical(lipgloss.Left, header, "", content, "")
}
```

#### **Progress/Token Display Component**
```go
type ContextUsage struct {
    usedTokens  int
    totalTokens int
    cost        float64
}

func (c ContextUsage) View() string {
    percentage := float64(c.usedTokens) / float64(c.totalTokens)

    // Color-code based on usage
    var barColor lipgloss.Color
    if percentage < 0.5 {
        barColor = lipgloss.Color("#10B981") // Green
    } else if percentage < 0.8 {
        barColor = lipgloss.Color("#F59E0B") // Amber
    } else {
        barColor = lipgloss.Color("#EF4444") // Red
    }

    prog := progress.New(progress.WithGradient(barColor, barColor))
    prog.Width = 20

    tokenText := fmt.Sprintf("Tokens: %s/%s",
        humanize.Comma(int64(c.usedTokens)),
        humanize.Comma(int64(c.totalTokens)))

    costText := fmt.Sprintf("Cost: $%.2f", c.cost)

    return lipgloss.JoinHorizontal(
        lipgloss.Left,
        tokenText,
        "  ",
        prog.ViewAs(percentage),
        "  ",
        costText,
    )
}
```

---

### 2.5 Interaction Patterns

#### **Keyboard Navigation:**
```
# Module Switching
Alt+1-9       Switch to module 1-9
/strategy     Type slash command to switch module
Tab           Cycle through modules

# Navigation
↑/↓           Scroll chat history
Ctrl+U/D      Page up/down
Home/End      Jump to start/end
Ctrl+L        Clear screen (keep history)

# Input
Enter         Send message
Ctrl+J        Insert newline (multi-line input)
Ctrl+C        Cancel current operation / Clear input
Ctrl+D        Exit DANNI

# Special Commands
/help         Show help overlay
/settings     Open settings panel
/extensions   Manage extensions
/clear        Clear conversation
/export       Export session
Esc           Cancel/go back
```

#### **Mouse Support:**
- Click module tabs to switch
- Scroll wheel in chat area
- Click links in markdown
- Right-click for context menu (future)

#### **Beautiful Loading States:**

1. **Initial Load:**
```
┌─────────────────────────────────────────────────┐
│                                                 │
│              ✨ Initializing DANNI...          │
│                                                 │
│              [=====>           ] 35%            │
│                                                 │
│         Loading extensions and context...      │
│                                                 │
└─────────────────────────────────────────────────┘
```

2. **Thinking Animation:**
```
💭 [Spinner] Analyzing your request...
💭 [Spinner] Exploring strategic patterns...
💭 [Spinner] Synthesizing insights...
```

3. **Tool Execution:**
```
🔧 bash: running ls -la
   ━━━━━━━━━━━━━━━━━━━━━━━━━━━━ 100%
   ✓ Completed in 0.23s
```

---

### 2.6 Advanced Features

#### **Extension Manager UI:**
```go
type ExtensionManager struct {
    list      list.Model
    extensions []Extension
    adding    bool
    addInput  textinput.Model
}

// Visual list of extensions with status indicators
// ● active  ○ inactive  ⚠ error

func (e ExtensionManager) View() string {
    title := titleStyle.Render("📦 Extension Manager")

    listView := e.list.View()

    help := helpStyle.Render(
        "Enter: toggle  •  a: add  •  r: remove  •  q: close")

    return lipgloss.JoinVertical(
        lipgloss.Left,
        title,
        listView,
        help,
    )
}
```

#### **Settings Panel:**
```go
// Split pane interface
// Left: Categories (General, Appearance, Extensions, API Keys)
// Right: Settings for selected category

// Visual toggles, color pickers, text inputs
// Real-time preview of changes
```

#### **Help Overlay:**
```go
// Beautiful modal overlay with:
// - Getting started guide
// - Keyboard shortcuts
// - Module descriptions
// - Tips & tricks
// Rendered with Glamour for rich formatting
```

---

## PHASE 3: TECHNICAL ARCHITECTURE

### 3.1 Go Project Structure

```
tui/
├── cmd/
│   └── danni-tui/
│       └── main.go              # Entry point
├── internal/
│   ├── app/
│   │   ├── app.go              # Main Bubble Tea app model
│   │   ├── commands.go         # Tea commands
│   │   └── update.go           # Update logic
│   ├── components/
│   │   ├── chat.go             # Chat area component
│   │   ├── input.go            # Input component
│   │   ├── modules.go          # Module selector
│   │   ├── header.go           # Header bar
│   │   ├── footer.go           # Footer/status
│   │   ├── thinking.go         # Thinking indicator
│   │   └── messages.go         # Message rendering
│   ├── bridge/
│   │   ├── danni.go            # Bridge to Rust CLI
│   │   ├── protocol.go         # JSON protocol definitions
│   │   └── events.go           # Event streaming
│   ├── styles/
│   │   ├── theme.go            # Danni color palette
│   │   ├── components.go       # Component styles
│   │   └── layouts.go          # Layout definitions
│   └── config/
│       ├── config.go           # App configuration
│       └── modules.go          # Module definitions
├── pkg/
│   └── glamour-danni/
│       └── theme.json          # Custom Glamour theme
├── go.mod
├── go.sum
└── README.md
```

### 3.2 Main App Structure (Bubble Tea Model)

```go
package app

import (
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/bubbles/viewport"
    "github.com/charmbracelet/bubbles/textarea"
    "github.com/charmbracelet/bubbles/spinner"
)

type Model struct {
    // Core state
    width  int
    height int
    ready  bool

    // Components
    chat         viewport.Model
    input        textarea.Model
    modules      ModuleSelector
    thinking     ThinkingIndicator

    // App state
    messages     []Message
    activeModule string
    sessionID    string

    // Bridge to Rust
    danni        *DanniBridge

    // UI state
    showHelp     bool
    showSettings bool
    err          error
}

// Bubble Tea interface
func (m Model) Init() tea.Cmd {
    return tea.Batch(
        textarea.Blink,
        m.danni.Connect(),
        m.thinking.spinner.Tick,
    )
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmds []tea.Cmd

    switch msg := msg.(type) {
    case tea.WindowSizeMsg:
        m.width = msg.Width
        m.height = msg.Height
        m.updateLayout()

    case tea.KeyMsg:
        return m.handleKeyPress(msg)

    case DanniEventMsg:
        return m.handleDanniEvent(msg)

    case spinner.TickMsg:
        m.thinking.spinner, cmd = m.thinking.spinner.Update(msg)
        cmds = append(cmds, cmd)
    }

    return m, tea.Batch(cmds...)
}

func (m Model) View() string {
    if !m.ready {
        return m.renderLoading()
    }

    if m.showHelp {
        return m.renderHelp()
    }

    if m.showSettings {
        return m.renderSettings()
    }

    return m.renderMain()
}
```

### 3.3 Rust CLI Bridge

```go
package bridge

import (
    "bufio"
    "encoding/json"
    "os/exec"
    tea "github.com/charmbracelet/bubbletea"
)

type DanniBridge struct {
    cmd    *exec.Cmd
    stdin  io.WriteCloser
    stdout io.ReadCloser
    stderr io.ReadCloser
}

func NewDanniBridge(sessionID string) *DanniBridge {
    cmd := exec.Command("danni",
        "--output-format", "json",
        "--tui-mode",
        "session", sessionID)

    stdin, _ := cmd.StdinPipe()
    stdout, _ := cmd.StdoutPipe()
    stderr, _ := cmd.StderrPipe()

    return &DanniBridge{
        cmd: cmd,
        stdin: stdin,
        stdout: stdout,
        stderr: stderr,
    }
}

func (d *DanniBridge) Connect() tea.Cmd {
    return func() tea.Msg {
        if err := d.cmd.Start(); err != nil {
            return ErrMsg{err}
        }
        return ConnectedMsg{}
    }
}

func (d *DanniBridge) SendMessage(content string) tea.Cmd {
    return func() tea.Msg {
        msg := UserMessage{
            Type: "message",
            Content: content,
        }

        if err := json.NewEncoder(d.stdin).Encode(msg); err != nil {
            return ErrMsg{err}
        }

        return MessageSentMsg{}
    }
}

func (d *DanniBridge) StreamEvents() tea.Cmd {
    return func() tea.Msg {
        scanner := bufio.NewScanner(d.stdout)
        for scanner.Scan() {
            var event AgentEvent
            if err := json.Unmarshal(scanner.Bytes(), &event); err != nil {
                continue
            }
            return DanniEventMsg{event}
        }
        return DisconnectedMsg{}
    }
}

// Event types
type AgentEvent struct {
    Type    string          `json:"type"`
    Data    json.RawMessage `json:"data"`
}

type UserMessage struct {
    Type    string `json:"type"`
    Content string `json:"content"`
}

// Bubble Tea messages
type (
    ConnectedMsg    struct{}
    DisconnectedMsg struct{}
    MessageSentMsg  struct{}
    DanniEventMsg   struct{ Event AgentEvent }
    ErrMsg          struct{ Err error }
)
```

### 3.4 Custom Glamour Theme

```json
{
  "document": {
    "block_prefix": "\n",
    "block_suffix": "\n",
    "color": "#F7F3E9",
    "background_color": "#1E3A5F"
  },
  "heading": {
    "block_suffix": "\n",
    "color": "#9F7AEA",
    "bold": true
  },
  "h1": {
    "prefix": " ✨ ",
    "suffix": " ",
    "color": "#9F7AEA",
    "bold": true
  },
  "h2": {
    "prefix": " 🎯 ",
    "color": "#B76E79"
  },
  "paragraph": {
    "color": "#F7F3E9"
  },
  "code_block": {
    "color": "#3B82F6",
    "background_color": "#2D2D2D",
    "margin": 2,
    "chroma": {
      "text": {
        "color": "#F7F3E9"
      },
      "keyword": {
        "color": "#9F7AEA",
        "bold": true
      },
      "string": {
        "color": "#10B981"
      }
    }
  },
  "list": {
    "color": "#F7F3E9"
  },
  "link": {
    "color": "#3B82F6",
    "underline": true
  },
  "link_text": {
    "color": "#3B82F6"
  },
  "emph": {
    "italic": true,
    "color": "#B76E79"
  },
  "strong": {
    "bold": true,
    "color": "#9F7AEA"
  }
}
```

---

## PHASE 4: IMPLEMENTATION ROADMAP

### 4.1 Week 1-2: Foundation
**Goal:** Basic TUI skeleton with Rust bridge

**Tasks:**
- [ ] Set up Go project structure
- [ ] Create basic Bubble Tea app shell
- [ ] Implement Rust CLI bridge (subprocess communication)
- [ ] Test JSON event streaming
- [ ] Create Danni color theme in Lip Gloss
- [ ] Build basic layout (header, chat, input, footer)

**Deliverable:** TUI that can send/receive messages to/from Rust CLI

### 4.2 Week 3-4: Core Components
**Goal:** Main chat interface with markdown rendering

**Tasks:**
- [ ] Implement chat viewport with scrolling
- [ ] Integrate Glamour with custom Danni theme
- [ ] Build message rendering component
- [ ] Create thinking indicator with spinner
- [ ] Add input component with multi-line support
- [ ] Implement basic keyboard navigation

**Deliverable:** Functional chat interface with beautiful message rendering

### 4.3 Week 5-6: Module System & Polish
**Goal:** Module switching and advanced UI features

**Tasks:**
- [ ] Build module selector component
- [ ] Implement module switching logic
- [ ] Add context/token usage display
- [ ] Create extension manager UI
- [ ] Build settings panel
- [ ] Add help overlay

**Deliverable:** Full-featured TUI with all major components

### 4.4 Week 7-8: Animation & Refinement
**Goal:** Smooth animations and UX polish

**Tasks:**
- [ ] Add smooth transitions between states
- [ ] Implement progressive reveal animations
- [ ] Optimize rendering performance
- [ ] Add mouse support
- [ ] Create loading states and skeleton screens
- [ ] Polish color gradients and effects
- [ ] Write comprehensive documentation

**Deliverable:** Production-ready TUI with polished UX

---

## PHASE 5: RISK ASSESSMENT & MITIGATION

### 5.1 Technical Risks

| Risk | Probability | Impact | Mitigation |
|------|-------------|--------|------------|
| Rust-Go communication instability | Low | High | Robust error handling, fallback to direct Rust CLI |
| Performance issues with large chat histories | Medium | Medium | Implement virtual scrolling, message pagination |
| Terminal compatibility problems | Medium | Medium | Test on multiple terminals, graceful degradation |
| Glamour rendering issues | Low | Low | Fallback to plain text rendering |
| Learning curve for Bubble Tea | Medium | Low | Follow official tutorials, reference examples |

### 5.2 UX Risks

| Risk | Probability | Impact | Mitigation |
|------|-------------|--------|------------|
| Too complex interface | Low | Medium | Progressive disclosure, hide advanced features initially |
| Color scheme not universally appealing | Medium | Low | Make theme configurable, provide alternatives |
| Keyboard shortcuts conflict | Low | Low | Follow terminal conventions, make customizable |
| Module switching confusion | Medium | Medium | Clear visual indicators, contextual help |

### 5.3 Timeline Risks

| Risk | Probability | Impact | Mitigation |
|------|-------------|--------|------------|
| Underestimating animation polish time | High | Low | Start simple, iterate with animations |
| Rust CLI modifications needed | Medium | Medium | Design protocol to minimize CLI changes |
| Go learning curve | Low | Low | Team has Go experience, well-documented libraries |

---

## PHASE 6: ALTERNATIVE APPROACHES (CONSIDERED & REJECTED)

### 6.1 Pure Rust TUI (ratatui/cursive)

**Why Considered:**
- Same language as existing codebase
- No process boundary

**Why Rejected:**
- Rust TUI libraries less mature than Charmbracelet
- Harder to achieve sophisticated styling
- Less ecosystem support
- Team preference for Go TUI development

### 6.2 Web-based Terminal UI (xterm.js)

**Why Considered:**
- Rich features, extensive styling
- Already have desktop Electron app

**Why Rejected:**
- Not a true TUI experience
- Requires web server
- More resource-intensive
- Goal is native terminal feel

### 6.3 Text-based UI in Rust CLI itself

**Why Considered:**
- Simplest approach
- No new languages

**Why Rejected:**
- Limited styling capabilities
- Hard to achieve Danni's sophisticated aesthetic
- Rustyline is very basic
- Difficult to build complex layouts

---

## PHASE 7: SUCCESS METRICS

### 7.1 Technical Metrics
- [ ] TUI starts in < 500ms
- [ ] Message rendering latency < 16ms (60fps)
- [ ] Handles 1000+ message history without lag
- [ ] Works on all major terminals (iTerm2, Alacritty, Windows Terminal, etc.)
- [ ] Memory usage < 50MB baseline

### 7.2 UX Metrics
- [ ] Users can send first message within 5 seconds of launch
- [ ] Module switching is intuitive (no documentation needed)
- [ ] Keyboard navigation feels natural
- [ ] Colors are legible in both dark and light terminals
- [ ] Help is discoverable and useful

### 7.3 Aesthetic Metrics
- [ ] Color palette feels cohesive and sophisticated
- [ ] Typography hierarchy is clear
- [ ] Animations feel smooth and purposeful
- [ ] Interface matches Danni's personality (warm, intelligent, elegant)
- [ ] Visual polish comparable to modern CLI tools (gh, glow, lazygit)

---

## PHASE 8: FUTURE ENHANCEMENTS (POST-MVP)

### 8.1 Advanced Features
- **Session Management UI:** Browse, switch, archive sessions visually
- **Split Panes:** Side-by-side comparison, documentation viewer
- **Inline File Editing:** Edit files directly in TUI
- **Graph Visualizations:** Token usage over time, module distribution
- **Export Formatting:** Beautiful PDF/HTML exports with branding
- **Themes:** User-created color schemes
- **Plugin System:** Extend TUI with custom components

### 8.2 Integration Enhancements
- **Git Integration:** Show git status, diffs in TUI
- **File Browser:** Navigate project files with tree view
- **Search:** Full-text search across all sessions
- **Notifications:** Desktop notifications for long-running tasks
- **Remote Sessions:** Connect to Danni running on remote machine

### 8.3 AI-Powered Features
- **Smart Autocomplete:** Context-aware suggestions in input
- **Session Summarization:** Visual summaries of long conversations
- **Insight Highlighting:** Auto-highlight key insights
- **Personalization:** UI adapts to user preferences over time

---

## CONCLUSION

### Why This Approach Will Succeed

1. **Proven Technology Stack:**
   - Bubble Tea: Production-ready, used by major projects
   - Lip Gloss: Sophisticated styling capabilities
   - Glamour: Beautiful markdown rendering
   - Bubbles: Pre-built components save time

2. **Clean Architecture:**
   - Subprocess bridge is simple and robust
   - Clear separation of concerns
   - Each component can be developed/tested independently
   - Minimal changes to existing Rust codebase

3. **Aligned with Danni's Identity:**
   - Color palette directly from brand guidelines
   - Sophisticated yet warm aesthetic
   - Progressive revelation of features
   - Smooth, magical interactions

4. **Manageable Risk:**
   - Low technical risk (well-documented libraries)
   - Fallback to Rust CLI always available
   - Incremental development possible
   - No breaking changes to existing functionality

5. **Long-term Vision:**
   - Foundation for advanced features
   - Extensible architecture
   - Community can contribute themes/components
   - Sets Danni apart in the AI agent space

### Next Steps

1. **Decision Point:** Get approval on overall approach
2. **Design Refinement:** Create ASCII mockups or terminal screenshots
3. **Prototype:** Build 2-week proof-of-concept
4. **Iterate:** Gather feedback, refine design
5. **Build:** Full 6-8 week implementation
6. **Polish:** 2 weeks of refinement and testing
7. **Launch:** Ship beautiful TUI with Danni v0.1.0

---

**This TUI will make Danni feel as sophisticated and warm as her personality—a terminal interface that sparks joy and inspires creativity.**

*Welcome to the future of developer-AI collaboration.*
