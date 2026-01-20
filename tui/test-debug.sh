#!/bin/bash
# Debug script to test TUI thoroughly

echo "🔍 Debug Test: Step-by-step TUI Validation"
echo "==========================================="
echo ""

# Test 1: Check binary path resolution
echo "Test 1: Binary Path Resolution"
cd /home/dom/danni-goose-fork/tui
go run cmd/danni-tui/main.go 2>&1 &
PID=$!
sleep 3

if ps -p $PID > /dev/null 2>&1; then
    echo "✅ TUI process is running (PID: $PID)"
    echo "   Terminating..."
    kill $PID 2>/dev/null
    wait $PID 2>/dev/null
else
    echo "❌ TUI process died immediately"
    echo "   This indicates initialization failure"
fi

echo ""
echo "Test 2: CLI Standalone JSON Output"
timeout 2 bash -c 'echo "{\"type\":\"ping\"}" | ../target/release/danni session --tui-mode' 2>&1 | head -10

echo ""
echo "Test 3: Check what binary TUI would find"
cd /home/dom/danni-goose-fork/tui
if [ -f "../target/release/danni" ]; then
    echo "✅ Found: $(realpath ../target/release/danni)"
    ls -lh ../target/release/danni
else
    echo "❌ Not found: ../target/release/danni"
fi

echo ""
echo "Test 4: Verify TUI can spawn subprocess"
cat > /tmp/test-spawn.go << 'EOF'
package main
import (
    "fmt"
    "os/exec"
)
func main() {
    cmd := exec.Command("echo", "test")
    out, err := cmd.Output()
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    fmt.Printf("Output: %s\n", out)
}
EOF

go run /tmp/test-spawn.go
rm /tmp/test-spawn.go

echo ""
echo "✅ Debug tests complete"
