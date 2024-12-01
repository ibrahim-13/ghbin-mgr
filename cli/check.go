package cli

import (
	"fmt"
	"gbm/core/release"
	"gbm/util"
	"os"
)

func Check() {
	flags := util.NewFlagSet("info")
	var user, repo, tag string
	flags.StringVar(&user, "u", "", "github user name")
	flags.StringVar(&repo, "r", "", "github repository name")
	flags.StringVar(&tag, "t", "", "release tag to compare with")
	flags.ParseCmdFlags()

	flags.ValidateStringNotEmpty(user, "user not provided")
	flags.ValidateStringNotEmpty(repo, "repository not provided")
	flags.ValidateStringNotEmpty(tag, "tag not provided")

	gh_release := release.NewRelease()
	info, err := gh_release.GetReleaseResponse(user, repo)
	if err != nil {
		panic(err)
	}
	if info.TagName != tag {
		fmt.Println("update available")
		os.Exit(0)
	} else {
		fmt.Println("no update available")
		os.Exit(1)
	}
}
