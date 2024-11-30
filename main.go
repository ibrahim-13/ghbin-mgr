package main

import (
	"fmt"
	"gbm/util"
	"os"
)

func print_help() {
	fmt.Print("ghbin-mgr manage binaries of github releases\n\n")
	fmt.Println("comands:")
	fmt.Println("    info      get release information")
	fmt.Println("    install   install binary")
}

func main() {
	if len(os.Args) < 2 {
		print_help()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "info":
		flags := util.NewFlagSet("info")
		flags.ParseCmdFlags()
	case "install":
		flags := util.NewFlagSet("install")
		var installDir, repo string
		flags.StringVar(&installDir, "dir", "", "installation directory")
		flags.StringVar(&repo, "repo", "", "repository (ex. user/repo)")

		flags.ValidateStringNotEmpty(installDir, "installation directoy not provided")
		flags.ValidateStringNotEmpty(repo, "repository not provided")

		appCtx := util.NewAppCtx(installDir)
		fmt.Println(appCtx.Conf.InstallDir)
	default:
		print_help()
	}

	fmt.Println("exiting")
	// mgr := manager.NewManager(appCtx)
	// jesseduffield/lazygit
}
