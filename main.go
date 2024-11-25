package main

import (
	"flag"
	"gbm/util"
	"os"
)

func main() {
	flags := flag.NewFlagSet("ghbin-mgr", flag.ExitOnError)
	var showUi bool
	flags.BoolVar(&showUi, "ui", false, "show terminal ui")

	if len(os.Args) < 2 {
		flags.Usage()
	}
	if err := flags.Parse(os.Args[1:]); err != nil {
		panic(err)
	}

	appCtx := util.NewAppCtx(showUi)
	defer appCtx.Cleanup()
	appCtx.Log.Println("gg")
}
