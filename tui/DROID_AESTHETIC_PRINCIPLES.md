# DROID-Inspired Aesthetic Principles for Danni TUI

## What Makes DROID Look Professional

### 1. **Color Palette**
- **Background**: Dark blue-gray (#2B3E50 or similar) - NOT pure black
- **Primary Text**: Light gray/white (#E0E0E0) - softer than pure white
- **Accent Color**: ONE accent sparingly used (yellow/gold for highlights)
- **Secondary**: Muted blue-grays for less important text

### 2. **Typography**
- **ASCII Logo**: Big, blocky, monospace - takes up space confidently
- **Body Text**: Clean, no decorations, high contrast
- **NO GRADIENTS on body text** - Hard to read, looks amateur
- **Subtle dim text** for metadata/status

### 3. **Layout Principles**
- **Breathing Room**: Generous spacing around elements
- **Clear Hierarchy**: Important info stands out through position, not color
- **Minimal Borders**: Maybe one subtle divider line, not boxes everywhere
- **Status Bar**: Clean, informative, stays out of the way

### 4. **What to AVOID**
- ❌ Rainbow gradients on text
- ❌ Particle effects during normal operation
- ❌ Heavy borders and boxes
- ❌ Multiple bright colors competing
- ❌ Animations that distract from content

## Applying to Danni via Crush

### Start with Crush's Foundation
```go
// Crush already provides:
- Claude integration ✓
- Session management ✓
- Clean architecture ✓
- MCP support ✓
```

### Add DROID-style Visual Layer

#### 1. Welcome Screen
```
╭────────────────────────────────────────────╮
│                                            │
│     ██████   █████  ███   ██ ███   ██ ██  │
│     ██   ██ ██   ██ ████  ██ ████  ██ ██  │
│     ██   ██ ███████ ██ ██ ██ ██ ██ ██ ██  │
│     ██   ██ ██   ██ ██  ████ ██  ████ ██  │
│     ██████  ██   ██ ██   ███ ██   ███ ██  │
│                                            │
│              Strategic Intelligence         │
│                    v1.0.0                   │
╰────────────────────────────────────────────╯

  Current module: /strategy
  Session: abc-123  [connected]

> Ready for input...
```

#### 2. Color Usage
```go
var (
    // Base palette - muted and professional
    Background = lipgloss.Color("#1a1f2e")  // Dark blue-gray
    Surface    = lipgloss.Color("#2b3447")  // Slightly lighter
    TextPrimary = lipgloss.Color("#e0e6ed") // Soft white
    TextSecondary = lipgloss.Color("#8892a8") // Muted gray

    // ONE accent - used sparingly
    AccentGold = lipgloss.Color("#ffd369")  // Warm gold for highlights

    // Semantic - subtle versions
    Success = lipgloss.Color("#4ade80")  // Soft green
    Warning = lipgloss.Color("#fb923c")  // Soft orange
    Error   = lipgloss.Color("#f87171")  // Soft red
)
```

#### 3. Message Display
```
┌─ Danni (/strategy) ─────────────────────────
│ I've noticed an interesting pattern in your
│ approach. The strategic framework you're
│ building has three key leverage points...
└──────────────────────────────────────────────

> What specific outcomes are you optimizing for?

[tokens: 1,234 | cost: $0.02]
```

#### 4. Thinking State
```
◐ Analyzing strategic patterns...
```
Simple, no particles, no rainbow gradients.

## The Key Insight

**DROID looks professional because it shows RESTRAINT.**

- It has the confidence to use lots of empty space
- It doesn't need to prove it's "cool" with effects
- The interface gets out of the way of the content
- One accent color makes things pop when needed

## Implementation Strategy

1. **Fork Crush** - Get the solid foundation
2. **Strip out** any excessive styling
3. **Apply DROID palette** - Dark blue-gray base
4. **Add Danni personality** through:
   - Thoughtful responses
   - Strategic insights
   - NOT through visual noise
5. **Test readability** - Can you use it for hours without eye strain?

## Examples from Other Professional TUIs

### lazygit
- Muted colors
- Clear sections
- Keyboard hints subtle

### k9s (Kubernetes)
- Dark background
- Color coding that means something
- Dense information, still readable

### gh-dash
- Clean panels
- Subtle borders
- Accent color for current selection

## The Bottom Line

Danni should feel like a **high-end consultant** who shows up in understated designer clothes, not a Vegas showgirl. The intelligence and insights should dazzle, not the colors.