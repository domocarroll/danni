#!/bin/bash
# Manual test to trace the exact message flow

echo "🔍 Manual Message Flow Test"
echo "============================"
echo ""

# Test: Send a message directly to CLI in TUI mode
echo "Sending message to CLI and watching output..."
echo ""

# Create a test script that sends a message after the CLI is ready
(
    # Wait for ready signal
    sleep 1
    # Send a message
    echo '{"type":"message","content":"Say hello in exactly 3 words"}'
    # Wait for response
    sleep 10
) | timeout 12 ../target/release/danni session --tui-mode 2>&1

echo ""
echo "Test complete. Check above for:"
echo "  1. session_info event"
echo "  2. ready event"
echo "  3. thinking event (after message sent)"
echo "  4. message_chunk events (response)"
echo "  5. complete event"
