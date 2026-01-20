package components

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/lipgloss"
	"github.com/subfracture/danni/tui/internal/styles"
)

type ThinkingSpinner struct {
	spinner spinner.Model
}

func NewThinkingSpinner() ThinkingSpinner {
	s := spinner.New()
	s.Spinner = spinner.Dot
	// Gold spinner to suggest valuable processing
	s.Style = lipgloss.NewStyle().Foreground(styles.Gold)
	return ThinkingSpinner{
		spinner: s,
	}
}

func (s ThinkingSpinner) View() string {
	// More elegant, Danni-appropriate thinking message
	return styles.ThinkingStyle.Render(s.spinner.View() + " Danni is thinking deeply...")
}
