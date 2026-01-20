package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/subfracture/danni/tui/internal/styles"
)

type Header struct {
	Width     int
	Module    string
	SessionID string
}

func (h Header) View() string {
	// Ensure minimum width
	width := h.Width
	if width < 40 {
		width = 80 // sensible default
	}

	// Left side - "DANNI" bold + module path muted
	brand := lipgloss.NewStyle().
		Foreground(styles.White).
		Bold(true).
		Render("DANNI")

	module := lipgloss.NewStyle().
		Foreground(styles.GrayMed).
		Render(h.Module)

	left := brand + " " + module

	// Right side - Gold dot + status
	dot := lipgloss.NewStyle().
		Foreground(styles.Gold).
		Render("●")

	status := lipgloss.NewStyle().
		Foreground(styles.GrayMed).
		Render(h.SessionID)

	right := dot + " " + status

	// Calculate spacing
	leftLen := lipgloss.Width(left)
	rightLen := lipgloss.Width(right)
	spacerLen := width - leftLen - rightLen - 4 // account for padding

	if spacerLen < 1 {
		spacerLen = 1
	}

	spacer := strings.Repeat(" ", spacerLen)

	// Main line
	mainLine := left + spacer + right

	// Separator line - subtle but visible
	separator := lipgloss.NewStyle().
		Foreground(styles.GrayMed).
		Render(strings.Repeat("─", width-4))

	// Combine with padding
	return lipgloss.NewStyle().
		Padding(0, 2).
		Render(mainLine + "\n" + separator)
}
