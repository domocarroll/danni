# DANNI TUI ARCHITECTURE DIAGRAMS
**Visual Reference Guide**

---

## SYSTEM ARCHITECTURE

### High-Level Overview
```
┌──────────────────────────────────────────────────────────────┐
│                         USER                                 │
│                           │                                  │
│                           ▼                                  │
│  ┌────────────────────────────────────────────────────┐     │
│  │         DANNI TUI (Go + Charmbracelet)             │     │
│  │                                                     │     │
│  │  ┌─────────────┐  ┌──────────────┐  ┌──────────┐ │     │
│  │  │  Bubble Tea │  │  Lip Gloss   │  │ Glamour  │ │     │
│  │  │  Framework  │  │   Styling    │  │ Markdown │ │     │
│  │  └─────────────┘  └──────────────┘  └──────────┘ │     │
│  │                                                     │     │
│  │  ┌──────────────────────────────────────────────┐ │     │
│  │  │         Bubbles Components                   │ │     │
│  │  │  Viewport • TextArea • Spinner • List        │ │     │
│  │  └──────────────────────────────────────────────┘ │     │
│  └────────────────────────────────────────────────────┘     │
│                           │                                  │
│                  JSON over stdin/stdout                      │
│                           │                                  │
│  ┌────────────────────────────────────────────────────┐     │
│  │       DANNI CLI (Rust) --tui-mode                  │     │
│  │                                                     │     │
│  │  ┌─────────────┐  ┌──────────────┐  ┌──────────┐ │     │
│  │  │   Session   │  │   Agent      │  │Extension │ │     │
│  │  │  Manager    │  │   System     │  │  System  │ │     │
│  │  └─────────────┘  └──────────────┘  └──────────┘ │     │
│  └────────────────────────────────────────────────────┘     │
│                           │                                  │
│                           ▼                                  │
│  ┌────────────────────────────────────────────────────┐     │
│  │        goose-server (Rust)                         │     │
│  │                                                     │     │
│  │  Agent • Providers • Tools • Extensions            │     │
│  └────────────────────────────────────────────────────┘     │
│                           │                                  │
│                           ▼                                  │
│  ┌────────────────────────────────────────────────────┐     │
│  │        LLM Provider (Anthropic, etc.)              │     │
│  └────────────────────────────────────────────────────┘     │
└──────────────────────────────────────────────────────────────┘
```

---

## COMMUNICATION FLOW

### User Message Flow
```
┌─────────────┐                                    ┌─────────────┐
│             │  1. User types message             │             │
│   User      │───────────────────────────────────>│  TUI Input  │
│             │                                    │  Component  │
└─────────────┘                                    └──────┬──────┘
                                                          │
                                                          │ 2. Capture
                                                          │    input
                                                          ▼
                                                   ┌──────────────┐
                                                   │ Bubble Tea   │
                                                   │ Update()     │
                                                   └──────┬───────┘
                                                          │
                                                          │ 3. Send JSON
                                                          │    message
                                                          ▼
┌─────────────────────────────────────────────────────────────────┐
│                        Bridge Layer                             │
│                                                                 │
│  stdin.Write({"type":"message","content":"help me"})           │
└────────────────────────────────┬────────────────────────────────┘
                                 │
                                 │ 4. JSON via stdin
                                 ▼
┌─────────────────────────────────────────────────────────────────┐
│                      Danni CLI (Rust)                           │
│                                                                 │
│  Parse JSON → Create Message → Agent.reply()                   │
└────────────────────────────────┬────────────────────────────────┘
                                 │
                                 │ 5. Stream events
                                 │    to stdout
                                 ▼
┌─────────────────────────────────────────────────────────────────┐
│                        Bridge Layer                             │
│                                                                 │
│  stdout.Read() → Parse JSON events → Send to Bubble Tea        │
└────────────────────────────────┬────────────────────────────────┘
                                 │
                                 │ 6. Update UI
                                 ▼
                          ┌──────────────┐
                          │ Bubble Tea   │
                          │ Update()     │
                          └──────┬───────┘
                                 │
                                 │ 7. Re-render
                                 ▼
                          ┌──────────────┐
                          │ Bubble Tea   │
                          │ View()       │
                          └──────┬───────┘
                                 │
                                 │ 8. Display
                                 ▼
                          ┌──────────────┐
                          │   Terminal   │
                          │   Screen     │
                          └──────────────┘
```

---

## EVENT STREAMING

### Danni Response Events
```
Time ──────────────────────────────────────────────────────────>

t=0s    {"type":"thinking","message":"Analyzing your request..."}
        │
        └──> TUI shows spinner: 💭 ⠋ Analyzing your request...

t=0.5s  {"type":"thinking","message":"Exploring patterns..."}
        │
        └──> TUI updates spinner: 💭 ⠙ Exploring patterns...

t=1s    {"type":"message","role":"assistant","content":"I've noticed..."}
        │
        └──> TUI hides spinner, renders message with Glamour

t=1.2s  {"type":"tool_request","name":"bash","args":{"cmd":"ls"}}
        │
        └──> TUI shows tool execution box

t=1.5s  {"type":"tool_response","result":"file1\nfile2\n..."}
        │
        └──> TUI updates tool box with result

t=2s    {"type":"message","role":"assistant","content":"Based on..."}
        │
        └──> TUI appends final message

t=2.1s  {"type":"session_update","tokens":1500}
        │
        └──> TUI updates token count in footer
```

---

## BUBBLE TEA MODEL-VIEW-UPDATE

### The Elm Architecture in Action
```
┌─────────────────────────────────────────────────────────────┐
│                      INIT                                   │
│                                                             │
│  • Create model with initial state                         │
│  • Set up components (viewport, textarea, spinner)         │
│  • Start bridge connection                                 │
│  • Return initial commands                                 │
└──────────────────┬──────────────────────────────────────────┘
                   │
                   ▼
┌─────────────────────────────────────────────────────────────┐
│                    UPDATE                                   │
│                                                             │
│  msg switch {                                               │
│    KeyMsg:        Handle keyboard input                    │
│    WindowSizeMsg: Adjust layout                            │
│    AgentEventMsg: Process Danni response                   │
│    SpinnerTickMsg: Animate spinner                         │
│  }                                                          │
│                                                             │
│  Return: (updated model, commands to run)                  │
└──────────────────┬──────────────────────────────────────────┘
                   │
                   ▼
┌─────────────────────────────────────────────────────────────┐
│                     VIEW                                    │
│                                                             │
│  • Render header                                            │
│  • Render thinking indicator (if active)                   │
│  • Render chat messages (with Glamour)                     │
│  • Render module selector                                  │
│  • Render input area                                       │
│  • Render footer                                           │
│                                                             │
│  Return: string (terminal output)                          │
└──────────────────┬──────────────────────────────────────────┘
                   │
                   ▼
              [Terminal Display]
                   │
                   │ User interaction
                   │ or system event
                   │
                   ▼
              [Back to UPDATE]
```

---

## COMPONENT HIERARCHY

### UI Component Tree
```
Model (Bubble Tea root)
├── Header
│   ├── Logo ("✨ DANNI")
│   ├── Module Badge (/strategy)
│   └── Session Info (● Session #42)
│
├── Thinking Indicator (conditional)
│   ├── Spinner (bubbles.spinner)
│   └── Message text
│
├── Chat Viewport (bubbles.viewport)
│   ├── Message 1
│   │   ├── Role label [You]/[Danni]
│   │   ├── Timestamp
│   │   └── Content (Glamour-rendered markdown)
│   │
│   ├── Message 2
│   │   └── ...
│   │
│   └── Tool Execution Box (conditional)
│       ├── Tool name & command
│       ├── Progress bar
│       └── Result
│
├── Module Selector
│   ├── Tab: /strategy ●
│   ├── Tab: /creative
│   ├── Tab: /design
│   └── ... (9 modules total)
│
├── Input Area (bubbles.textarea)
│   ├── Prompt (">")
│   ├── Text content
│   └── Cursor
│
└── Footer
    ├── Token usage (with progress bar)
    ├── Cost tracking
    └── Keyboard shortcuts
```

---

## DATA FLOW

### Message Lifecycle
```
1. USER INPUT
   ┌──────────────────────┐
   │ User types message   │
   │ "help me analyze X"  │
   └──────────┬───────────┘
              │
              ▼
2. TUI CAPTURE
   ┌──────────────────────┐
   │ TextArea.Value()     │
   │ → "help me..."       │
   └──────────┬───────────┘
              │
              ▼
3. MODEL UPDATE
   ┌─────────────────────────┐
   │ Append to messages[]    │
   │ {role: "user", ...}     │
   └──────────┬──────────────┘
              │
              ▼
4. BRIDGE SEND
   ┌────────────────────────────────┐
   │ JSON.Encode(UserMessage)       │
   │ → stdin.Write()                │
   └──────────┬─────────────────────┘
              │
              ▼
5. RUST CLI
   ┌────────────────────────────────┐
   │ Parse JSON                     │
   │ → Message::user()              │
   │ → agent.reply()                │
   └──────────┬─────────────────────┘
              │
              │ (processing...)
              │
              ▼
6. RUST EVENTS
   ┌────────────────────────────────┐
   │ Stream AgentEvent              │
   │ → stdout.Write(JSON)           │
   └──────────┬─────────────────────┘
              │
              ▼
7. BRIDGE RECEIVE
   ┌────────────────────────────────┐
   │ stdout.Read()                  │
   │ → JSON.Decode(AgentEvent)      │
   └──────────┬─────────────────────┘
              │
              ▼
8. MODEL UPDATE
   ┌─────────────────────────────────┐
   │ Process event type              │
   │ • thinking → show spinner       │
   │ • message → append to messages[]│
   │ • tool → update tool box        │
   └──────────┬──────────────────────┘
              │
              ▼
9. VIEW RENDER
   ┌─────────────────────────────────┐
   │ Glamour.Render(markdown)        │
   │ → LipGloss.Style()              │
   │ → Join components               │
   └──────────┬──────────────────────┘
              │
              ▼
10. DISPLAY
   ┌─────────────────────────────────┐
   │ Terminal updates with new UI    │
   └─────────────────────────────────┘
```

---

## COLOR SYSTEM

### Danni Palette Application
```
┌────────────────────────────────────────────────────────────────┐
│ ✨ DANNI v0.1.0        /strategy           ● Session #42      │ ← Header
│ ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓  │   Deep Purple (#6B46C1)
│ ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░  │   ↓ gradient ↓
├────────────────────────────────────────────────────────────────┤   Midnight Blue (#1E3A5F)
│                                                                │
│ ┌────────────────────────────────────────────────────────────┐│
│ │ 💭 Thinking...                                             ││ ← Status Bar
│ └────────────────────────────────────────────────────────────┘│   Purple border (#9F7AEA)
│ ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓ │   Blue background (#1E3A5F)
│                                                                │
│ [You]                                          12:34 PM       │ ← User Message
│ ░░░░░░░░░░░░                                   ░░░░░░         │   Cream text (#F7F3E9)
│ Message content here...                                       │   Gray timestamp (#6B7280)
│                                                                │
│ [Danni] /strategy                              12:35 PM       │ ← Assistant Message
│ ▓▓▓▓▓▓▓ ▓▓▓▓▓▓▓▓▓                              ░░░░░░         │   Purple name (#9F7AEA)
│                                                                │   Cream text (#F7F3E9)
│ I've noticed something fascinating...                         │
│                                                                │
│ 🎯 Strategic Pattern                                          │ ← Formatted content
│ ▓▓                                                             │   Rose gold (#B76E79)
│                                                                │   for emoji/headers
│                                                                │
│ ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░  │   Midnight Blue background
├────────────────────────────────────────────────────────────────┤
│ ● /strategy  /creative  /design  /technology  /gravity        │ ← Module Selector
│ ▓            ░░░░░░░░░  ░░░░░░   ░░░░░░░░░░░  ░░░░░░         │   Active: Purple (#9F7AEA)
│                                                                │   Inactive: Gray (#6B7280)
│ ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓  │   Deep Charcoal (#2D2D2D)
├────────────────────────────────────────────────────────────────┤
│ > Message input here_                                          │ ← Input Area
│ ░             ░░░░░                                            │   Cream text (#F7F3E9)
│ ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓  │   Charcoal background
├────────────────────────────────────────────────────────────────┤
│ Tokens: 1.2K/200K [████░░░░] 0.6% • Cost: $0.03 • /help      │ ← Footer
│ ░░░░░░                ▓▓▓▓                      ░░░░          │   Gray text (#6B7280)
│ ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓  │   Charcoal background
└────────────────────────────────────────────────────────────────┘

Legend:
▓ = Primary/accent colors (Purple, Blue, Rose Gold)
░ = Text/secondary colors (Cream, Gray)
```

---

## MODULE SYSTEM

### Module Architecture
```
                    ┌───────────────────┐
                    │   User selects    │
                    │   /strategy       │
                    └─────────┬─────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                    TUI Module Router                        │
│                                                             │
│  • Update UI highlighting                                  │
│  • Store active module in state                            │
│  • Include module in message context                       │
└──────────────────────┬──────────────────────────────────────┘
                       │
                       ▼
┌─────────────────────────────────────────────────────────────┐
│              Send message with module context               │
│                                                             │
│  {"type":"message","content":"...","module":"strategy"}    │
└──────────────────────┬──────────────────────────────────────┘
                       │
                       ▼
┌─────────────────────────────────────────────────────────────┐
│                   Rust CLI receives                         │
│                                                             │
│  • Parse module from context                                │
│  • Load module-specific prompt                             │
│  • Apply module personality                                │
└──────────────────────┬──────────────────────────────────────┘
                       │
                       ▼
┌─────────────────────────────────────────────────────────────┐
│                Agent processes with module                  │
│                                                             │
│  /strategy  → Analytical, market-focused                   │
│  /creative  → Playful, culturally aware                    │
│  /design    → Aesthetic, philosophical                     │
│  /technology→ Human-centered technical                     │
│  etc.                                                       │
└──────────────────────┬──────────────────────────────────────┘
                       │
                       ▼
┌─────────────────────────────────────────────────────────────┐
│           Response tagged with module                       │
│                                                             │
│  {"type":"message","role":"assistant",                     │
│   "content":"...","module":"strategy"}                     │
└──────────────────────┬──────────────────────────────────────┘
                       │
                       ▼
┌─────────────────────────────────────────────────────────────┐
│            TUI renders with module context                  │
│                                                             │
│  [Danni] /strategy    ← Shows which module responded       │
└─────────────────────────────────────────────────────────────┘
```

---

## ERROR HANDLING

### Graceful Degradation
```
┌────────────────────────────────────────────────────────────┐
│                   Happy Path                               │
│                                                            │
│  TUI ◄──JSON──► Rust CLI ◄──► Agent ◄──► LLM             │
│  ✓              ✓             ✓          ✓                │
└────────────────────────────────────────────────────────────┘

┌────────────────────────────────────────────────────────────┐
│                Bridge Connection Fails                     │
│                                                            │
│  TUI ✗ Rust CLI                                           │
│   │                                                        │
│   └──> Show error overlay                                 │
│   └──> Offer retry                                        │
│   └──> Fallback: Launch Rust CLI directly (legacy mode)  │
└────────────────────────────────────────────────────────────┘

┌────────────────────────────────────────────────────────────┐
│                Rust CLI Crashes                            │
│                                                            │
│  TUI ◄──JSON──► Rust CLI ✗                               │
│   │               │                                        │
│   │               └──> stderr: error message              │
│   │                                                        │
│   └──> Detect disconnection                               │
│   └──> Save session state                                 │
│   └──> Show error with details                            │
│   └──> Offer restart or fallback                          │
└────────────────────────────────────────────────────────────┘

┌────────────────────────────────────────────────────────────┐
│                LLM Provider Error                          │
│                                                            │
│  TUI ◄──JSON──► Rust CLI ◄──► Agent ◄✗─► LLM            │
│                         │                                  │
│                         └──> {"type":"error",              │
│                               "message":"Rate limit"}      │
│   │                                                        │
│   └──> Display error message                              │
│   └──> Keep conversation history                          │
│   └──> User can retry                                     │
└────────────────────────────────────────────────────────────┘

┌────────────────────────────────────────────────────────────┐
│                Markdown Rendering Fails                    │
│                                                            │
│  TUI → Glamour.Render() ✗                                │
│   │                                                        │
│   └──> Catch error                                        │
│   └──> Fallback to plain text                            │
│   └──> Log warning                                        │
│   └──> Continue normally                                  │
└────────────────────────────────────────────────────────────┘
```

---

## PERFORMANCE OPTIMIZATIONS

### Rendering Pipeline
```
┌────────────────────────────────────────────────────────────┐
│              Bubble Tea Update Cycle                       │
│                                                            │
│  Event → Update() → View() → Terminal                     │
│          ▲         ▲                                       │
│          │         └── Only if model changed              │
│          │                                                 │
│          └── Batch multiple updates                       │
│                                                            │
│  Optimizations:                                            │
│  • Debounce rapid events                                  │
│  • Batch keyboard input                                   │
│  • Virtual scrolling for large chat history              │
│  • Memoize rendered markdown                              │
│  • Lazy load messages outside viewport                    │
└────────────────────────────────────────────────────────────┘

┌────────────────────────────────────────────────────────────┐
│              Message Rendering                             │
│                                                            │
│  Message → Glamour → Cache → LipGloss → Display          │
│             ▲        │                                     │
│             │        └── Store in map[hash]string         │
│             │                                              │
│             └── Only render new/changed messages          │
│                                                            │
│  Cache Strategy:                                           │
│  • Hash message content                                   │
│  • Check cache before rendering                           │
│  • Evict old entries (LRU)                                │
│  • Limit cache size (100 messages)                        │
└────────────────────────────────────────────────────────────┘

┌────────────────────────────────────────────────────────────┐
│              Virtual Scrolling                             │
│                                                            │
│  1000 messages in history                                 │
│  Only render:                                              │
│    • Visible messages (30)                                │
│    • +10 above viewport (buffer)                          │
│    • +10 below viewport (buffer)                          │
│                                                            │
│  = 50 messages rendered instead of 1000                   │
│  = 20x performance improvement                            │
└────────────────────────────────────────────────────────────┘
```

---

These diagrams provide visual references for understanding the Danni TUI architecture. Use them alongside the main documentation for implementation.
