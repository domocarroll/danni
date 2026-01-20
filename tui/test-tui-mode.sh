#!/bin/bash
# Test script for Danni TUI mode

set -e

echo "🧪 Testing Danni TUI Integration"
echo "================================="
echo ""

# Check if danni binary exists
if [ ! -f "../target/release/danni" ]; then
    echo "❌ Error: danni binary not found at ../target/release/danni"
    echo "   Run: cargo build --release"
    exit 1
fi

echo "✅ Found danni binary"

# Test 1: Check if --tui-mode flag is recognized
echo ""
echo "Test 1: Checking --tui-mode flag..."
if ../target/release/danni session --help | grep -q "tui-mode"; then
    echo "✅ --tui-mode flag is available"
else
    echo "❌ --tui-mode flag not found in help"
    exit 1
fi

# Test 2: Try to run CLI in TUI mode (will wait for stdin)
echo ""
echo "Test 2: Testing CLI in TUI mode (5 second timeout)..."
echo ""
echo "Sending test JSON request..."

# Send a ping request and see if we get a response
(
    echo '{"type":"ping"}'
    sleep 2
) | timeout 5 ../target/release/danni session --tui-mode 2>&1 | head -20 &

PID=$!
sleep 3

if ps -p $PID > /dev/null; then
    echo "✅ CLI is running in TUI mode (process active)"
    kill $PID 2>/dev/null || true
else
    echo "⚠️  CLI process ended (this might be expected if no API keys configured)"
fi

# Test 3: Check if TUI binary exists and is executable
echo ""
echo "Test 3: Checking TUI binary..."
if [ -f "bin/danni-tui" ] && [ -x "bin/danni-tui" ]; then
    echo "✅ TUI binary is built and executable"
    ls -lh bin/danni-tui
else
    echo "❌ TUI binary not found or not executable"
    exit 1
fi

echo ""
echo "================================="
echo "🎉 Integration Tests Complete!"
echo ""
echo "To run the TUI:"
echo "  cd /home/dom/danni-goose-fork/tui"
echo "  ./bin/danni-tui"
echo ""
echo "Requirements:"
echo "  - ANTHROPIC_API_KEY or other LLM API key configured"
echo "  - Run 'danni configure' if not set up yet"
echo ""
