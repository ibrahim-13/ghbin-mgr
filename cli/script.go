package cli

import (
	"gbm/util"
)

func Script() {
	flags := util.NewFlagSet("script")
	var scriptFile string
	flags.StringVar(&scriptFile, "f", "", "binary")
	flags.ParseCmdFlags()

	flags.ValidateStringNotEmpty(scriptFile, "script file location")
}
