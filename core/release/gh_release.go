package release

import (
	"os/exec"
)

type GhRelease interface {
	GetRelease(user, repo string, pattern ...string) (*GhReleaseInfo, error)
	GetName() string
}

const (
	__gh_api_endpoint string = "https://api.github.com"
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
