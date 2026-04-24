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
echo $SPONGEBOB_GITHUB_TOKEN
echo $SPONGEBOB_GITHUB_WORK_TOKEN
```

If neither is set, you must tell the user to configure their token(s) before proceeding. Present this information clearly:

1. **Get a Token**: Go to https://github.com/settings/tokens and generate a token with only the `gist` scope.
2. **Single Account**: Tell the user to run `export SPONGEBOB_GITHUB_TOKEN=your_token`
3. **Multiple Accounts**: Explain that if they use separate personal and work accounts, they can instead set:
   - `export SPONGEBOB_GITHUB_TOKEN=personal_token`
   - `export SPONGEBOB_GITHUB_WORK_TOKEN=work_token`
4. **Persisting**: Tell them to add these `export` lines to their `~/.zshrc` or `~/.bashrc`.
5. **CRITICAL WARNING**: Warn the user that if they add it to their `~/.zshrc`, they **must** either restart Claude or run the `export` command directly in the current terminal, otherwise Claude will not see the updated environment variables.

Do **not** proceed until a token is confirmed to be set in the current environment.

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
