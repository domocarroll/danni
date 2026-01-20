#!/bin/bash
# Final integration test - wait for actual response

SESSION="danni-final-test-$$"
PURPLE='\033[0;35m'
GREEN='\033[0;32m'
NC='\033[0m'

cleanup() {
    tmux kill-session -t $SESSION 2>/dev/null || true
}
trap cleanup EXIT

echo -e "${PURPLE}═══════════════════════════════════════════${NC}"
echo -e "${PURPLE}  DANNI TUI FINAL INTEGRATION TEST${NC}"
echo -e "${PURPLE}═══════════════════════════════════════════${NC}\n"

cd /home/dom/danni-goose-fork/tui

# Start tmux
tmux new-session -d -s $SESSION -x 120 -y 30
tmux send-keys -t $SESSION "cd /home/dom/danni-goose-fork/tui" C-m
sleep 1

# Launch TUI
echo "Launching TUI..."
tmux send-keys -t $SESSION "./bin/danni-tui" C-m
sleep 3

echo "Sending test message..."
tmux send-keys -t $SESSION "What is 2+2? Reply with just the number." C-m
sleep 2

echo "Waiting for response (15 seconds)..."
for i in {1..15}; do
    sleep 1
    tmux capture-pane -t $SESSION -p > /tmp/tui-test-$i.txt

    if grep -q "DANNI:" /tmp/tui-test-$i.txt 2>/dev/null; then
        echo -e "${GREEN}✓ Response received after ${i} seconds!${NC}"
        echo ""
        echo "Response:"
        grep "DANNI:" /tmp/tui-test-$i.txt | head -5
        echo ""

        # Check tokens
        if grep -q "Tokens:" /tmp/tui-test-$i.txt; then
            grep "Tokens:" /tmp/tui-test-$i.txt | tail -1
        fi

        break
    fi
    echo -n "."
done

echo ""
echo -e "\n${PURPLE}Final TUI State:${NC}"
tmux capture-pane -t $SESSION -p | tail -30

echo -e "\n${GREEN}✓ Test complete${NC}"
