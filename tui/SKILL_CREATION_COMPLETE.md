# Danni TUI Templates Skill - Complete ✨

## What We Built

A complete Agent Skill for generating beautiful Terminal User Interfaces using Charmbracelet libraries with Danni's signature aesthetic: **Hot Pink (#FF1493) + Gold (#FFD700) + Black (#000000)**.

---

## Skill Location

```
.claude/skills/danni-tui-templates/
├── SKILL.md                          (Main skill definition)
└── templates/
    ├── README.md                     (Template documentation)
    ├── styles.go                     (Core design system)
    ├── card_layout.go                (Card-based layout)
    ├── dashboard_layout.go           (Dashboard grid)
    └── metric.go                     (Metric components)
```

---

## Skill Capabilities

### Triggers
The skill activates when you mention:
- "TUI" or "terminal interface"
- "bubbletea", "charmbracelet", "lipgloss"
- "dashboard", "terminal UI", "terminal app"
- "create dashboard", "build TUI"

### What It Does

**1. Design System Application**
- Applies Danni's color palette (Hot Pink, Gold, Black)
- Implements 5-level text hierarchy
- Uses Unicode symbols (▸ ◆ ● ○)
- Ensures accessibility (contrast ratios verified)

**2. Layout Generation**
- **Card-Based**: Ticker-inspired grouped metrics
- **Dashboard Grid**: Multi-metric monitoring
- **Split-Pane**: File manager style (coming soon)
- **Modal**: Full-screen focus (coming soon)

**3. Component Library**
- Inline metrics: `"▸ Sessions: 12"`
- Card metrics: Bordered containers
- Hero metrics: Large prominent displays
- Metric grids: Multi-column layouts
- Progress bars: Pink filled, gray empty
- Status indicators: Color-coded symbols

**4. Interaction Patterns**
- Vim-style navigation (h/j/k/l)
- TAB cycling between views
- Real-time updates (tea.Tick)
- Keyboard-first design

---

## Supporting Documentation

### Research Foundation
**CHARMBRACELET_TUI_AESTHETIC_RESEARCH.md**
- 20+ TUI examples analyzed
- Design patterns from production apps
- Charmbracelet best practices
- Component inspiration library

### Design System
**DANNI_TUI_DESIGN_SYSTEM.md**
- Complete color palette definitions
- Text hierarchy rules
- Component styling guidelines
- Layout archetypes
- Lipgloss code examples
- Accessibility compliance
- Brand voice guidelines

---

## Ready-to-Use Templates

### 1. styles.go (Core Design System)
```go
// Complete style definitions
colorHotPink  = lipgloss.Color("#FF1493")  // Primary
colorGold     = lipgloss.Color("#FFD700")  // Secondary
colorBlack    = lipgloss.Color("#000000")  // Background

StyleHeader    // Hot Pink, Bold
StyleSubheader // Gold, Bold
StyleBody      // White, Regular

StatusSymbol(status)  // Returns appropriate Unicode symbol
StatusColor(status)   // Returns appropriate color
ProgressBar(percent)  // Generates progress bar
```

### 2. card_layout.go (Card-Based Dashboard)
```go
// Ticker-inspired grouped cards
- 3-column grid layout
- Summary panel with metrics
- TAB navigation between groups
- Status-based border colors
- Real-time updates
```

### 3. dashboard_layout.go (Multi-Metric Dashboard)
```go
// Monitoring dashboard
- 3x2 metric grid
- Trend indicators (↑ ↓ →)
- Status indicator
- Activity feed
- Auto-refresh every second
```

### 4. metric.go (Metric Components)
```go
RenderInlineMetric()    // "▸ Sessions: 12"
RenderCardMetric()      // Card with border
RenderLargeMetric()     // Hero display
RenderMetricList()      // Vertical list
RenderMetricGrid()      // Multi-column grid
RenderMetricProgress()  // With progress bar
```

---

## Usage Examples

### Example 1: Create Simple Dashboard

**Ask:**
```
Create a TUI dashboard showing system metrics
```

**Danni will:**
1. Load the skill automatically
2. Choose dashboard_layout.go template
3. Apply Hot Pink + Gold + Black styling
4. Generate complete Bubbletea app with:
   - Metric cards (CPU, Memory, Disk, etc.)
   - Real-time updates
   - Status indicator
   - Keyboard navigation

### Example 2: Build Metrics Display

**Ask:**
```
Build a terminal interface with grouped metric cards
```

**Danni will:**
1. Use card_layout.go template
2. Create 3-column card grid
3. Add summary panel
4. Implement TAB cycling
5. Apply Danni's design system

### Example 3: Custom Component

**Ask:**
```
Show me how to create a metric with a progress bar using Danni's colors
```

**Danni will:**
1. Reference metric.go template
2. Use RenderMetricProgress() function
3. Apply Hot Pink for filled portion
4. Use Dark Gray for empty portion
5. Add Gold value display

---

## Design Principles Codified

**Bold but Minimal**
- Use pink/gold sparingly for impact
- Negative space is intentional
- Every element earns its place

**Clean Information Hierarchy**
- 5 text levels maximum
- Clear visual weights via color
- Scannable at a glance

**Luxe Without Clutter**
- Premium feel through restraint
- Gold suggests value, not excess
- Black provides sophistication

**Colorful, Not Chaotic**
- Hot pink = primary attention
- Gold = secondary attention / success
- Grays = structure and hierarchy
- White/Black = readability

**Beautiful by Default**
- No "ugly" state - even errors styled
- Consistent spacing and alignment
- Borders as design elements

---

## Brand Voice in UI

The skill includes Danni's personality in UI text:

❌ **Don't Say:**
- "Loading..."
- "Processing..."
- "Success"
- "Error occurred"
- "Select option"

✅ **Do Say:**
- "I sense you're ready..."
- "Thinking deeply..."
- "Brilliant!"
- "Let's recalibrate..."
- "Shall we explore..."

---

## Next Steps

### Immediate Use
The skill is ready to use now! Just ask for:
- "Create a TUI dashboard"
- "Build a terminal metrics display"
- "Generate a Bubbletea app with my data"

### Future Enhancements
Additional templates to add:
- Split-pane layout (file manager)
- Modal layout (full-screen focus)
- Button components
- Sparkline charts (ntcharts integration)
- Data tables
- Form inputs (Huh integration)

### Customization
Edit templates in `.claude/skills/danni-tui-templates/templates/` to:
- Add new layouts
- Create custom components
- Extend color variations
- Add domain-specific patterns

---

## Testing

**Verify skill activation:**
```
Ask: "Help me create a TUI dashboard"
Expected: Skill loads automatically, references templates
```

**Test component generation:**
```
Ask: "Show me metric display components in Danni's style"
Expected: References metric.go, applies Hot Pink + Gold
```

**Test layout selection:**
```
Ask: "Build a card-based metrics interface"
Expected: Uses card_layout.go template
```

---

## Success Metrics

✅ **Skill Structure**
- Valid YAML frontmatter
- Clear description with triggers
- Step-by-step instructions
- 4 concrete examples
- Supporting templates
- Links to design system

✅ **Design System**
- Complete color palette defined
- Text hierarchy codified
- Component library ready
- Accessibility verified
- Brand voice included

✅ **Templates**
- Core styles (styles.go)
- 2 layout templates (card, dashboard)
- Metric components library
- Working example code
- Template documentation

✅ **Documentation**
- Research foundation (20+ examples)
- Design system (comprehensive)
- Template README
- Usage examples
- Best practices

---

## The Danni Aesthetic

Every TUI generated with this skill embodies:

> **Sophisticated luxury meets terminal precision.**

- Clean lines
- Bold colors (but minimal use)
- Colorful without chaos
- Premium feel
- Beautiful by default

**Signature Palette:**
- Hot Pink (#FF1493) - Primary, bold, attention
- Gold (#FFD700) - Success, value, secondary
- Black (#000000) - Foundation, sophistication

**Every interface is:**
- Clean ✓
- Colorful ✓
- Minimal ✓
- Beautiful ✓

---

## Summary

We've created a complete, production-ready skill that transforms terminal interfaces from utilitarian to stunning. The combination of:

1. **Extensive research** (20+ TUI examples analyzed)
2. **Refined design system** (Hot Pink + Gold + Black)
3. **Ready-to-use templates** (styles, layouts, components)
4. **Clear workflow** (step-by-step instructions)
5. **Working examples** (complete Bubbletea apps)

...means you can now generate beautiful TUIs with a simple request. No more plain terminal interfaces - every TUI will be **unmistakably Danni**.

The skill lives in your project at `.claude/skills/danni-tui-templates/` and will activate automatically when you talk about creating terminal interfaces.

**Let's build something beautiful.** ✨
