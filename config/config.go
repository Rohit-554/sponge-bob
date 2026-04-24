package config

import (
	"fmt"
	"os"
)

func ResolveToken(isWork bool) (string, error) {
	if isWork {
		if token := os.Getenv("GITHUB_WORK_TOKEN"); token != "" {
			return token, nil
		}
		if token := os.Getenv("SPONGEBOB_WORK_TOKEN"); token != "" {
			return token, nil
		}
		return "", missingTokenError(true)
	}

	if token := os.Getenv("GITHUB_TOKEN"); token != "" {
		return token, nil
	}
	if token := os.Getenv("SPONGEBOB_TOKEN"); token != "" {
		return token, nil
	}
	return "", missingTokenError(false)
}

func missingTokenError(isWork bool) error {
	if isWork {
		return fmt.Errorf(`no work GitHub token found.

Set one of the following environment variables:

  export GITHUB_WORK_TOKEN=your_work_token
  # or
  export SPONGEBOB_WORK_TOKEN=your_work_token

Add it to ~/.zshrc or ~/.bashrc to persist across sessions.
Generate a token at: https://github.com/settings/tokens (needs "gist" scope)`)
	}

	return fmt.Errorf(`no GitHub token found.

Set one of the following environment variables:

  export GITHUB_TOKEN=your_token
  # or
  export SPONGEBOB_TOKEN=your_token

Add it to ~/.zshrc or ~/.bashrc to persist across sessions.
Generate a token at: https://github.com/settings/tokens (needs "gist" scope)`)
}
