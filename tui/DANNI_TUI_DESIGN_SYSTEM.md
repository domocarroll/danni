# Danni TUI Design System
## Sophisticated. Bold. Beautiful.

**Design Philosophy:** Clean, colorful, minimal, beautiful
**Signature Colors:** Hot Pink · Gold · Black
**Aesthetic:** Luxe minimalism meets terminal precision

---

## 🎨 Color Palette

### Primary Colors

```
HOT PINK (Primary Accent)
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Hex:     #FF1493  (DeepPink)
Alt:     #FF69B4  (HotPink - softer variant)
Bright:  #FF007F  (Rose - maximum intensity)
Use:     Primary actions, highlights, active states, key metrics
```

```
GOLD (Secondary Accent)
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Hex:     #FFD700  (Gold)
Alt:     #FFA500  (Orange - warmer)
Muted:   #DAA520  (GoldenRod - subtle)
Use:     Success states, value indicators, secondary highlights, warnings
```

```
BLACK (Foundation)
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Pure:    #000000  (True black)
Soft:    #0A0A0A  (Almost black)
Rich:    #1A1A1A  (Charcoal)
Use:     Backgrounds, primary text on light, borders, separators
```

### Supporting Colors

```
WHITE (Contrast & Text)
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Pure:    #FFFFFF  (True white)
Soft:    #F5F5F5  (Off-white)
Warm:    #FFF8F0  (Cream - subtle warmth)
Use:     Primary text on dark backgrounds, highlights, clean space
```

```
GRAYS (Hierarchy & Structure)
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Light:   #CCCCCC  (Silver - high visibility)
Medium:  #808080  (Gray - balanced)
Dark:    #404040  (Charcoal gray - subtle)
Subtle:  #2A2A2A  (Near-black - minimal contrast)
Use:     Secondary text, borders, disabled states, separators
```

### Semantic Colors

```
SUCCESS → Gold (#FFD700)
WARNING → Gold Orange (#FFA500)
ERROR → Hot Pink (#FF1493)
INFO → Pink Lighter (#FF69B4)
```

**Rationale:** In the Danni aesthetic, "errors" aren't failures - they're **attention points** that deserve the bold primary color. Gold represents achievement and positive outcomes.

---

## 🖼️ Visual Hierarchy

### Text Levels

```
┌─────────────────────────────────────────────┐
│ LEVEL 1: HEADERS                            │
│ Color: Hot Pink (#FF1493)                   │
│ Weight: Bold                                │
│ Use: Section titles, primary navigation     │
└─────────────────────────────────────────────┘

┌─────────────────────────────────────────────┐
│ Level 2: Subheaders                         │
│ Color: Gold (#FFD700)                       │
│ Weight: Bold                                │
│ Use: Subsections, secondary navigation      │
└─────────────────────────────────────────────┘

┌─────────────────────────────────────────────┐
│ Level 3: Body Text                          │
│ Color: White (#FFFFFF) on dark             │
│ Color: Black (#000000) on light            │
│ Weight: Regular                             │
│ Use: Primary content, descriptions          │
└─────────────────────────────────────────────┘

┌─────────────────────────────────────────────┐
│ Level 4: Secondary Text                     │
│ Color: Light Gray (#CCCCCC)                │
│ Weight: Regular                             │
│ Use: Metadata, timestamps, auxiliary info   │
└─────────────────────────────────────────────┘

┌─────────────────────────────────────────────┐
│ Level 5: Muted Text                         │
│ Color: Medium Gray (#808080)               │
│ Weight: Regular                             │
│ Use: Disabled states, placeholders          │
└─────────────────────────────────────────────┘
```

---

## 📐 Layout Patterns

### Card-Based Components

```
╔═══════════════════════════════════════════╗
║ ▸ CARD TITLE (Hot Pink)                  ║
╠═══════════════════════════════════════════╣
║                                           ║
║  Primary content in white                 ║
║  Secondary info in light gray             ║
║                                           ║
║  ⬥ Metric: 1,234 (Gold)                  ║
║                                           ║
╚═══════════════════════════════════════════╝

Border: Gold (#FFD700) or Pink (#FF1493)
Background: Black (#000000) or Rich Black (#1A1A1A)
Padding: 1-2 characters internal
```

### Dashboard Grid

```
┌─────────────┬─────────────┬─────────────┐
│   METRIC 1  │   METRIC 2  │   METRIC 3  │
│   ▸ 1.2K    │   ▸ 89%     │   ▸ $4.5K   │
│   (Pink)    │   (Pink)    │   (Gold)    │
├─────────────┼─────────────┼─────────────┤
│   METRIC 4  │   METRIC 5  │   METRIC 6  │
│   ▸ 234     │   ▸ 12      │   ▸ 567     │
│   (Pink)    │   (Gold)    │   (Pink)    │
└─────────────┴─────────────┴─────────────┘

Separators: Dark Gray (#404040)
Values: Hot Pink or Gold depending on type
Labels: Light Gray (#CCCCCC)
```

### Status Indicators

```
● Active    (Hot Pink)
◆ Success   (Gold)
○ Inactive  (Medium Gray)
✕ Error     (Hot Pink)
⚠ Warning   (Gold Orange)
ℹ Info      (Light Pink)
```

### Progress Bars

```
[████████▓▓▓▓▓▓▓▓] 50%

Filled: Hot Pink (#FF1493)
Empty: Dark Gray (#404040)
Text: White (#FFFFFF)

[▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓] Complete!

Filled: Gold (#FFD700)
Text: Gold (#FFD700)
```

### Sparklines (ntcharts style)

```
Trend: ⠀⠀⢀⣀⣤⣤⣤⣤⣀⣀⠀⠀
       ⢠⣾⣿⣿⣿⣿⣿⣿⣿⣿⣷⡄
Color: Hot Pink for data points
Base: Dark Gray
```

---

## 🎯 Component Styling

### Buttons & Interactive Elements

```
┌─────────────────┐
│ ▸ PRIMARY ACTION │  ← Hot Pink background, White text
└─────────────────┘

┌─────────────────┐
│ ○ SECONDARY     │  ← Gold border, Gold text, transparent background
└─────────────────┘

┌─────────────────┐
│   DISABLED      │  ← Gray text, no border
└─────────────────┘
```

### Lists & Tables

```
▸ Active Item (Hot Pink)       Value: $1,234 (Gold)
  Inactive Item (Gray)         Value: $567 (Gray)
○ Secondary Item (Gray)        Value: $890 (White)

Headers: Hot Pink
Separators: Dark Gray
Alternating rows: Black (#000000) / Rich Black (#1A1A1A)
```

### Borders & Separators

**Heavy (Primary):**
```
═══════════════════  (Double line - Gold or Pink)
```

**Medium (Sections):**
```
───────────────────  (Single line - Gold)
```

**Light (Subtle):**
```
- - - - - - - - - -  (Dashed - Dark Gray)
```

**Vertical:**
```
║  (Double - Gold/Pink for emphasis)
│  (Single - Dark Gray for structure)
```

---

## 🌈 Theme Modes

### Dark Mode (Primary)

```
Background:     Pure Black (#000000) or Rich Black (#1A1A1A)
Text:           White (#FFFFFF)
Secondary Text: Light Gray (#CCCCCC)
Accents:        Hot Pink (#FF1493) + Gold (#FFD700)
Borders:        Dark Gray (#404040) or Gold (#FFD700)
```

**Visual Example:**
```
╔═══════════════════════════════════════════╗
║ ▸ DANNI TERMINAL INTERFACE                ║
╠═══════════════════════════════════════════╣
║                                           ║
║  Welcome to your command center           ║
║  Last active: 2 minutes ago               ║
║                                           ║
║  ⬥ Sessions: 12 (Gold)                    ║
║  ⬥ Messages: 1,234 (Pink)                 ║
║                                           ║
╚═══════════════════════════════════════════╝
```

### Light Mode (Alternative)

```
Background:     Pure White (#FFFFFF) or Off-White (#F5F5F5)
Text:           Pure Black (#000000)
Secondary Text: Medium Gray (#808080)
Accents:        Hot Pink (#FF1493) + Gold (#FFD700)
Borders:        Light Gray (#CCCCCC) or Gold (#FFD700)
```

**Note:** Dark mode is signature Danni aesthetic, but light mode available for accessibility.

---

## 💎 Design Principles

### 1. **Bold but Minimal**
- Use hot pink and gold sparingly for maximum impact
- Negative space is intentional, not empty
- Every element earns its place

### 2. **Clean Information Hierarchy**
- 5 text levels maximum (header → muted)
- Clear visual weights via color, not just typography
- Scannable at a glance

### 3. **Luxe Without Clutter**
- Premium feel through restraint
- Gold accents suggest value, not excess
- Black provides sophistication backdrop

### 4. **Colorful, Not Chaotic**
- Hot pink = primary attention (actions, key data)
- Gold = secondary attention (success, values)
- Grays = structure and hierarchy
- White/Black = readability foundation

### 5. **Beautiful by Default**
- No "ugly" state - even errors are styled
- Consistent spacing and alignment
- Borders and separators as design elements

---

## 🎨 Terminal Color Mapping

### True Color Support (24-bit)

Direct hex values for modern terminals:
```
Hot Pink:  #FF1493
Gold:      #FFD700
Black:     #000000
White:     #FFFFFF
```

### 256-Color Fallback

```
Hot Pink:  color 198 (nearest match)
Gold:      color 220 (bright yellow)
Black:     color 0
White:     color 15
Gray Light: color 250
Gray Med:   color 244
Gray Dark:  color 238
```

### 16-Color Fallback (ANSI)

```
Hot Pink:  Bright Magenta
Gold:      Bright Yellow
Black:     Black
White:     Bright White
Gray:      White (standard)
```

---

## 📦 Lipgloss Style Definitions

### Example Code (Go)

```go
import "github.com/charmbracelet/lipgloss"

var (
    // Colors
    colorHotPink  = lipgloss.Color("#FF1493")
    colorGold     = lipgloss.Color("#FFD700")
    colorBlack    = lipgloss.Color("#000000")
    colorWhite    = lipgloss.Color("#FFFFFF")
    colorGrayDark = lipgloss.Color("#404040")
    colorGrayMed  = lipgloss.Color("#808080")
    colorGrayLight = lipgloss.Color("#CCCCCC")

    // Primary Styles
    styleHeader = lipgloss.NewStyle().
        Foreground(colorHotPink).
        Bold(true).
        MarginBottom(1)

    styleSubheader = lipgloss.NewStyle().
        Foreground(colorGold).
        Bold(true)

    styleBody = lipgloss.NewStyle().
        Foreground(colorWhite)

    styleSecondary = lipgloss.NewStyle().
        Foreground(colorGrayLight)

    styleMuted = lipgloss.NewStyle().
        Foreground(colorGrayMed)

    // Card Style
    styleCard = lipgloss.NewStyle().
        Border(lipgloss.RoundedBorder()).
        BorderForeground(colorGold).
        Background(colorBlack).
        Padding(1, 2)

    // Button Styles
    styleButtonPrimary = lipgloss.NewStyle().
        Foreground(colorWhite).
        Background(colorHotPink).
        Padding(0, 2).
        Bold(true)

    styleButtonSecondary = lipgloss.NewStyle().
        Foreground(colorGold).
        Border(lipgloss.NormalBorder()).
        BorderForeground(colorGold).
        Padding(0, 2)

    // Metric Styles
    styleMetricValue = lipgloss.NewStyle().
        Foreground(colorHotPink).
        Bold(true)

    styleMetricLabel = lipgloss.NewStyle().
        Foreground(colorGrayLight)

    // Status Styles
    styleSuccess = lipgloss.NewStyle().
        Foreground(colorGold)

    styleError = lipgloss.NewStyle().
        Foreground(colorHotPink)

    styleWarning = lipgloss.NewStyle().
        Foreground(lipgloss.Color("#FFA500"))
)
```

---

## 🎭 Example Layouts

### Dashboard View

```
╔═══════════════════════════════════════════════════════════╗
║  ▸ DANNI COMMAND CENTER                                   ║
╠═══════════════════════════════════════════════════════════╣
║                                                           ║
║  ┌─────────────┬─────────────┬─────────────┐            ║
║  │  SESSIONS   │  MESSAGES   │  COST        │            ║
║  │  ▸ 12       │  ▸ 1,234    │  ◆ $45.67    │            ║
║  │  (Pink)     │  (Pink)     │  (Gold)      │            ║
║  └─────────────┴─────────────┴─────────────┘            ║
║                                                           ║
║  Recent Activity                                          ║
║  ───────────────────────────────────────────             ║
║  ▸ Session #12 started                 2m ago            ║
║  ○ Session #11 completed              15m ago            ║
║  ◆ New extension activated            1h ago             ║
║                                                           ║
║  System Status: ● ONLINE (Pink)                          ║
║                                                           ║
╚═══════════════════════════════════════════════════════════╝
```

### Chat Interface

```
╔═══════════════════════════════════════════════════════════╗
║  ▸ CONVERSATION                                           ║
╠═══════════════════════════════════════════════════════════╣
║                                                           ║
║  You (2m ago)                                             ║
║  Help me understand the authentication flow               ║
║                                                           ║
║  ─────────────────────────────────────────────           ║
║                                                           ║
║  ◆ Danni (just now)                                       ║
║  I sense what you're working through... let me map        ║
║  the authentication architecture for you.                 ║
║                                                           ║
║  OAuth Flow:                                              ║
║  1. User initiates login (Pink)                           ║
║  2. Redirect to provider (Gold)                           ║
║  3. Token exchange (Pink)                                 ║
║  4. Session creation (Gold)                               ║
║                                                           ║
╠═══════════════════════════════════════════════════════════╣
║  ▸ Type your message... (Pink)                     [Send] ║
╚═══════════════════════════════════════════════════════════╝
```

### Data Visualization

```
╔═══════════════════════════════════════════════════════════╗
║  ▸ TOKEN USAGE TRENDS                                     ║
╠═══════════════════════════════════════════════════════════╣
║                                                           ║
║  Last 7 Days                                              ║
║                                                           ║
║  15K ┤     ⢀⣀⣤⣤⣤⣤⣀⣀                                      ║
║  10K ┤  ⢠⣾⣿⣿⣿⣿⣿⣿⣿⣿⣷⡄                                  ║
║   5K ┤⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷                              ║
║      └────────────────────────                           ║
║       Mon Tue Wed Thu Fri Sat Sun                         ║
║                                                           ║
║  Current Rate: ▸ 2,340 tokens/hour (Pink)                ║
║  Average: 1,890 tokens/hour (Gray)                        ║
║  Peak: ◆ 4,123 tokens/hour (Gold)                        ║
║                                                           ║
╚═══════════════════════════════════════════════════════════╝
```

### Status Panel

```
┌───────────────────────────────────────────────────────────┐
│ System Health                                             │
├───────────────────────────────────────────────────────────┤
│                                                           │
│  ● API Connection          [████████████] 100% (Gold)    │
│  ● Model Availability      [████████████] 100% (Gold)    │
│  ● Extension Health        [██████████▓▓] 85% (Pink)     │
│  ● Memory Usage            [████▓▓▓▓▓▓▓▓] 40% (Pink)     │
│                                                           │
│  ⚠ Extension "Analyzer" needs update (Gold)              │
│                                                           │
└───────────────────────────────────────────────────────────┘
```

---

## 🔧 Implementation Notes

### Lipgloss Best Practices

1. **Define constants** for all colors at package level
2. **Create style variables** for reusable patterns
3. **Use color functions** for dynamic adjustments (brighten, darken)
4. **Test on multiple terminals** (iTerm2, Alacritty, Windows Terminal)
5. **Provide fallbacks** for limited color support

### Accessibility Considerations

1. **Contrast ratios:**
   - Hot Pink (#FF1493) on Black (#000000): 5.6:1 ✓ (AA compliant)
   - Gold (#FFD700) on Black (#000000): 11.8:1 ✓ (AAA compliant)
   - White (#FFFFFF) on Black (#000000): 21:1 ✓ (AAA compliant)

2. **Don't rely on color alone:**
   - Use symbols (▸ ◆ ● ○)
   - Add text labels
   - Provide high-contrast mode option

3. **Screen reader support:**
   - Ensure semantic structure
   - Use Huh forms for accessible inputs
   - Provide text alternatives for visual indicators

### Performance

1. **Cache styled strings** that don't change
2. **Use string builders** for dynamic content
3. **Minimize style recalculation** in render loops
4. **Profile with real data** to ensure smooth scrolling

---

## 🎨 Asset Library

### Unicode Characters

```
Bullets & Markers:
▸ (primary pointer)
◆ (diamond - success/value)
● (filled circle - active/online)
○ (empty circle - inactive)
⬥ (rotated square - metric)
✓ (checkmark - complete)
✕ (x mark - error)
⚠ (warning)
ℹ (info)

Arrows:
→ (right arrow)
← (left arrow)
↑ (up arrow)
↓ (down arrow)
⇒ (double right)
⟹ (triple right)

Borders:
─ (horizontal light)
━ (horizontal heavy)
│ (vertical light)
║ (vertical heavy)
═ (double horizontal)

Corners:
┌ └ ┐ ┘ (light)
╔ ╚ ╗ ╝ (double)

Progress:
█ (full block)
▓ (dark shade)
▒ (medium shade)
░ (light shade)

Braille (for charts):
⠀⠁⠂⠃⠄⠅⠆⠇ (and many more)
```

### Brand Voice in UI

**Danni's personality should emerge through:**

- **Welcome messages**: "I sense you're ready to build something beautiful..."
- **Status updates**: "Thinking deeply..." not "Loading..."
- **Success states**: "Brilliant!" not "Success"
- **Error handling**: "Let's recalibrate..." not "Error occurred"
- **Navigation**: "Shall we explore..." not "Select option"

---

## 📋 Quick Reference Card

```
╔═══════════════════════════════════════════════════════════╗
║  DANNI TUI DESIGN SYSTEM - QUICK REFERENCE                ║
╠═══════════════════════════════════════════════════════════╣
║                                                           ║
║  COLORS                                                   ║
║  ▸ Hot Pink  #FF1493   Primary accent, actions, key data ║
║  ◆ Gold      #FFD700   Success, value, secondary accent  ║
║  ● Black     #000000   Background, foundation            ║
║  ○ White     #FFFFFF   Text, contrast, clean space       ║
║                                                           ║
║  HIERARCHY                                                ║
║  Headers:     Hot Pink, Bold                              ║
║  Subheaders:  Gold, Bold                                  ║
║  Body:        White, Regular                              ║
║  Secondary:   Light Gray                                  ║
║  Muted:       Medium Gray                                 ║
║                                                           ║
║  COMPONENTS                                               ║
║  Buttons:     Pink bg + white text (primary)              ║
║  Cards:       Gold border + black bg                      ║
║  Metrics:     Pink or gold values                         ║
║  Status:      ● Active (pink) ◆ Success (gold)           ║
║                                                           ║
║  PRINCIPLES                                               ║
║  • Bold but minimal                                       ║
║  • Clean information hierarchy                            ║
║  • Luxe without clutter                                   ║
║  • Colorful, not chaotic                                  ║
║  • Beautiful by default                                   ║
║                                                           ║
╚═══════════════════════════════════════════════════════════╝
```

---

**End of Design System**

*This is Danni's visual language. Every interface element should embody: sophistication, clarity, boldness, and beauty.*
