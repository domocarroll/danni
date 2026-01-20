package effects

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/subfracture/danni/tui/internal/styles"
)

// ═══════════════════════════════════════════════════════════════
// ANIMATED BACKGROUND ENGINE - Living, breathing environments
// ═══════════════════════════════════════════════════════════════

type BackgroundType int

const (
	BGNone BackgroundType = iota
	BGStars
	BGMatrix
	BGPlasma
	BGAurora
	BGCircuit
	BGWaves
	BGFire
	BGRain
	BGNebula
	BGHologram
)

// AnimatedBackground manages dynamic background effects
type AnimatedBackground struct {
	Type          BackgroundType
	Width         int
	Height        int
	Frame         int
	Particles     *ParticleSwarm
	GridData      [][]float64
	ColorData     [][]string
	LastUpdate    time.Time
	Intensity     float64
	ScrollOffset  float64
}

// NewAnimatedBackground creates a new background
func NewAnimatedBackground(bgType BackgroundType, width, height int) *AnimatedBackground {
	bg := &AnimatedBackground{
		Type:      bgType,
		Width:     width,
		Height:    height,
		Frame:     0,
		Intensity: 0.3,
	}

	// Initialize data structures based on type
	switch bgType {
	case BGStars, BGNebula:
		bg.Particles = NewParticleSwarm(width, height, SchemeGalaxy)
		bg.initStarField()

	case BGMatrix:
		bg.GridData = make([][]float64, height)
		bg.ColorData = make([][]string, height)
		for i := range bg.GridData {
			bg.GridData[i] = make([]float64, width)
			bg.ColorData[i] = make([]string, width)
			for j := range bg.GridData[i] {
				bg.GridData[i][j] = rand.Float64()
			}
		}

	case BGPlasma, BGAurora:
		bg.GridData = make([][]float64, height)
		for i := range bg.GridData {
			bg.GridData[i] = make([]float64, width)
		}

	case BGHologram:
		bg.Particles = NewParticleSwarm(width, height, SchemeHologram)

	default:
		// Simple backgrounds don't need initialization
	}

	return bg
}

// Update advances the background animation
func (bg *AnimatedBackground) Update() {
	bg.Frame++
	bg.ScrollOffset += 0.02

	switch bg.Type {
	case BGStars:
		bg.updateStars()

	case BGMatrix:
		bg.updateMatrix()

	case BGPlasma:
		bg.updatePlasma()

	case BGAurora:
		bg.updateAurora()

	case BGCircuit:
		// Circuit patterns update slowly
		if bg.Frame%10 == 0 {
			bg.updateCircuit()
		}

	case BGWaves:
		bg.updateWaves()

	case BGFire:
		bg.updateFire()

	case BGRain:
		bg.updateRain()

	case BGNebula:
		bg.updateNebula()

	case BGHologram:
		bg.updateHologram()
	}
}

// Render creates the visual representation
func (bg *AnimatedBackground) Render() string {
	switch bg.Type {
	case BGNone:
		return bg.renderEmpty()
	case BGStars:
		return bg.renderStars()
	case BGMatrix:
		return bg.renderMatrix()
	case BGPlasma:
		return bg.renderPlasma()
	case BGAurora:
		return bg.renderAurora()
	case BGCircuit:
		return bg.renderCircuit()
	case BGWaves:
		return bg.renderWaves()
	case BGFire:
		return bg.renderFire()
	case BGRain:
		return bg.renderRain()
	case BGNebula:
		return bg.renderNebula()
	case BGHologram:
		return bg.renderHologram()
	default:
		return bg.renderEmpty()
	}
}

// ═══════════════════════════════════════════════════════════════
// BACKGROUND UPDATE METHODS
// ═══════════════════════════════════════════════════════════════

func (bg *AnimatedBackground) initStarField() {
	// Create initial stars
	starCount := (bg.Width * bg.Height) / 50
	for i := 0; i < starCount; i++ {
		x := rand.Float64() * float64(bg.Width)
		y := rand.Float64() * float64(bg.Height)
		bg.Particles.EmitParticle(ParticleStar, x, y)
	}
}

func (bg *AnimatedBackground) updateStars() {
	// Twinkle effect
	for i := range bg.Particles.Particles {
		p := &bg.Particles.Particles[i]
		p.GlowIntensity = (math.Sin(float64(bg.Frame)*0.1+p.PulsePhase) + 1) / 2
	}

	// Occasional shooting star
	if rand.Float64() < 0.01 {
		y := rand.Float64() * float64(bg.Height/2)
		particle := Particle{
			X:      float64(bg.Width),
			Y:      y,
			VX:     -5.0,
			VY:     0.5,
			Life:   1.0,
			Type:   ParticleStar,
			Symbol: "⟶",
			Color:  "#FFFFFF",
		}
		bg.Particles.Particles = append(bg.Particles.Particles, particle)
	}

	bg.Particles.Update()
}

func (bg *AnimatedBackground) updateMatrix() {
	// Matrix rain effect
	for col := 0; col < bg.Width; col++ {
		// Random chance to start a new drop
		if rand.Float64() < 0.01 {
			bg.GridData[0][col] = 1.0
		}

		// Cascade down
		for row := bg.Height - 1; row > 0; row-- {
			if bg.GridData[row-1][col] > 0.1 {
				bg.GridData[row][col] = bg.GridData[row-1][col] * 0.9
			} else {
				bg.GridData[row][col] *= 0.95
			}
		}

		// Fade top row
		bg.GridData[0][col] *= 0.9
	}
}

func (bg *AnimatedBackground) updatePlasma() {
	time := float64(bg.Frame) * 0.02
	for y := 0; y < bg.Height; y++ {
		for x := 0; x < bg.Width; x++ {
			// Classic plasma algorithm
			value := math.Sin(float64(x)*0.1 + time)
			value += math.Sin(float64(y)*0.1 + time*1.5)
			value += math.Sin(math.Sqrt(float64(x*x+y*y))*0.05 + time*2)
			value = (value + 3) / 6 // Normalize to 0-1
			bg.GridData[y][x] = value
		}
	}
}

func (bg *AnimatedBackground) updateAurora() {
	time := float64(bg.Frame) * 0.01
	for y := 0; y < bg.Height; y++ {
		for x := 0; x < bg.Width; x++ {
			// Aurora waves
			wave1 := math.Sin(float64(x)*0.05 + time) * math.Cos(float64(y)*0.02)
			wave2 := math.Sin(float64(x)*0.03 + time*1.5) * math.Sin(float64(y)*0.04 + time)
			value := (wave1 + wave2 + 2) / 4
			bg.GridData[y][x] = value * bg.Intensity
		}
	}
}

func (bg *AnimatedBackground) updateCircuit() {
	// Occasionally light up circuit paths
	if rand.Float64() < 0.1 {
		startX := rand.Intn(bg.Width)
		startY := rand.Intn(bg.Height)
		// Create a path
		x, y := startX, startY
		for steps := 0; steps < 20; steps++ {
			if x >= 0 && x < bg.Width && y >= 0 && y < bg.Height {
				if bg.GridData == nil {
					bg.GridData = make([][]float64, bg.Height)
					for i := range bg.GridData {
						bg.GridData[i] = make([]float64, bg.Width)
					}
				}
				bg.GridData[y][x] = 1.0
			}

			// Random walk
			switch rand.Intn(4) {
			case 0:
				x++
			case 1:
				x--
			case 2:
				y++
			case 3:
				y--
			}
		}
	}

	// Fade existing paths
	if bg.GridData != nil {
		for y := range bg.GridData {
			for x := range bg.GridData[y] {
				bg.GridData[y][x] *= 0.95
			}
		}
	}
}

func (bg *AnimatedBackground) updateWaves() {
	// Water wave simulation
	time := float64(bg.Frame) * 0.05
	for y := 0; y < bg.Height; y++ {
		amplitude := math.Sin(float64(y)*0.1 + time) * 5
		offset := int(amplitude + float64(bg.Width)/2)
		if offset >= 0 && offset < bg.Width {
			if bg.GridData == nil {
				bg.GridData = make([][]float64, bg.Height)
				for i := range bg.GridData {
					bg.GridData[i] = make([]float64, bg.Width)
				}
			}
			bg.GridData[y][offset] = 1.0
		}
	}
}

func (bg *AnimatedBackground) updateFire() {
	// Fire effect - heat rises and spreads
	if bg.GridData == nil {
		bg.GridData = make([][]float64, bg.Height)
		for i := range bg.GridData {
			bg.GridData[i] = make([]float64, bg.Width)
		}
	}

	// Add heat at bottom
	for x := bg.Width/3; x < 2*bg.Width/3; x++ {
		if rand.Float64() < 0.8 {
			bg.GridData[bg.Height-1][x] = 1.0
		}
	}

	// Propagate heat upward
	for y := 0; y < bg.Height-1; y++ {
		for x := 1; x < bg.Width-1; x++ {
			heat := bg.GridData[y+1][x]
			heat += bg.GridData[y+1][x-1] * 0.2
			heat += bg.GridData[y+1][x+1] * 0.2
			bg.GridData[y][x] = heat * 0.95
		}
	}
}

func (bg *AnimatedBackground) updateRain() {
	// Rain drops
	if bg.Particles == nil {
		bg.Particles = NewParticleSwarm(bg.Width, bg.Height, SchemeDanni)
		bg.Particles.GravityY = 0.5
		bg.Particles.WindStrength = 0.02
	}

	// Emit rain drops
	for x := 0; x < bg.Width; x++ {
		if rand.Float64() < 0.05 {
			bg.Particles.EmitParticle(ParticleSnow, float64(x), 0)
		}
	}

	bg.Particles.Update()
}

func (bg *AnimatedBackground) updateNebula() {
	// Nebula clouds with particles
	if bg.Particles != nil {
		bg.Particles.PlasmaField()
		bg.Particles.Update()
	}
}

func (bg *AnimatedBackground) updateHologram() {
	// Holographic glitch effect
	if bg.Particles != nil {
		if rand.Float64() < 0.1 {
			x := rand.Float64() * float64(bg.Width)
			y := rand.Float64() * float64(bg.Height)
			bg.Particles.EmitParticle(ParticleHologram, x, y)
		}
		bg.Particles.Update()
	}
}

// ═══════════════════════════════════════════════════════════════
// BACKGROUND RENDER METHODS
// ═══════════════════════════════════════════════════════════════

func (bg *AnimatedBackground) renderEmpty() string {
	lines := make([]string, bg.Height)
	for i := range lines {
		lines[i] = strings.Repeat(" ", bg.Width)
	}
	return strings.Join(lines, "\n")
}

func (bg *AnimatedBackground) renderStars() string {
	if bg.Particles != nil {
		return bg.Particles.Render()
	}
	return bg.renderEmpty()
}

func (bg *AnimatedBackground) renderMatrix() string {
	if bg.GridData == nil {
		return bg.renderEmpty()
	}

	chars := "ﾊﾐﾋｰｳｼﾅﾓﾆｻﾜﾂｵﾘｱﾎﾃﾏｹﾒｴｶｷﾑﾕﾗｾﾈｽﾀﾇ01"
	lines := make([]string, bg.Height)

	for y := 0; y < bg.Height; y++ {
		line := ""
		for x := 0; x < bg.Width; x++ {
			intensity := bg.GridData[y][x]
			if intensity > 0.1 {
				char := string(chars[rand.Intn(len(chars))])
				// Color based on intensity
				if intensity > 0.8 {
					char = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF00")).Render(char)
				} else if intensity > 0.4 {
					char = lipgloss.NewStyle().Foreground(lipgloss.Color("#00AA00")).Render(char)
				} else {
					char = lipgloss.NewStyle().Foreground(lipgloss.Color("#005500")).Render(char)
				}
				line += char
			} else {
				line += " "
			}
		}
		lines[y] = line
	}

	return strings.Join(lines, "\n")
}

func (bg *AnimatedBackground) renderPlasma() string {
	if bg.GridData == nil {
		return bg.renderEmpty()
	}

	chars := []string{" ", "░", "▒", "▓", "█"}
	lines := make([]string, bg.Height)

	for y := 0; y < bg.Height; y++ {
		line := ""
		for x := 0; x < bg.Width; x++ {
			intensity := bg.GridData[y][x]
			charIndex := int(intensity * float64(len(chars)-1))
			if charIndex >= len(chars) {
				charIndex = len(chars) - 1
			}
			if charIndex < 0 {
				charIndex = 0
			}

			char := chars[charIndex]
			// Apply color gradient
			if intensity > 0.7 {
				char = lipgloss.NewStyle().Foreground(styles.HotPink).Render(char)
			} else if intensity > 0.4 {
				char = lipgloss.NewStyle().Foreground(styles.Gold).Render(char)
			} else if intensity > 0.2 {
				char = lipgloss.NewStyle().Foreground(styles.HotPinkAlt).Render(char)
			}
			line += char
		}
		lines[y] = line
	}

	return strings.Join(lines, "\n")
}

func (bg *AnimatedBackground) renderAurora() string {
	return bg.renderPlasma() // Similar rendering to plasma
}

func (bg *AnimatedBackground) renderCircuit() string {
	if bg.GridData == nil {
		return bg.renderEmpty()
	}

	lines := make([]string, bg.Height)
	circuitChars := []string{"─", "│", "┌", "┐", "└", "┘", "├", "┤", "┬", "┴", "┼", "•", "◦"}

	for y := 0; y < bg.Height; y++ {
		line := ""
		for x := 0; x < bg.Width; x++ {
			if bg.GridData[y][x] > 0.1 {
				char := circuitChars[rand.Intn(len(circuitChars))]
				intensity := bg.GridData[y][x]
				if intensity > 0.7 {
					char = lipgloss.NewStyle().Foreground(styles.Gold).Render(char)
				} else {
					char = lipgloss.NewStyle().Foreground(styles.GrayDark).Render(char)
				}
				line += char
			} else {
				line += " "
			}
		}
		lines[y] = line
	}

	return strings.Join(lines, "\n")
}

func (bg *AnimatedBackground) renderWaves() string {
	waveChars := []string{"∿", "∾", "≈", "≋", "～"}
	lines := make([]string, bg.Height)

	for y := 0; y < bg.Height; y++ {
		line := ""
		waveOffset := int(math.Sin(float64(y)*0.2+bg.ScrollOffset) * 3)
		for x := 0; x < bg.Width; x++ {
			if x == bg.Width/2+waveOffset {
				char := waveChars[(y+bg.Frame/5)%len(waveChars)]
				line += lipgloss.NewStyle().Foreground(styles.Gold).Render(char)
			} else {
				line += " "
			}
		}
		lines[y] = line
	}

	return strings.Join(lines, "\n")
}

func (bg *AnimatedBackground) renderFire() string {
	if bg.GridData == nil {
		return bg.renderEmpty()
	}

	lines := make([]string, bg.Height)
	fireChars := []string{" ", ".", ":", "^", "*", "🔥"}

	for y := 0; y < bg.Height; y++ {
		line := ""
		for x := 0; x < bg.Width; x++ {
			heat := bg.GridData[y][x]
			if heat > 0.1 {
				charIndex := int(heat * float64(len(fireChars)-1))
				if charIndex >= len(fireChars) {
					charIndex = len(fireChars) - 1
				}
				char := fireChars[charIndex]

				// Color based on heat
				if heat > 0.8 {
					char = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFF00")).Render(char)
				} else if heat > 0.5 {
					char = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF8800")).Render(char)
				} else {
					char = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Render(char)
				}
				line += char
			} else {
				line += " "
			}
		}
		lines[y] = line
	}

	return strings.Join(lines, "\n")
}

func (bg *AnimatedBackground) renderRain() string {
	if bg.Particles != nil {
		return bg.Particles.Render()
	}
	return bg.renderEmpty()
}

func (bg *AnimatedBackground) renderNebula() string {
	if bg.Particles != nil {
		// Render particles over a subtle cloud background
		particleRender := bg.Particles.Render()
		lines := strings.Split(particleRender, "\n")

		// Add nebula clouds
		for y := range lines {
			if y < len(lines) {
				runes := []rune(lines[y])
				for x := range runes {
					if runes[x] == ' ' && rand.Float64() < 0.02 {
						// Add subtle cloud effect
						if rand.Float64() < 0.5 {
							runes[x] = '·'
						} else {
							runes[x] = '∙'
						}
					}
				}
				lines[y] = string(runes)
			}
		}

		return strings.Join(lines, "\n")
	}
	return bg.renderEmpty()
}

func (bg *AnimatedBackground) renderHologram() string {
	if bg.Particles != nil {
		// Add scan lines for holographic effect
		render := bg.Particles.Render()
		lines := strings.Split(render, "\n")

		for y := range lines {
			if (y+bg.Frame/2)%4 == 0 {
				// Add scan line
				line := lines[y]
				styledLine := lipgloss.NewStyle().
					Foreground(lipgloss.Color("#00FFFF")).
					Faint(true).
					Render(line)
				lines[y] = styledLine
			}
		}

		return strings.Join(lines, "\n")
	}
	return bg.renderEmpty()
}

// ═══════════════════════════════════════════════════════════════
// TEA INTEGRATION
// ═══════════════════════════════════════════════════════════════

type BackgroundMsg time.Time

func AnimateBackground() tea.Cmd {
	return tea.Tick(time.Millisecond*50, func(t time.Time) tea.Msg {
		return BackgroundMsg(t)
	})
}