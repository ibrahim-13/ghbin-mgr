package cli

import (
	"gbm/core/manager"
	"gbm/core/release"
	"gbm/util"
	"path/filepath"
)

func Install() {
	flags := util.NewFlagSet("install")
	var binName, installDir, user, repo, pattern string
	flags.StringVar(&binName, "n", "", "binary")
	flags.StringVar(&installDir, "d", "", "installation directory")
	flags.StringVar(&user, "u", "", "github user name")
	flags.StringVar(&repo, "r", "", "github repository name")
	flags.StringVar(&pattern, "p", "", "pattern to filter asset (comma separated, case-insensitive)")
	flags.ParseCmdFlags()

	flags.ValidateStringNotEmpty(binName, "binary name not provided")
	flags.ValidateStringNotEmpty(installDir, "installation directoy not provided")
	flags.ValidateStringNotEmpty(user, "user not provided")
	flags.ValidateStringNotEmpty(repo, "repository not provided")
	flags.ValidateStringNotEmpty(pattern, "filter pattern for assets not provided")

	gh_release := release.NewRelease()
	info, err := gh_release.GetRelease(user, repo, util.ParsePatternsFromString(pattern)...)
	if err != nil {
		panic(err)
	}
	err = manager.Download(info.AssetDownloadLink, filepath.Join(installDir, binName))
	if err != nil {
		panic(err)
	}
}
