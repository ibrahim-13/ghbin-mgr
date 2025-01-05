package main

import (
	"fmt"
	"gbm/cli"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		cli.PrintHelp()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "info":
		cli.Info()
	case "check":
		cli.Check()
	case "install":
		cli.Install()
	case "installx":
		cli.InstallExtract()
	case "code":
		cli.InstructionSet()
	default:
		cli.PrintHelp()
	}

	fmt.Println("exiting")
	// mgr := manager.NewManager(appCtx)
	// jesseduffield/lazygit
}
