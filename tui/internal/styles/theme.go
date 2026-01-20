package styles

import "github.com/charmbracelet/lipgloss"

// ═══════════════════════════════════════════════════════════════
// DANNI COLOR PALETTE - Hot Pink + Gold + Black
// Sophisticated. Bold. Beautiful.
// ═══════════════════════════════════════════════════════════════

// Primary Colors
var (
	// Hot Pink - Primary accent, actions, key data, attention points
	HotPink       = lipgloss.Color("#FF1493") // DeepPink - primary
	HotPinkAlt    = lipgloss.Color("#FF69B4") // HotPink - softer variant
	HotPinkBright = lipgloss.Color("#FF007F") // Rose - maximum intensity

	// Gold - Success, value, secondary accent, warmth
	Gold       = lipgloss.Color("#FFD700") // Gold - primary success
	GoldAlt    = lipgloss.Color("#FFA500") // Orange - warmer variant
	GoldMuted  = lipgloss.Color("#DAA520") // GoldenRod - subtle elegance
	GoldPale   = lipgloss.Color("#F4E4C1") // Pale gold - whisper of luxury

	// Black - Foundation, sophistication, depth
	Black     = lipgloss.Color("#000000") // True black - pure depth
	BlackSoft = lipgloss.Color("#0A0A0A") // Almost black - breathing room
	BlackRich = lipgloss.Color("#1A1A1A") // Charcoal - rich background

	// White - Clarity, contrast, clean space
	White     = lipgloss.Color("#FFFFFF") // Pure white - maximum clarity
	WhiteSoft = lipgloss.Color("#F5F5F5") // Off-white - gentle
	WhiteWarm = lipgloss.Color("#FFF8F0") // Cream - subtle warmth
)

// Supporting Colors - Hierarchy & Structure
var (
	GrayLight  = lipgloss.Color("#CCCCCC") // Silver - high visibility
	GrayMed    = lipgloss.Color("#808080") // Gray - balanced
	GrayDark   = lipgloss.Color("#404040") // Charcoal gray - subtle structure
	GraySubtle = lipgloss.Color("#2A2A2A") // Near-black - minimal contrast
)

// Semantic Colors
var (
	ColorSuccess = Gold           // Achievement, positive outcomes
	ColorWarning = GoldAlt         // Caution, needs attention
	ColorError   = HotPink         // Attention points (not failures)
	ColorInfo    = HotPinkAlt      // Informational highlights
	ColorActive  = HotPink         // Active states, online status
	ColorMuted   = GrayMed         // Disabled, inactive states
	Accent       = Gold            // Primary accent color (warm gold)
)

// Legacy aliases for smooth transition
var (
	PrimaryPurple = HotPink       // Maps to hot pink
	DeepPurple    = HotPinkBright // Maps to bright pink
	MidnightBlue  = BlackRich     // Maps to rich black
	RoseGold      = Gold          // Maps to gold
	SoftCream     = White         // Maps to white
	DeepCharcoal  = BlackSoft     // Maps to soft black
	SubtleGray    = GrayMed       // Maps to medium gray
)

// ═══════════════════════════════════════════════════════════════
// COMPONENT STYLES - Luxe Minimalism
// ═══════════════════════════════════════════════════════════════

var (
	// Header - Bold presence with hot pink accent
	HeaderStyle = lipgloss.NewStyle().
		Foreground(White).
		Background(Black).
		BorderStyle(lipgloss.NormalBorder()).
		BorderBottom(true).
		BorderForeground(HotPink).
		Bold(true).
		Padding(0, 2)

	// Chat viewport - Clean black canvas for content
	ChatStyle = lipgloss.NewStyle().
		Foreground(White).
		Background(Black).
		Padding(1, 2)

	// Input area - Subtle depth with gold accent when active
	InputStyle = lipgloss.NewStyle().
		Foreground(White).
		Background(BlackSoft).
		BorderStyle(lipgloss.NormalBorder()).
		BorderTop(true).
		BorderForeground(Gold).
		Padding(1, 2)

	// Footer - Minimal presence with key information
	FooterStyle = lipgloss.NewStyle().
		Foreground(GrayMed).
		Background(BlackSoft).
		Padding(0, 2)

	// User messages - Clean white text, bold for visibility
	UserMessageStyle = lipgloss.NewStyle().
		Foreground(White).
		Bold(true)

	// User label style - Bold white for "You:"
	UserLabelStyle = lipgloss.NewStyle().
		Foreground(White).
		Bold(true)

	// Assistant messages - Hot pink to establish Danni's presence
	AssistantMessageStyle = lipgloss.NewStyle().
		Foreground(HotPink).
		Bold(true)

	// Assistant label style - Bold hot pink for "Danni:"
	AssistantLabelStyle = lipgloss.NewStyle().
		Foreground(HotPink).
		Bold(true)

	// Tool usage style - Gold italic for tool messages
	ToolUsageStyle = lipgloss.NewStyle().
		Foreground(Gold).
		Italic(true)

	// Thinking indicator - Gold to suggest valuable processing
	ThinkingStyle = lipgloss.NewStyle().
		Foreground(Gold).
		Italic(true)

	// Error states - Hot pink for attention (not negative)
	ErrorStyle = lipgloss.NewStyle().
		Foreground(HotPink).
		Bold(true)

	// Success states - Gold for achievement
	SuccessStyle = lipgloss.NewStyle().
		Foreground(Gold).
		Bold(true)

	// Warning states - Warm gold/orange
	WarningStyle = lipgloss.NewStyle().
		Foreground(GoldAlt).
		Bold(true)

	// Info states - Softer pink
	InfoStyle = lipgloss.NewStyle().
		Foreground(HotPinkAlt)

	// Muted/secondary text
	MutedStyle = lipgloss.NewStyle().
		Foreground(GrayMed)

	// Highlight style for important information
	HighlightStyle = lipgloss.NewStyle().
		Foreground(Black).
		Background(Gold).
		Bold(true).
		Padding(0, 1)

	// Accent text (hot pink without bold)
	AccentStyle = lipgloss.NewStyle().
		Foreground(HotPink)

	// Secondary accent text (gold without bold)
	SecondaryAccentStyle = lipgloss.NewStyle().
		Foreground(Gold)
)

// ═══════════════════════════════════════════════════════════════
// CARD & CONTAINER STYLES
// ═══════════════════════════════════════════════════════════════

var (
	// Primary card - Gold border for elegance
	CardStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(Gold).
		Background(Black).
		Padding(1, 2)

	// Active/highlighted card - Hot pink border
	CardActiveStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(HotPink).
		Background(Black).
		Padding(1, 2)

	// Subtle card - Minimal border
	CardSubtleStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(GrayDark).
		Background(BlackSoft).
		Padding(1, 2)
)

// ═══════════════════════════════════════════════════════════════
// BORDER STYLES
// ═══════════════════════════════════════════════════════════════

var (
	RoundedBorder = lipgloss.RoundedBorder()
	ThickBorder   = lipgloss.ThickBorder()
	DoubleBorder  = lipgloss.DoubleBorder()
	NormalBorder  = lipgloss.NormalBorder()
)

// ═══════════════════════════════════════════════════════════════
// UNICODE SYMBOLS - Visual Language
// ═══════════════════════════════════════════════════════════════

const (
	// Bullets & Markers
	SymbolPointer      = "▸" // Primary pointer
	SymbolDiamond      = "◆" // Diamond - success/value
	SymbolFilledCircle = "●" // Filled circle - active/online
	SymbolEmptyCircle  = "○" // Empty circle - inactive
	SymbolSquare       = "⬥" // Rotated square - metric
	SymbolCheck        = "✓" // Checkmark - complete
	SymbolX            = "✕" // X mark - error
	SymbolWarning      = "⚠" // Warning
	SymbolInfo         = "ℹ" // Info
	SymbolStar         = "★" // Star - featured/favorite
	SymbolSparkle      = "✨" // Sparkle - Danni's signature

	// Arrows
	SymbolRightArrow  = "→"
	SymbolLeftArrow   = "←"
	SymbolUpArrow     = "↑"
	SymbolDownArrow   = "↓"
	SymbolDoubleRight = "⇒"

	// Progress
	SymbolBlockFull   = "█"
	SymbolBlockDark   = "▓"
	SymbolBlockMedium = "▒"
	SymbolBlockLight  = "░"

	// Separators
	SeparatorHeavy  = "═══════════════════"
	SeparatorMedium = "───────────────────"
	SeparatorLight  = "- - - - - - - - - -"
)

// ═══════════════════════════════════════════════════════════════
// HELPER FUNCTIONS
// ═══════════════════════════════════════════════════════════════

// StatusSymbol returns the appropriate symbol for a status
func StatusSymbol(status string) string {
	switch status {
	case "online", "active", "running":
		return SymbolFilledCircle
	case "success", "completed", "done":
		return SymbolDiamond
	case "offline", "inactive", "stopped":
		return SymbolEmptyCircle
	case "error", "failed", "attention":
		return SymbolX
	case "warning", "caution":
		return SymbolWarning
	case "info":
		return SymbolInfo
	case "featured", "favorite":
		return SymbolStar
	default:
		return SymbolPointer
	}
}

// StatusColor returns the appropriate color for a status
func StatusColor(status string) lipgloss.Color {
	switch status {
	case "online", "active", "running", "error", "failed", "attention":
		return HotPink
	case "success", "completed", "done", "warning", "caution":
		return Gold
	case "offline", "inactive", "stopped":
		return GrayMed
	case "info":
		return HotPinkAlt
	default:
		return White
	}
}

// ProgressBar creates a progress bar string with hot pink fill
func ProgressBar(percent int, width int) string {
	filled := int(float64(width) * float64(percent) / 100.0)
	empty := width - filled

	bar := ""
	for i := 0; i < filled; i++ {
		bar += SymbolBlockFull
	}
	for i := 0; i < empty; i++ {
		bar += SymbolBlockDark
	}

	return bar
}
