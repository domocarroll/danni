# DANNI TUI QUICK START GUIDE
**Get the Beautiful TUI Running in 2 Weeks**

---

## OVERVIEW

This is your practical, step-by-step guide to building the Danni TUI proof-of-concept. We'll create a minimal but beautiful version that demonstrates the core concepts, then iterate from there.

**Goal:** Working TUI with basic chat interface in 2 weeks
**Focus:** Core functionality + Danni aesthetic
**Tech:** Go + Charmbracelet + Rust CLI bridge

---

## WEEK 1: FOUNDATION

### Day 1-2: Project Setup & Hello World

#### 1. Create Go Project Structure
```bash
cd /home/dom/danni-goose-fork
mkdir -p tui/{cmd/danni-tui,internal/{app,bridge,styles,components}}
cd tui

# Initialize Go module
go mod init github.com/subfracture/danni/tui

# Add Charmbracelet dependencies
go get github.com/charmbracelet/bubbletea
go get github.com/charmbracelet/lipgloss
go get github.com/charmbracelet/glamour
go get github.com/charmbracelet/bubbles/viewport
go get github.com/charmbracelet/bubbles/textarea
go get github.com/charmbracelet/bubbles/spinner
```

#### 2. Create Hello World TUI

**File:** `cmd/danni-tui/main.go`
```go
package main

import (
    "fmt"
    "os"

    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
)

// Danni purple color palette
var (
    primaryPurple = lipgloss.Color("#9F7AEA")
    midnightBlue  = lipgloss.Color("#1E3A5F")
    softCream     = lipgloss.Color("#F7F3E9")
)

type model struct {
    width  int
    height int
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        if msg.String() == "q" || msg.String() == "ctrl+c" {
            return m, tea.Quit
        }
    case tea.WindowSizeMsg:
        m.width = msg.Width
        m.height = msg.Height
    }
    return m, nil
}

func (m model) View() string {
    style := lipgloss.NewStyle().
        Foreground(primaryPurple).
        Background(midnightBlue).
        Padding(2).
        Width(m.width).
        Height(m.height).
        Align(lipgloss.Center, lipgloss.Center)

    return style.Render("✨ Hello, DANNI! ✨\n\nPress 'q' to quit")
}

func main() {
    p := tea.NewProgram(model{}, tea.WithAltScreen())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Error: %v\n", err)
        os.Exit(1)
    }
}
```

#### 3. Test It
```bash
go run cmd/danni-tui/main.go
```

**Expected:** Full-screen purple window with centered "Hello, DANNI!" text

---

### Day 3-4: Layout Structure

#### 1. Create Danni Theme

**File:** `internal/styles/theme.go`
```go
package styles

import "github.com/charmbracelet/lipgloss"

// Danni Color Palette
var (
    PrimaryPurple = lipgloss.Color("#9F7AEA")
    DeepPurple    = lipgloss.Color("#6B46C1")
    MidnightBlue  = lipgloss.Color("#1E3A5F")
    RoseGold      = lipgloss.Color("#B76E79")
    SoftCream     = lipgloss.Color("#F7F3E9")
    DeepCharcoal  = lipgloss.Color("#2D2D2D")
    SubtleGray    = lipgloss.Color("#6B7280")
)

// Component Styles
var (
    HeaderStyle = lipgloss.NewStyle().
        Foreground(SoftCream).
        Background(DeepPurple).
        Bold(true).
        Padding(0, 2)

    ChatStyle = lipgloss.NewStyle().
        Foreground(SoftCream).
        Background(MidnightBlue).
        Padding(1, 2)

    InputStyle = lipgloss.NewStyle().
        Foreground(SoftCream).
        Background(DeepCharcoal).
        Padding(0, 1)

    FooterStyle = lipgloss.NewStyle().
        Foreground(SubtleGray).
        Background(DeepCharcoal).
        Padding(0, 2)

    UserMessageStyle = lipgloss.NewStyle().
        Foreground(SoftCream).
        Bold(true)

    AssistantMessageStyle = lipgloss.NewStyle().
        Foreground(PrimaryPurple).
        Bold(true)
)
```

#### 2. Create Header Component

**File:** `internal/components/header.go`
```go
package components

import (
    "fmt"
    "github.com/charmbracelet/lipgloss"
    "github.com/subfracture/danni/tui/internal/styles"
)

type Header struct {
    Width     int
    Module    string
    SessionID string
}

func (h Header) View() string {
    left := lipgloss.NewStyle().
        Foreground(styles.SoftCream).
        Bold(true).
        Render("✨ DANNI v0.1.0")

    center := lipgloss.NewStyle().
        Foreground(styles.PrimaryPurple).
        Bold(true).
        Render(h.Module)

    right := lipgloss.NewStyle().
        Foreground(styles.SubtleGray).
        Render(fmt.Sprintf("● %s", h.SessionID))

    // Calculate spacing
    spacing := h.Width - lipgloss.Width(left) - lipgloss.Width(center) - lipgloss.Width(right) - 4

    content := lipgloss.JoinHorizontal(
        lipgloss.Top,
        left,
        strings.Repeat(" ", spacing/2),
        center,
        strings.Repeat(" ", spacing/2),
        right,
    )

    return styles.HeaderStyle.Width(h.Width).Render(content)
}
```

#### 3. Create Footer Component

**File:** `internal/components/footer.go`
```go
package components

import (
    "fmt"
    "github.com/charmbracelet/lipgloss"
    "github.com/subfracture/danni/tui/internal/styles"
)

type Footer struct {
    Width      int
    Tokens     int
    MaxTokens  int
    Cost       float64
}

func (f Footer) View() string {
    percentage := float64(f.Tokens) / float64(f.MaxTokens) * 100

    tokensText := fmt.Sprintf("Tokens: %d/%d (%.1f%%)",
        f.Tokens, f.MaxTokens, percentage)

    costText := fmt.Sprintf("Cost: $%.2f", f.Cost)

    helpText := "Ctrl+J: newline • /help • Ctrl+D: exit"

    content := lipgloss.JoinHorizontal(
        lipgloss.Top,
        tokensText,
        "  •  ",
        costText,
        "  •  ",
        helpText,
    )

    return styles.FooterStyle.Width(f.Width).Render(content)
}
```

#### 4. Update Main App with Layout

**File:** `internal/app/app.go`
```go
package app

import (
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
    "github.com/subfracture/danni/tui/internal/components"
    "github.com/subfracture/danni/tui/internal/styles"
)

type Model struct {
    width     int
    height    int
    header    components.Header
    footer    components.Footer
    ready     bool
}

func New() Model {
    return Model{
        header: components.Header{
            Module:    "/strategy",
            SessionID: "danni-42",
        },
        footer: components.Footer{
            Tokens:    1234,
            MaxTokens: 200000,
            Cost:      0.03,
        },
    }
}

func (m Model) Init() tea.Cmd {
    return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "q", "ctrl+c", "ctrl+d":
            return m, tea.Quit
        }
    case tea.WindowSizeMsg:
        m.width = msg.Width
        m.height = msg.Height
        m.header.Width = m.width
        m.footer.Width = m.width
        m.ready = true
    }
    return m, nil
}

func (m Model) View() string {
    if !m.ready {
        return "Initializing..."
    }

    header := m.header.View()
    footer := m.footer.View()

    // Calculate available space for content
    contentHeight := m.height - lipgloss.Height(header) - lipgloss.Height(footer)

    // Temporary content placeholder
    content := styles.ChatStyle.
        Width(m.width).
        Height(contentHeight).
        Render("Chat content will go here...")

    return lipgloss.JoinVertical(
        lipgloss.Left,
        header,
        content,
        footer,
    )
}
```

**Update:** `cmd/danni-tui/main.go`
```go
package main

import (
    "fmt"
    "os"

    tea "github.com/charmbracelet/bubbletea"
    "github.com/subfracture/danni/tui/internal/app"
)

func main() {
    m := app.New()
    p := tea.NewProgram(m, tea.WithAltScreen())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Error: %v\n", err)
        os.Exit(1)
    }
}
```

#### 5. Test Layout
```bash
go run cmd/danni-tui/main.go
```

**Expected:** Three-section layout with purple header, blue content area, dark footer

---

### Day 5-7: Chat Interface & Input

#### 1. Create Message Type

**File:** `internal/app/message.go`
```go
package app

import "time"

type Message struct {
    Role      string    // "user" or "assistant"
    Content   string
    Timestamp time.Time
    Module    string
}
```

#### 2. Add Chat Viewport

**Update:** `internal/app/app.go`
```go
package app

import (
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/bubbles/viewport"
    "github.com/charmbracelet/bubbles/textarea"
    "github.com/charmbracelet/lipgloss"
    "time"
    // ... other imports
)

type Model struct {
    width     int
    height    int
    ready     bool

    // Components
    header    components.Header
    footer    components.Footer
    chat      viewport.Model
    input     textarea.Model

    // State
    messages  []Message
}

func New() Model {
    ta := textarea.New()
    ta.Placeholder = "Ask Danni anything..."
    ta.Focus()
    ta.Prompt = "> "
    ta.CharLimit = 4000
    ta.SetWidth(80)
    ta.SetHeight(3)
    ta.ShowLineNumbers = false

    return Model{
        header: components.Header{
            Module:    "/strategy",
            SessionID: "danni-42",
        },
        footer: components.Footer{
            Tokens:    1234,
            MaxTokens: 200000,
            Cost:      0.03,
        },
        input:    ta,
        messages: []Message{
            {
                Role:      "assistant",
                Content:   "I've noticed something fascinating... you're exploring Danni's new interface. How can I help you today?",
                Timestamp: time.Now(),
                Module:    "/strategy",
            },
        },
    }
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var (
        cmdChat  tea.Cmd
        cmdInput tea.Cmd
    )

    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c", "ctrl+d":
            return m, tea.Quit
        case "enter":
            // Send message
            content := m.input.Value()
            if content != "" {
                m.messages = append(m.messages, Message{
                    Role:      "user",
                    Content:   content,
                    Timestamp: time.Now(),
                })
                m.input.Reset()
                // TODO: Send to Rust CLI
            }
        }
    case tea.WindowSizeMsg:
        m.width = msg.Width
        m.height = msg.Height
        m.header.Width = m.width
        m.footer.Width = m.width

        if !m.ready {
            m.chat = viewport.New(m.width, m.height-10)
            m.ready = true
        } else {
            m.chat.Width = m.width
            m.chat.Height = m.height - 10
        }
    }

    m.chat, cmdChat = m.chat.Update(msg)
    m.input, cmdInput = m.input.Update(msg)

    return m, tea.Batch(cmdChat, cmdInput)
}

func (m Model) View() string {
    if !m.ready {
        return "Initializing..."
    }

    // Render messages
    var chatContent string
    for _, msg := range m.messages {
        chatContent += m.renderMessage(msg) + "\n\n"
    }
    m.chat.SetContent(chatContent)

    header := m.header.View()
    chat := styles.ChatStyle.Width(m.width).Render(m.chat.View())
    input := styles.InputStyle.Width(m.width).Render(m.input.View())
    footer := m.footer.View()

    return lipgloss.JoinVertical(
        lipgloss.Left,
        header,
        chat,
        input,
        footer,
    )
}

func (m Model) renderMessage(msg Message) string {
    var roleStyle lipgloss.Style
    var roleLabel string

    if msg.Role == "user" {
        roleStyle = styles.UserMessageStyle
        roleLabel = "[You]"
    } else {
        roleStyle = styles.AssistantMessageStyle
        roleLabel = "[Danni]"
    }

    timestamp := lipgloss.NewStyle().
        Foreground(styles.SubtleGray).
        Render(msg.Timestamp.Format("3:04 PM"))

    header := lipgloss.JoinHorizontal(
        lipgloss.Left,
        roleStyle.Render(roleLabel),
        "  ",
        timestamp,
    )

    content := lipgloss.NewStyle().
        Foreground(styles.SoftCream).
        Render(msg.Content)

    return lipgloss.JoinVertical(lipgloss.Left, header, "", content)
}
```

#### 3. Test Chat Interface
```bash
go run cmd/danni-tui/main.go
```

**Expected:**
- Can type in input area
- Press Enter to send message
- Messages appear in chat area
- Can scroll with arrow keys

---

## WEEK 2: RUST BRIDGE & POLISH

### Day 8-10: Rust CLI Bridge

#### 1. Create Bridge Package

**File:** `internal/bridge/danni.go`
```go
package bridge

import (
    "bufio"
    "encoding/json"
    "io"
    "os/exec"

    tea "github.com/charmbracelet/bubbletea"
)

type Bridge struct {
    cmd    *exec.Cmd
    stdin  io.WriteCloser
    stdout io.ReadCloser
}

type UserMessage struct {
    Type    string `json:"type"`
    Content string `json:"content"`
}

type AgentEvent struct {
    Type    string          `json:"type"`
    Role    string          `json:"role,omitempty"`
    Content string          `json:"content,omitempty"`
    Data    json.RawMessage `json:"data,omitempty"`
}

// Bubble Tea messages
type (
    ConnectedMsg    struct{}
    DisconnectedMsg struct{}
    AgentEventMsg   struct{ Event AgentEvent }
    ErrorMsg        struct{ Err error }
)

func New(sessionID string) *Bridge {
    // For now, we'll use a mock. Later, replace with actual danni CLI
    cmd := exec.Command("danni",
        "--output-format", "json",
        "session", sessionID)

    stdin, _ := cmd.StdinPipe()
    stdout, _ := cmd.StdoutPipe()

    return &Bridge{
        cmd:    cmd,
        stdin:  stdin,
        stdout: stdout,
    }
}

func (b *Bridge) Connect() tea.Cmd {
    return func() tea.Msg {
        if err := b.cmd.Start(); err != nil {
            return ErrorMsg{err}
        }
        return ConnectedMsg{}
    }
}

func (b *Bridge) SendMessage(content string) tea.Cmd {
    return func() tea.Msg {
        msg := UserMessage{
            Type:    "message",
            Content: content,
        }

        if err := json.NewEncoder(b.stdin).Encode(msg); err != nil {
            return ErrorMsg{err}
        }

        return nil
    }
}

func (b *Bridge) StreamEvents() tea.Cmd {
    return func() tea.Msg {
        scanner := bufio.NewScanner(b.stdout)
        for scanner.Scan() {
            var event AgentEvent
            if err := json.Unmarshal(scanner.Bytes(), &event); err != nil {
                continue
            }
            return AgentEventMsg{Event: event}
        }
        return DisconnectedMsg{}
    }
}
```

#### 2. Integrate Bridge into App

**Update:** `internal/app/app.go`
```go
type Model struct {
    // ... existing fields
    bridge    *bridge.Bridge
}

func New() Model {
    // ... existing setup

    m := Model{
        // ... existing initialization
        bridge: bridge.New("danni-42"),
    }

    return m
}

func (m Model) Init() tea.Cmd {
    return tea.Batch(
        textarea.Blink,
        m.bridge.Connect(),
        m.bridge.StreamEvents(),
    )
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmds []tea.Cmd

    switch msg := msg.(type) {
    case bridge.ConnectedMsg:
        // Connected to Danni
    case bridge.DisconnectedMsg:
        // Handle disconnect
        return m, tea.Quit
    case bridge.AgentEventMsg:
        // Handle agent response
        if msg.Event.Type == "message" && msg.Event.Role == "assistant" {
            m.messages = append(m.messages, Message{
                Role:      "assistant",
                Content:   msg.Event.Content,
                Timestamp: time.Now(),
                Module:    m.header.Module,
            })
            cmds = append(cmds, m.bridge.StreamEvents())
        }
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c", "ctrl+d":
            return m, tea.Quit
        case "enter":
            content := m.input.Value()
            if content != "" {
                m.messages = append(m.messages, Message{
                    Role:      "user",
                    Content:   content,
                    Timestamp: time.Now(),
                })
                m.input.Reset()
                cmds = append(cmds, m.bridge.SendMessage(content))
            }
        }
    // ... rest of update logic
    }

    // ... component updates

    return m, tea.Batch(cmds...)
}
```

#### 3. Create Mock Danni CLI (for testing)

**File:** `scripts/mock-danni.sh`
```bash
#!/bin/bash
# Mock Danni CLI for testing TUI
# Replace with actual danni once bridge protocol is finalized

while read -r line; do
    # Parse input
    type=$(echo "$line" | jq -r '.type')
    content=$(echo "$line" | jq -r '.content')

    # Simulate thinking
    sleep 0.5

    # Send back response
    cat <<EOF
{"type":"message","role":"assistant","content":"I've received your message: '$content'. This is a mock response from the Rust CLI bridge. The actual Danni will provide sophisticated strategic insights here."}
EOF
done
```

```bash
chmod +x scripts/mock-danni.sh
```

---

### Day 11-12: Glamour Integration

#### 1. Add Glamour Markdown Rendering

**Update:** `internal/app/app.go`
```go
import (
    "github.com/charmbracelet/glamour"
    // ... other imports
)

func (m Model) renderMessage(msg Message) string {
    var roleStyle lipgloss.Style
    var roleLabel string

    if msg.Role == "user" {
        roleStyle = styles.UserMessageStyle
        roleLabel = "[You]"
    } else {
        roleStyle = styles.AssistantMessageStyle
        roleLabel = "[Danni]"
    }

    timestamp := lipgloss.NewStyle().
        Foreground(styles.SubtleGray).
        Render(msg.Timestamp.Format("3:04 PM"))

    header := lipgloss.JoinHorizontal(
        lipgloss.Left,
        roleStyle.Render(roleLabel),
        "  ",
        timestamp,
    )

    // Render markdown with Glamour
    renderer, _ := glamour.NewTermRenderer(
        glamour.WithAutoStyle(),
        glamour.WithWordWrap(m.width-4),
    )

    content, err := renderer.Render(msg.Content)
    if err != nil {
        content = msg.Content // Fallback to plain text
    }

    return lipgloss.JoinVertical(lipgloss.Left, header, "", content)
}
```

---

### Day 13-14: Polish & Testing

#### 1. Add Spinner for Thinking State

**Update:** `internal/app/app.go`
```go
import (
    "github.com/charmbracelet/bubbles/spinner"
)

type Model struct {
    // ... existing fields
    thinking bool
    spinner  spinner.Model
}

func New() Model {
    s := spinner.New()
    s.Spinner = spinner.Dot
    s.Style = lipgloss.NewStyle().Foreground(styles.PrimaryPurple)

    // ... existing setup with spinner added
}

func (m Model) Init() tea.Cmd {
    return tea.Batch(
        textarea.Blink,
        m.spinner.Tick,
        m.bridge.Connect(),
        m.bridge.StreamEvents(),
    )
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case spinner.TickMsg:
        var cmd tea.Cmd
        m.spinner, cmd = m.spinner.Update(msg)
        return m, cmd
    // ... handle other messages
    }
}

func (m Model) View() string {
    // ... existing view code

    var thinking string
    if m.thinking {
        thinking = lipgloss.NewStyle().
            Foreground(styles.PrimaryPurple).
            Background(styles.MidnightBlue).
            Padding(1, 2).
            Render(fmt.Sprintf("💭 %s Thinking...", m.spinner.View()))
    }

    return lipgloss.JoinVertical(
        lipgloss.Left,
        header,
        thinking,
        chat,
        input,
        footer,
    )
}
```

#### 2. Add Module Selector (Simple Version)

**File:** `internal/components/modules.go`
```go
package components

import (
    "github.com/charmbracelet/lipgloss"
    "github.com/subfracture/danni/tui/internal/styles"
)

type ModuleSelector struct {
    Width  int
    Active int
}

var modules = []string{
    "🎯 /strategy",
    "🎨 /creative",
    "✨ /design",
    "⚡ /technology",
    "🌌 /gravity",
}

func (m ModuleSelector) View() string {
    var tabs []string

    for i, mod := range modules {
        var style lipgloss.Style
        if i == m.Active {
            style = lipgloss.NewStyle().
                Foreground(styles.PrimaryPurple).
                Bold(true).
                Underline(true)
        } else {
            style = lipgloss.NewStyle().
                Foreground(styles.SubtleGray)
        }
        tabs = append(tabs, style.Render(mod))
    }

    content := lipgloss.JoinHorizontal(lipgloss.Top, tabs...)

    return lipgloss.NewStyle().
        Width(m.Width).
        Padding(0, 2).
        Background(styles.DeepCharcoal).
        Render(content)
}
```

#### 3. Final Integration & Testing

**Create test script:** `scripts/test-tui.sh`
```bash
#!/bin/bash
echo "Testing Danni TUI..."
cd tui
go build -o ../bin/danni-tui cmd/danni-tui/main.go
../bin/danni-tui
```

---

## FINAL CHECKLIST

### ✅ Core Features
- [ ] TUI starts and displays properly
- [ ] Header shows DANNI branding
- [ ] Chat area scrolls correctly
- [ ] Input accepts multi-line text
- [ ] Messages display with correct styling
- [ ] Markdown renders beautifully (via Glamour)
- [ ] Footer shows token/cost info
- [ ] Module selector displays
- [ ] Thinking spinner appears during processing
- [ ] Can quit gracefully (Ctrl+D)

### ✅ Danni Aesthetic
- [ ] Purple gradient in header
- [ ] Midnight blue background
- [ ] Rose gold accents
- [ ] Cream text, easy to read
- [ ] Proper spacing and padding
- [ ] Borders look polished
- [ ] Colors match brand guidelines

### ✅ Technical
- [ ] No crashes on window resize
- [ ] Handles terminal of various sizes
- [ ] Bridge communicates with Rust CLI
- [ ] Events stream correctly
- [ ] Error handling works
- [ ] Performance is smooth (60fps)

---

## NEXT STEPS AFTER POC

Once you have the basic TUI working:

1. **Enhance Bridge Protocol:**
   - Add support for tool execution events
   - Handle streaming responses properly
   - Implement error recovery

2. **Add Advanced Features:**
   - Extension manager UI
   - Settings panel
   - Help overlay
   - Session management

3. **Polish Animations:**
   - Smooth module switching
   - Fade in/out effects
   - Progress bars
   - Loading states

4. **Testing:**
   - Test on multiple terminals
   - Test on different OS
   - Performance profiling
   - User testing

5. **Documentation:**
   - User guide
   - Developer docs
   - Video demos

---

## TROUBLESHOOTING

### TUI doesn't render colors
- Check terminal supports 256 colors or True Color
- Try: `echo $TERM` (should be `xterm-256color` or similar)
- Set: `export TERM=xterm-256color`

### Text wrapping issues
- Ensure viewport width matches terminal width
- Use `lipgloss.Width()` to measure rendered text
- Account for padding/margins in calculations

### Keyboard shortcuts don't work
- Check terminal captures key events
- Some terminals intercept certain key combos
- Provide alternative shortcuts

### Bridge doesn't connect
- Verify Rust CLI is in PATH
- Check CLI supports `--output-format json`
- Test CLI manually first
- Check stderr for error messages

---

## RESOURCES

### Official Docs
- [Bubble Tea Tutorial](https://github.com/charmbracelet/bubbletea/tree/master/tutorials)
- [Lip Gloss Examples](https://github.com/charmbracelet/lipgloss/tree/master/examples)
- [Glamour Styles](https://github.com/charmbracelet/glamour/tree/master/styles)
- [Bubbles Components](https://github.com/charmbracelet/bubbles)

### Example Projects
- [Glow](https://github.com/charmbracelet/glow) - Markdown reader
- [Soft Serve](https://github.com/charmbracelet/soft-serve) - Git server
- [VHS](https://github.com/charmbracelet/vhs) - Terminal recorder

### Community
- [Charm Discord](https://charm.sh/chat)
- [r/golang](https://reddit.com/r/golang)
- [GitHub Discussions](https://github.com/charmbracelet/bubbletea/discussions)

---

**You're ready to build! Start with Day 1 and work through systematically. The beautiful Danni TUI awaits! ✨**
