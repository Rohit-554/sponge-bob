package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/Rohit-554/sponge-bob/config"
	"github.com/Rohit-554/sponge-bob/gist"
)

const maxUploadBytes = 10 * 1024 * 1024

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "❌ "+err.Error())
		os.Exit(1)
	}
}

func run() error {
	upload := parseUploadRequest()

	token, err := config.ResolveToken()
	if err != nil {
		return err
	}

	plan, err := readPlan(upload.sourceArg)
	if err != nil {
		return err
	}

	fmt.Fprintln(os.Stderr, "⏳ Uploading to GitHub Gist…")

	link, err := gist.Share(plan, gist.Publication{
		Token:       token,
		Description: upload.description,
		Filename:    resolveFilename(upload.sourceArg, upload.filenameOverride),
		Secret:      !upload.public,
	})
	if err != nil {
		return err
	}

	fmt.Fprintln(os.Stderr, "✅ Done!")
	fmt.Println("🔗 " + link)
	return nil
}

type uploadRequest struct {
	sourceArg        string
	description      string
	filenameOverride string
	public           bool
}

func parseUploadRequest() uploadRequest {
	req := uploadRequest{}

	flag.BoolVar(&req.public, "public", false, "Make gist public (default: secret)")
	flag.StringVar(&req.description, "desc", "Shared via spongebob", "Gist description")
	flag.StringVar(&req.filenameOverride, "filename", "", "Filename shown in Gist (default: source filename or plan.md)")
	flag.Usage = printUsage
	flag.Parse()

	if len(flag.Args()) > 0 {
		req.sourceArg = flag.Args()[0]
	}

	return req
}

func readPlan(sourcePath string) (string, error) {
	if sourcePath != "" {
		return readFromFile(sourcePath)
	}
	return readFromStdin()
}

func readFromFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("cannot read %q: %w", path, err)
	}
	if len(data) > maxUploadBytes {
		return "", fmt.Errorf("file %q exceeds the 10 MB upload limit", path)
	}
	content := string(data)
	if isBlank(content) {
		return "", fmt.Errorf("file %q is empty — nothing to upload", path)
	}
	return content, nil
}

func readFromStdin() (string, error) {
	if isInteractiveTerminal() {
		printUsage()
		os.Exit(1)
	}
	data, err := io.ReadAll(io.LimitReader(os.Stdin, maxUploadBytes+1))
	if err != nil {
		return "", fmt.Errorf("failed to read stdin: %w", err)
	}
	if len(data) > maxUploadBytes {
		return "", fmt.Errorf("piped content exceeds the 10 MB upload limit")
	}
	content := string(data)
	if isBlank(content) {
		return "", fmt.Errorf("piped content is empty — nothing to upload")
	}
	return content, nil
}

func resolveFilename(sourcePath, override string) string {
	if override != "" {
		return override
	}
	if sourcePath != "" {
		return filepath.Base(sourcePath)
	}
	return "plan.md"
}

func isInteractiveTerminal() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) != 0
}

func isBlank(content string) bool {
	for _, ch := range content {
		if ch != ' ' && ch != '\t' && ch != '\n' && ch != '\r' {
			return false
		}
	}
	return true
}

func printUsage() {
	fmt.Fprintln(os.Stderr, `spongebob — absorbs your plans, spits out a link

Usage:
  spongebob <file>              Upload file as secret gist (default)
  cat <file> | spongebob        Pipe content as secret gist
  spongebob <file> --public     Upload as public gist

Flags:`)
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, `
Environment:
  GITHUB_TOKEN      GitHub personal access token (needs "gist" scope)
  SPONGEBOB_TOKEN   Alternative token env var

Generate a token at: https://github.com/settings/tokens`)
}
