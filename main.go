package main

import (
	"encoding/json"
	"fmt"
	"gbm/core/release"
	"gbm/util"
	"os"
	"runtime"
	"strings"
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

		var formatted_patterns []string
		for _, v := range strings.Split(pattern, ",") {
			_v := strings.ToLower(v)
			if _v == "__os__" {
				formatted_patterns = append(formatted_patterns, runtime.GOOS)
			} else if _v == "__arch__" {
				formatted_patterns = append(formatted_patterns, runtime.GOARCH)
			} else {
				formatted_patterns = append(formatted_patterns, v)
			}
		}

		gh_release := release.NewRelease()
		info, err := gh_release.GetRelease(user, repo, formatted_patterns...)
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
			fmt.Println("WIP")
		}
	case "install":
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
	default:
		print_help()
	}

	fmt.Println("exiting")
	// mgr := manager.NewManager(appCtx)
	// jesseduffield/lazygit
}
