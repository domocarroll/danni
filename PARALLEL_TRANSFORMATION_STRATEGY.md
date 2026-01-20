# PARALLEL TRANSFORMATION STRATEGY
## Comprehensive Danni Rebranding - 6 Parallel Workstreams

*Generated: 2025-11-24*

---

## EXECUTIVE SUMMARY

The reconnaissance phase has revealed a **complete, production-ready Electron application** requiring systematic transformation across 6 parallel dimensions:

| Workstream | Scope | Impact | Files | Changes |
|-----------|-------|--------|-------|---------|
| **A: Code Rebranding** | Component/file renames | High | 40+ | 684 refs |
| **B: Design System** | Visual aesthetic transformation | Critical | 15+ | Complete overhaul |
| **C: Copy & Messaging** | Voice transformation | High | 30+ | 80+ items |
| **D: Asset Pipeline** | Visual assets replacement | Critical | 50+ | All icons/images |
| **E: Build & Config** | Configuration files | Medium | 10+ | Complete audit |
| **F: Backend Integration** | Daemon/server layer | High | 20+ | 200+ lines |

**Total Estimated Impact:** 165+ files, 1000+ changes

---

## SWARM A: CODE REBRANDING
*Agent: Explore (Code Audit)*

### Critical Findings
- **684 total references** to "goose/Goose/GOOSE"
- **38 protocol references** (goose:// → danni://)
- **13 component files** requiring rename
- **55+ environment variables** in daemon file

### File Renames Required

```
CRITICAL PATH (Direct User Visibility):
├── src/components/GooseLogo.tsx → DanniLogo.tsx
├── src/components/GooseMessage.tsx → DanniMessage.tsx
├── src/components/LoadingGoose.tsx → LoadingDanni.tsx
├── src/components/GoosehintsModal.tsx → DanniInsightsModal.tsx
├── src/components/WelcomeGooseLogo.tsx → WelcomeDanniLogo.tsx
├── src/components/icons/Goose.tsx → Danni.tsx
├── src/goosed.ts → dannid.ts
└── src/components/GooseSidebar/ → DanniSidebar/
```

### Protocol Transformation
```typescript
// 38 instances across 10 files
goose://sessions/{shareToken} → danni://sessions/{shareToken}
goose://recipe?config=...     → danni://recipe?config=...
goose://extension?cmd=...     → danni://extension?cmd=...
```

### Environment Variables
```bash
GOOSE_PORT                    → DANNI_PORT
GOOSE_SERVER__SECRET_KEY      → DANNI_SERVER__SECRET_KEY
GOOSE_EXTERNAL_BACKEND        → DANNI_EXTERNAL_BACKEND
GOOSE_DEFAULT_PROVIDER        → DANNI_DEFAULT_PROVIDER
GOOSE_DEFAULT_MODEL           → DANNI_DEFAULT_MODEL
# ... 9 total unique variables
```

### Execution Strategy
**Phase 1:** Foundation (no dependencies)
- Rename `icons/Goose.tsx` first
- Update all 13 imports

**Phase 2:** Core Components
- Logo, Message, Loading components
- Update cascade: ProgressiveMessageList, BaseChat, BaseChat2

**Phase 3:** Sidebar & Settings
- Rename directory
- Update import paths

**Phase 4:** Backend Integration
- Rename goosed.ts → dannid.ts
- Update main.ts import

**Phase 5:** Protocols
- Update all goose:// → danni:// across 10 files

---

## SWARM B: DESIGN SYSTEM TRANSFORMATION
*Agent: goose-aesthetic-architect*

### Color Palette Revolution

```css
/* Danni Signature Colors */
--color-danni-pink-vivid: #FF1493;      /* DeepPink - Primary brand */
--color-danni-pink-bright: #FF69B4;     /* HotPink - Interactive states */
--color-danni-pink-soft: #FFB6D9;       /* Soft pink - Backgrounds */

--color-danni-gold-vivid: #FFD700;      /* Gold - Premium accents */
--color-danni-gold-warm: #FFA500;       /* Orange-gold - Hover states */
--color-danni-gold-subtle: #FFE4B3;     /* Pale gold - Highlights */

--color-danni-black-deep: #000000;      /* Pure black - Contrast */
--color-danni-black-rich: #1a1a1a;      /* Rich black - Backgrounds */
--color-danni-black-soft: #2d2d2d;      /* Soft black - Cards */
```

### Typography Elevation

```css
/* Explicit Hierarchy */
.text-display      → 36px, bold, tight tracking
.text-heading-1    → 30px, bold, tight tracking
.text-heading-2    → 24px, medium, snug leading
.text-heading-3    → 20px, medium, normal leading
.text-body-lg      → 18px, regular, relaxed leading
.text-body         → 16px, regular, normal leading
.text-label        → 14px, medium, wide tracking, uppercase
.text-caption      → 12px, regular, muted color
```

### Component Redesign Priorities

**Phase 1: Foundation (Week 1)**
1. Color System (main.css)
2. Typography Scale
3. Button Components (gradients, shadows)
4. Card Components (hover states, premium variants)

**Phase 2: Core Experience (Week 2)**
5. Loading States (shimmer, glow effects)
6. Modals & Dialogs (backdrop blur, animations)
7. Input Components (pink accent focus states)
8. Navigation Sidebar (gradient active states)

**Phase 3: Delight (Week 3)**
9. Danni Logo (geometric D with neural network)
10. Micro-Interactions (sparkle, ripple, magnetic effects)
11. Page Transitions (enhanced routing)
12. Easter Eggs (hidden goose appearances)

### Animation Signatures

```css
/* Signature Easing */
--ease-danni-smooth: cubic-bezier(0.4, 0.0, 0.2, 1);
--ease-danni-bounce: cubic-bezier(0.68, -0.55, 0.265, 1.55);
--ease-danni-elegant: cubic-bezier(0.16, 1, 0.3, 1);

/* New Animations */
@keyframes shimmer-danni { ... }
@keyframes pulse-glow-danni { ... }
@keyframes sparkle { ... }
```

---

## SWARM C: COPY & MESSAGING TRANSFORMATION
*Agent: Explore (Text Audit)*

### Critical Text Transformations

#### Error Messages (12 instances)
```
BEFORE: "Honk! Goose experienced an error while responding"
AFTER:  "I sense there was an interruption in my response. Allow me to try again."

BEFORE: "Failed to Load Session"
AFTER:  "Unable to Retrieve Session"

BEFORE: "Image too large (X MB). Maximum 5MB allowed."
AFTER:  "This image exceeds our format constraints. Please select an image under 5MB."
```

#### Loading States (11 instances)
```
BEFORE: "goose is thinking…"
AFTER:  "I'm considering the possibilities…"

BEFORE: "goose is working on it…"
AFTER:  "Let me craft a thoughtful response…"

BEFORE: "goose is compacting the conversation..."
AFTER:  "Organizing our conversation for continuity…"
```

#### Button Labels (11 instances)
```
BEFORE: "Allow Once" | "Always Allow" | "Deny"
AFTER:  "Permit This Time" | "Always Permit" | "Decline"

BEFORE: "Ask goose"
AFTER:  "Get Assistance" or "Seek Guidance"
```

#### Placeholders (2 instances)
```
BEFORE: "Ask goose anything..."
AFTER:  "What's on your mind?" or "Share your vision…"
```

#### Greeting Messages (19 variants)
```
BEFORE: "Hello! Ready to get started?"
AFTER:  "Welcome. What discovery awaits?"

BEFORE: "What would you like to work on?"
AFTER:  "What challenges shall we address?"
```

### Voice Guidelines
- **Sophisticated but not pretentious**
- **Warm professionalism**
- Signature phrases: "I sense...", "What intrigues me most...", "Shall we..."
- **No "Honk!"** - elegant observations instead
- **Removes playfulness, adds mystique**

---

## SWARM D: ASSET PIPELINE & BRANDING
*Agent: Explore (Asset Inventory)*

### Asset Replacement Checklist

#### Priority 1: Critical Brand Assets
- [ ] Design new master icon SVG (2048x2048)
  - Geometric "D" logomark with neural network hints
  - Solid color version
  - Template version (white on transparent)
- [ ] Generate icon.png (1024x1024)
- [ ] Generate icon@2x.png (2048x2048)
- [ ] Generate icon.ico (Windows multi-res)
- [ ] Generate icon.icns (macOS multi-res)
- [ ] Create template icons (18x18, 36x36 for menu bar)

#### Priority 2: Branding Assets
- [ ] Replace block-lockup_black.png with Danni branding
- [ ] Replace block-lockup_white.png with Danni branding

#### Priority 3: Loading/Animation Assets
- [ ] Design 7-frame Danni loading animation
- [ ] Regenerate loading-danni SVG files (1-7)

#### Priority 4: Game Assets (Optional)
- [ ] Update battle-game/goose.png → danni.png
- [ ] Update background.png
- [ ] Update/replace llama.png

### Asset Locations
```
src/images/
├── icon.svg, icon.png, icon@2x.png
├── icon.ico, icon.icns (platform-specific)
├── icon-light.png, icon-light.icns
├── iconTemplate.png, iconTemplate@2x.png (menu bar)
├── iconTemplateUpdate.png, iconTemplateUpdate@2x.png
├── glyph.svg (master character vector)
└── loading-goose/1-7.svg → loading-danni/1-7.svg

src/components/settings/app/icons/
├── block-lockup_black.png → danni-lockup_black.png
└── block-lockup_white.png → danni-lockup_white.png

src/assets/
├── clock-icon.svg (review for theme consistency)
└── battle-game/* (optional updates)
```

### Technical Specifications
- **SVG Master:** 2048x2048 viewBox
- **PNG Exports:** 1024x1024 (1x), 2048x2048 (2x)
- **ICO Bundle:** 16, 32, 48, 64, 256px
- **ICNS Bundle:** 512x512pt @ 1x and 2x
- **Template Icons:** Pure white + alpha transparency

---

## SWARM E: BUILD & CONFIGURATION
*Agent: architecture-strategist (FAILED - needs manual audit)*

### Files Requiring Updates

#### Package Configuration
- [x] package.json (already correct: "productName": "Danni")
- [ ] forge.config.ts (verify protocol: 'danni')
- [ ] Desktop entry files (forge.deb.desktop, forge.rpm.desktop)

#### Linux Desktop Integration
```ini
# forge.deb.desktop & forge.rpm.desktop
Name=Goose → Name=Danni
Exec=/usr/lib/goose/Goose → Exec=/usr/lib/danni/Danni
Icon=/usr/share/pixmaps/goose.png → Icon=/usr/share/pixmaps/danni.png
MimeType=x-scheme-handler/goose → MimeType=x-scheme-handler/danni
```

#### Build Scripts
- [ ] scripts/prepare-platform-binaries.js (update binary patterns)
- [ ] Update any CI/CD environment variables
- [ ] Update GitHub Actions (if present in ../.github/workflows/)

---

## SWARM F: BACKEND INTEGRATION
*Agent: Explore (Backend Audit)*

### Critical Refactoring

#### File Rename
```
src/goosed.ts → src/dannid.ts
```

#### Function Renames
```typescript
startGoosed() → startDannid()
getGoosedBinaryPath() → getDannidBinaryPath()
checkServerStatus() → // references to "goosed" updated
```

#### Binary Search
```typescript
// Line 222 in goosed.ts
let executableName = process.platform === 'win32' ? 'goosed.exe' : 'goosed';
// BECOMES:
let executableName = process.platform === 'win32' ? 'dannid.exe' : 'dannid';
```

#### Environment Variables (9 unique)
```bash
GOOSE_PORT                    → DANNI_PORT
GOOSE_SERVER__SECRET_KEY      → DANNI_SERVER__SECRET_KEY
GOOSE_EXTERNAL_BACKEND        → DANNI_EXTERNAL_BACKEND
GOOSE_DEFAULT_PROVIDER        → DANNI_DEFAULT_PROVIDER
GOOSE_DEFAULT_MODEL           → DANNI_DEFAULT_MODEL
GOOSE_PREDEFINED_MODELS       → DANNI_PREDEFINED_MODELS
GOOSE_BASE_URL_SHARE          → DANNI_BASE_URL_SHARE
GOOSE_VERSION                 → DANNI_VERSION
GOOSE_API_HOST                → DANNI_API_HOST
GOOSE_WORKING_DIR             → DANNI_WORKING_DIR
GOOSE_ALLOWLIST_WARNING       → DANNI_ALLOWLIST_WARNING
GOOSE_SERVER__MEMORY          → DANNI_SERVER__MEMORY
GOOSE_SERVER__COMPUTER_CONTROLLER → DANNI_SERVER__COMPUTER_CONTROLLER
```

#### Protocol Registration
```typescript
// main.ts line 148, 163
app.setAsDefaultProtocolClient('goose') → app.setAsDefaultProtocolClient('danni')
```

#### Persistent Storage Paths
```typescript
// Session persistence (MIGRATION RISK)
partition: 'persist:goose' → partition: 'persist:danni'

// Temp directory
const gooseTempDir = path.join(app.getPath('temp'), 'goose-pasted-images');
// BECOMES:
const danniTempDir = path.join(app.getPath('temp'), 'danni-pasted-images');
```

#### IPC Handlers
```typescript
// preload.ts line 185
getGoosedHostPort → getDannidHostPort
```

### Architectural Risks
1. **Session Persistence:** Changing partition will invalidate existing user sessions
2. **Protocol Conflict:** Old goose:// links won't work after update
3. **Temp Directory Migration:** Old pasted images won't migrate automatically
4. **Binary Name Verification:** Must confirm backend binary is actually renamed to 'dannid'

---

## PARALLEL EXECUTION PLAN

### Week 1: Foundation & Core
**Execute in Parallel:**
- SWARM A Phase 1-2 (icon renames, core components)
- SWARM B Phase 1 (color system, typography, buttons)
- SWARM D Priority 1 (design & generate master icons)

**Sequential Dependencies:**
- Can't apply colors until CSS variables defined
- Can't update components until icons renamed

### Week 2: Integration & Polish
**Execute in Parallel:**
- SWARM A Phase 3-4 (sidebar, backend integration)
- SWARM B Phase 2 (loading states, modals, inputs)
- SWARM C Phase 1-2 (critical text, loading states)
- SWARM F Phase 1-3 (daemon rename, env vars, protocols)

**Sequential Dependencies:**
- Backend protocol changes before protocol text updates
- Component renames before copy updates

### Week 3: Assets & Testing
**Execute in Parallel:**
- SWARM A Phase 5-8 (protocols, tests, assets)
- SWARM B Phase 3 (micro-interactions, easter eggs)
- SWARM C Phase 3 (remaining text transformations)
- SWARM D Priority 2-4 (branding assets, animations)
- SWARM E (build configuration updates)

**Sequential Dependencies:**
- All code changes before build testing
- Asset generation before asset replacement

---

## SUCCESS METRICS

### Code Quality
- [ ] Zero TypeScript errors
- [ ] All tests passing
- [ ] No console warnings
- [ ] Linter passing

### Visual Quality
- [ ] Color palette consistently applied
- [ ] Typography hierarchy clear and elegant
- [ ] All hover states working
- [ ] Animations smooth and performant
- [ ] Icons display correctly at all DPIs

### Functional Quality
- [ ] App launches successfully
- [ ] Deep links work (danni://)
- [ ] Sessions persist correctly
- [ ] Extensions install properly
- [ ] All platforms build (macOS, Windows, Linux)

### Brand Quality
- [ ] No "Goose" references visible to users
- [ ] Voice consistently sophisticated
- [ ] Visual identity cohesive and distinctive
- [ ] Premium feel throughout experience

---

## ROLLBACK STRATEGY

If critical issues arise:
1. **Git Branch:** All work on `danni-transformation` branch
2. **Commit Checkpoints:** After each swarm phase completion
3. **Testing Gates:** Don't proceed to next week without passing tests
4. **Asset Backups:** Keep original goose assets in `src/images/legacy/`
5. **Config Backup:** Keep old env var names as fallbacks during transition

---

## NEXT STEPS

1. **Confirm Strategy:** Review this document and approve approach
2. **Create Branch:** `git checkout -b danni-transformation`
3. **Launch Swarms:** Deploy parallel agents to execute phases
4. **Monitor Progress:** Use todo tracking to monitor all 6 swarms
5. **Integration Testing:** After each week, test integrated changes
6. **User Acceptance:** Final review before merge to main

---

*"The space between playful and professional holds the essence of Danni's transformation..."*
