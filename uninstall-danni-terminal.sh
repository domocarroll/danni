#!/usr/bin/env bash
# ✨ DANNI TERMINAL UNINSTALLER
# Removes Danni Terminal configuration and binaries
#
# Usage: ./uninstall-danni-terminal.sh

set -e

# Color codes
PURPLE='\033[0;35m'
BLUE='\033[0;34m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
CREAM='\033[0;37m'
BOLD='\033[1m'
NC='\033[0m'

print_header() {
    echo -e "${PURPLE}${BOLD}"
    echo "✨ ═══════════════════════════════════════════════════════════"
    echo "   DANNI TERMINAL UNINSTALLER"
    echo "═══════════════════════════════════════════════════════════ ✨"
    echo -e "${NC}"
}

print_step() {
    echo -e "${BLUE}${BOLD}▸ $1${NC}"
}

print_success() {
    echo -e "${GREEN}✓ $1${NC}"
}

print_warning() {
    echo -e "${YELLOW}⚠ $1${NC}"
}

print_info() {
    echo -e "${CREAM}  $1${NC}"
}

# Detect platform
case "$(uname -s)" in
    Darwin*|Linux*)
        INSTALL_DIR="$HOME/.local/bin"
        CONFIG_DIR="$HOME/.config/ghostty"
        ;;
    *)
        echo "Unsupported platform"
        exit 1
        ;;
esac

main() {
    print_header
    echo ""

    print_warning "This will remove:"
    print_info "• Danni TUI binary ($INSTALL_DIR/danni-tui)"
    print_info "• Danni CLI binary ($INSTALL_DIR/danni)"
    print_info "• Danni Terminal launcher ($INSTALL_DIR/danni-terminal)"
    print_info "• Ghostty configuration ($CONFIG_DIR/config)"
    print_info "  (A backup will be created if it exists)"
    echo ""
    print_warning "Ghostty itself will NOT be uninstalled"
    echo ""

    read -p "Are you sure you want to uninstall Danni Terminal? (y/N) " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        print_info "Uninstall cancelled"
        exit 0
    fi

    echo ""
    print_step "Uninstalling Danni Terminal..."

    # Remove binaries
    if [[ -f "$INSTALL_DIR/danni-tui" ]]; then
        rm "$INSTALL_DIR/danni-tui"
        print_success "Removed danni-tui"
    fi

    if [[ -f "$INSTALL_DIR/danni" ]]; then
        rm "$INSTALL_DIR/danni"
        print_success "Removed danni CLI"
    fi

    if [[ -f "$INSTALL_DIR/danni-terminal" ]]; then
        rm "$INSTALL_DIR/danni-terminal"
        print_success "Removed danni-terminal launcher"
    fi

    # Backup and remove config
    if [[ -f "$CONFIG_DIR/config" ]]; then
        BACKUP_FILE="$CONFIG_DIR/config.pre-danni-backup.$(date +%s)"
        cp "$CONFIG_DIR/config" "$BACKUP_FILE"
        rm "$CONFIG_DIR/config"
        print_success "Removed Ghostty config (backed up to: $BACKUP_FILE)"
    fi

    echo ""
    echo -e "${GREEN}${BOLD}✓ Danni Terminal uninstalled successfully${NC}"
    echo ""
    print_info "To restore Ghostty to default settings, delete: $CONFIG_DIR/"
    print_info "To reinstall Danni Terminal, run: ./install-danni-terminal.sh"
    echo ""
}

main "$@"
