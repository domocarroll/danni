# 🚀 LAUNCH DANNI TUI

## Quick Start

### 1. Configure Danni (First Time Only)
```bash
cd /home/dom/danni-goose-fork
./target/release/danni configure
```

Follow the prompts to set up your LLM provider (Anthropic, OpenAI, etc.)

### 2. Launch the TUI
```bash
cd /home/dom/danni-goose-fork/tui
./bin/danni-tui
```

### 3. Start Chatting!
- Type your message in the input area
- Press **Enter** to send
- Watch DANNI respond in real-time with the beautiful purple/blue interface

---

## Alternative: Use Environment Variable

```bash
export ANTHROPIC_API_KEY="sk-ant-your-key-here"
cd /home/dom/danni-goose-fork/tui
./bin/danni-tui
```

---

## Features

- **Scrollable chat**: Use ↑/↓ or k/j
- **Multi-line input**: Type freely, Enter to send
- **Thinking indicator**: Rose gold animated spinner
- **Token tracking**: Real-time in footer
- **Module support**: /strategy, /creative, /design, etc.
- **Beautiful markdown**: (Coming soon with Glamour)

---

## Troubleshooting

**Binary not found?**
```bash
cargo build --release --package danni-cli
```

**TUI won't start?**
```bash
go build -o bin/danni-tui cmd/danni-tui/main.go
```

**No response?**
- Check API key is configured
- Try: `../target/release/danni session` first
- Verify provider is working

---

✨ **Enjoy the most beautiful AI terminal interface!** ✨
