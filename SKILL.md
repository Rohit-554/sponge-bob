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

Check if any of these tokens are set in the environment:

```bash
echo $GITHUB_TOKEN
echo $SPONGEBOB_TOKEN
echo $GITHUB_WORK_TOKEN
echo $SPONGEBOB_WORK_TOKEN
```

If neither is set → tell the user:

> "Please set your GitHub token first:
> ```
> export GITHUB_TOKEN=your_token
> ```
> *(If you have multiple accounts, you can also set `export GITHUB_WORK_TOKEN=your_work_token`)*
> 
> Add to `~/.zshrc` or `~/.bashrc` to persist across sessions.
> Generate a token at: https://github.com/settings/tokens (needs **gist** scope)"

Do **not** proceed until a token is confirmed to be set.

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
spongebob <file> --work       # use the GITHUB_WORK_TOKEN instead
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
