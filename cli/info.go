package cli

import (
	"encoding/json"
	"fmt"
	"gbm/core/release"
	"gbm/util"
)

func Info() {
	flags := util.NewFlagSet("info")
	var user, repo, pattern string
	var outputJson bool
	flags.StringVar(&user, "u", "", "github user name")
	flags.StringVar(&repo, "r", "", "github repository name")
	flags.StringVar(&pattern, "p", "", "pattern to filter asset (comma separated, case-insensitive)")
	flags.BoolVar(&outputJson, "json", false, "print output to json")
	flags.ParseCmdFlags()

	flags.ValidateStringNotEmpty(user, "user not provided")
	flags.ValidateStringNotEmpty(repo, "repository not provided")
	flags.ValidateStringNotEmpty(pattern, "filter pattern for assets not provided")

	gh_release := release.NewRelease()
	info, err := gh_release.GetRelease(user, repo, util.ParsePatternsFromString(pattern)...)
	if err != nil {
		panic(err)
	}
	if outputJson {
		bytes, err := json.MarshalIndent(&info, "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(bytes))
	} else {
		fmt.Println("tag_name            : " + info.TagName)
		fmt.Println("name                : " + info.Name)
		fmt.Println("created_at          : " + info.CreatedAt.String())
		fmt.Println("published_at        : " + info.PublishedAt.String())
		fmt.Println("asset_name          : " + info.AssetName)
		fmt.Println("asset_download_link : " + info.AssetDownloadLink)
	}
}
