# ✨ DANNI TERMINAL
**The Sophisticated AI Terminal Experience**

Danni Terminal combines the speed and beauty of [Ghostty](https://ghostty.org) with the intelligence of Danni's AI assistant, creating a seamless, elegant terminal experience.

---

## What You Get

🎨 **Beautiful Aesthetic**
- Danni's signature purple/midnight blue theme
- GPU-accelerated rendering via Ghostty
- Sophisticated color palette optimized for long coding sessions

🧠 **Intelligent Assistant**
- Full Danni AI capabilities built into your terminal
- Module-based expertise (strategy, creative, design, technology, etc.)
- Context-aware assistance while you work

⚡ **Performance**
- Blazing fast terminal emulation
- Native multiplexing (tabs, splits)
- Smooth scrolling and rendering

🎯 **Zero Configuration**
- Launches Danni TUI automatically
- Pre-configured theme and keybindings
- Works out of the box

---

## Quick Start

### Installation

```bash
# Clone the repository (if you haven't already)
git clone https://github.com/subfracture/danni-goose-fork.git
cd danni-goose-fork

# Run the installer
./install-danni-terminal.sh
```

The installer will:
1. Check for Ghostty (guide you to install if needed)
2. Build the Danni TUI interface
3. Configure everything with Danni's theme
4. Create launcher scripts
5. Set up PATH

### Launch

After installation, start Danni Terminal with:

```bash
danni-terminal
```

Or just open Ghostty normally - it will automatically launch Danni TUI:

```bash
ghostty
```

---

## Requirements

### Supported Platforms
- ✅ macOS 12+ (Monterey or later)
- ✅ Linux (Ubuntu 20.04+, Fedora 35+, Arch, etc.)
- ⏳ Windows (coming when Ghostty adds support)

### Dependencies

**Automatically installed/checked:**
- Rust 1.70+ (for Danni CLI)
- Go 1.21+ (for TUI)

**You need to install:**
- [Ghostty](https://ghostty.org/download) terminal emulator

#### Installing Ghostty

**macOS:**
```bash
brew install ghostty
```

**Linux (Ubuntu/Debian):**
```bash
curl -fsSL https://ghostty.org/install/ubuntu.sh | sh
```

**Other Linux:**
See [Ghostty download page](https://ghostty.org/download)

---

## Usage

### Basic Commands

Once Danni Terminal is running, you have access to all Danni modules:

```
/strategy    - Strategic analysis and market insights
/creative    - Creative direction and cultural insights
/design      - Design systems and visual strategy
/technology  - Technical architecture and implementation
/gravity     - Deep pattern recognition across domains
/validate    - Validation and quality assurance
/synthesize  - Cross-domain synthesis and innovation
/recall      - Search and recall from conversation history
/upload      - Import and analyze documents/assets
```

### Multiplexing

Ghostty has native tabs and splits - no tmux needed!

**Keyboard Shortcuts:**

| Action | macOS | Linux |
|--------|-------|-------|
| New Tab | `Cmd+T` | `Ctrl+Shift+T` |
| New Split (Right) | `Cmd+D` | `Ctrl+Shift+D` |
| New Split (Down) | `Cmd+Shift+D` | `Ctrl+Alt+D` |
| Close Tab/Split | `Cmd+W` | `Ctrl+Shift+W` |
| Next Tab | `Cmd+Shift+]` | `Ctrl+Page Down` |
| Previous Tab | `Cmd+Shift+[` | `Ctrl+Page Up` |
| Zoom In | `Cmd+=` | `Ctrl+=` |
| Zoom Out | `Cmd+-` | `Ctrl+-` |

### Multi-Session Workflows

The real power: run multiple Danni sessions simultaneously

**Example: Full-Stack Development**
```
┌─────────────┬─────────────┐
│  /strategy  │  /creative  │  ← Two Danni sessions
│  Planning   │  Branding   │    working in parallel
├─────────────┴─────────────┤
│  /technology                │  ← Third session for
│  Implementation             │    technical work
└─────────────────────────────┘
```

Just open splits and each will have its own Danni instance!

---

## Configuration

### Ghostty Config Location

```
~/.config/ghostty/config
```

The installer creates this automatically with Danni's theme.

### Customizing

Want to tweak colors, fonts, or behavior? Edit the config file:

```bash
# macOS/Linux
nano ~/.config/ghostty/config
```

**Common customizations:**

```conf
# Change font
font-family = "JetBrains Mono"
font-size = 14

# Adjust window padding
window-padding-x = 16
window-padding-y = 16

# Modify cursor style
cursor-style = underline
cursor-style-blink = false
```

See [Ghostty docs](https://ghostty.org/docs/config/reference) for all options.

### Themes

Want to switch themes temporarily?

```bash
# Use Ghostty's built-in themes
ghostty --config-file=~/.config/ghostty/config --theme=nord

# Or create custom themes in:
~/.config/ghostty/themes/
```

---

## Advanced Usage

### Multiple Profiles

Create different Danni Terminal profiles:

```bash
# Create alternate config
cp ~/.config/ghostty/config ~/.config/ghostty/config-dev

# Edit config-dev for development-specific settings
# Then launch with:
ghostty --config-file=~/.config/ghostty/config-dev
```

### Integration with IDEs

Use Danni Terminal as your IDE's integrated terminal:

**VS Code:**
```json
{
  "terminal.external.osxExec": "Ghostty.app",
  "terminal.external.linuxExec": "ghostty"
}
```

**JetBrains IDEs:**
Settings → Tools → Terminal → Shell path: `/Users/you/.local/bin/danni-terminal`

### Shell Integration

Ghostty provides rich shell integration:

- **Smart cursor:** Different shapes for vim modes
- **Sudo password:** Secure password prompts
- **Title updates:** Shows current command/directory
- **Jump to previous command:** Cmd+Up/Down

This works automatically with zsh, bash, fish, and elvish.

---

## Troubleshooting

### "ghostty: command not found"

Ghostty isn't installed or not in PATH.

**Fix:**
```bash
# macOS
brew install ghostty

# Linux - check Ghostty install docs
# Then restart terminal
```

### "danni-tui: command not found"

Installation didn't complete or PATH not set.

**Fix:**
```bash
# Check if binary exists
ls -la ~/.local/bin/danni-tui

# If it exists, add to PATH:
echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.zshrc  # or ~/.bashrc
source ~/.zshrc
```

### Colors look wrong

Your terminal might not support True Color.

**Fix:**
```bash
# Check color support
echo $COLORTERM  # Should show "truecolor"

# If not, set it:
export COLORTERM=truecolor
```

### TUI crashes on launch

Usually a build issue.

**Fix:**
```bash
# Rebuild from source
cd ~/danni-goose-fork/tui
go clean
go build -o ~/.local/bin/danni-tui cmd/danni-tui/main.go

# Test directly
~/.local/bin/danni-tui
```

### Keybindings conflict with other apps

Customize in Ghostty config:

```conf
# Disable a keybinding
keybind = super+t=ignore

# Or remap it
keybind = super+shift+t=new_tab
```

---

## Uninstall

If you need to remove Danni Terminal:

```bash
cd ~/danni-goose-fork
./uninstall-danni-terminal.sh
```

This removes:
- Danni TUI binary
- Danni CLI binary
- danni-terminal launcher
- Ghostty configuration (with backup)

Ghostty itself remains installed (uninstall separately if desired).

---

## What's Next?

### Enhance Your Setup

1. **Custom Themes:** Create your own color schemes
2. **Aliases:** Add shell aliases for quick Danni commands
3. **Automation:** Use Danni in scripts via `danni-terminal -e "your command"`

### Learn More

- **Danni Modules:** See `DANNI_TRANSFORMATION_BLUEPRINT.md`
- **TUI Architecture:** See `DANNI_TUI_ARCHITECTURE.md`
- **Ghostty Docs:** https://ghostty.org/docs
- **Community:** Join discussions at https://github.com/subfracture/danni

---

## Architecture

Curious how it works?

```
┌─────────────────────────────────┐
│   Ghostty (Terminal Emulator)  │  ← Fast, GPU-accelerated
│   ┌─────────────────────────┐  │     Native UI (Swift/GTK)
│   │  Danni TUI (Go)         │  │  ← Beautiful interface
│   │  ┌───────────────────┐  │  │     Charmbracelet stack
│   │  │ Danni Core (Rust) │  │  │  ← AI agent core
│   │  └───────────────────┘  │  │     MCP extensions
│   └─────────────────────────┘  │
└─────────────────────────────────┘
```

**Clean separation:**
- Ghostty: Terminal rendering & multiplexing
- TUI (Go): User interface & interaction
- Core (Rust): AI logic, LLM communication, tools

This architecture means:
- No fork maintenance burden
- Benefit from Ghostty's ongoing development
- TUI can evolve independently
- Core remains language-agnostic

---

## Philosophy

Danni Terminal embodies our design principles:

**Sophistication through Simplicity**
- One command to install
- Zero configuration needed
- Beautiful by default
- Powerful when needed

**Progressive Revelation**
- Start with basic chat
- Discover modules naturally
- Advanced features available when ready
- Multiplexing for power users

**Warmth & Intelligence**
- Colors create calm, focused environment
- Responses feel personal and insightful
- Technical excellence without coldness
- Joy in the details

---

## Credits

**Danni Terminal** is built on the shoulders of giants:

- **[Ghostty](https://ghostty.org)** by Mitchell Hashimoto - The foundation
- **[Charmbracelet](https://charm.sh)** - Beautiful TUI components
- **[Goose](https://github.com/square/goose)** - Original agent architecture
- **Anthropic Claude** - The intelligence behind Danni

Created with care by the [Subfracture](https://subfracture.com) team.

---

## Support

### Getting Help

- **Issues:** https://github.com/subfracture/danni/issues
- **Discussions:** https://github.com/subfracture/danni/discussions
- **Email:** support@subfracture.com

### Contributing

We welcome contributions! See `CONTRIBUTING.md` for guidelines.

Areas we'd love help with:
- Additional themes
- Platform-specific packaging
- Documentation improvements
- Testing on different environments

---

**✨ Welcome to Danni Terminal. Let's build something beautiful together.**
