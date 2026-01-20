#!/bin/bash
# Trace CLI subprocess lifecycle

SESSION="trace-cli-$$"

cleanup() {
    tmux kill-session -t $SESSION 2>/dev/null || true
}
trap cleanup EXIT

tmux new-session -d -s $SESSION
tmux send-keys -t $SESSION "cd /home/dom/danni-goose-fork/tui" C-m
tmux send-keys -t $SESSION "./bin/danni-tui 2>/tmp/tui-stderr.log" C-m

sleep 3

CLIP=$(pgrep -f "danni session --tui-mode" | head -1)
echo "CLI PID after TUI launch: $CLIP"

if [ -n "$CLIP" ]; then
    echo "CLI is running"
    ps -p $CLIP -o pid,stat,cmd
fi

tmux send-keys -t $SESSION "Test message" C-m
echo "Sent message, waiting..."
sleep 10

CLIP2=$(pgrep -f "danni session --tui-mode" | head -1)
echo "CLI PID after message: $CLIP2"

if [ -n "$CLIP2" ]; then
    echo "CLI still running"
    ps -p $CLIP2 -o pid,stat,cmd
else
    echo "CLI died"
fi

echo ""
echo "TUI stderr:"
cat /tmp/tui-stderr.log 2>/dev/null | tail -30

echo ""
echo "TUI State:"
tmux capture-pane -t $SESSION -p | tail -25
