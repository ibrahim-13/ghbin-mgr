package release

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type GhCli struct {
	command string
}

func NewGhCli(command string) *GhCli {
	return &GhCli{
		command: command,
	}
}

func (ctx *GhCli) GetRelease(user, repo string) (*GhReleaseInfo, error) {
	url := fmt.Sprintf("/repos/%s/%s/releases/latest", user, repo)
	cmd := exec.Command(ctx.command, "api", "-H", "Accept: application/vnd.github+json", "-H", "X-GitHub-Api-Version: 2022-11-28", url)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	var info GhReleaseInfo
	err = json.Unmarshal(output, &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

func (ctx *GhCli) GetName() string {
	return "ghcli"
}
