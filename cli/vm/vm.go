package vm

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction = string

type InstructionSet struct {
	Instruction Instruction
	LineNumber  int
	IsLabel     bool
	Params      []Data
}

const (
	INST_PUSH   Instruction = "push"
	INST_POP    Instruction = "pop"
	INST_NONE   Instruction = "none"
	INST_CALL   Instruction = "call"
	INST_GOTO   Instruction = "goto"
	INST_RETURN Instruction = "return"
	INST_PRINT  Instruction = "print"
)

type StackVm struct {
	Stack        *Stack
	instructions []InstructionSet
	labelMap     map[string]int
}

func NewStackVm() *StackVm {
	return &StackVm{
		Stack:    NewStack(),
		labelMap: make(map[string]int),
	}
}

func (vm *StackVm) Load(instructionFilePath string) error {
	file, err := os.Open(instructionFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	counter := 1
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, ":") && strings.HasSuffix(line, ":") {
			vm.instructions = append(vm.instructions, InstructionSet{
				Instruction: INST_NONE,
				LineNumber:  counter,
				IsLabel:     true,
			})
			vm.labelMap[line] = len(vm.instructions) - 1
		} else {
			inst := strings.SplitN(line, " ", 2)
			if len(inst) > 0 {
				switch strings.ToLower(inst[0]) {
				case INST_PUSH:
					if len(inst) < 2 {
						break
					}
					params, err := ParseParameters(inst[1])
					if err != nil {
						return fmt.Errorf("line %d: %w", counter, err)
					}
					vm.instructions = append(vm.instructions, InstructionSet{
						Instruction: INST_PUSH,
						LineNumber:  counter,
						Params:      params,
					})
				case INST_POP:
					pop_count := 1
					if len(inst) > 1 {
						count, err := strconv.Atoi(inst[1])
						if err != nil {
							return fmt.Errorf("line %d : invalid pop param, int required: %s", counter, inst[1])
						}
						pop_count = count
					}
					vm.instructions = append(vm.instructions, InstructionSet{
						Instruction: INST_POP,
						LineNumber:  counter,
						Params:      []Data{NewData(DT_INT, pop_count)},
					})
				case INST_PRINT:
					pop_count := 1
					if len(inst) > 1 {
						count, err := strconv.Atoi(inst[1])
						if err != nil {
							return fmt.Errorf("line %d : invalid print param, int required: %s", counter, inst[1])
						}
						pop_count = count
					}
					vm.instructions = append(vm.instructions, InstructionSet{
						Instruction: INST_PRINT,
						LineNumber:  counter,
						Params:      []Data{NewData(DT_INT, pop_count)},
					})
				}
			}
		}

	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func (vm *StackVm) Exec() error {
	for _, inst := range vm.instructions {
		switch inst.Instruction {
		case INST_PUSH:
			for _, v := range inst.Params {
				vm.Stack.Push(v)
			}
		case INST_POP:
			for i := 0; i < inst.Params[0].data.(int); i++ {
				_, err := vm.Stack.Pop()
				if err != nil {
					return fmt.Errorf("line %d : %w", inst.LineNumber, err)
				}
			}
		case INST_PRINT:
			var params []any
			for i := 0; i < inst.Params[0].data.(int); i++ {
				val, err := vm.Stack.Pop()
				if err != nil {
					return fmt.Errorf("line %d : %w", inst.LineNumber, err)
				}
				params = append(params, val.data)
			}
			fmt.Println(params...)
		}
	}
	return nil
}
