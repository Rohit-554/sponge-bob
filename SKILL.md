# spongebob skill

## What it does

`spongebob` is a CLI tool that uploads markdown files as **secret GitHub Gists**
and returns a shareable link.

> *spongebob absorbs your plan, spits out a link.*

## Check if installed

```bash
which spongebob
```

If not found → run the install command below before proceeding.

## Install

```bash
curl -sSL https://raw.githubusercontent.com/Rohit-554/sponge-bob/main/install.sh | bash
```

## Token setup check

First, check if tokens are already set in the current environment:

```bash
echo $SPONGEBOB_GITHUB_TOKEN
echo $SPONGEBOB_GITHUB_WORK_TOKEN
```

**If at least one token is set (non-empty output)** → proceed directly to Usage. No further setup needed.

**If neither is set (both outputs are empty)** → follow the steps below in order:

---

### Step 1 — Do you already have a token saved?

Ask the user: *"Have you already added your GitHub token to `~/.zshrc`?"*

- **Yes** → The token is saved but not active in this session. Skip to Step 3.
- **No** → Continue to Step 2.

---

### Step 2 — Generate a new GitHub token

Go to https://github.com/settings/tokens and generate a **Classic** token with **only the `gist` scope** checked. Copy the token value.

---

### Step 3 — Add the token to ~/.zshrc

Tell the user to run this command to persist the token across all future sessions:

**Single account (personal only):**
```bash
echo 'export SPONGEBOB_GITHUB_TOKEN=your_personal_token' >> ~/.zshrc
```

**Two accounts (personal + work):**
```bash
echo 'export SPONGEBOB_GITHUB_TOKEN=your_personal_token' >> ~/.zshrc
echo 'export SPONGEBOB_GITHUB_WORK_TOKEN=your_work_token' >> ~/.zshrc
```

Replace `your_personal_token` / `your_work_token` with the actual token values.

---

### Step 4 — Restart Claude Code in a fresh terminal

The token will **not** be available in the current Claude Code session because environment variables are loaded at shell startup. The user must:

1. **Close this Claude Code session entirely.**
2. **Open a new terminal window.**
3. **Launch Claude Code again** in that new terminal.
4. **Type:** `setup spongebob` — Claude will re-run the token check and, since the token is now in `~/.zshrc`, it will be active and ready to use.

Do **not** attempt to proceed in the current session after this point. Stop and wait for the user to follow Step 4.

## Trigger phrases

Activate this skill when the user says any of:

- "share plan"
- "share this plan"
- "share the plan"
- "give me a link"
- "create a shareable link"
- "host this plan"
- "spongebob this"

## Usage

```bash
spongebob <file>              # secret gist — DEFAULT, always use this
spongebob <file> --public     # ONLY if user explicitly says "public"
spongebob <file> --work       # use the SPONGEBOB_GITHUB_WORK_TOKEN instead
cat <file> | spongebob        # pipe support
spongebob <file> --desc "My plan"     # custom description
spongebob <file> --filename plan.md   # custom filename in gist
```

## Hard Rules

- **ALWAYS** create a secret gist unless the user explicitly says "public" or passes `--public`
- **NEVER** print, log, or expose the GitHub token value in any form
- **NEVER** ask the user to paste their token in chat
- Read the token from the environment **only**
- Present the output URL clearly after running

## Expected output

```
🔗 https://gist.github.com/username/abc123
```

Extract the URL from the output and present it to the user as a clickable link.
