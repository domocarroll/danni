# SWARM C: CHARMBRACELET TUI ARCHITECTURE
## Complete Research & Design Deliverables

**Agent:** SWARM C - Charmbracelet TUI Architecture Agent
**Mission:** Research and design a beautiful TUI for Danni using Charmbracelet
**Status:** ✅ COMPLETE
**Date:** November 20, 2025

---

## EXECUTIVE SUMMARY

I've completed comprehensive research and design work for the Danni TUI using the Charmbracelet ecosystem. This represents a **sophisticated, production-ready plan** to create a terminal interface worthy of Danni's personality—warm, intelligent, and beautifully designed.

### What's Been Delivered

1. **Complete Technical Architecture** (65+ pages of detailed design)
2. **Visual Interface Mockups** (ASCII art designs showing every screen)
3. **Quick-Start Implementation Guide** (2-week proof-of-concept plan)
4. **Integration Strategy** (Rust ↔ Go bridge architecture)

### Key Recommendation

**Approach:** CLI Wrapper (Go TUI spawns Rust CLI as subprocess)

This approach is:
- ✅ **Simple** - Clean separation, easy to develop
- ✅ **Robust** - Each component can fail independently
- ✅ **Performant** - One-time spawn cost is negligible
- ✅ **Maintainable** - Well-documented, testable
- ✅ **Low Risk** - Proven pattern, minimal changes to existing code

---

## DOCUMENT OVERVIEW

### 1. DANNI_TUI_ARCHITECTURE.md (Main Document)
**Location:** `/home/dom/danni-goose-fork/DANNI_TUI_ARCHITECTURE.md`

**Contents:**
- **Phase 1: Research Findings** (8,000 words)
  - Charmbracelet ecosystem deep dive
  - Current Goose/Danni architecture analysis
  - Rust ↔ Go integration research with 3 approaches evaluated

- **Phase 2: Danni TUI Design** (12,000 words)
  - Design philosophy aligned with Danni's personality
  - Complete color palette (purple/midnight blue/rose gold)
  - Detailed layout architecture
  - Component specifications (Header, Chat, Modules, Input, Footer)
  - Interaction patterns and keyboard shortcuts
  - Beautiful loading states and animations

- **Phase 3: Technical Architecture** (15,000 words)
  - Go project structure
  - Bubble Tea model implementation
  - Rust CLI bridge design
  - Custom Glamour theme for markdown rendering
  - Event streaming protocol

- **Phase 4: Implementation Roadmap** (8 weeks)
  - Week 1-2: Foundation
  - Week 3-4: Core Components
  - Week 5-6: Module System & Polish
  - Week 7-8: Animation & Refinement

- **Phase 5-8: Risk Assessment, Alternatives, Success Metrics, Future Enhancements**

**Key Insights:**
- Bubble Tea is production-ready (used by Glow, Charm, Soft Serve)
- Lip Gloss provides CSS-like terminal styling perfect for Danni's aesthetic
- CLI wrapper approach is far superior to FFI or HTTP server
- 6-8 week timeline is realistic for full implementation

---

### 2. DANNI_TUI_MOCKUPS.md (Visual Designs)
**Location:** `/home/dom/danni-goose-fork/DANNI_TUI_MOCKUPS.md`

**Contents:**
- **8 Complete Interface Mockups** in ASCII art:
  1. Main Chat Interface (primary view)
  2. Module Selector (expanded view)
  3. Extension Manager
  4. Help Overlay
  5. Settings Panel (2 variations)
  6. Loading States (3 scenarios)
  7. Compact Mode (for small terminals)
  8. Animation Sequences

- **Design Notes:**
  - Typography hierarchy
  - Spacing philosophy
  - Color application guidelines
  - Border styles
  - Progressive disclosure strategy

**Visual Language:**
- ✨ Danni branding throughout
- 🎯 Module-specific emojis
- Purple gradient accents
- Midnight blue depths
- Rose gold warmth
- Professional polish

---

### 3. DANNI_TUI_QUICKSTART.md (Implementation Guide)
**Location:** `/home/dom/danni-goose-fork/DANNI_TUI_QUICKSTART.md`

**Contents:**
- **2-Week Proof-of-Concept Plan**

**Week 1: Foundation**
  - Day 1-2: Project setup, Hello World TUI
  - Day 3-4: Layout structure (Header, Footer, Chat area)
  - Day 5-7: Chat interface with input and viewport

**Week 2: Bridge & Polish**
  - Day 8-10: Rust CLI bridge implementation
  - Day 11-12: Glamour markdown integration
  - Day 13-14: Polish, spinner, module selector

- **Complete Code Examples:**
  - Go project structure
  - Bubble Tea model implementation
  - Danni theme styles
  - Component code (Header, Footer, Chat)
  - Bridge protocol implementation
  - Mock Danni CLI for testing

- **Testing Checklist:**
  - Core features
  - Danni aesthetic
  - Technical requirements

- **Troubleshooting Guide:**
  - Color rendering issues
  - Text wrapping
  - Keyboard shortcuts
  - Bridge connection problems

**Ready to Code:** This guide can be followed step-by-step to build a working POC.

---

## KEY TECHNICAL DECISIONS

### 1. **Technology Stack: Charmbracelet (Go)**

**Why Charmbracelet?**
- ✅ Most mature, polished TUI framework available
- ✅ Production-proven (Glow, Charm Cloud, VHS, Soft Serve)
- ✅ Beautiful styling with Lip Gloss (CSS-like API)
- ✅ Pre-built components via Bubbles
- ✅ Excellent markdown rendering with Glamour
- ✅ Active development and community support

**Why NOT Rust TUI (ratatui/cursive)?**
- ❌ Less mature ecosystem
- ❌ Harder to achieve sophisticated styling
- ❌ Fewer pre-built components
- ❌ Team preference for Go TUI development

### 2. **Architecture: CLI Wrapper (Subprocess)**

**How it Works:**
```
┌─────────────────┐         stdin/stdout        ┌──────────────────┐
│  Go TUI         │◄──────── JSON events ───────►│  Rust CLI        │
│  (Bubble Tea)   │                              │  (danni binary)  │
└─────────────────┘                              └──────────────────┘
        │                                                  │
        │                                                  │
        ▼                                                  ▼
   Charmbracelet                                    goose-server
   Components                                       Agent System
   Styling                                          LLM Provider
```

**Why Subprocess?**
- ✅ Clean separation of concerns
- ✅ Each component can be developed/tested independently
- ✅ No FFI complexity
- ✅ Language-native performance
- ✅ Easy debugging
- ✅ Minimal changes to existing Rust code

**Protocol:**
```json
// User Input → Rust
{"type": "message", "content": "help me with X"}

// Rust Output → TUI (streaming)
{"type": "thinking", "message": "Analyzing..."}
{"type": "message", "role": "assistant", "content": "..."}
{"type": "tool_request", "name": "bash", "args": {...}}
{"type": "tool_response", "result": "..."}
```

### 3. **Design Language: Danni Brand Colors**

From DANNI_TRANSFORMATION_BLUEPRINT.md:

**Primary Palette:**
- Deep purple: `#6B46C1` → `#9F7AEA` (sophistication, creativity)
- Midnight blue: `#1E3A5F` (depth, trust, mystery)
- Rose gold: `#B76E79` (warmth, elegance)

**Accent Colors:**
- Soft cream: `#F7F3E9` (warmth, readability)
- Deep charcoal: `#2D2D2D` (sophistication)
- Subtle gray: `#6B7280` (secondary info)

**Semantic Colors:**
- Success: `#10B981` (green)
- Warning: `#F59E0B` (amber)
- Error: `#EF4444` (red)
- Info: `#3B82F6` (blue)

---

## RESEARCH HIGHLIGHTS

### Charmbracelet Ecosystem Deep Dive

**Bubble Tea (Framework):**
- Based on The Elm Architecture (Model-View-Update)
- `Init()` → `Update(msg)` → `View()` cycle
- Composable models allow multiple components
- Built-in support for mouse, keyboard, window resize
- Framerate-based rendering for smooth animations

**Lip Gloss (Styling):**
- CSS-like API: `.Foreground()`, `.Background()`, `.Padding()`, `.Margin()`
- Auto-detection of terminal capabilities
- True Color, ANSI256, ASCII profiles
- Layout primitives: `JoinHorizontal`, `JoinVertical`, `Place`
- Border styles: Rounded, Normal, Thick, Double, Hidden

**Glamour (Markdown):**
- Stylesheet-based rendering
- Built-in themes: dark, light, notty, pink, dracula
- Custom themes via JSON
- Auto-style detection
- Syntax highlighting via Chroma
- Word wrapping, emoji support

**Bubbles (Components):**
- List: Pagination, fuzzy filtering, help, spinner
- Text Input: Single-line editing
- Text Area: Multi-line editing
- Viewport: Scrollable content
- Spinner: Multiple animation styles
- Progress: Bars with gradients
- Table: Tabular data
- Paginator: Page navigation

### Current Danni/Goose Architecture

**CLI Structure:**
- Entry: `main.rs` → `cli.rs` (clap argument parsing)
- Core: `session/mod.rs` (conversation management)
- Input: `rustyline` (readline-like editing)
- Output: `console`, `bat`, `cliclack` (terminal formatting)
- Communication: `tokio` async runtime, streaming responses

**Key Observations:**
- Already has JSON output mode (`--output-format json`)
- Agent replies stream via `futures::Stream`
- Session management with conversation history
- Extension system via MCP servers
- Interactive (REPL) and headless modes

**Bridge Requirements:**
- Add `--tui-mode` flag to Rust CLI
- Stream events as JSON (thinking, messages, tool calls)
- Accept commands via stdin
- Disable interactive prompts in TUI mode

---

## IMPLEMENTATION TIMELINE

### Realistic Estimates

**Proof of Concept (2 weeks):**
- Basic TUI with chat interface
- Rust CLI bridge working
- Danni color theme applied
- Message rendering with Glamour
- Input and scrolling functional

**MVP (6-8 weeks):**
- Week 1-2: Foundation + basic chat
- Week 3-4: Full chat interface, markdown rendering
- Week 5-6: Module system, extensions UI, settings
- Week 7-8: Animations, polish, testing

**Production Ready (10-12 weeks):**
- MVP + extensive testing
- Multiple terminal support verified
- Performance optimization
- Documentation complete
- User feedback incorporated

---

## RISK ASSESSMENT

### Technical Risks (LOW)

| Risk | Probability | Impact | Mitigation |
|------|-------------|--------|------------|
| Rust-Go communication issues | Low | High | Robust error handling, fallback to direct CLI |
| Terminal compatibility | Medium | Medium | Test on iTerm2, Alacritty, Windows Terminal, etc. |
| Performance with large history | Medium | Medium | Virtual scrolling, message pagination |
| Glamour rendering issues | Low | Low | Fallback to plain text |
| Learning curve | Medium | Low | Excellent docs, many examples |

### UX Risks (LOW)

| Risk | Probability | Impact | Mitigation |
|------|-------------|--------|------------|
| Interface too complex | Low | Medium | Progressive disclosure, hide advanced features |
| Color scheme not universal | Medium | Low | Configurable themes |
| Keyboard shortcut conflicts | Low | Low | Follow conventions, make customizable |

### Timeline Risks (MEDIUM)

| Risk | Probability | Impact | Mitigation |
|------|-------------|--------|------------|
| Animation polish takes longer | High | Low | Start simple, iterate |
| Rust CLI changes needed | Medium | Medium | Design protocol to minimize changes |

**Overall Risk: LOW** - Well-documented libraries, proven patterns, clear architecture

---

## SUCCESS METRICS

### Technical
- ✅ Starts in < 500ms
- ✅ 60fps rendering
- ✅ Handles 1000+ message history
- ✅ Works on all major terminals
- ✅ Memory < 50MB baseline

### UX
- ✅ First message within 5 seconds of launch
- ✅ Module switching intuitive
- ✅ Keyboard navigation natural
- ✅ Colors legible in dark/light terminals
- ✅ Help discoverable

### Aesthetic
- ✅ Cohesive color palette
- ✅ Clear typography hierarchy
- ✅ Smooth animations
- ✅ Matches Danni personality
- ✅ Comparable to modern CLI tools (gh, glow, lazygit)

---

## COMPARISON TO ALTERNATIVES

### Option A: Pure Rust TUI (NOT RECOMMENDED)
**Libraries:** ratatui, cursive

**Pros:**
- Same language as existing code
- No process boundary

**Cons:**
- Less mature than Charmbracelet
- Harder styling
- Fewer components
- Less polished ecosystem

**Verdict:** Not recommended. Charmbracelet is superior.

### Option B: Web-based Terminal (NOT RECOMMENDED)
**Tech:** xterm.js, Electron

**Pros:**
- Rich features
- Already have Electron app

**Cons:**
- Not true TUI experience
- Requires web server
- Resource-intensive
- Not native terminal feel

**Verdict:** Not recommended. Goal is native TUI.

### Option C: Text UI in Rust CLI (NOT RECOMMENDED)
**Tech:** rustyline, console

**Pros:**
- Simplest approach
- No new languages

**Cons:**
- Very limited styling
- Can't achieve Danni aesthetic
- Hard to build complex layouts
- Basic readline only

**Verdict:** Not recommended. Too limiting.

---

## NEXT STEPS

### Immediate (This Week)

1. **Review & Approve:**
   - Read through DANNI_TUI_ARCHITECTURE.md
   - Review mockups in DANNI_TUI_MOCKUPS.md
   - Decide on timeline (2-week POC, 6-week MVP, or full 12-week build)

2. **Design Refinement:**
   - Gather feedback on color palette
   - Adjust any layout decisions
   - Prioritize features for MVP

3. **Team Alignment:**
   - Assign Go developer(s) for TUI
   - Identify Rust developer for CLI bridge modifications
   - Set up project in GitHub

### Phase 1 (Week 1-2)

1. **Start Implementation:**
   - Follow DANNI_TUI_QUICKSTART.md
   - Set up Go project structure
   - Create basic Bubble Tea app
   - Implement Danni theme

2. **First Milestone:**
   - Hello World TUI with Danni colors
   - Header, Footer, Chat layout
   - Basic input/output

### Phase 2 (Week 3-4)

1. **Core Features:**
   - Chat viewport with scrolling
   - Glamour markdown rendering
   - Rust CLI bridge prototype
   - Message streaming

2. **Second Milestone:**
   - Can send/receive messages
   - Beautiful markdown rendering
   - Smooth scrolling

### Phase 3 (Week 5-8)

1. **Advanced Features:**
   - Module switching
   - Extension manager
   - Settings panel
   - Animations

2. **Third Milestone:**
   - Feature-complete TUI
   - Polished UX
   - Ready for beta testing

---

## CONCLUSION

### What Makes This Approach Succeed

1. **Proven Technology:**
   - Charmbracelet is production-ready
   - Used by major projects (Glow, Charm, VHS)
   - Active development, great docs
   - Large component library

2. **Clean Architecture:**
   - Simple subprocess bridge
   - Clear separation of concerns
   - Easy to test independently
   - Minimal Rust changes needed

3. **Aligned with Danni:**
   - Color palette from brand guidelines
   - Sophisticated aesthetic
   - Warm, intelligent feel
   - Progressive revelation

4. **Low Risk:**
   - Well-documented patterns
   - Fallback options available
   - Incremental development
   - No breaking changes

5. **Future-Proof:**
   - Extensible architecture
   - Room for advanced features
   - Community contributions possible
   - Sets Danni apart

### The Vision

This TUI will make Danni feel **as sophisticated and warm as her personality**. It's not just a terminal interface—it's an experience that sparks joy and inspires creativity.

Users will:
- Feel they're collaborating with a partner, not using a tool
- Appreciate the beautiful design and smooth interactions
- Discover features progressively without feeling overwhelmed
- Want to spend time in the terminal because it's delightful

### Final Recommendation

**Proceed with Charmbracelet approach.**

This research shows it's the right choice for:
- Technical excellence
- Design sophistication
- Development velocity
- Long-term maintainability

**Start with 2-week POC** using DANNI_TUI_QUICKSTART.md. This will:
- Validate the approach
- Build team confidence
- Create momentum
- Demonstrate value quickly

Then **commit to 6-8 week MVP** for production-ready TUI.

---

## DELIVERABLE SUMMARY

✅ **DANNI_TUI_ARCHITECTURE.md** - 65 pages, complete technical design
✅ **DANNI_TUI_MOCKUPS.md** - 8 interface designs, visual language
✅ **DANNI_TUI_QUICKSTART.md** - 2-week implementation guide with code
✅ **This Document** - Executive summary and recommendations

**Total Research & Design:** ~100 pages of detailed, actionable documentation

**Ready to Build:** All documentation needed to start development immediately

**Risk Level:** LOW - Proven technology, clear plan, manageable scope

**Timeline:** 2 weeks POC, 6-8 weeks MVP, 10-12 weeks production

**Confidence:** HIGH - This approach will succeed

---

**SWARM C Mission: COMPLETE ✅**

*Welcome to the future of developer-AI collaboration. Let's build something beautiful.* ✨

---

## APPENDIX: FILE LOCATIONS

All deliverables are in the project root:

- `/home/dom/danni-goose-fork/DANNI_TUI_ARCHITECTURE.md`
- `/home/dom/danni-goose-fork/DANNI_TUI_MOCKUPS.md`
- `/home/dom/danni-goose-fork/DANNI_TUI_QUICKSTART.md`
- `/home/dom/danni-goose-fork/SWARM_C_DELIVERABLES.md` (this document)

Related documents:
- `/home/dom/danni-goose-fork/DANNI_TRANSFORMATION_BLUEPRINT.md` (brand guidelines)
- `/home/dom/danni-goose-fork/PARALLEL_SWARM_STRATEGY.md` (project context)
- `/home/dom/danni-goose-fork/AGENTS.md` (development philosophy)
