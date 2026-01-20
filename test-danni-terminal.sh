#!/usr/bin/env bash
# Quick test script for Danni Terminal installation

echo "✨ Testing Danni Terminal Setup..."
echo ""

# Check binaries
echo "Checking binaries..."
command -v ghostty >/dev/null 2>&1 && echo "✓ ghostty found" || echo "✗ ghostty not found"
command -v danni-tui >/dev/null 2>&1 && echo "✓ danni-tui found" || echo "✗ danni-tui not found"
command -v danni-terminal >/dev/null 2>&1 && echo "✓ danni-terminal launcher found" || echo "✗ danni-terminal launcher not found"

echo ""
echo "Checking configuration..."
if [ -f "$HOME/.config/ghostty/config" ]; then
    echo "✓ Ghostty config found"
    echo "  Location: $HOME/.config/ghostty/config"
else
    echo "✗ Ghostty config not found"
fi

echo ""
echo "Checking PATH..."
if [[ ":$PATH:" == *":$HOME/.local/bin:"* ]]; then
    echo "✓ ~/.local/bin in PATH"
else
    echo "⚠ ~/.local/bin not in PATH (may need shell restart)"
fi

echo ""
echo "Test complete!"
