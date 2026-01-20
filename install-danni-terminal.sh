#!/usr/bin/env bash
# ✨ DANNI TERMINAL INSTALLER
# Installs and configures Danni Terminal (Ghostty + Danni TUI)
#
# Usage: ./install-danni-terminal.sh
#
# This script will:
# 1. Check for Ghostty terminal emulator
# 2. Build and install Danni TUI
# 3. Configure Ghostty with Danni theme
# 4. Create danni-terminal launcher
# 5. Set up everything for a beautiful experience

set -e  # Exit on error

# Color codes for beautiful output
PURPLE='\033[0;35m'
BLUE='\033[0;34m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
CREAM='\033[0;37m'
BOLD='\033[1m'
NC='\033[0m' # No Color

# Print header
print_header() {
    # Read version if available
    VERSION="dev"
    if [[ -f "$(dirname "$0")/VERSION" ]]; then
        VERSION=$(cat "$(dirname "$0")/VERSION")
    fi

    echo -e "${PURPLE}${BOLD}"
    echo "✨ ═══════════════════════════════════════════════════════════"
    echo "   DANNI TERMINAL INSTALLER v${VERSION}"
    echo "   Sophisticated AI assistant terminal experience"
    echo "═══════════════════════════════════════════════════════════ ✨"
    echo -e "${NC}"
}

# Print step
print_step() {
    echo -e "${BLUE}${BOLD}▸ $1${NC}"
}

# Print success
print_success() {
    echo -e "${GREEN}✓ $1${NC}"
}

# Print warning
print_warning() {
    echo -e "${YELLOW}⚠ $1${NC}"
}

# Print error
print_error() {
    echo -e "${RED}✗ $1${NC}"
}

# Print info
print_info() {
    echo -e "${CREAM}  $1${NC}"
}

# Detect platform
detect_platform() {
    case "$(uname -s)" in
        Darwin*)
            PLATFORM="macos"
            INSTALL_DIR="$HOME/.local/bin"
            CONFIG_DIR="$HOME/.config/ghostty"
            ;;
        Linux*)
            PLATFORM="linux"
            INSTALL_DIR="$HOME/.local/bin"
            CONFIG_DIR="$HOME/.config/ghostty"
            ;;
        *)
            print_error "Unsupported platform: $(uname -s)"
            exit 1
            ;;
    esac
}

# Check if Ghostty is installed
check_ghostty() {
    print_step "Checking for Ghostty terminal emulator..."

    if command -v ghostty &> /dev/null; then
        GHOSTTY_VERSION=$(ghostty --version 2>/dev/null || echo "unknown")
        print_success "Ghostty found: $GHOSTTY_VERSION"
        return 0
    else
        print_warning "Ghostty not found"
        return 1
    fi
}

# Install Ghostty (with instructions)
install_ghostty() {
    print_step "Ghostty installation required"
    echo ""
    print_info "Danni Terminal requires Ghostty (the fast, GPU-accelerated terminal)."
    print_info "Please install Ghostty first:"
    echo ""

    if [[ "$PLATFORM" == "macos" ]]; then
        print_info "macOS:"
        echo -e "${CREAM}  brew install ghostty${NC}"
        echo -e "${CREAM}  # OR download from: https://ghostty.org/download${NC}"
    else
        print_info "Linux:"
        echo -e "${CREAM}  # Ubuntu/Debian:${NC}"
        echo -e "${CREAM}  curl -fsSL https://ghostty.org/install/ubuntu.sh | sh${NC}"
        echo -e "${CREAM}  # OR download from: https://ghostty.org/download${NC}"
    fi

    echo ""
    read -p "Would you like to open the Ghostty website? (y/n) " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        if [[ "$PLATFORM" == "macos" ]]; then
            open "https://ghostty.org/download"
        else
            xdg-open "https://ghostty.org/download" 2>/dev/null || echo "Please visit: https://ghostty.org/download"
        fi
    fi

    echo ""
    print_info "After installing Ghostty, run this installer again."
    exit 0
}

# Check dependencies
check_dependencies() {
    print_step "Checking dependencies..."

    # Check for Rust/Cargo (needed to build danni-cli)
    if ! command -v cargo &> /dev/null; then
        print_warning "Rust/Cargo not found"
        print_info "Installing Rust..."
        curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y
        source "$HOME/.cargo/env"
    fi

    # Check for Go (needed for TUI)
    if ! command -v go &> /dev/null; then
        print_warning "Go not found"
        print_info "Please install Go 1.21+ from: https://go.dev/dl/"
        exit 1
    fi

    print_success "All dependencies satisfied"
}

# Build Danni TUI
build_danni_tui() {
    print_step "Building Danni TUI..."

    # Check if pre-built binary exists (from release)
    if [[ -f "$(dirname "$0")/bin/danni-tui" ]]; then
        print_info "Using pre-built Danni TUI binary"
        cp "$(dirname "$0")/bin/danni-tui" "$INSTALL_DIR/danni-tui"
        chmod +x "$INSTALL_DIR/danni-tui"
        print_success "Danni TUI installed"
        return 0
    fi

    # Otherwise build from source
    cd "$(dirname "$0")/tui"

    # Check if TUI is ready
    if [[ ! -f "cmd/danni-tui/main.go" ]]; then
        print_warning "Danni TUI source not found - TUI is in development"
        print_info "You can still install the theme and launcher"
        print_info "The TUI will be available in future releases"
        return 1
    fi

    # Build the TUI
    if go build -o "$INSTALL_DIR/danni-tui" cmd/danni-tui/main.go; then
        print_success "Danni TUI built successfully"
    else
        print_warning "Failed to build Danni TUI (development version)"
        return 1
    fi

    cd - > /dev/null
}

# Build Danni CLI (Rust)
build_danni_cli() {
    print_step "Building Danni CLI..."

    cd "$(dirname "$0")"

    # Check if we're in the right directory
    if [[ ! -f "Cargo.toml" ]]; then
        print_error "Cannot find Cargo.toml. Are you in the danni-goose-fork directory?"
        exit 1
    fi

    # Build the CLI
    if cargo build --release --bin danni-cli 2>/dev/null || cargo build --release --bin danni 2>/dev/null; then
        # Find the binary (could be danni or danni-cli)
        if [[ -f "target/release/danni-cli" ]]; then
            cp target/release/danni-cli "$INSTALL_DIR/danni"
        elif [[ -f "target/release/danni" ]]; then
            cp target/release/danni "$INSTALL_DIR/danni"
        fi
        print_success "Danni CLI built successfully"
    else
        print_warning "Could not build Danni CLI (this is optional for TUI)"
        print_info "The TUI will still work, but some features may be limited"
    fi

    cd - > /dev/null
}

# Configure Ghostty
configure_ghostty() {
    print_step "Configuring Ghostty with Danni theme..."

    # Create config directory
    mkdir -p "$CONFIG_DIR"

    # Backup existing config if present
    if [[ -f "$CONFIG_DIR/config" ]]; then
        print_warning "Existing Ghostty config found"
        cp "$CONFIG_DIR/config" "$CONFIG_DIR/config.backup.$(date +%s)"
        print_info "Backed up to: $CONFIG_DIR/config.backup.*"
    fi

    # Copy Danni theme
    cp "$(dirname "$0")/tui/config/ghostty-theme.conf" "$CONFIG_DIR/config"

    print_success "Ghostty configured with Danni theme"
}

# Create launcher
create_launcher() {
    print_step "Creating danni-terminal launcher..."

    cat > "$INSTALL_DIR/danni-terminal" <<'EOF'
#!/usr/bin/env bash
# ✨ Danni Terminal Launcher
# Launches Ghostty with Danni configuration

# Set Danni title and launch
exec ghostty --title "✨ DANNI" "$@"
EOF

    chmod +x "$INSTALL_DIR/danni-terminal"

    print_success "Launcher created at: $INSTALL_DIR/danni-terminal"
}

# Add to PATH
setup_path() {
    print_step "Setting up PATH..."

    # Check if install dir is in PATH
    if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
        print_warning "$INSTALL_DIR not in PATH"

        # Add to appropriate shell config
        if [[ -f "$HOME/.zshrc" ]]; then
            SHELL_CONFIG="$HOME/.zshrc"
        elif [[ -f "$HOME/.bashrc" ]]; then
            SHELL_CONFIG="$HOME/.bashrc"
        else
            SHELL_CONFIG="$HOME/.profile"
        fi

        echo "" >> "$SHELL_CONFIG"
        echo "# Danni Terminal" >> "$SHELL_CONFIG"
        echo "export PATH=\"$INSTALL_DIR:\$PATH\"" >> "$SHELL_CONFIG"

        print_info "Added to PATH in: $SHELL_CONFIG"
        print_warning "Please restart your shell or run: source $SHELL_CONFIG"
    else
        print_success "PATH already configured"
    fi
}

# Print completion message
print_completion() {
    echo ""
    echo -e "${PURPLE}${BOLD}"
    echo "✨ ═══════════════════════════════════════════════════════════"
    echo "   INSTALLATION COMPLETE!"
    echo "═══════════════════════════════════════════════════════════ ✨"
    echo -e "${NC}"
    echo ""
    print_info "Danni Terminal is now installed and ready to use."
    echo ""
    print_step "To launch Danni Terminal:"
    echo -e "${CREAM}  ${BOLD}danni-terminal${NC}"
    echo ""
    print_step "Or open Ghostty normally (it will auto-launch Danni TUI):"
    echo -e "${CREAM}  ${BOLD}ghostty${NC}"
    echo ""
    print_step "Quick tips:"
    print_info "• Use Cmd+T (macOS) or Ctrl+Shift+T (Linux) for new tabs"
    print_info "• Use Cmd+D for splits"
    print_info "• Type /help in Danni for available modules"
    print_info "• Config location: $CONFIG_DIR/config"
    echo ""
    print_step "Documentation:"
    print_info "• Danni: /home/dom/danni-goose-fork/README.md"
    print_info "• Ghostty: https://ghostty.org/docs"
    echo ""
    echo -e "${PURPLE}Enjoy your sophisticated AI assistant experience! 💜${NC}"
    echo ""
}

# Main installation flow
main() {
    print_header

    # Detect platform
    detect_platform
    print_info "Platform: $PLATFORM"
    echo ""

    # Check for Ghostty
    if ! check_ghostty; then
        install_ghostty
    fi

    # Create install directory
    mkdir -p "$INSTALL_DIR"

    # Check dependencies
    check_dependencies
    echo ""

    # Build components
    build_danni_tui
    build_danni_cli
    echo ""

    # Configure
    configure_ghostty
    create_launcher
    setup_path
    echo ""

    # Done!
    print_completion
}

# Run installer
main "$@"
