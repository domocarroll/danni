# Danni Terminal Release Strategy

## Option B: Pre-Built Binary Releases

This document explains how to distribute Danni Terminal using GitHub Releases with pre-compiled binaries.

---

## Why Option B?

| Aspect | Option A (Source) | Option B (Binary) |
|--------|------------------|-------------------|
| **User Needs** | Git, Go, Rust | Nothing (just Ghostty) |
| **Install Time** | 5-10 minutes | 30 seconds |
| **User Skill** | Developer | Anyone |
| **Download Size** | Full repo (~50MB) | ~5MB per platform |
| **Experience** | Good | **Excellent** |

**Bottom line:** Option B is vastly better for end users while only slightly more work for you.

---

## How It Works

### The Workflow

```
┌─────────────────────────────────────────────────┐
│ 1. You push a version tag                      │
│    $ git tag v0.1.0                            │
│    $ git push origin v0.1.0                    │
└────────────────┬────────────────────────────────┘
                 ↓
┌─────────────────────────────────────────────────┐
│ 2. GitHub Actions automatically:               │
│    • Checks out code                           │
│    • Builds for macOS (Intel + Apple Silicon)  │
│    • Builds for Linux (x64 + ARM64)            │
│    • Runs tests                                │
│    • Creates tarballs                          │
│    • Publishes GitHub Release                  │
│                                                 │
│    Takes ~10-15 minutes, zero effort from you  │
└────────────────┬────────────────────────────────┘
                 ↓
┌─────────────────────────────────────────────────┐
│ 3. Release appears on GitHub                   │
│    https://github.com/you/danni/releases       │
│                                                 │
│    Contains:                                   │
│    • danni-terminal-macos-arm64.tar.gz         │
│    • danni-terminal-macos-x64.tar.gz           │
│    • danni-terminal-linux-x64.tar.gz           │
│    • danni-terminal-linux-arm64.tar.gz         │
│    • Release notes (auto-generated)            │
└────────────────┬────────────────────────────────┘
                 ↓
┌─────────────────────────────────────────────────┐
│ 4. Users install with one command              │
│    $ curl -fsSL https://danni.sh/install | sh  │
│                                                 │
│    Or manually download & run install.sh       │
└─────────────────────────────────────────────────┘
```

### What's in Each Release Tarball

```
danni-terminal-macos-arm64.tar.gz (5MB)
├── bin/
│   └── danni-tui              # ← Pre-compiled Go binary
├── config/
│   └── ghostty-theme.conf     # ← Danni theme
├── install.sh                 # ← Modified installer
├── uninstall.sh               # ← Uninstaller
└── README.md                  # ← Documentation
```

The `install.sh` is simpler than the source version - it just copies files instead of building.

---

## Implementation

### 1. GitHub Actions Workflow (Already Created)

File: `.github/workflows/release.yml`

**What it does:**
- Triggers on version tags (`v*`) or manual workflow dispatch
- Builds binaries for all platforms in parallel
- Packages everything into tarballs
- Creates GitHub Release with all artifacts
- Adds nice release notes

### 2. Quick Install Script (Already Created)

File: `scripts/install.sh`

**What it does:**
- Detects user's platform automatically
- Downloads correct tarball from latest release
- Extracts and runs the bundled installer
- One-liner installation experience

### 3. Modified Installer

The existing `install-danni-terminal.sh` already handles both cases:
- If binaries are in `bin/` directory → copy them
- If source code is present → build them

No changes needed!

---

## How to Create a Release

### First Release (v0.1.0)

```bash
# 1. Make sure everything is committed
git status

# 2. Create and push a version tag
git tag -a v0.1.0 -m "Initial Danni Terminal release"
git push origin v0.1.0

# 3. Watch GitHub Actions
# Go to: https://github.com/your-username/danni/actions
# The workflow will run automatically (~10-15 mins)

# 4. Check the release
# Go to: https://github.com/your-username/danni/releases
# You should see v0.1.0 with 4 downloadable tarballs
```

### Subsequent Releases

```bash
# Same process with new version
git tag -a v0.2.0 -m "Add new features"
git push origin v0.2.0

# GitHub Actions handles everything else
```

---

## User Experience

### For Power Users (macOS)

```bash
# One-liner install
curl -fsSL https://raw.githubusercontent.com/subfracture/danni/main/scripts/install.sh | sh

# Or with wget
wget -qO- https://raw.githubusercontent.com/subfracture/danni/main/scripts/install.sh | sh

# Done! No build tools needed.
```

### For Cautious Users

```bash
# Download release manually
wget https://github.com/subfracture/danni/releases/download/v0.1.0/danni-terminal-macos-arm64.tar.gz

# Extract
tar xzf danni-terminal-macos-arm64.tar.gz

# Inspect contents
ls -la
cat README.md

# Install when ready
./install.sh
```

### For Package Manager Users (Future)

Eventually you could add:

```bash
# Homebrew (macOS/Linux)
brew install subfracture/tap/danni-terminal

# AUR (Arch Linux)
yay -S danni-terminal

# Nix
nix-env -iA nixpkgs.danni-terminal
```

But starting with GitHub Releases is perfect.

---

## Distribution Channels Comparison

### Option A: Source Only
```
GitHub Repo → User clones → User builds → User installs
(Requires: Git, Go, Rust, 5-10 min build time)
```

### Option B: GitHub Releases (This)
```
GitHub Release → User downloads binary → User installs
(Requires: Nothing, 30 second install)
```

### Option C: Full Fork + Homebrew
```
Ghostty Fork → Custom build → Sign & notarize → Homebrew tap
(Requires: Maintaining fork, code signing, tap repo)
```

**Option B is the sweet spot.**

---

## Advantages of Option B

### For Users
✅ **No build tools required** - Just download and run
✅ **Fast installation** - 30 seconds vs 5-10 minutes
✅ **Consistent experience** - Pre-tested binaries
✅ **Lower barrier to entry** - Non-developers can try it
✅ **Smaller downloads** - 5MB vs 50MB+ repo clone

### For You (Maintainer)
✅ **Automated** - GitHub Actions does all the work
✅ **No fork maintenance** - Still using stock Ghostty
✅ **Version control** - Clear release history
✅ **Analytics** - Can track download counts
✅ **Professional** - Looks like real software product

---

## What Happens When You Create a Release

### The GitHub Release Page Shows:

```
✨ Danni Terminal v0.1.0
────────────────────────────────

The Sophisticated AI Terminal Experience

Installation
────────────

Quick Install (Recommended):
  curl -fsSL https://danni.sh/install | sh

Manual Install:
  1. Download the appropriate tarball below
  2. Extract: tar xzf danni-terminal-*.tar.gz
  3. Run: cd danni-terminal && ./install.sh

Assets
──────

📦 danni-terminal-macos-arm64.tar.gz     5.2 MB   ⬇ 142 downloads
📦 danni-terminal-macos-x64.tar.gz       5.3 MB   ⬇ 89 downloads
📦 danni-terminal-linux-x64.tar.gz       4.8 MB   ⬇ 234 downloads
📦 danni-terminal-linux-arm64.tar.gz     4.7 MB   ⬇ 12 downloads

Source code (zip)
Source code (tar.gz)
```

Users see a professional software release page.

---

## Testing Releases

### Before Your First Release

```bash
# Test the workflow manually
cd .github/workflows
act push -j build-macos-arm64  # Requires 'act' tool

# Or just push a test tag
git tag v0.0.1-test
git push origin v0.0.1-test

# Check if workflow succeeds
# Delete test tag after:
git tag -d v0.0.1-test
git push origin :refs/tags/v0.0.1-test
```

### After Release

```bash
# Download your own release
curl -L https://github.com/YOU/danni/releases/download/v0.1.0/danni-terminal-macos-arm64.tar.gz -o test.tar.gz

# Extract and test
tar xzf test.tar.gz
./install.sh

# Verify it works
danni-terminal
```

---

## Cost

**GitHub Actions minutes:**
- Public repos: ✅ **Free** (2000 minutes/month)
- Private repos: Costs money

**Estimated per release:**
- ~15 minutes for all 4 platforms
- For public repo: **$0.00**

**Storage:**
- GitHub Releases: ✅ **Free** (unlimited)

**Total cost: $0** for public repositories.

---

## Roadmap: Beyond GitHub Releases

Once you have releases working, you can expand to:

### 1. Homebrew Tap (Week 1)
```ruby
# subfracture/homebrew-tap/Formula/danni-terminal.rb
class DanniTerminal < Formula
  desc "Sophisticated AI terminal experience"
  homepage "https://danni.ai"
  url "https://github.com/subfracture/danni/releases/download/v0.1.0/danni-terminal-macos-arm64.tar.gz"
  sha256 "abc123..."

  depends_on "ghostty"

  def install
    bin.install "bin/danni-tui"
    # ...
  end
end
```

Then users: `brew install subfracture/tap/danni-terminal`

### 2. AUR Package (Week 2)
```bash
# Arch Linux user repository
# Creates PKGBUILD that downloads from releases
```

### 3. Debian/RPM Packages (Week 3-4)
```bash
# Convert .tar.gz to .deb and .rpm
# Upload to packagecloud.io or similar
```

### 4. Website with Install Script (Week 5)
```bash
# https://danni.ai → "Get Started" button
# Detects platform, shows install command
```

### 5. Update Checker (Week 6)
```bash
# danni-tui checks for new releases
# Prompts user to upgrade
```

---

## Summary

**Option B gives you:**

1. ✅ **Professional distribution** - Real software product
2. ✅ **User-friendly** - Anyone can install, not just devs
3. ✅ **Automated** - GitHub Actions does the work
4. ✅ **Free** - No cost for public repos
5. ✅ **Scalable** - Foundation for package managers later
6. ✅ **No fork burden** - Still using stock Ghostty

**Next steps:**

```bash
# 1. Push to GitHub (if not already)
git remote add origin https://github.com/subfracture/danni.git
git push -u origin main

# 2. Create first release
git tag -a v0.1.0 -m "Initial release"
git push origin v0.1.0

# 3. Watch the magic happen
# GitHub Actions builds and publishes everything

# 4. Test the install
curl -fsSL https://raw.githubusercontent.com/subfracture/danni/main/scripts/install.sh | sh
```

**That's it. You now have professional software distribution.**

---

Questions? See the workflow file at `.github/workflows/release.yml`
