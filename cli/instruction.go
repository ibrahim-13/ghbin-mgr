package cli

import (
	"gbm/stackmachine"
	"gbm/util"
)

func InstructionSet() {
	flags := util.NewFlagSet("code")
	var instructionFilePath string
	flags.StringVar(&instructionFilePath, "f", "", "instruction file path")
	flags.ParseCmdFlags()

	flags.ValidateStringNotEmpty(instructionFilePath, "instruction file not provided")

	machine := stackmachine.NewStackMachine()

	if err := machine.Load(instructionFilePath); err != nil {
		panic(err)
	}

	if err := machine.Exec(); err != nil {
		panic(err)
	}
}
