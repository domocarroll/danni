package components

import (
	"fmt"
	"math"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/subfracture/danni/tui/internal/styles"
)

// ═══════════════════════════════════════════════════════════════
// ANIMATED SPINNER - Sophisticated loading with gradient effects
// ═══════════════════════════════════════════════════════════════

// AnimatedSpinner represents an enhanced spinner with gradient and pulse effects
type AnimatedSpinner struct {
	spinner     spinner.Model
	pulsePhase  float64
	gradientPos float64
	message     string
	style       SpinnerStyle
	startTime   time.Time
}

// SpinnerStyle defines the visual style of the spinner
type SpinnerStyle int

const (
	SpinnerDanniPulse SpinnerStyle = iota // Hot Pink/Gold pulse
	SpinnerGradientWave                    // Gradient wave effect
	SpinnerElectric                        // Electric neon effect
	SpinnerLuxury                          // Luxury gold shimmer
	SpinnerMatrix                          // Matrix-style cascade
)

// Custom spinner shapes for different moods
var (
	// Elegant dots with varying sizes
	DanniDots = spinner.Spinner{
		Frames: []string{
			"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏",
		},
		FPS: time.Second / 10,
	}

	// Constellation pattern for "thinking deeply"
	ConstellationSpinner = spinner.Spinner{
		Frames: []string{
			"✦     ", " ✦    ", "  ✦   ", "   ✦  ", "    ✦ ", "     ✦",
			"    ✦ ", "   ✦  ", "  ✦   ", " ✦    ", "✦     ",
		},
		FPS: time.Second / 12,
	}

	// Pulse effect using block characters
	PulseSpinner = spinner.Spinner{
		Frames: []string{
			"█▁▁▁▁", "▁█▁▁▁", "▁▁█▁▁", "▁▁▁█▁", "▁▁▁▁█",
			"▁▁▁█▁", "▁▁█▁▁", "▁█▁▁▁", "█▁▁▁▁",
		},
		FPS: time.Second / 15,
	}

	// Diamond morph for luxury feel
	DiamondSpinner = spinner.Spinner{
		Frames: []string{
			"◇", "◈", "◆", "◈", "◇", "◈", "◆", "◈",
		},
		FPS: time.Second / 8,
	}

	// Wave pattern
	WaveSpinner = spinner.Spinner{
		Frames: []string{
			"≈≈≈", "≋≈≈", "≈≋≈", "≈≈≋", "≈≈≈",
		},
		FPS: time.Second / 6,
	}
)

// NewAnimatedSpinner creates a new animated spinner
func NewAnimatedSpinner(style SpinnerStyle, message string) AnimatedSpinner {
	s := spinner.New()

	// Set spinner based on style
	switch style {
	case SpinnerDanniPulse:
		s.Spinner = PulseSpinner
	case SpinnerGradientWave:
		s.Spinner = WaveSpinner
	case SpinnerElectric:
		s.Spinner = DanniDots
	case SpinnerLuxury:
		s.Spinner = DiamondSpinner
	case SpinnerMatrix:
		s.Spinner = ConstellationSpinner
	default:
		s.Spinner = DanniDots
	}

	return AnimatedSpinner{
		spinner:   s,
		message:   message,
		style:     style,
		startTime: time.Now(),
	}
}

// Init initializes the spinner
func (a AnimatedSpinner) Init() tea.Cmd {
	return tea.Batch(
		a.spinner.Tick,
		tickAnimation(),
	)
}

// Update handles spinner updates and animations
func (a AnimatedSpinner) Update(msg tea.Msg) (AnimatedSpinner, tea.Cmd) {
	switch msg := msg.(type) {
	case spinner.TickMsg:
		var cmd tea.Cmd
		a.spinner, cmd = a.spinner.Update(msg)
		return a, cmd

	case animationTickMsg:
		// Update animation phases
		a.pulsePhase += 0.1
		a.gradientPos += 0.05

		return a, tickAnimation()

	default:
		var cmd tea.Cmd
		a.spinner, cmd = a.spinner.Update(msg)
		return a, cmd
	}
}

// View renders the animated spinner
func (a AnimatedSpinner) View() string {
	elapsed := time.Since(a.startTime).Seconds()

	switch a.style {
	case SpinnerDanniPulse:
		return a.renderPulseSpinner(elapsed)
	case SpinnerGradientWave:
		return a.renderGradientWaveSpinner(elapsed)
	case SpinnerElectric:
		return a.renderElectricSpinner(elapsed)
	case SpinnerLuxury:
		return a.renderLuxurySpinner(elapsed)
	case SpinnerMatrix:
		return a.renderMatrixSpinner(elapsed)
	default:
		return a.renderPulseSpinner(elapsed)
	}
}

// renderPulseSpinner renders a pulsing spinner with color transitions
func (a AnimatedSpinner) renderPulseSpinner(elapsed float64) string {
	// Calculate pulse intensity (sine wave)
	intensity := (math.Sin(elapsed*2) + 1) / 2

	// Interpolate between hot pink and gold based on pulse
	var spinnerColor lipgloss.Color
	if intensity > 0.5 {
		spinnerColor = styles.HotPink
	} else {
		spinnerColor = styles.Gold
	}

	// Apply color and render
	styledSpinner := lipgloss.NewStyle().
		Foreground(spinnerColor).
		Render(a.spinner.View())

	// Add gradient message
	gradientMsg := styles.DanniWaveGradient(a.message)

	return fmt.Sprintf("%s %s", styledSpinner, gradientMsg)
}

// renderGradientWaveSpinner renders a wave-effect spinner
func (a AnimatedSpinner) renderGradientWaveSpinner(elapsed float64) string {
	// Apply wave gradient to spinner
	spinnerText := a.spinner.View()
	gradientSpinner := styles.WaveGradientText(spinnerText, "#FF1493", "#FFD700", elapsed)

	// Animate message with reverse wave
	animatedMsg := styles.WaveGradientText(a.message, "#FFD700", "#FF1493", elapsed+math.Pi)

	return fmt.Sprintf("%s %s", gradientSpinner, animatedMsg)
}

// renderElectricSpinner renders an electric neon effect
func (a AnimatedSpinner) renderElectricSpinner(elapsed float64) string {
	spinnerText := a.spinner.View()

	// Create electric effect with shimmer
	shimmerPositions := []int{}
	if int(elapsed*10)%3 == 0 {
		shimmerPositions = append(shimmerPositions, 0)
	}

	electricSpinner := styles.ShimmerEffect(
		spinnerText,
		"#FF1493",
		"#FFFFFF",
		shimmerPositions,
	)

	electricMsg := styles.ElectricGradient(a.message)

	return fmt.Sprintf("%s %s", electricSpinner, electricMsg)
}

// renderLuxurySpinner renders a luxury gold shimmer effect
func (a AnimatedSpinner) renderLuxurySpinner(elapsed float64) string {
	spinnerText := a.spinner.View()

	// Gold with shimmer
	shimmerPos := int(elapsed*5) % len(spinnerText)
	luxurySpinner := styles.ShimmerEffect(
		spinnerText,
		"#FFD700",
		"#FFFFFF",
		[]int{shimmerPos},
	)

	luxuryMsg := styles.LuxuryGradient(a.message)

	return fmt.Sprintf("%s %s", luxurySpinner, luxuryMsg)
}

// renderMatrixSpinner renders a Matrix-style cascading effect
func (a AnimatedSpinner) renderMatrixSpinner(elapsed float64) string {
	spinnerText := a.spinner.View()

	// Create cascading green effect
	matrixSpinner := styles.GradientText(spinnerText, "#00FF00", "#003300")

	// Add glitch effect to message occasionally
	msg := a.message
	if int(elapsed*4)%7 == 0 {
		// Brief glitch
		msg = styles.GradientText(msg, "#FF0000", "#00FF00")
	} else {
		msg = styles.DanniGradient(msg)
	}

	return fmt.Sprintf("%s %s", matrixSpinner, msg)
}

// SetMessage updates the spinner message
func (a *AnimatedSpinner) SetMessage(message string) {
	a.message = message
}

// SetStyle changes the spinner style
func (a *AnimatedSpinner) SetStyle(style SpinnerStyle) {
	a.style = style

	// Update spinner type based on new style
	switch style {
	case SpinnerDanniPulse:
		a.spinner.Spinner = PulseSpinner
	case SpinnerGradientWave:
		a.spinner.Spinner = WaveSpinner
	case SpinnerElectric:
		a.spinner.Spinner = DanniDots
	case SpinnerLuxury:
		a.spinner.Spinner = DiamondSpinner
	case SpinnerMatrix:
		a.spinner.Spinner = ConstellationSpinner
	}
}

// Animation tick message
type animationTickMsg time.Time

// tickAnimation creates a command for animation updates
func tickAnimation() tea.Cmd {
	return tea.Tick(time.Millisecond*50, func(t time.Time) tea.Msg {
		return animationTickMsg(t)
	})
}