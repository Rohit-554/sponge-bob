#!/usr/bin/env bash
set -euo pipefail

REPO="Rohit-554/sponge-bob"
BINARY="spongebob"
INSTALL_DIR="${INSTALL_DIR:-$HOME/.local/bin}"
mkdir -p "$INSTALL_DIR"

OS_RAW=$(uname -s | tr '[:upper:]' '[:lower:]')

case "$OS_RAW" in
  linux)                OS="linux"   ;;
  darwin)               OS="darwin"  ;;
  mingw*|msys*|cygwin*) OS="windows" ;;
  *)
    echo "Unsupported OS: $OS_RAW"
    exit 1
    ;;
esac

ARCH=$(uname -m)
case "$ARCH" in
  x86_64)  ARCH="amd64" ;;
  arm64)   ARCH="arm64" ;;
  aarch64) ARCH="arm64" ;;
  *)
    echo "Unsupported architecture: $ARCH"
    exit 1
    ;;
esac

echo "Fetching latest release..."
TAG=$(curl -sSf "https://api.github.com/repos/${REPO}/releases/latest" \
  | grep '"tag_name"' \
  | head -1 \
  | sed 's/.*"tag_name": "\(.*\)".*/\1/')

# Validate the tag only contains safe characters (e.g. v1.2.3)
if [ -z "$TAG" ] || ! echo "$TAG" | grep -qE '^v?[0-9]+\.[0-9]+\.[0-9]+$'; then
  echo "Could not determine a valid release tag. Got: ${TAG:-empty}"
  exit 1
fi

if [ "$OS" = "windows" ]; then
  FILENAME="${BINARY}_${OS}_${ARCH}.exe"
  INSTALL_DIR="$HOME/bin"
  mkdir -p "$INSTALL_DIR"
else
  FILENAME="${BINARY}_${OS}_${ARCH}"
fi

BASE_URL="https://github.com/${REPO}/releases/download/${TAG}"
DOWNLOAD_URL="${BASE_URL}/${FILENAME}"
CHECKSUM_URL="${BASE_URL}/checksums.txt"

TMP_FILE=$(mktemp)
TMP_CHECKSUMS=$(mktemp)
trap 'rm -f "$TMP_FILE" "$TMP_CHECKSUMS"' EXIT

echo "Downloading ${FILENAME} (${TAG})..."
curl -sSfL "$DOWNLOAD_URL" -o "$TMP_FILE"

echo "Verifying checksum..."
if curl -sSfL "$CHECKSUM_URL" -o "$TMP_CHECKSUMS" 2>/dev/null; then
  EXPECTED=$(grep "${FILENAME}" "$TMP_CHECKSUMS" | awk '{print $1}')
  if [ -z "$EXPECTED" ]; then
    echo "Warning: no checksum entry found for ${FILENAME}. Skipping verification."
  else
    if command -v sha256sum >/dev/null 2>&1; then
      ACTUAL=$(sha256sum "$TMP_FILE" | awk '{print $1}')
    elif command -v shasum >/dev/null 2>&1; then
      ACTUAL=$(shasum -a 256 "$TMP_FILE" | awk '{print $1}')
    else
      echo "Warning: no sha256 tool found. Skipping checksum verification."
      ACTUAL="$EXPECTED"
    fi

    if [ "$ACTUAL" != "$EXPECTED" ]; then
      echo "Checksum mismatch. Download may be corrupted or tampered with."
      echo "Expected: $EXPECTED"
      echo "Actual:   $ACTUAL"
      exit 1
    fi
    echo "Checksum verified."
  fi
else
  echo "Warning: checksums.txt not found for this release. Skipping verification."
fi

chmod +x "$TMP_FILE"

if [ "$OS" = "windows" ]; then
  mv "$TMP_FILE" "${INSTALL_DIR}/${BINARY}.exe"
  echo ""
  echo "Installed to ${INSTALL_DIR}/${BINARY}.exe"
  echo "Make sure ${INSTALL_DIR} is in your PATH."
elif [ -w "$INSTALL_DIR" ]; then
  mv "$TMP_FILE" "${INSTALL_DIR}/${BINARY}"
  echo "Installed to ${INSTALL_DIR}/${BINARY}"
else
  sudo mv "$TMP_FILE" "${INSTALL_DIR}/${BINARY}"
  echo "Installed to ${INSTALL_DIR}/${BINARY}"
fi

echo ""
echo "Quick start:"
echo "  export GITHUB_TOKEN=your_token"
echo "  spongebob plan.md"
echo ""
echo "Get a token at: https://github.com/settings/tokens (needs gist scope)"
