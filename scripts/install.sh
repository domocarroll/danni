#!/usr/bin/env bash
# Quick installer for Danni Terminal from GitHub Releases
set -e

PURPLE='\033[0;35m'
GREEN='\033[0;32m'
NC='\033[0m'

echo -e "${PURPLE}✨ Installing Danni Terminal...${NC}"

# Detect platform
OS="$(uname -s)"
ARCH="$(uname -m)"

case "$OS" in
    Darwin)
        PLATFORM="macos"
        [[ "$ARCH" == "arm64" ]] && ARCH_SUFFIX="arm64" || ARCH_SUFFIX="x64"
        ;;
    Linux)
        PLATFORM="linux"
        [[ "$ARCH" == "aarch64" ]] && ARCH_SUFFIX="arm64" || ARCH_SUFFIX="x64"
        ;;
    *)
        echo "Unsupported OS: $OS"
        exit 1
        ;;
esac

# Get latest release
RELEASE_URL="https://github.com/subfracture/danni/releases/latest/download"
TARBALL="danni-terminal-${PLATFORM}-${ARCH_SUFFIX}.tar.gz"

echo "Downloading for ${PLATFORM}-${ARCH_SUFFIX}..."
curl -fsSL "${RELEASE_URL}/${TARBALL}" -o "/tmp/${TARBALL}"

echo "Extracting..."
mkdir -p /tmp/danni-terminal
tar -xzf "/tmp/${TARBALL}" -C /tmp/danni-terminal

echo "Installing..."
cd /tmp/danni-terminal
./install.sh

echo -e "${GREEN}✓ Danni Terminal installed!${NC}"
echo "Run: danni-terminal"
