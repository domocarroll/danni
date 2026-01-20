package styles

import (
	"fmt"
	"math"
	"strings"
)

// ═══════════════════════════════════════════════════════════════
// GRADIENT ENGINE - Smooth color transitions for Danni's presence
// ═══════════════════════════════════════════════════════════════

// RGB represents a color in RGB space
type RGB struct {
	R, G, B uint8
}

// HexToRGB converts a hex color string to RGB
func HexToRGB(hex string) RGB {
	// Remove # if present
	hex = strings.TrimPrefix(hex, "#")

	// Parse hex values
	var r, g, b uint8
	fmt.Sscanf(hex, "%02x%02x%02x", &r, &g, &b)

	return RGB{R: r, G: g, B: b}
}

// ToANSI converts RGB to ANSI escape code
func (c RGB) ToANSI() string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", c.R, c.G, c.B)
}

// Interpolate creates a color between two RGB colors
func Interpolate(start, end RGB, t float64) RGB {
	// Clamp t between 0 and 1
	if t < 0 {
		t = 0
	}
	if t > 1 {
		t = 1
	}

	return RGB{
		R: uint8(float64(start.R) + (float64(end.R)-float64(start.R))*t),
		G: uint8(float64(start.G) + (float64(end.G)-float64(start.G))*t),
		B: uint8(float64(start.B) + (float64(end.B)-float64(start.B))*t),
	}
}

// ═══════════════════════════════════════════════════════════════
// GRADIENT GENERATORS
// ═══════════════════════════════════════════════════════════════

// GradientText creates a linear gradient across text
func GradientText(text string, startHex, endHex string) string {
	if text == "" {
		return ""
	}

	start := HexToRGB(startHex)
	end := HexToRGB(endHex)

	// Build the gradient string
	var result strings.Builder
	runes := []rune(text)
	length := len(runes)

	for i, r := range runes {
		// Calculate position in gradient (0.0 to 1.0)
		t := float64(i) / float64(length-1)
		if length == 1 {
			t = 0.5
		}

		// Get interpolated color
		color := Interpolate(start, end, t)

		// Apply color to character
		result.WriteString(color.ToANSI())
		result.WriteRune(r)
	}

	// Reset color at the end
	result.WriteString("\033[0m")

	return result.String()
}

// TriGradientText creates a gradient with three colors
func TriGradientText(text string, startHex, midHex, endHex string) string {
	if text == "" {
		return ""
	}

	start := HexToRGB(startHex)
	mid := HexToRGB(midHex)
	end := HexToRGB(endHex)

	var result strings.Builder
	runes := []rune(text)
	length := len(runes)

	for i, r := range runes {
		t := float64(i) / float64(length-1)
		if length == 1 {
			t = 0.5
		}

		var color RGB
		if t < 0.5 {
			// First half: start to mid
			color = Interpolate(start, mid, t*2)
		} else {
			// Second half: mid to end
			color = Interpolate(mid, end, (t-0.5)*2)
		}

		result.WriteString(color.ToANSI())
		result.WriteRune(r)
	}

	result.WriteString("\033[0m")
	return result.String()
}

// WaveGradientText creates a wave/sine gradient effect
func WaveGradientText(text string, color1Hex, color2Hex string, frequency float64) string {
	if text == "" {
		return ""
	}

	color1 := HexToRGB(color1Hex)
	color2 := HexToRGB(color2Hex)

	var result strings.Builder
	runes := []rune(text)

	for i, r := range runes {
		// Create sine wave position
		t := (math.Sin(float64(i)*frequency) + 1) / 2

		color := Interpolate(color1, color2, t)
		result.WriteString(color.ToANSI())
		result.WriteRune(r)
	}

	result.WriteString("\033[0m")
	return result.String()
}

// RainbowText creates a rainbow gradient effect
func RainbowText(text string) string {
	if text == "" {
		return ""
	}

	// Define rainbow colors
	colors := []RGB{
		HexToRGB("#FF0000"), // Red
		HexToRGB("#FF7F00"), // Orange
		HexToRGB("#FFFF00"), // Yellow
		HexToRGB("#00FF00"), // Green
		HexToRGB("#0000FF"), // Blue
		HexToRGB("#4B0082"), // Indigo
		HexToRGB("#9400D3"), // Violet
	}

	var result strings.Builder
	runes := []rune(text)
	length := len(runes)

	for i, r := range runes {
		// Position in the rainbow (0 to 6)
		pos := float64(i) / float64(length) * float64(len(colors)-1)

		// Find which two colors we're between
		colorIndex := int(pos)
		if colorIndex >= len(colors)-1 {
			colorIndex = len(colors) - 2
		}

		// Calculate position between the two colors
		localPos := pos - float64(colorIndex)

		// Interpolate between the colors
		color := Interpolate(colors[colorIndex], colors[colorIndex+1], localPos)

		result.WriteString(color.ToANSI())
		result.WriteRune(r)
	}

	result.WriteString("\033[0m")
	return result.String()
}

// ═══════════════════════════════════════════════════════════════
// DANNI SIGNATURE GRADIENTS
// ═══════════════════════════════════════════════════════════════

// DanniGradient creates Danni's signature Hot Pink to Gold gradient
func DanniGradient(text string) string {
	return GradientText(text, "#FF1493", "#FFD700")
}

// DanniReverseGradient creates Gold to Hot Pink gradient
func DanniReverseGradient(text string) string {
	return GradientText(text, "#FFD700", "#FF1493")
}

// DanniTriGradient creates Hot Pink → Gold → Hot Pink gradient
func DanniTriGradient(text string) string {
	return TriGradientText(text, "#FF1493", "#FFD700", "#FF1493")
}

// DanniWaveGradient creates a wave effect between Hot Pink and Gold
func DanniWaveGradient(text string) string {
	return WaveGradientText(text, "#FF1493", "#FFD700", 0.3)
}

// SunriseGradient creates a sunrise effect: Black → Hot Pink → Gold → White
func SunriseGradient(text string) string {
	if text == "" {
		return ""
	}

	colors := []RGB{
		HexToRGB("#000000"), // Black
		HexToRGB("#FF1493"), // Hot Pink
		HexToRGB("#FFD700"), // Gold
		HexToRGB("#FFFFFF"), // White
	}

	var result strings.Builder
	runes := []rune(text)
	length := len(runes)

	for i, r := range runes {
		pos := float64(i) / float64(length) * float64(len(colors)-1)
		colorIndex := int(pos)
		if colorIndex >= len(colors)-1 {
			colorIndex = len(colors) - 2
		}

		localPos := pos - float64(colorIndex)
		color := Interpolate(colors[colorIndex], colors[colorIndex+1], localPos)

		result.WriteString(color.ToANSI())
		result.WriteRune(r)
	}

	result.WriteString("\033[0m")
	return result.String()
}

// ElectricGradient creates an electric neon effect
func ElectricGradient(text string) string {
	return TriGradientText(text, "#FF1493", "#FF007F", "#FF69B4")
}

// LuxuryGradient creates a luxury gold effect
func LuxuryGradient(text string) string {
	return TriGradientText(text, "#FFD700", "#FFA500", "#DAA520")
}

// ═══════════════════════════════════════════════════════════════
// ANIMATED GRADIENT HELPERS
// ═══════════════════════════════════════════════════════════════

// AnimatedGradientFrame returns a single frame of animated gradient
// offset should be incremented each frame for animation effect
func AnimatedGradientFrame(text string, color1Hex, color2Hex string, offset float64) string {
	if text == "" {
		return ""
	}

	color1 := HexToRGB(color1Hex)
	color2 := HexToRGB(color2Hex)

	var result strings.Builder
	runes := []rune(text)

	for i, r := range runes {
		// Create animated position with offset
		t := math.Mod((float64(i)/10)+offset, 1.0)

		color := Interpolate(color1, color2, t)
		result.WriteString(color.ToANSI())
		result.WriteRune(r)
	}

	result.WriteString("\033[0m")
	return result.String()
}

// ShimmerEffect creates a shimmer/sparkle effect
func ShimmerEffect(text string, baseColorHex string, shimmerColorHex string, positions []int) string {
	if text == "" {
		return ""
	}

	baseColor := HexToRGB(baseColorHex)
	shimmerColor := HexToRGB(shimmerColorHex)

	var result strings.Builder
	runes := []rune(text)

	shimmerMap := make(map[int]bool)
	for _, pos := range positions {
		shimmerMap[pos] = true
	}

	for i, r := range runes {
		if shimmerMap[i] {
			result.WriteString(shimmerColor.ToANSI())
		} else {
			result.WriteString(baseColor.ToANSI())
		}
		result.WriteRune(r)
	}

	result.WriteString("\033[0m")
	return result.String()
}