package release

import (
	"encoding/json"
	"errors"
	"fmt"
	"gbm/util"
	"os/exec"
	"strings"
)

type GhCli struct {
	command string
}

func NewGhCli(command string) *GhCli {
	return &GhCli{
		command: command,
	}
}

func (ctx *GhCli) GetReleaseResponse(user, repo string) (*GhReleaseInfoResponse, error) {
	url := fmt.Sprintf("/repos/%s/%s/releases/latest", user, repo)
	cmd := exec.Command(ctx.command, "api", "-H", "Accept: application/vnd.github+json", "-H", "X-GitHub-Api-Version: 2022-11-28", url)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	var info GhReleaseInfoResponse
	err = json.Unmarshal(output, &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

func (ctx *GhCli) GetRelease(user, repo string, pattern ...string) (*GhReleaseInfo, error) {
	info, err := ctx.GetReleaseResponse(user, repo)
	if err != nil {
		return nil, err
	}
	for i := range info.Assets {
		if util.ContainsAllMatches(info.Assets[i].Name, pattern...) {
			return InfoFromResponse(info, &info.Assets[i]), nil
		}
	}
	return nil, errors.New("asset not found for pattern: " + strings.Join(pattern, ","))
}

func (ctx *GhCli) GetName() string {
	return "ghcli"
}
