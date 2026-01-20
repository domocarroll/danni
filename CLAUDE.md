# Danni Development Guide

## Building & Installing Danni

When asked to update, build, or install Danni:

```bash
# Pull latest changes
cd /path/to/danni
git fetch origin
git checkout fresh-clean
git pull origin fresh-clean

# Build release binary
cargo build --release

# Install to local bin
mkdir -p ~/.local/bin
cp target/release/danni ~/.local/bin/

# Verify installation
~/.local/bin/danni --version
```

## Zed Integration (ACP)

Danni runs as an ACP (Agent Communication Protocol) server for Zed. Configure in Zed settings:

```json
{
  "agent": {
    "profiles": {
      "danni": {
        "provider": "acp",
        "command": "~/.local/bin/danni",
        "args": ["acp"]
      }
    },
    "default_profile": "danni"
  }
}
```

## Architecture

- **ClaudeCodeProvider** (`crates/danni/src/providers/claude_code.rs`) - Wraps Claude Code CLI
  - Streaming via `--output-format stream-json --include-partial-messages`
  - Tool visibility shows file reads with line counts
  - Robust error handling skips unparseable events

## Troubleshooting

**Build fails:** Ensure Rust is installed: `curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh`

**Zed not seeing updates:** Restart Zed after installing new binary

**Streaming not working:** Check Claude Code CLI is installed and authenticated: `claude --version`
