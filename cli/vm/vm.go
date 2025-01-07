package vm

import (
	"bufio"
	"fmt"
	"math"
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
		case INST_LABEL:
			ic += 1
		case INST_PUSH:
			for _, v := range inst.PushParam() {
				vm.Stack.Push(v)
			}
			ic += 1
		case INST_POP:
			for i := 0; i < inst.PopParam(); i++ {
				_, err := vm.Stack.Pop()
				if err != nil {
					return fmt.Errorf("line %d : %w", inst.LineNumber, err)
				}
			}
			ic += 1
		case INST_PRINT:
			var params []any
			for i := 0; i < inst.PrintParam(); i++ {
				val, err := vm.Stack.Peak(i)
				if err != nil {
					return fmt.Errorf("line %d : %w", inst.LineNumber, err)
				}
				params = append(params, val.data)
			}
			fmt.Println(params...)
			ic += 1
		case INST_GOTO:
			hasFoundLabelIc := false
			for i, v := range vm.instructions {
				if v.Type == INST_LABEL && v.LabelName() == inst.GotoLabel() {
					vm.Stack.PushRet(ic + 1)
					ic = i
					hasFoundLabelIc = true
					break
				}
			}
			if !hasFoundLabelIc {
				return fmt.Errorf("line %d : could not find label: %s", inst.LineNumber, inst.GotoLabel())
			}
			ic += 1
		case INST_RETURN:
			new_ic, err := vm.Stack.PopRet()
			if err != nil {
				return fmt.Errorf("line %d : %w", inst.LineNumber, err)
			}
			if ic < 0 || ic >= len(vm.instructions) {
				return fmt.Errorf("line %d : ic is out of bounds", inst.LineNumber)
			}
			ic = new_ic
		case INST_EXIT:
			ic = math.MaxInt - 1
		}
	}
	return nil
}

func (vm *StackVm) processLine(line string, counter int) error {
	if strings.HasPrefix(line, ":") && strings.HasSuffix(line, ":") {
		vm.AddInstruction(NewInstructionLabel(counter, line))
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
		for i, j := 0, len(params)-1; i < j; i, j = i+1, j-1 {
			params[i], params[j] = params[j], params[i]
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
	case INST_GOTO:
		if len(inst) < 2 || inst[1] == "" {
			return fmt.Errorf("line %d : invalid goto label name", counter)
		}
		vm.AddInstruction(NewInstructionGoto(counter, inst[1]))
	case INST_RETURN:
		if len(inst) > 1 {
			return fmt.Errorf("line %d : return can not have param", counter)
		}
		vm.AddInstruction(NewInstructionReturn(counter))
	case INST_EXIT:
		if len(inst) > 1 {
			return fmt.Errorf("line %d : exit can not have param", counter)
		}
		vm.AddInstruction(NewInstructionExit(counter))
	}

	return nil
}
