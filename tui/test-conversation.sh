#!/bin/bash
# Test multi-turn conversation

SESSION="conv-test-$$"

cleanup() {
    tmux kill-session -t $SESSION 2>/dev/null || true
}
trap cleanup EXIT

echo "🧪 Testing Multi-Turn Conversation"
echo "===================================="
echo ""

cd /home/dom/danni-goose-fork/tui

# Create session
tmux new-session -d -s $SESSION -x 120 -y 30
tmux send-keys -t $SESSION "cd /home/dom/danni-goose-fork/tui" C-m
sleep 1

# Launch TUI
tmux send-keys -t $SESSION "./bin/danni-tui" C-m
sleep 3

echo "Turn 1: Asking a question..."
tmux send-keys -t $SESSION "What is 5+3?" C-m
sleep 6

tmux capture-pane -t $SESSION -p | grep -E "User:|DANNI:" | tail -4

echo ""
echo "Turn 2: Follow-up question..."
tmux send-keys -t $SESSION "What about 10+10?" C-m
sleep 6

tmux capture-pane -t $SESSION -p | grep -E "User:|DANNI:" | tail -6

echo ""
echo "Turn 3: Different topic..."
tmux send-keys -t $SESSION "Name a color" C-m
sleep 6

echo ""
echo "Full conversation history:"
tmux capture-pane -t $SESSION -p | grep -E "User:|DANNI:"

echo ""
echo "✅ Conversation test complete"
