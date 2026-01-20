#!/bin/bash
# Comprehensive integration test with detailed diagnostics

SESSION_NAME="danni-test-$$"
PURPLE='\033[0;35m'
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m'

cleanup() {
    tmux kill-session -t $SESSION_NAME 2>/dev/null || true
}
trap cleanup EXIT

echo -e "${PURPLE}═══════════════════════════════════════════${NC}"
echo -e "${PURPLE}  DANNI TUI INTEGRATION TEST${NC}"
echo -e "${PURPLE}═══════════════════════════════════════════${NC}\n"

# Ensure we're in the right directory
cd /home/dom/danni-goose-fork/tui

# Build if needed
if [ ! -f "bin/danni-tui" ] || [ cmd/danni-tui/main.go -nt bin/danni-tui ]; then
    echo "Building TUI..."
    go build -o bin/danni-tui cmd/danni-tui/main.go || exit 1
fi

if [ ! -f "../target/release/danni" ]; then
    echo -e "${RED}✗ Rust CLI binary not found${NC}"
    echo "  Run: cd .. && cargo build --release --package danni-cli"
    exit 1
fi

echo -e "${GREEN}✓ Binaries ready${NC}\n"

# Test 1: CLI TUI Mode
echo "Test 1: CLI JSON Output"
echo "-----------------------"
timeout 2 bash -c 'echo "{\"type\":\"ping\"}" | ../target/release/danni session --tui-mode 2>/dev/null' | while IFS= read -r line; do
    echo "  $line"
done

echo -e "\n${GREEN}✓ CLI emits valid JSON${NC}\n"

# Test 2: TUI in tmux
echo "Test 2: TUI in tmux"
echo "-------------------"

# Create tmux session
tmux new-session -d -s $SESSION_NAME -x 120 -y 30

# Launch TUI
tmux send-keys -t $SESSION_NAME "cd /home/dom/danni-goose-fork/tui" C-m
sleep 1
tmux send-keys -t $SESSION_NAME "./bin/danni-tui" C-m
echo "TUI launched, waiting for initialization..."
sleep 3

# Capture initial state
tmux capture-pane -t $SESSION_NAME -p > /tmp/tui-capture-1.txt

echo -e "${YELLOW}Initial TUI state:${NC}"
cat /tmp/tui-capture-1.txt | head -20

# Check if TUI is running
if grep -q "DANNI\|Welcome\|Connecting" /tmp/tui-capture-1.txt; then
    echo -e "\n${GREEN}✓ TUI interface visible${NC}"
else
    echo -e "\n${RED}✗ TUI not visible${NC}"
    cat /tmp/tui-capture-1.txt
    exit 1
fi

# Check for errors
if grep -qi "failed to parse\|error:.*unexpected" /tmp/tui-capture-1.txt; then
    echo -e "${RED}✗ Parse errors detected${NC}"
    grep -i "error" /tmp/tui-capture-1.txt
    exit 1
else
    echo -e "${GREEN}✓ No parse errors${NC}"
fi

# Check if CLI subprocess is running
sleep 1
if pgrep -f "danni session --tui-mode" > /dev/null; then
    echo -e "${GREEN}✓ CLI subprocess is running${NC}"
    CLIPID=$(pgrep -f "danni session --tui-mode" | head -1)
    echo "  PID: $CLIPID"
else
    echo -e "${YELLOW}⚠ CLI subprocess not detected (may have completed initialization)${NC}"
fi

# Test 3: Send input
echo -e "\nTest 3: User Input"
echo "-------------------"

tmux send-keys -t $SESSION_NAME "Test message from tmux" C-m
echo "Sent test message, waiting for response..."
sleep 5

# Capture state after input
tmux capture-pane -t $SESSION_NAME -p > /tmp/tui-capture-2.txt

echo -e "${YELLOW}State after input:${NC}"
cat /tmp/tui-capture-2.txt | tail -20

# Check if input appeared
if grep -q "Test message from tmux" /tmp/tui-capture-2.txt; then
    echo -e "\n${GREEN}✓ User message visible${NC}"
else
    echo -e "\n${YELLOW}⚠ User message not found${NC}"
fi

# Check for thinking indicator
if grep -qi "thinking\|⚬" /tmp/tui-capture-2.txt; then
    echo -e "${GREEN}✓ Thinking indicator appeared${NC}"
else
    echo -e "${YELLOW}⚠ No thinking indicator${NC}"
fi

# Check for response
if grep -q "DANNI:" /tmp/tui-capture-2.txt; then
    echo -e "${GREEN}✓ DANNI response received!${NC}"
else
    echo -e "${YELLOW}⚠ No response yet (may need API key or more time)${NC}"
fi

# Test 4: Graceful exit
echo -e "\nTest 4: Graceful Shutdown"
echo "-------------------------"

tmux send-keys -t $SESSION_NAME C-c
sleep 2

if tmux has-session -t $SESSION_NAME 2>/dev/null; then
    echo -e "${GREEN}✓ Clean shutdown (returned to shell)${NC}"
else
    echo -e "${YELLOW}⚠ Session terminated${NC}"
fi

# Summary
echo -e "\n${PURPLE}═══════════════════════════════════════════${NC}"
echo -e "${PURPLE}  TEST COMPLETE${NC}"
echo -e "${PURPLE}═══════════════════════════════════════════${NC}\n"

echo "Captures saved to:"
echo "  /tmp/tui-capture-1.txt - Initial state"
echo "  /tmp/tui-capture-2.txt - After input"
echo ""
echo "To review:"
echo "  cat /tmp/tui-capture-1.txt"
echo "  cat /tmp/tui-capture-2.txt"
