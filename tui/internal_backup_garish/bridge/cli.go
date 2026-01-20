package bridge

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"
)

// CLI represents a connection to the Danni Rust CLI
type CLI struct {
	cmd      *exec.Cmd
	stdin    io.WriteCloser
	stdout   io.ReadCloser
	stderr   io.ReadCloser
	mu       sync.Mutex
	running  bool
	responses chan Response
	errors   chan error
}

// NewCLI creates a new CLI bridge instance
func NewCLI(binaryPath string) (*CLI, error) {
	cmd := exec.Command(binaryPath, "session", "--tui-mode")
	cmd.Env = append(os.Environ(), "DANNI_DEBUG=false")

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to create stdin pipe: %w", err)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to create stdout pipe: %w", err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to create stderr pipe: %w", err)
	}

	cli := &CLI{
		cmd:       cmd,
		stdin:     stdin,
		stdout:    stdout,
		stderr:    stderr,
		running:   false,
		responses: make(chan Response, 100),
		errors:    make(chan error, 10),
	}

	return cli, nil
}

// Start launches the CLI process and begins reading responses
func (c *CLI) Start() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.running {
		return fmt.Errorf("CLI already running")
	}

	if err := c.cmd.Start(); err != nil {
		return fmt.Errorf("failed to start CLI: %w", err)
	}

	c.running = true

	// Start reading stdout
	go c.readStdout()

	// Start reading stderr (for debugging)
	go c.readStderr()

	return nil
}

// Stop terminates the CLI process
func (c *CLI) Stop() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if !c.running {
		return nil
	}

	// Send stop request
	stopReq := Request{
		Type: RequestTypeStop,
	}
	if err := c.sendRequest(stopReq); err != nil {
		// If we can't send stop, force kill
		return c.cmd.Process.Kill()
	}

	// Wait for process to exit
	if err := c.cmd.Wait(); err != nil {
		return fmt.Errorf("CLI process error: %w", err)
	}

	c.running = false
	close(c.responses)
	close(c.errors)

	return nil
}

// SendMessage sends a message to the CLI
func (c *CLI) SendMessage(content string, module string) error {
	req := Request{
		Type:    RequestTypeMessage,
		Content: content,
		Module:  module,
	}
	return c.sendRequest(req)
}

// Responses returns the channel for receiving responses
func (c *CLI) Responses() <-chan Response {
	return c.responses
}

// Errors returns the channel for receiving errors
func (c *CLI) Errors() <-chan error {
	return c.errors
}

// sendRequest sends a request to the CLI via stdin
func (c *CLI) sendRequest(req Request) error {
	data, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	data = append(data, '\n')

	_, err = c.stdin.Write(data)
	if err != nil {
		return fmt.Errorf("failed to write to CLI stdin: %w", err)
	}

	return nil
}

// readStdout reads responses from the CLI stdout
func (c *CLI) readStdout() {
	scanner := bufio.NewScanner(c.stdout)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		var resp Response
		if err := json.Unmarshal([]byte(line), &resp); err != nil {
			c.errors <- fmt.Errorf("failed to parse response: %w (line: %s)", err, line)
			continue
		}

		c.responses <- resp
	}

	if err := scanner.Err(); err != nil {
		c.errors <- fmt.Errorf("error reading stdout: %w", err)
	}
}

// readStderr reads error output from the CLI
func (c *CLI) readStderr() {
	scanner := bufio.NewScanner(c.stderr)
	for scanner.Scan() {
		line := scanner.Text()
		// Suppress informational stderr output (starting session, etc.)
		// Only report actual errors
		if line != "" {
			// Check if it's an actual error or just info
			if !strings.Contains(line, "starting session") &&
				!strings.Contains(line, "resuming session") &&
				!strings.Contains(line, "session id:") &&
				!strings.Contains(line, "working directory:") &&
				!strings.Contains(line, "provider:") &&
				!strings.Contains(line, "model:") {
				// This looks like a real error
				c.errors <- fmt.Errorf("CLI error: %s", line)
			}
			// Otherwise silently ignore informational messages
		}
	}
}
