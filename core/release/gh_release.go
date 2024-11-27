package release

import (
	"os/exec"
	"time"
)

type GhReleaseInfo struct {
	TagName     string    `json:"tag_name"`
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"created_at"`
	PublishedAt time.Time `json:"published_at"`
	Assets      []struct {
		Name               string `json:"name"`
		BrowserDownloadURL string `json:"browser_download_url"`
	} `json:"assets"`
}

type GhRelease interface {
	GetRelease(user, repo string) (*GhReleaseInfo, error)
	GetName() string
}

const (
	__gh_api_endpoint string = "https://api.github.com"
	__gh_cli_command  string = "gh"
)

func checkIfGhCliInstalled() bool {
	cmd := exec.Command(__gh_cli_command, "version")
	_, err := cmd.Output()
	return err == nil
}

func NewRelease() GhRelease {
	if checkIfGhCliInstalled() {
		return NewGhCli(__gh_cli_command)
	}
	return NewGhApi(__gh_api_endpoint)
}
