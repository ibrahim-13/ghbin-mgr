package release

import (
	"time"
)

type GhReleaseInfoAssetResponse struct {
	Name               string `json:"name"`
	BrowserDownloadURL string `json:"browser_download_url"`
}

type GhReleaseInfoResponse struct {
	TagName     string                       `json:"tag_name"`
	Name        string                       `json:"name"`
	CreatedAt   time.Time                    `json:"created_at"`
	PublishedAt time.Time                    `json:"published_at"`
	Assets      []GhReleaseInfoAssetResponse `json:"assets"`
}

type GhReleaseInfo struct {
	TagName           string    `json:"tag_name"`
	Name              string    `json:"name"`
	CreatedAt         time.Time `json:"created_at"`
	PublishedAt       time.Time `json:"published_at"`
	AssetName         string    `json:"asset_name"`
	AssetDownloadLink string    `json:"asset_download_link"`
}

func InfoFromResponse(response *GhReleaseInfoResponse, asset *GhReleaseInfoAssetResponse) *GhReleaseInfo {
	return &GhReleaseInfo{
		TagName:           response.TagName,
		Name:              response.Name,
		CreatedAt:         response.CreatedAt,
		PublishedAt:       response.PublishedAt,
		AssetName:         asset.Name,
		AssetDownloadLink: asset.BrowserDownloadURL,
	}
}
