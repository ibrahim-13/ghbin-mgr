package vm

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type StackVm struct {
	Stack        *Stack
	instructions []Instruction
}

func NewStackVm() *StackVm {
	return &StackVm{
		Stack: NewStack(),
	}
}

func (vm *StackVm) AddInstruction(inst Instruction) {
	vm.instructions = append(vm.instructions, inst)
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
		vm.processLine(line, counter)
		counter += 1
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func (vm *StackVm) Exec() error {
	ic := 0
	for ic < len(vm.instructions) {
		inst := vm.instructions[ic]
		switch inst.Type {
		case INST_PUSH:
			for _, v := range inst.PushParam() {
				vm.Stack.Push(v)
			}
		case INST_POP:
			for i := 0; i < inst.PopParam(); i++ {
				_, err := vm.Stack.Pop()
				if err != nil {
					return fmt.Errorf("line %d : %w", inst.LineNumber, err)
				}
			}
		case INST_PRINT:
			var params []any
			for i := 0; i < inst.PrintParam(); i++ {
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

func (vm *StackVm) processLine(line string, counter int) error {
	if strings.HasPrefix(line, ":") && strings.HasSuffix(line, ":") {
		vm.AddInstruction(NewInstructionLabel(counter, len(vm.instructions)-1))
		return nil
	}

	inst := strings.SplitN(line, " ", 2)
	if len(inst) < 1 {
		return nil
	}

	inst_cmd := strings.ToLower(inst[0])
	switch InstructionType(inst_cmd) {
	case INST_PUSH:
		if len(inst) < 2 {
			break
		}
		params, err := ParseParameters(inst[1])
		if err != nil {
			return fmt.Errorf("line %d: %w", counter, err)
		}
		vm.AddInstruction(NewInstructionPush(counter, params))
	case INST_POP:
		pop_count := 1
		if len(inst) > 1 {
			count, err := strconv.Atoi(inst[1])
			if err != nil {
				return fmt.Errorf("line %d : invalid pop param, int required: %s", counter, inst[1])
			}
			pop_count = count
		}
		vm.AddInstruction(NewInstructionPop(counter, pop_count))
	case INST_PRINT:
		pop_count := 1
		if len(inst) > 1 {
			count, err := strconv.Atoi(inst[1])
			if err != nil {
				return fmt.Errorf("line %d : invalid print param, int required: %s", counter, inst[1])
			}
			pop_count = count
		}
		vm.AddInstruction(NewInstructionPrint(counter, pop_count))
	}

	return nil
}
