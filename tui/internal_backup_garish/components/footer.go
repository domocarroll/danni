package components

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/subfracture/danni/tui/internal/styles"
)

type Footer struct {
	Width         int
	TokensUsed    int
	TokensTotal   int
	EstimatedCost float64
}

func (f Footer) View() string {
	// Left - Token usage with hot pink accent for current usage
	tokenDisplay := fmt.Sprintf("Tokens: %s%d%s/%d",
		lipgloss.NewStyle().Foreground(styles.HotPink).Render(""),
		f.TokensUsed,
		lipgloss.NewStyle().Foreground(styles.GrayMed).Render(""),
		f.TokensTotal)

	left := lipgloss.NewStyle().
		Foreground(styles.GrayMed).
		Render(tokenDisplay)

	// Center - Cost with gold accent for value
	costText := fmt.Sprintf("$%.4f", f.EstimatedCost)
	center := lipgloss.NewStyle().
		Foreground(styles.GrayMed).
		Render(fmt.Sprintf("Cost: %s",
			lipgloss.NewStyle().Foreground(styles.Gold).Render(costText)))

	// Right - Help text with subtle presence
	right := lipgloss.NewStyle().
		Foreground(styles.GrayMed).
		Render("^C quit • ? help • Tab modules")

	// Calculate spacing
	leftWidth := lipgloss.Width(left)
	centerWidth := lipgloss.Width(center)
	rightWidth := lipgloss.Width(right)

	totalContentWidth := leftWidth + centerWidth + rightWidth
	if totalContentWidth >= f.Width {
		return styles.FooterStyle.Width(f.Width).Render(left + " " + center + " " + right)
	}

	// Distribute space evenly
	availableSpace := f.Width - totalContentWidth
	leftPad := availableSpace / 2
	rightPad := availableSpace - leftPad

	line := left + lipgloss.NewStyle().Width(leftPad).Render("") +
		center +
		lipgloss.NewStyle().Width(rightPad).Render("") + right

	return styles.FooterStyle.Width(f.Width).Render(line)
}
