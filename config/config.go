package config

import (
	"fmt"
	"os"
)

func ResolveToken() (string, error) {
	if token := os.Getenv("GITHUB_TOKEN"); token != "" {
		return token, nil
	}
	if token := os.Getenv("SPONGEBOB_TOKEN"); token != "" {
		return token, nil
	}
	return "", missingTokenError()
}

func missingTokenError() error {
	return fmt.Errorf(`no GitHub token found.

Set one of the following environment variables:

  export GITHUB_TOKEN=your_token
  # or
  export SPONGEBOB_TOKEN=your_token

Add it to ~/.zshrc or ~/.bashrc to persist across sessions.
Generate a token at: https://github.com/settings/tokens (needs "gist" scope)`)
}
