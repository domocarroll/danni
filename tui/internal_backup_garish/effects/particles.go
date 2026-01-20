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
// PARTICLE SWARM ENGINE - Living, breathing visual effects
// ═══════════════════════════════════════════════════════════════

type ParticleType int

const (
	ParticleSparkle ParticleType = iota
	ParticleStar
	ParticleOrb
	ParticleFlame
	ParticleSnow
	ParticleMatrix
	ParticlePlasma
	ParticleHologram
)

// Particle represents a single animated particle
type Particle struct {
	X, Y           float64
	VX, VY         float64  // Velocity
	Life           float64  // 0.0 to 1.0
	Type           ParticleType
	Color          string
	Symbol         string
	GlowIntensity  float64
	Trail          []Position
	PulsePhase     float64
}

// Position tracks historical positions for trails
type Position struct {
	X, Y float64
	Age  float64
}

// ParticleSwarm manages a collection of particles
type ParticleSwarm struct {
	Particles      []Particle
	Width, Height  int
	Time           float64
	EmissionRate   float64
	GravityX       float64
	GravityY       float64
	WindStrength   float64
	Turbulence     float64
	AttractorX     float64
	AttractorY     float64
	AttractorForce float64
	ColorScheme    ColorScheme
}

type ColorScheme int

const (
	SchemeDanni ColorScheme = iota // Hot Pink & Gold
	SchemeNeon                      // Electric colors
	SchemeAurora                    // Northern lights
	SchemePlasma                    // Plasma field
	SchemeHologram                  // Cyan & Magenta
	SchemeMatrix                    // Green cascade
	SchemeSunset                    // Warm gradients
	SchemeGalaxy                    // Deep space
)

// Particle symbols for different effects
var particleSymbols = map[ParticleType][]string{
	ParticleSparkle:  {"✦", "✧", "⋆", "✨", "★", "☆", "✯", "✰"},
	ParticleStar:     {"★", "☆", "✯", "✰", "✭", "✮"},
	ParticleOrb:      {"●", "○", "◉", "◎", "◌", "◍"},
	ParticleFlame:    {"🔥", "▲", "△", "▴", "▵"},
	ParticleSnow:     {"❄", "❅", "❆", "✻", "✼", "❉"},
	ParticleMatrix:   {"0", "1", "ﾊ", "ﾐ", "ﾋ", "ｰ", "ｳ", "ｼ", "ﾅ", "ﾓ", "ﾆ", "ｻ", "ﾜ", "ﾂ", "ｵ", "ﾘ", "ｱ", "ﾎ", "ﾃ", "ﾏ", "ｹ", "ﾒ", "ｴ", "ｶ", "ｷ", "ﾑ", "ﾕ", "ﾗ", "ｾ", "ﾈ", "ｽ", "ﾀ", "ﾇ"},
	ParticlePlasma:   {"◈", "◆", "◇", "◊", "⬡", "⬢"},
	ParticleHologram: {"▓", "▒", "░", "█", "▄", "▀"},
}

// NewParticleSwarm creates a new particle system
func NewParticleSwarm(width, height int, scheme ColorScheme) *ParticleSwarm {
	return &ParticleSwarm{
		Particles:      make([]Particle, 0, 100),
		Width:          width,
		Height:         height,
		EmissionRate:   0.5,
		GravityY:       0.1,
		WindStrength:   0.05,
		Turbulence:     0.02,
		AttractorX:     float64(width) / 2,
		AttractorY:     float64(height) / 2,
		AttractorForce: 0.01,
		ColorScheme:    scheme,
	}
}

// EmitParticle creates a new particle
func (ps *ParticleSwarm) EmitParticle(pType ParticleType, x, y float64) {
	symbols := particleSymbols[pType]
	symbol := symbols[rand.Intn(len(symbols))]

	particle := Particle{
		X:             x,
		Y:             y,
		VX:            (rand.Float64() - 0.5) * 2,
		VY:            (rand.Float64() - 0.5) * 2,
		Life:          1.0,
		Type:          pType,
		Symbol:        symbol,
		GlowIntensity: rand.Float64(),
		Trail:         make([]Position, 0, 10),
		PulsePhase:    rand.Float64() * math.Pi * 2,
	}

	// Set color based on scheme
	particle.Color = ps.getParticleColor(pType)

	ps.Particles = append(ps.Particles, particle)
}

// getParticleColor returns color based on scheme and type
func (ps *ParticleSwarm) getParticleColor(pType ParticleType) string {
	switch ps.ColorScheme {
	case SchemeDanni:
		if rand.Float64() > 0.5 {
			return "#FF1493" // Hot Pink
		}
		return "#FFD700" // Gold

	case SchemeNeon:
		colors := []string{"#FF00FF", "#00FFFF", "#FFFF00", "#FF00AA", "#00FFAA"}
		return colors[rand.Intn(len(colors))]

	case SchemeAurora:
		colors := []string{"#00FF00", "#00FFFF", "#FF00FF", "#9400D3", "#4B0082"}
		return colors[rand.Intn(len(colors))]

	case SchemePlasma:
		colors := []string{"#FF1493", "#FF00FF", "#FF69B4", "#FF007F", "#C71585"}
		return colors[rand.Intn(len(colors))]

	case SchemeHologram:
		colors := []string{"#00FFFF", "#FF00FF", "#00CED1", "#FF1493", "#7B68EE"}
		return colors[rand.Intn(len(colors))]

	case SchemeMatrix:
		intensity := []string{"#003300", "#006600", "#009900", "#00CC00", "#00FF00"}
		return intensity[rand.Intn(len(intensity))]

	case SchemeSunset:
		colors := []string{"#FF6B6B", "#FF8E53", "#FFD93D", "#6BCB77", "#4D96FF"}
		return colors[rand.Intn(len(colors))]

	case SchemeGalaxy:
		colors := []string{"#4B0082", "#8A2BE2", "#9400D3", "#9932CC", "#BA55D3"}
		return colors[rand.Intn(len(colors))]

	default:
		return "#FFFFFF"
	}
}

// Update updates all particles
func (ps *ParticleSwarm) Update() {
	ps.Time += 0.016 // ~60fps

	// Emit new particles
	if rand.Float64() < ps.EmissionRate {
		x := rand.Float64() * float64(ps.Width)
		y := 0.0
		pType := ParticleType(rand.Intn(int(ParticleHologram) + 1))
		ps.EmitParticle(pType, x, y)
	}

	// Update existing particles
	aliveParticles := make([]Particle, 0, len(ps.Particles))
	for i := range ps.Particles {
		p := &ps.Particles[i]

		// Add to trail
		if len(p.Trail) > 5 {
			p.Trail = p.Trail[1:]
		}
		p.Trail = append(p.Trail, Position{X: p.X, Y: p.Y, Age: p.Life})

		// Apply forces
		p.VX += ps.GravityX
		p.VY += ps.GravityY

		// Wind
		p.VX += math.Sin(ps.Time*2+p.Y*0.1) * ps.WindStrength

		// Turbulence
		p.VX += (rand.Float64() - 0.5) * ps.Turbulence
		p.VY += (rand.Float64() - 0.5) * ps.Turbulence

		// Attractor
		dx := ps.AttractorX - p.X
		dy := ps.AttractorY - p.Y
		dist := math.Sqrt(dx*dx + dy*dy)
		if dist > 0.1 {
			p.VX += (dx / dist) * ps.AttractorForce
			p.VY += (dy / dist) * ps.AttractorForce
		}

		// Update position
		p.X += p.VX
		p.Y += p.VY

		// Update life and effects
		p.Life -= 0.01
		p.PulsePhase += 0.1
		p.GlowIntensity = (math.Sin(p.PulsePhase) + 1) / 2

		// Keep alive particles
		if p.Life > 0 && p.Y < float64(ps.Height) && p.X >= 0 && p.X < float64(ps.Width) {
			aliveParticles = append(aliveParticles, *p)
		}
	}
	ps.Particles = aliveParticles
}

// Render creates a visual representation of the particle system
func (ps *ParticleSwarm) Render() string {
	// Create a 2D grid
	grid := make([][]rune, ps.Height)
	for i := range grid {
		grid[i] = make([]rune, ps.Width)
		for j := range grid[i] {
			grid[i][j] = ' '
		}
	}

	// Render particles with trails
	for _, p := range ps.Particles {
		// Render trail
		for i, pos := range p.Trail {
			x := int(pos.X)
			y := int(pos.Y)
			if x >= 0 && x < ps.Width && y >= 0 && y < ps.Height {
				alpha := float64(i) / float64(len(p.Trail)) * pos.Age
				if alpha > 0.3 {
					grid[y][x] = '·'
				}
			}
		}

		// Render particle
		x := int(p.X)
		y := int(p.Y)
		if x >= 0 && x < ps.Width && y >= 0 && y < ps.Height {
			if len(p.Symbol) > 0 {
				grid[y][x] = []rune(p.Symbol)[0]
			}
		}
	}

	// Convert grid to string
	var result strings.Builder
	for _, row := range grid {
		result.WriteString(string(row))
		result.WriteString("\n")
	}

	return result.String()
}

// ═══════════════════════════════════════════════════════════════
// SPECIAL EFFECTS
// ═══════════════════════════════════════════════════════════════

// FireworkBurst creates an explosion of particles
func (ps *ParticleSwarm) FireworkBurst(x, y float64, count int) {
	for i := 0; i < count; i++ {
		angle := float64(i) * (2 * math.Pi / float64(count))
		speed := 3.0 + rand.Float64()*2

		particle := Particle{
			X:             x,
			Y:             y,
			VX:            math.Cos(angle) * speed,
			VY:            math.Sin(angle) * speed,
			Life:          1.0,
			Type:          ParticleSparkle,
			Symbol:        "✨",
			GlowIntensity: 1.0,
			Color:         ps.getParticleColor(ParticleSparkle),
		}

		ps.Particles = append(ps.Particles, particle)
	}
}

// MatrixRain creates the Matrix-style falling code effect
func (ps *ParticleSwarm) MatrixRain() {
	for x := 0; x < ps.Width; x += 2 {
		if rand.Float64() < 0.1 {
			ps.EmitParticle(ParticleMatrix, float64(x), 0)
		}
	}
}

// AuroraBorealis creates northern lights effect
func (ps *ParticleSwarm) AuroraBorealis() {
	wave := math.Sin(ps.Time * 0.5)
	for x := 0; x < ps.Width; x++ {
		y := float64(ps.Height/3) + wave*10 + math.Sin(float64(x)*0.1+ps.Time)*5
		if rand.Float64() < 0.3 {
			particle := Particle{
				X:             float64(x),
				Y:             y,
				VX:            0,
				VY:            (rand.Float64() - 0.5) * 0.5,
				Life:          0.5,
				Type:          ParticlePlasma,
				Symbol:        "░",
				GlowIntensity: rand.Float64(),
				Color:         ps.getParticleColor(ParticlePlasma),
			}
			ps.Particles = append(ps.Particles, particle)
		}
	}
}

// Constellation creates connected star patterns
func (ps *ParticleSwarm) Constellation() {
	// Create star points
	stars := 5 + rand.Intn(5)
	centerX := float64(ps.Width) / 2
	centerY := float64(ps.Height) / 2
	radius := math.Min(float64(ps.Width), float64(ps.Height)) / 3

	for i := 0; i < stars; i++ {
		angle := float64(i) * (2 * math.Pi / float64(stars)) + ps.Time*0.02
		x := centerX + math.Cos(angle)*radius
		y := centerY + math.Sin(angle)*radius

		ps.EmitParticle(ParticleStar, x, y)
	}
}

// PlasmaField creates an undulating plasma effect
func (ps *ParticleSwarm) PlasmaField() {
	for x := 0; x < ps.Width; x += 3 {
		for y := 0; y < ps.Height; y += 3 {
			value := math.Sin(float64(x)*0.1+ps.Time) +
					 math.Sin(float64(y)*0.1+ps.Time*1.5) +
					 math.Sin(math.Sqrt(float64(x*x+y*y))*0.05+ps.Time*2)

			if value > 1.5 && rand.Float64() < 0.1 {
				ps.EmitParticle(ParticlePlasma, float64(x), float64(y))
			}
		}
	}
}

// ═══════════════════════════════════════════════════════════════
// TEA INTEGRATION
// ═══════════════════════════════════════════════════════════════

type ParticleMsg time.Time

func AnimateParticles() tea.Cmd {
	return tea.Tick(time.Millisecond*16, func(t time.Time) tea.Msg {
		return ParticleMsg(t)
	})
}