package cli

import (
	"bufio"
	"gbm/cli/vm"
	"gbm/util"
	"os"
)

func InstructionSet() {
	flags := util.NewFlagSet("code")
	var instructionFilePath string
	flags.StringVar(&instructionFilePath, "f", "", "instruction file path")
	flags.ParseCmdFlags()

	flags.ValidateStringNotEmpty(instructionFilePath, "instruction file not provided")

	file, err := os.Open(instructionFilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	vm := vm.NewStackVm()

	for scanner.Scan() {
		line := scanner.Text()
		if err := vm.Load(line); err != nil {
			panic(err)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	if err := vm.Exec(); err != nil {
		panic(err)
	}
}
