# spongebob

[![Go Version](https://img.shields.io/badge/go-1.22+-00ADD8?style=flat-square&logo=go)](https://go.dev)
[![License](https://img.shields.io/badge/license-MIT-blue?style=flat-square)](LICENSE)
[![Platform](https://img.shields.io/badge/platform-macOS%20%7C%20Linux%20%7C%20Windows-lightgrey?style=flat-square)]()
[![Zero Dependencies](https://img.shields.io/badge/dependencies-zero-brightgreen?style=flat-square)]()

A command-line tool that uploads a markdown file to GitHub Gist and returns a shareable link. Pass it a file, get back a URL. That is all it does.

You use AI agents every day. They generate plans, write code, produce documents. At some point in the middle of a session you need to share one of those files, with a teammate, a reviewer, or another AI. That is where things fall apart.

- You have to leave the terminal, open a browser, and create a Gist or Pastebin by hand.
- You paste the entire file content into the chat just to share it with an AI.
- Sending a `.md` file over Slack or email forces the other person to download and render it themselves.
- You lose context switching out of your current session just to share one file.
- Most tools default to public, which is not what you want for a work plan.

spongebob fixes all of this. One command, one link, always secret by default.

```
spongebob plan.md
https://gist.github.com/Rohit-554/abc123def456
```

It is designed to work as a Claude Code skill so that you can say "share this plan" and Claude will run it automatically.

---

## Quick Start

**Option 1 — Let Claude handle it**

Download [SKILL.md](https://raw.githubusercontent.com/Rohit-554/sponge-bob/main/SKILL.md) and add it to your project's `CLAUDE.md`:

```sh
curl -sSL https://raw.githubusercontent.com/Rohit-554/sponge-bob/main/SKILL.md >> CLAUDE.md
```

Then tell Claude: `install spongebob`. Claude will check your environment, run the installer, verify your token is set, and confirm the tool is ready. After that, say "share this plan" whenever you want a shareable link.

**Option 2 — Install manually**

```sh
curl -sSL https://raw.githubusercontent.com/Rohit-554/sponge-bob/main/install.sh | bash
```

Then set your GitHub token and run it:

```sh
export GITHUB_TOKEN=your_token
spongebob plan.md
```

---

## Table of Contents

- [Requirements](#requirements)
- [Installation](#installation)
- [Token Setup](#token-setup)
- [Usage](#usage)
- [Flags](#flags)
- [Claude Code Integration](#claude-code-integration)
- [How It Works](#how-it-works)
- [Contributing](#contributing)
- [License](#license)

---

## Requirements

- macOS, Linux, or Windows
- A GitHub account
- A GitHub personal access token with the `gist` scope

---

## Installation

```sh
curl -sSL https://raw.githubusercontent.com/Rohit-554/sponge-bob/main/install.sh | bash
```

Detects your OS and architecture, downloads the correct binary, and places it on your PATH.

On Windows, run this inside Git Bash or WSL. The `curl` command exists in PowerShell natively, but `bash` does not, so the pipe will not work in a native PowerShell or CMD window.

### Build from source

```sh
git clone https://github.com/Rohit-554/sponge-bob.git
cd sponge-bob
go build -o spongebob .
```

Move the binary to any directory on your PATH. No external dependencies.

---

## Token Setup

spongebob needs a GitHub personal access token to create Gists on your behalf.

**Step 1.** Go to [https://github.com/settings/tokens](https://github.com/settings/tokens) and click **Generate new token (classic)**.

**Step 2.** Give it a name, then check only the `gist` scope.

**Step 3.** Copy the generated token.

**Step 4.** Export it in your shell:

```sh
export GITHUB_TOKEN=your_token_here
```

To keep it across sessions, add that line to your `~/.zshrc` or `~/.bashrc`.

spongebob checks for `GITHUB_TOKEN` first, then falls back to `SPONGEBOB_TOKEN` if you prefer to keep them separate.

```sh
export SPONGEBOB_TOKEN=your_token_here
```

The token is read from the environment only. It is never logged, printed, or accepted as a command-line argument.

---

## Usage

### Upload a file

```sh
spongebob plan.md
```

### Pipe content

```sh
cat plan.md | spongebob
```

### Custom description

```sh
spongebob plan.md --desc "Auth refactor plan"
```

### Custom filename shown inside the Gist

```sh
spongebob plan.md --filename auth-refactor.md
```

### Public Gist

By default every Gist is secret. Pass `--public` only when you explicitly want it to be publicly listed.

```sh
spongebob plan.md --public
```

### Output

The shareable URL is printed to stdout. Status messages go to stderr so the URL stays pipeable.

```
https://gist.github.com/Rohit-554/abc123def456
```

---

## Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--public` | false | Make the Gist public. Default is always secret. |
| `--desc` | `Shared via spongebob` | Description shown on the Gist page. |
| `--filename` | Source filename or `plan.md` | Filename shown inside the Gist. |

---

## Claude Code Integration

Add the contents of `SKILL.md` to your project's `CLAUDE.md` file.

```sh
cat SKILL.md >> CLAUDE.md
```

Once added, Claude Code will:

- Recognise phrases like "share plan", "give me a link", or "spongebob this"
- Check whether spongebob is installed before running it
- Verify the token is set in your environment
- Always create a secret Gist unless you explicitly say "public"
- Present the returned URL as a clickable link

---

## How It Works

1. spongebob reads content from a file path or stdin
2. It resolves a GitHub token from your environment
3. It sends a POST request to the GitHub Gist API with your content
4. GitHub returns a URL
5. spongebob prints that URL to stdout

```
your file
    |
    v
POST https://api.github.com/gists
Authorization: Bearer $GITHUB_TOKEN
public: false (default)
    |
    v
https://gist.github.com/Rohit-554/abc123
```

The binary is statically compiled with zero external Go dependencies. It uses only the standard library.

---

## Contributing

Pull requests are welcome. For significant changes, open an issue first to discuss what you want to change.

```sh
git clone https://github.com/Rohit-554/sponge-bob.git
cd sponge-bob
go build -o spongebob .
go vet ./...
```

Please keep the zero-dependency constraint. No third-party packages.

---

## License

[MIT](LICENSE)
