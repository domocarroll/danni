package effects

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/subfracture/danni/tui/internal/styles"
)

// ═══════════════════════════════════════════════════════════════
// ASCII ART GENERATOR - Epic visual headers and decorations
// ═══════════════════════════════════════════════════════════════

// DanniLogo returns the DANNI ASCII art logo with gradient
func DanniLogo() string {
	logo := []string{
		"╔═══════════════════════════════════════════════════════════╗",
		"║  ██████╗  █████╗ ███╗   ██╗███╗   ██╗██╗                ║",
		"║  ██╔══██╗██╔══██╗████╗  ██║████╗  ██║██║                ║",
		"║  ██║  ██║███████║██╔██╗ ██║██╔██╗ ██║██║                ║",
		"║  ██║  ██║██╔══██║██║╚██╗██║██║╚██╗██║██║                ║",
		"║  ██████╔╝██║  ██║██║ ╚████║██║ ╚████║██║                ║",
		"║  ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═══╝╚═╝  ╚═══╝╚═╝                ║",
		"║                                                           ║",
		"║  DEDICATED AUTONOMOUS NEURAL NETWORKED INTELLIGENCE      ║",
		"╚═══════════════════════════════════════════════════════════╝",
	}

	// Apply gradient to each line
	result := make([]string, len(logo))
	for i, line := range logo {
		if i == 0 || i == len(logo)-1 {
			// Borders in gold
			result[i] = lipgloss.NewStyle().Foreground(styles.Gold).Render(line)
		} else if i >= 1 && i <= 6 {
			// Logo text with gradient
			result[i] = styles.DanniTriGradient(line)
		} else {
			// Subtitle in white
			result[i] = lipgloss.NewStyle().Foreground(styles.White).Render(line)
		}
	}

	return strings.Join(result, "\n")
}

// ModuleHeader creates an ASCII header for modules
func ModuleHeader(module string) string {
	width := 60
	border := strings.Repeat("═", width-2)

	var icon string
	var title string
	var color lipgloss.Color

	switch module {
	case "/strategy":
		icon = "⚔"
		title = "STRATEGIC ANALYSIS"
		color = styles.HotPink
	case "/creative":
		icon = "✨"
		title = "CREATIVE IDEATION"
		color = styles.Gold
	case "/design":
		icon = "🎨"
		title = "DESIGN SYSTEMS"
		color = styles.HotPink
	case "/technology":
		icon = "⚡"
		title = "TECHNICAL SOLUTIONS"
		color = styles.Gold
	case "/gravity":
		icon = "🌌"
		title = "DATA PATTERNS"
		color = styles.HotPink
	case "/validate":
		icon = "✓"
		title = "BRAND VALIDATION"
		color = styles.Gold
	case "/synthesize":
		icon = "⬡"
		title = "BREAKTHROUGH SYNTHESIS"
		color = styles.HotPink
	case "/recall":
		icon = "📚"
		title = "INSTITUTIONAL MEMORY"
		color = styles.Gold
	case "/upload":
		icon = "⬆"
		title = "ASSET ARCHAEOLOGY"
		color = styles.HotPink
	default:
		icon = "◆"
		title = "DANNI SYSTEM"
		color = styles.Gold
	}

	// Build the header
	header := []string{
		fmt.Sprintf("╔%s╗", border),
		fmt.Sprintf("║ %s  %s", icon, centerText(title, width-8)),
		fmt.Sprintf("╚%s╝", border),
	}

	// Apply styling
	styledHeader := make([]string, len(header))
	for i, line := range header {
		styledHeader[i] = lipgloss.NewStyle().Foreground(color).Render(line)
	}

	return strings.Join(styledHeader, "\n")
}

// LoadingAnimation creates fancy loading animations
type LoadingStyle int

const (
	LoadingPulse LoadingStyle = iota
	LoadingWave
	LoadingMatrix
	LoadingGlitch
	LoadingHologram
)

func LoadingBar(progress float64, width int, style LoadingStyle) string {
	filled := int(float64(width) * progress)
	empty := width - filled

	var bar string
	var filledChar, emptyChar string
	var color lipgloss.Color

	switch style {
	case LoadingPulse:
		filledChar = "█"
		emptyChar = "░"
		color = styles.HotPink

	case LoadingWave:
		chars := []string{"▁", "▂", "▃", "▄", "▅", "▆", "▇", "█"}
		bar = ""
		for i := 0; i < width; i++ {
			if i < filled {
				idx := i % len(chars)
				bar += chars[idx]
			} else {
				bar += "░"
			}
		}
		return styles.DanniGradient(bar)

	case LoadingMatrix:
		filledChar = "0"
		emptyChar = "1"
		color = lipgloss.Color("#00FF00")

	case LoadingGlitch:
		glitchChars := []string{"░", "▒", "▓", "█", "▓", "▒"}
		bar = ""
		for i := 0; i < width; i++ {
			if i < filled {
				bar += glitchChars[i%len(glitchChars)]
			} else {
				bar += "░"
			}
		}
		return styles.ElectricGradient(bar)

	case LoadingHologram:
		filledChar = "▓"
		emptyChar = "░"
		// Create holographic effect with alternating colors
		bar = ""
		for i := 0; i < filled; i++ {
			if i%2 == 0 {
				bar += lipgloss.NewStyle().Foreground(styles.HotPink).Render(filledChar)
			} else {
				bar += lipgloss.NewStyle().Foreground(styles.Gold).Render(filledChar)
			}
		}
		for i := 0; i < empty; i++ {
			bar += emptyChar
		}
		return bar
	}

	if bar == "" {
		bar = strings.Repeat(filledChar, filled) + strings.Repeat(emptyChar, empty)
		return lipgloss.NewStyle().Foreground(color).Render(bar)
	}

	return bar
}

// ThinkingAnimation creates various thinking indicators
func ThinkingAnimation(frame int, style string) string {
	switch style {
	case "dots":
		dots := []string{
			"   ",
			".  ",
			".. ",
			"...",
			" ..",
			"  .",
		}
		return dots[frame%len(dots)]

	case "pulse":
		pulse := []string{
			"◯",
			"◉",
			"●",
			"◉",
		}
		return pulse[frame%len(pulse)]

	case "wave":
		wave := []string{
			"≈≈≈",
			"≋≈≈",
			"≈≋≈",
			"≈≈≋",
		}
		return wave[frame%len(wave)]

	case "constellation":
		stars := []string{
			"✦   ",
			" ✦  ",
			"  ✦ ",
			"   ✦",
			"  ✦ ",
			" ✦  ",
		}
		return stars[frame%len(stars)]

	case "neural":
		neural := []string{
			"⟨◯⟩",
			"⟨◉⟩",
			"⟨●⟩",
			"⟨◉⟩",
		}
		return neural[frame%len(neural)]

	default:
		return "..."
	}
}

// SuccessAnimation creates celebration effects
func SuccessAnimation() string {
	celebration := []string{
		"",
		"       ✨  ✨  ✨       ",
		"    🎉  SUCCESS!  🎉    ",
		"       ✨  ✨  ✨       ",
		"",
	}

	result := make([]string, len(celebration))
	for i, line := range celebration {
		if i == 2 {
			result[i] = styles.DanniTriGradient(line)
		} else {
			result[i] = styles.LuxuryGradient(line)
		}
	}

	return strings.Join(result, "\n")
}

// ErrorAnimation creates error indicators
func ErrorAnimation() string {
	error := []string{
		"╭─────────────────────╮",
		"│  ⚠  ATTENTION  ⚠   │",
		"│ Something needs     │",
		"│ your attention      │",
		"╰─────────────────────╯",
	}

	result := make([]string, len(error))
	for i, line := range error {
		result[i] = styles.ElectricGradient(line)
	}

	return strings.Join(result, "\n")
}

// DecorativeBorder creates fancy borders
func DecorativeBorder(width int, style string) (top, bottom string) {
	switch style {
	case "double":
		top = "╔" + strings.Repeat("═", width-2) + "╗"
		bottom = "╚" + strings.Repeat("═", width-2) + "╝"

	case "rounded":
		top = "╭" + strings.Repeat("─", width-2) + "╮"
		bottom = "╰" + strings.Repeat("─", width-2) + "╯"

	case "thick":
		top = "┏" + strings.Repeat("━", width-2) + "┓"
		bottom = "┗" + strings.Repeat("━", width-2) + "┛"

	case "dotted":
		top = "┌" + strings.Repeat("┄", width-2) + "┐"
		bottom = "└" + strings.Repeat("┄", width-2) + "┘"

	case "wave":
		wave := ""
		for i := 0; i < width-2; i++ {
			if i%2 == 0 {
				wave += "∿"
			} else {
				wave += "∾"
			}
		}
		top = "◈" + wave + "◈"
		bottom = top

	case "sparkle":
		sparkle := ""
		chars := []string{"✦", "✧", "⋆", "✦"}
		for i := 0; i < width-2; i++ {
			sparkle += chars[i%len(chars)]
		}
		top = "✨" + sparkle + "✨"
		bottom = top

	default:
		top = "┌" + strings.Repeat("─", width-2) + "┐"
		bottom = "└" + strings.Repeat("─", width-2) + "┘"
	}

	return styles.DanniGradient(top), styles.DanniReverseGradient(bottom)
}

// MatrixRainColumn creates a single column of Matrix rain
func MatrixRainColumn(height int, offset int) string {
	chars := "ﾊﾐﾋｰｳｼﾅﾓﾆｻﾜﾂｵﾘｱﾎﾃﾏｹﾒｴｶｷﾑﾕﾗｾﾈｽﾀﾇ01"
	column := make([]rune, height)

	for i := 0; i < height; i++ {
		if (i+offset)%10 < 7 {
			column[i] = []rune(chars)[(i+offset)%len(chars)]
		} else {
			column[i] = ' '
		}
	}

	return string(column)
}

// Utility function to center text
func centerText(text string, width int) string {
	padding := (width - len(text)) / 2
	if padding < 0 {
		return text[:width]
	}
	return strings.Repeat(" ", padding) + text + strings.Repeat(" ", width-len(text)-padding)
}