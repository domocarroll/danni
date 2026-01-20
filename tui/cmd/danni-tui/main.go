package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
	"github.com/subfracture/danni/tui/internal/bridge"
	"github.com/subfracture/danni/tui/internal/components"
	"github.com/subfracture/danni/tui/internal/styles"
)

type model struct {
	width           int
	height          int
	header          components.Header
	footer          components.Footer
	viewport        viewport.Model
	textarea        textarea.Model
	spinner         spinner.Model
	messages        []string
	ready           bool
	sessionID       string
	cli             *bridge.CLI
	thinking        bool
	thinkingStart   time.Time // When thinking began
	err             error
	tokensUsed      int
	tokensTotal     int
	estimatedCost   float64
	currentResponse strings.Builder // For streaming accumulation
	respondingNow   bool            // Whether we're mid-stream
}

// Danni's thinking phrases - in her voice
var thinkingPhrases = []string{
	"Danni is thinking...",
	"Finding the thread...",
	"Following a pattern...",
	"Something's taking shape...",
	"Weaving thoughts together...",
	"There's something here...",
	"Tracing the edges...",
	"The picture is forming...",
	"Sitting with this...",
	"Connecting what wants to connect...",
	"Sensing the shape of it...",
	"Pulling at something interesting...",
	"Almost there...",
	"Deeper still...",
}

// Message types from CLI
type cliEventMsg bridge.Response

// Error from CLI
type cliErrorMsg error

func findDanniBinary() (string, error) {
	// Check in release directory first (relative to tui directory)
	releasePath := filepath.Join("..", "target", "release", "danni")
	if _, err := os.Stat(releasePath); err == nil {
		absPath, _ := filepath.Abs(releasePath)
		return absPath, nil
	}

	// Try current directory
	localPath := "./danni"
	if _, err := os.Stat(localPath); err == nil {
		absPath, _ := filepath.Abs(localPath)
		return absPath, nil
	}

	// Try PATH
	path, err := exec.LookPath("danni")
	if err == nil {
		return path, nil
	}

	return "", fmt.Errorf("danni binary not found")
}

func initialModel() (model, error) {
	// Find the danni binary
	binaryPath, err := findDanniBinary()
	if err != nil {
		return model{}, fmt.Errorf("failed to find danni binary: %w", err)
	}

	// Create CLI bridge
	cli, err := bridge.NewCLI(binaryPath)
	if err != nil {
		return model{}, fmt.Errorf("failed to create CLI bridge: %w", err)
	}

	// Initialize textarea - single line, clean
	ta := textarea.New()
	ta.Placeholder = "Type your message..."
	ta.Focus()
	ta.Prompt = ""
	ta.CharLimit = 8192
	ta.SetWidth(80)
	ta.SetHeight(1)
	ta.ShowLineNumbers = false

	// Minimal styling - NO backgrounds
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()
	ta.FocusedStyle.Base = lipgloss.NewStyle().Foreground(styles.White)
	ta.FocusedStyle.Placeholder = lipgloss.NewStyle().Foreground(styles.GrayMed)
	ta.BlurredStyle.Base = lipgloss.NewStyle().Foreground(styles.GrayMed)

	// Initialize spinner with elegant constellation
	s := spinner.New()
	s.Spinner = components.ConstellationSpinner
	s.Style = lipgloss.NewStyle().Foreground(styles.Gold)

	return model{
		sessionID: "connecting...",
		textarea:  ta,
		spinner:   s,
		cli:       cli,
		thinking:  false,
		messages: []string{
			lipgloss.NewStyle().Foreground(styles.GrayMed).Render("  ✦ Finding the thread..."),
		},
		header: components.Header{
			Module:    "/strategy",
			SessionID: "connecting...",
		},
		footer: components.Footer{
			TokensUsed:    0,
			TokensTotal:   100000,
			EstimatedCost: 0.0000,
		},
		tokensTotal: 100000,
	}, nil
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		textarea.Blink,
		m.spinner.Tick,
		startCLI(m.cli),
		listenForEvents(m.cli),
	)
}

// Start the CLI process
func startCLI(cli *bridge.CLI) tea.Cmd {
	return func() tea.Msg {
		if err := cli.Start(); err != nil {
			return cliErrorMsg(err)
		}
		return nil
	}
}

// Listen for events from CLI
func listenForEvents(cli *bridge.CLI) tea.Cmd {
	return func() tea.Msg {
		select {
		case resp := <-cli.Responses():
			return cliEventMsg(resp)
		case err := <-cli.Errors():
			return cliErrorMsg(err)
		}
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		tiCmd      tea.Cmd
		vpCmd      tea.Cmd
		spinnerCmd tea.Cmd
		cmds       []tea.Cmd
	)

	m.textarea, tiCmd = m.textarea.Update(msg)
	m.viewport, vpCmd = m.viewport.Update(msg)
	m.spinner, spinnerCmd = m.spinner.Update(msg)

	cmds = append(cmds, tiCmd, vpCmd, spinnerCmd)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			if m.cli != nil {
				m.cli.Stop()
			}
			return m, tea.Quit

		case tea.KeyEnter:
			// Send message on plain enter
			value := strings.TrimSpace(m.textarea.Value())
			if value != "" && m.cli != nil {
				// Send to CLI
				if err := m.cli.SendMessage(value, m.header.Module); err != nil {
					m.err = err
					m.messages = append(m.messages, fmt.Sprintf("Error: %v", err))
				} else {
					// Add user message to display
					m.messages = append(m.messages, "")
					m.messages = append(m.messages, "You: "+value)
					m.thinking = true
					m.thinkingStart = time.Now()

					// Update viewport
					m.viewport.SetContent(m.renderMessages())
					m.viewport.GotoBottom()

					// Clear textarea
					m.textarea.Reset()
					m.textarea.Focus()
				}
			}
			return m, nil
		}

	case cliEventMsg:
		// Handle events from CLI
		m = m.handleCLIEvent(bridge.Response(msg))
		return m, listenForEvents(m.cli)

	case cliErrorMsg:
		m.err = error(msg)
		m.messages = append(m.messages, fmt.Sprintf("CLI Error: %v", msg))
		m.viewport.SetContent(m.renderMessages())
		return m, listenForEvents(m.cli)

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.header.Width = msg.Width
		m.footer.Width = msg.Width

		if !m.ready {
			// Initialize viewport
			headerHeight := 1
			footerHeight := 1
			inputHeight := 4
			verticalMarginHeight := headerHeight + footerHeight + inputHeight + 4

			m.viewport = viewport.New(msg.Width-4, msg.Height-verticalMarginHeight)
			m.viewport.YPosition = headerHeight + 2
			m.viewport.SetContent(m.renderMessages())
			m.ready = true

			// Update textarea width
			m.textarea.SetWidth(msg.Width - 6)
		} else {
			// Update viewport size
			headerHeight := 1
			footerHeight := 1
			inputHeight := 4
			verticalMarginHeight := headerHeight + footerHeight + inputHeight + 4

			m.viewport.Width = msg.Width - 4
			m.viewport.Height = msg.Height - verticalMarginHeight
			m.viewport.SetContent(m.renderMessages())

			// Update textarea width
			m.textarea.SetWidth(msg.Width - 6)
		}
	}

	return m, tea.Batch(cmds...)
}

func (m model) handleCLIEvent(event bridge.Response) model {
	switch event.Type {
	case bridge.ResponseTypeMessage: // Streaming chunks
		if event.Content != "" {
			// Start new response if not already accumulating
			if !m.respondingNow {
				m.respondingNow = true
				m.currentResponse.Reset()
				m.messages = append(m.messages, "")
				m.messages = append(m.messages, "Danni: ")
			}
			// Accumulate the chunk
			m.currentResponse.WriteString(event.Content)
			// Update the last message with accumulated response
			if len(m.messages) > 0 {
				m.messages[len(m.messages)-1] = "Danni: " + m.currentResponse.String()
			}
			m.viewport.SetContent(m.renderMessages())
			m.viewport.GotoBottom()
		}

	case bridge.ResponseTypeMessageComplete:
		// Final message - only add if we weren't streaming
		if event.Content != "" && !m.respondingNow {
			m.messages = append(m.messages, "")
			m.messages = append(m.messages, "Danni: "+event.Content)
			m.viewport.SetContent(m.renderMessages())
			m.viewport.GotoBottom()
		}
		// Reset streaming state
		m.respondingNow = false
		m.currentResponse.Reset()

	case bridge.ResponseTypeThinking:
		m.thinking = true
		if m.thinkingStart.IsZero() {
			m.thinkingStart = time.Now()
		}

	case bridge.ResponseTypeComplete:
		m.thinking = false
		m.respondingNow = false
		m.currentResponse.Reset()

	case bridge.ResponseTypeError:
		if event.Error != "" {
			m.messages = append(m.messages, "")
			m.messages = append(m.messages, "Error: "+event.Error)
			m.viewport.SetContent(m.renderMessages())
		}
		m.thinking = false

	case bridge.ResponseTypeToolUse:
		if toolName, ok := event.Data["tool_name"].(string); ok {
			m.messages = append(m.messages, "    Using: "+toolName)
			m.viewport.SetContent(m.renderMessages())
		}

	case bridge.ResponseTypeToken:
		if tokensUsed, ok := event.Data["tokens_used"].(float64); ok {
			m.tokensUsed = int(tokensUsed)
		}
		if tokensTotal, ok := event.Data["tokens_total"].(float64); ok {
			m.tokensTotal = int(tokensTotal)
		}
		if cost, ok := event.Data["estimated_cost"].(float64); ok {
			m.estimatedCost = cost
		}
		m.footer.TokensUsed = m.tokensUsed
		m.footer.TokensTotal = m.tokensTotal
		m.footer.EstimatedCost = m.estimatedCost

	case "session_info":
		if sessionID, ok := event.Data["session_id"].(string); ok && sessionID != "" {
			m.sessionID = sessionID
			m.header.SessionID = sessionID
		} else {
			m.sessionID = "active"
			m.header.SessionID = "active"
		}
		// Clear connecting message, set minimal welcome
		m.messages = []string{}
		m.viewport.SetContent(m.renderMessages())

	case "ready":
		m.thinking = false
		if m.sessionID == "connecting..." {
			m.sessionID = "active"
			m.header.SessionID = "active"
		}

	case "status":
		if msg, ok := event.Data["message"].(string); ok {
			m.messages = append(m.messages, "  "+msg)
			m.viewport.SetContent(m.renderMessages())
		}
	}

	return m
}

func (m model) renderMessages() string {
	var rendered []string

	// Calculate wrap width (viewport width minus padding)
	wrapWidth := m.width - 8
	if wrapWidth < 40 {
		wrapWidth = 72
	}

	// Content style with word wrap
	contentStyle := lipgloss.NewStyle().Width(wrapWidth)

	for _, msg := range m.messages {
		if msg == "" {
			rendered = append(rendered, "")
			continue
		}

		if strings.HasPrefix(msg, "You:") {
			// User messages - bold label, wrapped content
			label := lipgloss.NewStyle().Bold(true).Render("You:")
			content := strings.TrimPrefix(msg, "You:")
			wrappedContent := contentStyle.Render(strings.TrimSpace(content))
			rendered = append(rendered, "  "+label+" "+wrappedContent)

		} else if strings.HasPrefix(msg, "Danni:") {
			// Danni messages - pink label, wrapped content
			label := lipgloss.NewStyle().Foreground(styles.HotPink).Bold(true).Render("Danni:")
			content := strings.TrimPrefix(msg, "Danni:")
			wrappedContent := contentStyle.Render(strings.TrimSpace(content))
			rendered = append(rendered, "  "+label+" "+wrappedContent)

		} else if strings.HasPrefix(msg, "Error:") {
			rendered = append(rendered, "  "+lipgloss.NewStyle().Foreground(styles.HotPink).Width(wrapWidth).Render(msg))

		} else if strings.HasPrefix(msg, "    Using:") {
			// Tool usage - gold, indented
			rendered = append(rendered, lipgloss.NewStyle().Foreground(styles.Gold).Italic(true).Render(msg))

		} else {
			rendered = append(rendered, "  "+contentStyle.Render(msg))
		}
	}
	return strings.Join(rendered, "\n")
}

func (m model) View() string {
	if !m.ready {
		// Minimal loading state
		symbol := lipgloss.NewStyle().Foreground(styles.HotPink).Render("✦")
		message := lipgloss.NewStyle().Foreground(styles.GrayMed).Render("Finding the thread...")
		return "\n\n  " + symbol + " " + message
	}

	// Header
	headerView := m.header.View()

	// Chat area
	chatView := m.viewport.View()

	// Input area - simple prompt with underline
	var inputView string
	if m.thinking {
		// Calculate elapsed time and pick a rotating phrase
		elapsed := int(time.Since(m.thinkingStart).Seconds())
		phraseIndex := (elapsed / 5) % len(thinkingPhrases) // Rotate every 5 seconds
		phrase := thinkingPhrases[phraseIndex]

		// Format: ✦ Finding the thread... (12s)
		timeStr := lipgloss.NewStyle().Foreground(styles.GrayMed).Render(fmt.Sprintf("(%ds)", elapsed))
		phraseStr := lipgloss.NewStyle().Foreground(styles.Gold).Render(phrase)
		inputView = "  " + m.spinner.View() + " " + phraseStr + " " + timeStr
	} else {
		underline := lipgloss.NewStyle().Foreground(styles.Gold).Render(strings.Repeat("─", m.width-4))
		inputView = "  > " + m.textarea.View() + "\n  " + underline
	}

	// Clean layout - no footer
	return headerView + "\n" + chatView + "\n" + inputView
}

func main() {
	m, err := initialModel()
	if err != nil {
		fmt.Printf("Error initializing: %v\n", err)
		os.Exit(1)
	}

	p := tea.NewProgram(
		m,
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
