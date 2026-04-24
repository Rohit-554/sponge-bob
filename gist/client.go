package gist

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const githubGistEndpoint = "https://api.github.com/gists"
const maxResponseBytes = 1 * 1024 * 1024

var httpClient = &http.Client{Timeout: 30 * time.Second}

type Publication struct {
	Token       string
	Description string
	Filename    string
	Secret      bool
}

type gistPayload struct {
	Description string              `json:"description"`
	Public      bool                `json:"public"`
	Files       map[string]planFile `json:"files"`
}

type planFile struct {
	Content string `json:"content"`
}

type gistResponse struct {
	ShareableURL string `json:"html_url"`
	ErrorMessage string `json:"message"`
}

func Share(plan string, pub Publication) (string, error) {
	payload, err := buildPayload(plan, pub)
	if err != nil {
		return "", err
	}

	resp, err := postToGitHub(payload, pub.Token)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	return extractShareableLink(resp)
}

func buildPayload(plan string, pub Publication) ([]byte, error) {
	body := gistPayload{
		Description: pub.Description,
		Public:      !pub.Secret,
		Files:       map[string]planFile{pub.Filename: {Content: plan}},
	}
	data, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to build upload payload: %w", err)
	}
	return data, nil
}

func postToGitHub(payload []byte, token string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, githubGistEndpoint, bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to prepare upload request: %w", err)
	}
	attachAuthHeaders(req, token)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("upload request failed: %w", err)
	}
	return resp, nil
}

func attachAuthHeaders(req *http.Request, token string) {
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	req.Header.Set("Content-Type", "application/json")
}

func extractShareableLink(resp *http.Response) (string, error) {
	raw, err := io.ReadAll(io.LimitReader(resp.Body, maxResponseBytes))
	if err != nil {
		return "", fmt.Errorf("failed to read GitHub response: %w", err)
	}

	var result gistResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return "", fmt.Errorf("unexpected response from GitHub: %w", err)
	}

	if resp.StatusCode != http.StatusCreated {
		return "", githubAPIError(resp.StatusCode, result.ErrorMessage)
	}

	if result.ShareableURL == "" {
		return "", fmt.Errorf("GitHub returned no link — verify token has 'gist' scope")
	}

	return result.ShareableURL, nil
}

func githubAPIError(statusCode int, message string) error {
	if message != "" {
		return fmt.Errorf("GitHub API error: %s", message)
	}
	return fmt.Errorf("GitHub API error: unexpected status %d", statusCode)
}
