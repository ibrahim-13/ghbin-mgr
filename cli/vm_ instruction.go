package cli

import (
	"gbm/cli/vm"
	"gbm/util"
)

func InstructionSet() {
	flags := util.NewFlagSet("code")
	var instructionFilePath string
	flags.StringVar(&instructionFilePath, "f", "", "instruction file path")
	flags.ParseCmdFlags()

	flags.ValidateStringNotEmpty(instructionFilePath, "instruction file not provided")

	vm := vm.NewStackVm()

	if err := vm.Load(instructionFilePath); err != nil {
		panic(err)
	}

	if err := vm.Exec(); err != nil {
		panic(err)
	}
}
