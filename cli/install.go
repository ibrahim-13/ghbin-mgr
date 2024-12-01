package cli

import (
	"fmt"
	"gbm/util"
)

func Install() {
	flags := util.NewFlagSet("install")
	var installDir, user, repo string
	flags.StringVar(&installDir, "d", "", "installation directory")
	flags.StringVar(&user, "u", "", "github user name")
	flags.StringVar(&repo, "r", "", "github repository name")
	flags.ParseCmdFlags()

	flags.ValidateStringNotEmpty(installDir, "installation directoy not provided")
	flags.ValidateStringNotEmpty(user, "user not provided")
	flags.ValidateStringNotEmpty(repo, "repository not provided")

	appCtx := util.NewAppCtx(installDir)
	fmt.Println(appCtx.Conf.InstallDir)
}
