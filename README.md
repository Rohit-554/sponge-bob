# SpongeBob
[![Go Version](https://img.shields.io/badge/go-1.22+-00ADD8?style=flat-square&logo=go)](https://go.dev)
[![License](https://img.shields.io/badge/license-MIT-blue?style=flat-square)](LICENSE)
[![Platform](https://img.shields.io/badge/platform-macOS%20%7C%20Linux%20%7C%20Windows-lightgrey?style=flat-square)]()
[![Zero Dependencies](https://img.shields.io/badge/dependencies-zero-brightgreen?style=flat-square)]()

A command-line tool that turns any markdown file into a shareable link. Pass it a file, get back a URL. That is all it does.

<img width="1584" height="396" alt="Your paragraph text" src="https://github.com/user-attachments/assets/520aafd1-7119-4b49-b034-10d75b508cb6" />


## The Problem

You use AI agents every day. At some point mid-session you need to share a plan or file, and there is no fast way to do it.

- No way to get a shareable link without leaving the terminal and doing it by hand.
- Sharing with an AI means pasting the whole file into the chat.
- Sending a `.md` file over Slack forces the other person to download and render it.
- You lose context every time you context-switch just to share one file.
- Most tools default to public, which is wrong for a work plan.

## Demo 
<img width="2944" height="1448" alt="image" src="https://github.com/user-attachments/assets/1e4c7800-f4d8-47f1-87f3-768033f88eb7" />

<img width="1276" height="388" alt="image" src="https://github.com/user-attachments/assets/cfe3b563-682a-41b9-a016-cff01725fe8c" />


## The Solution

```
spongebob plan.md
https://gist.github.com/Rohit-554/abc123def456
```

One command. One link. Secret by default.

It also works as a Claude Code skill. Say "share this plan" and Claude runs it automatically.

---

## Quick Start

**Option 1 — Let Claude handle it**

Download [SKILL.md](https://raw.githubusercontent.com/Rohit-554/sponge-bob/main/SKILL.md) and add it to your project's `CLAUDE.md`:

```sh
curl -sSL https://raw.githubusercontent.com/Rohit-554/sponge-bob/main/SKILL.md >> CLAUDE.md
```

Then tell Claude: `install spongebob`. Claude will check your environment, run the installer, verify your token is set, and confirm the tool is ready. After that, say "share this plan" or "give me a link" whenever you want to share a file.

**Option 2 — Install manually**

```sh
curl -sSL https://raw.githubusercontent.com/Rohit-554/sponge-bob/main/install.sh | bash
```

Then persist your GitHub token and run it:

```sh
echo 'export SPONGEBOB_GITHUB_TOKEN=your_token' >> ~/.zshrc
source ~/.zshrc
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

spongebob uses GitHub Gist as the sharing backend. It needs a personal access token to create Gists on your behalf.

**Step 1.** Go to [https://github.com/settings/tokens](https://github.com/settings/tokens) and click **Generate new token (classic)**.

**Step 2.** Give it a name, then check only the `gist` scope.

**Step 3.** Copy the generated token.

**Step 4.** Persist it in your shell config so it is available in every session:

**Single account:**
```sh
echo 'export SPONGEBOB_GITHUB_TOKEN=your_token_here' >> ~/.zshrc
source ~/.zshrc
```

**Then verify it is set:**
```sh
echo $SPONGEBOB_GITHUB_TOKEN
```

### Multiple Accounts Setup

If you juggle personal and work GitHub accounts, run both commands:

```sh
echo 'export SPONGEBOB_GITHUB_TOKEN=your_personal_token' >> ~/.zshrc
echo 'export SPONGEBOB_GITHUB_WORK_TOKEN=your_work_token' >> ~/.zshrc
source ~/.zshrc
```

When you want to share something to your work account, simply use the `--work` flag:

```sh
spongebob --work plan.md
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

### Custom filename

```sh
spongebob plan.md --filename auth-refactor.md
```

### Share publicly

By default every share is secret. Pass `--public` only when you explicitly want the link to be publicly listed.

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
| `--public` | false | Share publicly. Default is always secret. |
| `--work` | false | Use your secondary/work token (`SPONGEBOB_GITHUB_WORK_TOKEN`). |
| `--desc` | `Shared via spongebob` | Description attached to the shared file. |
| `--filename` | Source filename or `plan.md` | Filename used in the shared view. |

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
- Always share secretly unless you explicitly say "public"
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
Authorization: Bearer $SPONGEBOB_GITHUB_TOKEN
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
