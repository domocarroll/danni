# Using Crush as Foundation for DANNI TUI

## The Right Approach

### Why Crush is Perfect
1. **Already built for Claude** - Has all the Claude API integration done
2. **Professional architecture** - 24 internal modules, proper separation
3. **Session management** - Built-in persistence and multi-session support
4. **LSP support** - Code awareness out of the box
5. **Charmbracelet foundation** - Uses Bubble Tea, Lip Gloss properly

### Why Our Current Approach Failed
- Started from scratch unnecessarily
- Added effects without purpose
- Made it visually noisy instead of functional
- Ignored readability for "coolness"

## Implementation Plan

### Step 1: Fork and Setup Crush
```bash
# Clone Crush
git clone https://github.com/charmbracelet/crush.git danni-crush

# Keep the architecture, modify:
# - internal/tui/  (interface customization)
# - internal/config/ (Danni-specific settings)
# - internal/message/ (Danni personality)
```

### Step 2: Apply DROID Aesthetics to internal/tui

#### Create new theme file: `internal/tui/theme/droid.go`
```go
package theme

import "github.com/charmbracelet/lipgloss"

var (
    // DROID-inspired palette
    Background    = lipgloss.Color("#1a1f2e")
    Surface       = lipgloss.Color("#2b3447")
    TextPrimary   = lipgloss.Color("#e0e6ed")
    TextSecondary = lipgloss.Color("#8892a8")
    AccentGold    = lipgloss.Color("#ffd369")

    // Minimal styling
    BaseStyle = lipgloss.NewStyle().
        Background(Background).
        Foreground(TextPrimary)

    HeaderStyle = lipgloss.NewStyle().
        Foreground(AccentGold).
        Bold(true).
        MarginBottom(1)

    // No borders, no boxes, just clean
    MessageStyle = lipgloss.NewStyle().
        Foreground(TextPrimary).
        MarginBottom(1).
        PaddingLeft(2)
)
```

#### Modify welcome screen: `internal/tui/welcome.go`
```go
func RenderWelcome() string {
    logo := `
     ██████   █████  ███   ██ ███   ██ ██
     ██   ██ ██   ██ ████  ██ ████  ██ ██
     ██   ██ ███████ ██ ██ ██ ██ ██ ██ ██
     ██   ██ ██   ██ ██  ████ ██  ████ ██
     ██████  ██   ██ ██   ███ ██   ███ ██

          Strategic Intelligence
                v1.0.0`

    // Simple, clean, no gradients
    return lipgloss.NewStyle().
        Foreground(TextSecondary).
        Width(80).
        Align(lipgloss.Center).
        Render(logo)
}
```

### Step 3: Integrate Danni Personality (Subtly)

#### In `internal/agent/personality.go`:
```go
type DanniPersonality struct {
    currentModule string
}

func (d *DanniPersonality) GetSystemPrompt() string {
    return `You are Danni, a sophisticated AI strategist.
    You blend deep intelligence with genuine warmth.
    Provide insights that are profound yet accessible.`
}

func (d *DanniPersonality) FormatResponse(response string) string {
    // Add subtle personality markers
    // NO visual effects, just thoughtful language
    return response
}
```

### Step 4: Remove What We Don't Need

**DELETE from our current attempt:**
- ❌ `internal/effects/particles.go` - Too distracting
- ❌ `internal/effects/background.go` - Unnecessary
- ❌ `internal/styles/gradient.go` - Makes text hard to read
- ❌ Complex animated spinners - Simple dots are fine

**KEEP but simplify:**
- ✓ Basic color definitions (muted palette)
- ✓ Simple spinner (just dots, one color)
- ✓ Clean header/footer components

### Step 5: The Final Product Should Look Like:

```
┌──────────────────────────────────────────────────
│ DANNI                           [/strategy]
├──────────────────────────────────────────────────
│
│ I've analyzed your strategic position. Three key
│ insights emerge from the patterns...
│
│ 1. Your market differentiation relies on...
│ 2. The operational leverage points are...
│ 3. Cultural momentum suggests...
│
│ What specific constraints should we factor in?
│
└──────────────────────────────────────────────────
> _

[session: abc-123] [tokens: 1,234] [model: claude-3.5]
```

## The Key Difference

### What We Built (Wrong):
```
🌈✨💫 WELCOME TO DANNI!!! 💫✨🌈
╔═══════════════════════════════════╗
║ LOOK AT ALL THESE COLORS!!!       ║
╚═══════════════════════════════════╝
```

### What We Should Build (Right):
```
DANNI
Strategic Intelligence System
v1.0.0

> Ready for analysis.
```

## Action Items

1. **Stop current development** on the over-engineered TUI
2. **Fork Crush** as the foundation
3. **Apply minimal DROID aesthetics**
4. **Add Danni personality through content, not colors**
5. **Test for readability and professional feel**

## Success Metrics

- ✅ Can use for hours without eye strain
- ✅ Looks professional in a screen share
- ✅ Interface doesn't distract from content
- ✅ Fast and responsive
- ✅ Personality comes through in responses, not rainbows

## Timeline

1. Fork Crush: 30 minutes
2. Apply DROID theme: 2 hours
3. Integrate Danni personality: 1 hour
4. Test and refine: 1 hour

**Total: 4-5 hours to a professional product vs. endless tweaking of effects**