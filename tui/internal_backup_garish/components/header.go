package components

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/subfracture/danni/tui/internal/styles"
)

type Header struct {
	Width     int
	Module    string
	SessionID string
}

func (h Header) View() string {
	// Left side - Brand identity with sparkle
	left := lipgloss.NewStyle().
		Foreground(styles.White).
		Bold(true).
		Render(fmt.Sprintf("%s DANNI v0.1.0", styles.SymbolSparkle))

	// Center - Module name in hot pink for prominence
	center := lipgloss.NewStyle().
		Foreground(styles.HotPink).
		Bold(true).
		Render(h.Module)

	// Right side - Session indicator with gold accent
	sessionDisplay := h.SessionID
	if len(sessionDisplay) > 8 {
		sessionDisplay = sessionDisplay[:8]
	}
	right := lipgloss.NewStyle().
		Foreground(styles.GrayMed).
		Render(fmt.Sprintf("%s %s", styles.SymbolFilledCircle, sessionDisplay))

	// Calculate spacing
	leftWidth := lipgloss.Width(left)
	centerWidth := lipgloss.Width(center)
	rightWidth := lipgloss.Width(right)

	// Center the module name
	totalContentWidth := leftWidth + centerWidth + rightWidth
	if totalContentWidth >= h.Width {
		// If content is too wide, just concatenate
		return styles.HeaderStyle.Width(h.Width).Render(left + " " + center + " " + right)
	}

	// Calculate padding to center the module name
	availableSpace := h.Width - totalContentWidth
	leftPad := availableSpace / 2
	rightPad := availableSpace - leftPad

	line := left + lipgloss.NewStyle().Width(leftPad).Render("") +
		center +
		lipgloss.NewStyle().Width(rightPad).Render("") + right

	return styles.HeaderStyle.Width(h.Width).Render(line)
}
