package release

import (
	"encoding/json"
	"errors"
	"fmt"
	"gbm/util"
	"io"
	"net/http"
	"strings"
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

func (ctx *GhApi) GetReleaseResponse(user, repo string) (*GhReleaseInfoResponse, error) {
	url := fmt.Sprintf("%s/repos/%s/%s/releases/latest", ctx.endpoint, user, repo)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	res, err := ctx.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	output, err := io.ReadAll(res.Body)
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

func (ctx *GhApi) GetRelease(user, repo string, pattern ...string) (*GhReleaseInfo, error) {
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

func (ctx *GhApi) GetName() string {
	return "ghapi"
}
