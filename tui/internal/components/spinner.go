package components

import (
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/lipgloss"
	"github.com/subfracture/danni/tui/internal/styles"
)

// Custom spinner definitions with personality
var (
	// Constellation - elegant, thoughtful, suggests depth and pattern recognition
	ConstellationSpinner = spinner.Spinner{
		Frames: []string{"✦  ", " ✦ ", "  ✦", " ✦ "},
		FPS:    time.Second / 4,
	}

	// Pulse - organic breathing feel, suggests life and presence
	PulseSpinner = spinner.Spinner{
		Frames: []string{"○", "◔", "◑", "◕", "●", "◕", "◑", "◔"},
		FPS:    time.Second / 8,
	}
)

type ThinkingSpinner struct {
	spinner spinner.Model
}

func NewThinkingSpinner() ThinkingSpinner {
	s := spinner.New()
	s.Spinner = ConstellationSpinner
	// Gold spinner to suggest valuable processing
	s.Style = lipgloss.NewStyle().Foreground(styles.Gold)
	return ThinkingSpinner{
		spinner: s,
	}
}

func (s ThinkingSpinner) View() string {
	// More elegant, Danni-appropriate thinking message
	return styles.ThinkingStyle.Render(s.spinner.View() + " Danni is thinking...")
}
