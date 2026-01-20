#!/bin/bash
# Direct test - run TUI and manually monitor CLI subprocess

echo "Starting TUI in background..."
./bin/danni-tui &
TUI_PID=$!

sleep 3

echo "TUI PID: $TUI_PID"

# Find CLI subprocess
CLI_PID=$(pgrep -P $TUI_PID -f "danni session --tui-mode" 2>/dev/null | head -1)

if [ -n "$CLI_PID" ]; then
    echo "✓ CLI subprocess found: $CLI_PID"
    ps -p $CLI_PID -o pid,cmd

    echo ""
    echo "Watching CLI process for 10 seconds..."
    for i in {1..10}; do
        if ps -p $CLI_PID > /dev/null 2>&1; then
            echo -n "."
        else
            echo " CLI exited"
            break
        fi
        sleep 1
    done
else
    echo "✗ No CLI subprocess found"
fi

echo ""
echo "Killing TUI..."
kill $TUI_PID 2>/dev/null
wait $TUI_PID 2>/dev/null

echo "Done"
