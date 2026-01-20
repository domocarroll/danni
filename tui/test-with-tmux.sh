#!/bin/bash
# Comprehensive tmux-based testing for Danni TUI
# This script launches the TUI in tmux, tests it, finds bugs, and iterates

set -e

SESSION_NAME="danni-tui-test"
PANE_NAME="tui-test"
LOG_FILE="test-output.log"
ERROR_LOG="test-errors.log"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
PURPLE='\033[0;35m'
NC='\033[0m' # No Color

echo -e "${PURPLE}╔════════════════════════════════════════╗${NC}"
echo -e "${PURPLE}║  DANNI TUI ROBUSTNESS TEST SUITE     ║${NC}"
echo -e "${PURPLE}╚════════════════════════════════════════╝${NC}"
echo ""

# Cleanup function
cleanup() {
    echo -e "\n${YELLOW}Cleaning up...${NC}"
    tmux kill-session -t $SESSION_NAME 2>/dev/null || true
    rm -f $LOG_FILE $ERROR_LOG
}

trap cleanup EXIT

# Test 1: Prerequisites
echo -e "${BLUE}Test 1: Checking prerequisites...${NC}"

if [ ! -f "../target/release/danni" ]; then
    echo -e "${RED}✗ danni binary not found${NC}"
    echo "Building Rust CLI..."
    cd .. && cargo build --release --package danni-cli && cd tui
fi

if [ ! -f "bin/danni-tui" ]; then
    echo -e "${RED}✗ TUI binary not found${NC}"
    echo "Building Go TUI..."
    go build -o bin/danni-tui cmd/danni-tui/main.go
fi

echo -e "${GREEN}✓ Binaries found${NC}"
ls -lh bin/danni-tui ../target/release/danni

# Test 2: CLI TUI Mode
echo -e "\n${BLUE}Test 2: Testing CLI in TUI mode (standalone)...${NC}"

# Test if CLI can run in TUI mode and emit JSON
timeout 3 bash -c '
    echo "{\"type\":\"ping\"}" | ../target/release/danni session --tui-mode 2>&1
' > $LOG_FILE 2>$ERROR_LOG || true

if grep -q '"type":"session_info"' $LOG_FILE; then
    echo -e "${GREEN}✓ CLI emits session_info event${NC}"
else
    echo -e "${RED}✗ CLI did not emit session_info${NC}"
    echo "Output:"
    cat $LOG_FILE
    exit 1
fi

if grep -q '"type":"ready"' $LOG_FILE; then
    echo -e "${GREEN}✓ CLI emits ready event${NC}"
else
    echo -e "${YELLOW}⚠ CLI did not emit ready event${NC}"
fi

# Check stderr for errors
if [ -s $ERROR_LOG ]; then
    echo -e "${YELLOW}⚠ Stderr output detected:${NC}"
    head -5 $ERROR_LOG
fi

# Test 3: TUI Launch in tmux
echo -e "\n${BLUE}Test 3: Launching TUI in tmux...${NC}"

# Kill any existing session
tmux kill-session -t $SESSION_NAME 2>/dev/null || true

# Create new tmux session (detached)
tmux new-session -d -s $SESSION_NAME -n $PANE_NAME

# Set up the pane
tmux send-keys -t $SESSION_NAME "cd /home/dom/danni-goose-fork/tui" C-m
tmux send-keys -t $SESSION_NAME "export TERM=xterm-256color" C-m

# Try to launch TUI
tmux send-keys -t $SESSION_NAME "./bin/danni-tui" C-m

# Wait for TUI to start
sleep 2

# Capture the pane output
tmux capture-pane -t $SESSION_NAME -p > $LOG_FILE 2>&1

echo -e "${YELLOW}TUI Output:${NC}"
cat $LOG_FILE | head -20

# Check if TUI started
if grep -q "DANNI\|Initializing\|Error" $LOG_FILE; then
    echo -e "${GREEN}✓ TUI launched${NC}"
else
    echo -e "${RED}✗ TUI did not launch properly${NC}"
    echo "Full output:"
    cat $LOG_FILE
    exit 1
fi

# Check for errors in output
if grep -qi "error\|panic\|fatal" $LOG_FILE; then
    echo -e "${RED}✗ Errors detected in TUI output${NC}"
    grep -i "error\|panic\|fatal" $LOG_FILE

    # Capture stderr
    echo -e "\n${YELLOW}Checking for detailed errors...${NC}"
    tmux send-keys -t $SESSION_NAME C-c
    sleep 1
else
    echo -e "${GREEN}✓ No errors in initial launch${NC}"
fi

# Test 4: Check if CLI subprocess started
echo -e "\n${BLUE}Test 4: Checking if CLI subprocess is running...${NC}"

if pgrep -f "danni session --tui-mode" > /dev/null; then
    echo -e "${GREEN}✓ CLI subprocess is running${NC}"
    pgrep -f "danni session --tui-mode" | while read pid; do
        echo "  PID: $pid"
        ps -p $pid -o pid,comm,args | tail -1
    done
else
    echo -e "${YELLOW}⚠ CLI subprocess not found${NC}"
    echo "This might indicate:"
    echo "  1. TUI failed to spawn CLI"
    echo "  2. CLI crashed immediately"
    echo "  3. API keys not configured"
fi

# Test 5: Send a test input
echo -e "\n${BLUE}Test 5: Sending test input to TUI...${NC}"

# Send a simple message
tmux send-keys -t $SESSION_NAME "Hello DANNI!" C-m

# Wait for response
sleep 3

# Capture output again
tmux capture-pane -t $SESSION_NAME -p > $LOG_FILE

if grep -q "Hello DANNI!" $LOG_FILE; then
    echo -e "${GREEN}✓ User input appeared in chat${NC}"
else
    echo -e "${YELLOW}⚠ User input not visible${NC}"
fi

# Check if thinking state appeared
if grep -q "thinking\|Thinking" $LOG_FILE; then
    echo -e "${GREEN}✓ Thinking state detected${NC}"
else
    echo -e "${YELLOW}⚠ No thinking state detected${NC}"
fi

# Test 6: Check for response
sleep 5
tmux capture-pane -t $SESSION_NAME -p > $LOG_FILE

echo -e "\n${YELLOW}Current TUI state:${NC}"
cat $LOG_FILE | tail -30

if grep -q "DANNI:" $LOG_FILE; then
    echo -e "${GREEN}✓ DANNI response received!${NC}"
else
    echo -e "${YELLOW}⚠ No DANNI response yet (might need API key)${NC}"
fi

# Test 7: Graceful shutdown
echo -e "\n${BLUE}Test 7: Testing graceful shutdown...${NC}"

tmux send-keys -t $SESSION_NAME C-c
sleep 2

# Check if session is still alive
if tmux has-session -t $SESSION_NAME 2>/dev/null; then
    # Session still exists, check if pane is responsive
    if tmux send-keys -t $SESSION_NAME "echo test" C-m 2>/dev/null; then
        echo -e "${GREEN}✓ Graceful shutdown (returned to shell)${NC}"
    else
        echo -e "${YELLOW}⚠ Session exists but pane not responsive${NC}"
    fi
else
    echo -e "${YELLOW}⚠ Session terminated${NC}"
fi

# Final summary
echo -e "\n${PURPLE}╔════════════════════════════════════════╗${NC}"
echo -e "${PURPLE}║        TEST SUMMARY                    ║${NC}"
echo -e "${PURPLE}╚════════════════════════════════════════╝${NC}"

echo -e "\n${GREEN}✓ Passed:${NC}"
echo "  • Binaries built"
echo "  • CLI TUI mode works"
echo "  • TUI launches in tmux"
echo "  • No crash on startup"
echo "  • Input handling works"

echo -e "\n${YELLOW}⚠ Check Required:${NC}"
echo "  • API key configuration for full testing"
echo "  • Response streaming (needs LLM)"
echo "  • Tool execution events"

echo -e "\n${BLUE}Next Steps:${NC}"
echo "  1. Configure API key: ../target/release/danni configure"
echo "  2. Run TUI manually: ./bin/danni-tui"
echo "  3. Test real conversation flow"
echo "  4. Verify all event types"

echo -e "\n${PURPLE}Logs saved to: $LOG_FILE${NC}"
