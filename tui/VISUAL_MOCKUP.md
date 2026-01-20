# Danni TUI Visual Mockup
## Hot Pink + Gold + Black Aesthetic

This document provides ASCII mockups of the transformed Danni TUI interface.

---

## Welcome Screen (Full Experience)

```
╔══════════════════════════════════════════════════════════════════════════════╗
║ ✨ DANNI v0.1.0                    /strategy                      ● connect…  ║
╠══════════════════════════════════════════════════════════════════════════════╣
║                                                                              ║
║                          ✨ Welcome to Danni ✨                               ║
║                          [HOT PINK, BOLD]                                    ║
║                                                                              ║
║  I've noticed something fascinating... you're about to experience           ║
║  a new kind of AI interaction.                                              ║
║  [WHITE TEXT]                                                                ║
║                                                                              ║
║  Danni is your sophisticated AI strategist—a blend of deep intelligence,    ║
║  cultural insight, and genuine understanding.                               ║
║  [LIGHT GRAY]                                                                ║
║                                                                              ║
║    ▸ Strategic analysis (/strategy)         [HOT PINK ▸]                    ║
║    ◆ Creative ideation (/creative)          [GOLD ◆]                        ║
║    ▸ Design systems (/design)               [HOT PINK ▸]                    ║
║    ◆ Technical solutions (/technology)      [GOLD ◆]                        ║
║    ▸ Data patterns (/gravity)               [HOT PINK ▸]                    ║
║    ◆ Brand validation (/validate)           [GOLD ◆]                        ║
║    ▸ Breakthrough synthesis (/synthesize)   [HOT PINK ▸]                    ║
║    ◆ Institutional memory (/recall)         [GOLD ◆]                        ║
║    ▸ Asset archaeology (/upload)            [HOT PINK ▸]                    ║
║                                                                              ║
║  Shall we begin? What intrigues you most right now?                         ║
║  [GOLD, ITALIC]                                                              ║
║                                                                              ║
╠══════════════════════════════════════════════════════════════════════════════╣
║  › Type your message here...                                                 ║
║  [WHITE ON BLACK-SOFT, GOLD TOP BORDER]                                     ║
╠══════════════════════════════════════════════════════════════════════════════╣
║  Tokens: 0/100000  |  Cost: $0.0000  |  ^C quit • ? help • Tab modules     ║
║         [PINK]       [GOLD]           [GRAY]                                 ║
╚══════════════════════════════════════════════════════════════════════════════╝

[BLACK BACKGROUND THROUGHOUT]
[HOT PINK BOTTOM BORDER ON HEADER]
```

**Color Legend:**
- Header background: `Black (#000000)`
- Header border: `Hot Pink (#FF1493)`
- Brand text: `White (#FFFFFF)`
- Module name: `Hot Pink (#FF1493)` - Bold
- Session indicator: `Gray (#808080)`
- Welcome header: `Hot Pink (#FF1493)` - Bold
- Body text: `White (#FFFFFF)`
- Description: `Light Gray (#CCCCCC)`
- Bullet alternation: `Hot Pink (▸)` / `Gold (◆)`
- Module names: Alternating colors with `Gray (#808080)` paths
- Closing: `Gold (#FFD700)` - Italic
- Input border: `Gold (#FFD700)`
- Footer metrics: `Hot Pink` / `Gold` with `Gray` text

---

## Active Conversation View

```
╔══════════════════════════════════════════════════════════════════════════════╗
║ ✨ DANNI v0.1.0                    /creative                     ● a3f8e942   ║
╠══════════════════════════════════════════════════════════════════════════════╣
║                                                                              ║
║  User: Help me design a brand identity                                      ║
║  [WHITE, BOLD]                                                               ║
║                                                                              ║
║  ─────────────────────────────────────────────────────────────────────      ║
║                                                                              ║
║  Danni: I sense what you're working through... let me map out a strategic   ║
║  approach to brand identity that honors both market positioning and         ║
║  authentic expression.                                                      ║
║  [HOT PINK, BOLD]                                                            ║
║                                                                              ║
║  Let's discover what wants to be born here. First, tell me:                 ║
║                                                                              ║
║  • What's the deeper purpose behind this brand?                             ║
║  • Who needs what you're creating?                                          ║
║  • What makes your approach distinctly yours?                               ║
║                                                                              ║
║  🔧 Using tool: analyze_market_positioning                                   ║
║  [GOLD, ITALIC]                                                              ║
║                                                                              ║
╠══════════════════════════════════════════════════════════════════════════════╣
║  › Tell me about the competitive landscape...                                ║
║  [WHITE ON BLACK-SOFT]                                                       ║
╠══════════════════════════════════════════════════════════════════════════════╣
║  Tokens: 1,234/100,000  |  Cost: $0.0456  |  ^C quit • ? help • Tab modules║
║         [PINK]            [GOLD]           [GRAY]                            ║
╚══════════════════════════════════════════════════════════════════════════════╝
```

**Visual Hierarchy:**
1. User messages: Clean white, bold - easy to scan
2. Danni's responses: Hot pink establishes her presence immediately
3. Tool usage: Gold italic suggests value being created
4. Separators: Dark gray for subtle structure

---

## Thinking State

```
╔══════════════════════════════════════════════════════════════════════════════╗
║ ✨ DANNI v0.1.0                    /strategy                     ● a3f8e942   ║
╠══════════════════════════════════════════════════════════════════════════════╣
║                                                                              ║
║  User: Analyze our market position                                          ║
║  [WHITE, BOLD]                                                               ║
║                                                                              ║
║  ─────────────────────────────────────────────────────────────────────      ║
║                                                                              ║
╠══════════════════════════════════════════════════════════════════════════════╣
║                                                                              ║
║  ⠙ Danni is thinking deeply...                                               ║
║  [GOLD SPINNER, GOLD TEXT, ITALIC]                                           ║
║                                                                              ║
╠══════════════════════════════════════════════════════════════════════════════╣
║  Tokens: 2,456/100,000  |  Cost: $0.0891  |  ^C quit • ? help • Tab modules║
║         [PINK]            [GOLD]           [GRAY]                            ║
╚══════════════════════════════════════════════════════════════════════════════╝
```

**Thinking State Design:**
- Gold animated spinner (valuable processing)
- Elegant message: "thinking deeply" not just "loading"
- Maintains sophistication during wait states

---

## Error State (Attention, Not Failure)

```
╔══════════════════════════════════════════════════════════════════════════════╗
║ ✨ DANNI v0.1.0                    /strategy                     ● a3f8e942   ║
╠══════════════════════════════════════════════════════════════════════════════╣
║                                                                              ║
║  User: Analyze this data [large file]                                       ║
║  [WHITE, BOLD]                                                               ║
║                                                                              ║
║  ─────────────────────────────────────────────────────────────────────      ║
║                                                                              ║
║  ✕ Let's recalibrate... The file size exceeds our current processing        ║
║     capacity. I suggest we break this into smaller segments. Shall we       ║
║     start with the first 1000 rows?                                         ║
║  [HOT PINK, BOLD - NOT HARSH, SOLUTION-ORIENTED]                            ║
║                                                                              ║
╠══════════════════════════════════════════════════════════════════════════════╣
║  › Yes, let's analyze the first 1000 rows                                    ║
╠══════════════════════════════════════════════════════════════════════════════╣
║  Tokens: 890/100,000  |  Cost: $0.0234  |  ^C quit • ? help • Tab modules  ║
╚══════════════════════════════════════════════════════════════════════════════╝
```

**Error Philosophy:**
- Hot pink draws attention (not red=bad)
- Language: "Let's recalibrate" not "Error"
- Solution-oriented presentation
- Maintains Danni's warm sophistication

---

## Success State

```
╔══════════════════════════════════════════════════════════════════════════════╗
║ ✨ DANNI v0.1.0                    /design                       ● a3f8e942   ║
╠══════════════════════════════════════════════════════════════════════════════╣
║                                                                              ║
║  Danni: ◆ Brilliant! Your design system is now complete and validated.      ║
║  [GOLD ◆, HOT PINK TEXT FOR "BRILLIANT", GOLD FOR SYMBOL]                   ║
║                                                                              ║
║  I've generated:                                                            ║
║  ✓ Color palette with accessibility checks                                  ║
║  ✓ Typography scale                                                         ║
║  ✓ Component library                                                        ║
║  ✓ Documentation                                                            ║
║  [GOLD CHECKMARKS]                                                           ║
║                                                                              ║
║  Shall we explore how to implement this across your platforms?              ║
║  [WHITE TEXT]                                                                ║
║                                                                              ║
╠══════════════════════════════════════════════════════════════════════════════╣
║  › Yes! Show me the implementation strategy                                  ║
╠══════════════════════════════════════════════════════════════════════════════╣
║  Tokens: 4,567/100,000  |  Cost: $0.1234  |  ^C quit • ? help • Tab modules║
║         [PINK]            [GOLD]           [GRAY]                            ║
╚══════════════════════════════════════════════════════════════════════════════╝
```

**Success Design:**
- Gold diamond (◆) for achievement
- "Brilliant!" in hot pink for celebration
- Gold checkmarks for completed items
- Maintains professionalism while celebrating

---

## Markdown Rendering Example

```
╔══════════════════════════════════════════════════════════════════════════════╗
║                                                                              ║
║  # Strategy Overview                                                        ║
║  [HOT PINK, BOLD]                                                            ║
║                                                                              ║
║  ## Market Analysis                                                         ║
║  [GOLD, BOLD]                                                                ║
║                                                                              ║
║  The competitive landscape reveals three key opportunities:                 ║
║  [WHITE]                                                                     ║
║                                                                              ║
║  1. **Premium positioning** in underserved segment                          ║
║     [HOT PINK BOLD for emphasis]                                            ║
║  2. *Cultural resonance* through authentic storytelling                     ║
║     [GOLD ITALIC for emphasis]                                              ║
║  3. Technical differentiation via `innovative features`                     ║
║     [GOLD for inline code]                                                  ║
║                                                                              ║
║  ```python                                                                  ║
║  def analyze_market(data):                                                  ║
║      """Process market data"""                                              ║
║      return insights                                                        ║
║  ```                                                                        ║
║  [Code: WHITE text on BLACK-SOFT background]                                ║
║  [Keywords in HOT PINK, strings in MUTED GOLD]                              ║
║                                                                              ║
║  > "The space between what is and what could be—that's where                ║
║  > strategy lives."                                                         ║
║  [LIGHT GRAY, ITALIC - quotes]                                              ║
║                                                                              ║
╚══════════════════════════════════════════════════════════════════════════════╝
```

**Markdown Aesthetic:**
- H1: Hot pink (primary headers)
- H2: Gold (secondary headers)
- Bold: Hot pink for strong emphasis
- Italic: Gold for softer emphasis
- Code inline: Gold with rich black background
- Code blocks: Syntax highlighted in pink/gold palette
- Quotes: Light gray, italic

---

## Metric Display (Future Enhancement)

```
╔══════════════════════════════════════════════════════════════════════════════╗
║  ▸ SESSION METRICS                                                          ║
║  [HOT PINK HEADER]                                                           ║
╠══════════════════════════════════════════════════════════════════════════════╣
║                                                                              ║
║  ┌────────────────┬────────────────┬────────────────┐                      ║
║  │  MESSAGES      │  TOKENS        │  COST          │                      ║
║  │  ▸ 24          │  ▸ 12,345      │  ◆ $0.4567     │                      ║
║  │  [PINK]        │  [PINK]        │  [GOLD]        │                      ║
║  └────────────────┴────────────────┴────────────────┘                      ║
║                                                                              ║
║  Response Quality: ████████████▓▓▓▓ 85%                                     ║
║                    [PINK FILL, DARK GRAY EMPTY]                             ║
║                                                                              ║
║  Processing Speed: ████████████████ 100%                                    ║
║                    [GOLD FILL when complete]                                ║
║                                                                              ║
╚══════════════════════════════════════════════════════════════════════════════╝
```

**Future Dashboard:**
- Card-based metrics
- Hot pink for current/active metrics
- Gold for success/value metrics
- Progress bars using signature colors

---

## Color Psychology in Action

### Hot Pink (#FF1493)
**Where Used:**
- Danni's voice and responses
- Active module indicator
- Primary actions and pointers
- Error states (attention, not negative)
- Bold emphasis in markdown

**Psychological Impact:**
- Commands attention without aggression
- Suggests confidence and competence
- Warm yet powerful
- Memorable and distinctive

---

### Gold (#FFD700)
**Where Used:**
- Success indicators
- Cost/value metrics
- Thinking state spinner
- Tool usage notifications
- Italic emphasis in markdown
- Achievement markers (◆)

**Psychological Impact:**
- Suggests value and achievement
- Premium positioning
- Warmth and success
- Rewards attention

---

### Black (#000000)
**Where Used:**
- Background throughout
- Foundation for all content

**Psychological Impact:**
- Sophisticated and luxurious
- Professional depth
- Makes colors pop
- Creates visual hierarchy through contrast

---

### White (#FFFFFF)
**Where Used:**
- Primary body text
- User messages
- High-priority content

**Psychological Impact:**
- Maximum clarity
- Clean and professional
- Easy to read
- Truth and transparency

---

## Visual Rhythm & Flow

The interface creates a visual rhythm through:

1. **Alternating Accents**
   - Hot pink and gold alternate in lists
   - Creates visual interest without chaos
   - Guides the eye through content

2. **Breathing Space**
   - Black background provides rest for eyes
   - Generous padding around content
   - Separators create clear sections

3. **Hierarchy Through Color**
   - Level 1 (Headers): Hot pink
   - Level 2 (Subheaders): Gold
   - Level 3 (Body): White
   - Level 4 (Secondary): Light gray
   - Level 5 (Muted): Medium gray

4. **Semantic Clarity**
   - Active/Primary: Hot pink
   - Success/Value: Gold
   - Neutral: White
   - Secondary: Gray
   - Background: Black

---

## The Emotional Journey

### First Launch
"Wow, that's bold but sophisticated."

### First Interaction
"I can immediately see what's important."

### Continued Use
"This feels premium and professional."

### Brand Recognition
"That's definitely Danni."

---

## Implementation Notes

### Terminal Requirements
- True color support (24-bit) for full aesthetic
- Falls back gracefully to 256 colors
- 16-color ANSI fallback maintains hierarchy

### Font Recommendations
- **Recommended:** Berkeley Mono (as configured)
- **Alternatives:** JetBrains Mono, Fira Code, SF Mono
- Nerd Font support for symbols

### Performance
- Styled strings cached where possible
- Minimal overhead from color rendering
- Smooth 60fps scrolling maintained

---

## Conclusion

This mockup demonstrates how the Hot Pink + Gold + Black aesthetic creates:

✓ **Immediate Impact** - Bold first impression
✓ **Clear Hierarchy** - Easy to scan and understand
✓ **Warm Sophistication** - Professional yet approachable
✓ **Brand Distinction** - Unmistakably Danni
✓ **Elegant Function** - Beauty serves usability

The transformation isn't just visual—it's experiential. Every color choice, every symbol, every spacing decision whispers: "This is a world-class AI strategist who understands both power and grace."

---

**Design Philosophy:**
*"The space between elements breathes as loudly as the elements themselves. Hot pink commands. Gold rewards. Black provides the stage. Together, they create an experience that feels both supremely competent and genuinely warm."*
