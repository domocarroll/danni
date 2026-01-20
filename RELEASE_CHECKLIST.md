# Danni Terminal Release Checklist

Quick reference for creating releases when you're ready.

---

## Current Status

**TUI Development:** 🚧 In Progress (Alpha)
**Release System:** ✅ Ready
**Next Milestone:** Complete TUI implementation

---

## Pre-Release Testing (Do This First)

Before your first real release, test the entire workflow:

### 1. Create Test Release

```bash
cd /home/dom/danni-goose-fork

# Make sure everything is committed
git add .
git commit -m "feat: prepare release infrastructure"
git push

# Create test tag (will trigger GitHub Actions)
git tag v0.0.1-test
git push origin v0.0.1-test
```

### 2. Watch GitHub Actions

1. Go to: https://github.com/YOUR-USERNAME/danni/actions
2. You should see "Build and Release Danni Terminal" workflow running
3. Wait ~10-15 minutes for all builds to complete
4. Check for any errors

### 3. Verify Release

1. Go to: https://github.com/YOUR-USERNAME/danni/releases
2. You should see "v0.0.1-test" release
3. Download one of the tarballs (your platform)
4. Test the installation:

```bash
# Extract
tar xzf danni-terminal-*.tar.gz
cd danni-terminal

# Check contents
ls -la
# Should see: bin/, config/, install.sh, uninstall.sh, README.md

# Test install
./install.sh
```

### 4. Clean Up Test Release

```bash
# Delete test tag locally and remotely
git tag -d v0.0.1-test
git push origin :refs/tags/v0.0.1-test

# Optionally delete test release from GitHub UI
```

---

## When TUI is Ready: v0.1.0 Release

### Pre-Release Checklist

- [ ] **TUI is functional**
  - [ ] Chat interface works
  - [ ] Module switching works
  - [ ] Messages display correctly
  - [ ] Input handling works
  - [ ] Bridge to Rust CLI works (or mock works)

- [ ] **Documentation updated**
  - [ ] README.md has current features
  - [ ] DANNI_TERMINAL_INSTALL.md is accurate
  - [ ] Screenshots/demos (optional but nice)

- [ ] **Testing completed**
  - [ ] Works on macOS
  - [ ] Works on Linux
  - [ ] Theme looks good in different terminals
  - [ ] No crashes on basic operations

- [ ] **Version numbers updated**
  - [ ] `/VERSION` → `0.1.0`
  - [ ] `/tui/VERSION` → `0.1.0`
  - [ ] Any version constants in code

### Release Steps

```bash
# 1. Final commit
git add .
git commit -m "chore: prepare v0.1.0 release"
git push

# 2. Create release tag
git tag -a v0.1.0 -m "First public release of Danni Terminal

Features:
- Beautiful TUI with hot pink + gold aesthetic
- Ghostty integration with custom theme
- Module-based AI assistance
- Fast, elegant terminal experience
"

# 3. Push tag (triggers release)
git push origin v0.1.0

# 4. Wait for GitHub Actions (~15 mins)
# Watch: https://github.com/YOUR-USERNAME/danni/actions

# 5. Verify release published
# Check: https://github.com/YOUR-USERNAME/danni/releases/tag/v0.1.0

# 6. Test the install
curl -fsSL https://raw.githubusercontent.com/YOUR-USERNAME/danni/main/scripts/install.sh | sh
```

### Post-Release

- [ ] Test the release yourself
- [ ] Share on social media/communities
- [ ] Monitor GitHub Issues for bug reports
- [ ] Plan next release

---

## Ongoing Releases

### Patch Release (0.1.x)

Bug fixes, small improvements, no breaking changes.

```bash
# Update VERSION files
echo "0.1.1" > VERSION
echo "0.1.1" > tui/VERSION

# Commit
git add VERSION tui/VERSION
git commit -m "chore: bump version to 0.1.1"

# Tag and push
git tag -a v0.1.1 -m "Fix XYZ bug, improve ABC"
git push origin v0.1.1
```

### Minor Release (0.x.0)

New features, enhancements, backwards compatible.

```bash
# Update VERSION files
echo "0.2.0" > VERSION
echo "0.2.0" > tui/VERSION

# Commit
git add VERSION tui/VERSION
git commit -m "chore: bump version to 0.2.0"

# Tag and push
git tag -a v0.2.0 -m "Add new module system, improve performance"
git push origin v0.2.0
```

### Major Release (x.0.0)

Breaking changes, major refactors.

```bash
# Update VERSION files
echo "1.0.0" > VERSION
echo "1.0.0" > tui/VERSION

# Commit
git add VERSION tui/VERSION
git commit -m "chore: bump version to 1.0.0"

# Tag and push
git tag -a v1.0.0 -m "Stable release - complete rewrite with new architecture"
git push origin v1.0.0
```

---

## Pre-Release Versions

For alpha/beta testing before official releases:

### Alpha (early testing)

```bash
echo "0.2.0-alpha.1" > VERSION
git tag v0.2.0-alpha.1
git push origin v0.2.0-alpha.1

# GitHub Actions will mark as "pre-release"
```

### Beta (feature complete, testing)

```bash
echo "0.2.0-beta.1" > VERSION
git tag v0.2.0-beta.1
git push origin v0.2.0-beta.1
```

### Release Candidate (final testing)

```bash
echo "0.2.0-rc.1" > VERSION
git tag v0.2.0-rc.1
git push origin v0.2.0-rc.1
```

---

## Troubleshooting

### GitHub Actions Failed

1. Check the Actions tab for error logs
2. Common issues:
   - Go/Rust build errors → Fix code, commit, push new tag
   - Missing files → Update workflow to include them
   - Permission errors → Check GitHub repository settings

### Release Assets Missing

If release exists but no files attached:
1. Check if workflow completed successfully
2. Look at "create-release" job logs
3. Files might have failed to upload

### Wrong Version Number

1. Delete the tag: `git tag -d vX.X.X && git push origin :refs/tags/vX.X.X`
2. Delete the GitHub release (via web UI)
3. Fix version, create new tag

### Need to Update Release

You can't modify releases, so:
1. Delete the tag/release
2. Fix issues
3. Create new tag (or increment patch version)

---

## Marketing Your Release

### When v0.1.0 Ships

**Social Media Posts:**
```
✨ Introducing Danni Terminal v0.1.0

The sophisticated AI assistant experience for your terminal.

Hot pink + gold aesthetic meets intelligent assistance.
Built on Ghostty for blazing speed & native feel.

Download: [link]
Docs: [link]

#AI #Terminal #DevTools #OpenSource
```

**Communities to Share:**
- Hacker News (Show HN)
- Reddit: r/programming, r/commandline, r/golang
- Twitter/X
- Lobsters
- Dev.to
- Product Hunt

**Email Announcement:**
- Send to early testers
- Include what's new
- Ask for feedback

---

## Version Strategy

| Version | Meaning | Example |
|---------|---------|---------|
| `0.0.x-pre` | Pre-alpha, infrastructure | Current state |
| `0.0.x-alpha` | Alpha, early TUI testing | When TUI builds |
| `0.1.0-beta` | Beta, feature complete | Before v0.1.0 |
| `0.1.0` | **First public release** | Target milestone |
| `0.x.0` | New features | Ongoing development |
| `0.x.x` | Bug fixes | Patches |
| `1.0.0` | Stable, production ready | Future goal |

---

## Quick Reference

**Test workflow:**
```bash
git tag v0.0.1-test && git push origin v0.0.1-test
```

**Real release:**
```bash
git tag -a v0.1.0 -m "Release message" && git push origin v0.1.0
```

**Delete tag:**
```bash
git tag -d vX.X.X && git push origin :refs/tags/vX.X.X
```

**Check workflow:**
```bash
# Visit: https://github.com/YOUR-USERNAME/danni/actions
```

**Download release:**
```bash
curl -fsSL https://raw.githubusercontent.com/YOUR-USERNAME/danni/main/scripts/install.sh | sh
```

---

## Notes

- Pre-release versions (alpha, beta, rc) are automatically marked as "Pre-release" on GitHub
- Semantic versioning: MAJOR.MINOR.PATCH
- Keep CHANGELOG.md updated (optional but professional)
- Tag messages become release notes
- You can edit release notes after creation via GitHub UI

---

**The release system is ready. Focus on building the TUI. When it's done, shipping is one command away.**
