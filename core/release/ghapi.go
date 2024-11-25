package release

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type GhApi struct {
	endpoint string
	client   *http.Client
}

func NewGhApi(endpoint string) *GhApi {
	return &GhApi{
		endpoint: endpoint,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (ctx *GhApi) GetRelease(user, repo string) (*GhReleaseInfo, error) {
	url := fmt.Sprintf("%s/repos/%s/%s/releases/latest", ctx.endpoint, user, repo)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	if err != nil {
		return nil, err
	}
	res, err := ctx.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	output, err := io.ReadAll(res.Body)
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

func (ctx *GhApi) GetName() string {
	return "ghapi"
}
