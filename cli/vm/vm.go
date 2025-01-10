package vm

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type StackVm struct {
	Stack        *Stack
	instructions []Instruction
	kv           map[string]string
	kv_filepath  string
}

func NewStackVm() *StackVm {
	return &StackVm{
		Stack: NewStack(),
		kv:    make(map[string]string),
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
		inst, incrementIc := vm.instructions[ic], true
		switch inst.Type {
		case INST_LABEL:
			ic += 1
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
				val, err := vm.Stack.Peak(i)
				if err != nil {
					return fmt.Errorf("line %d : %w", inst.LineNumber, err)
				}
				params = append(params, val.data)
			}
			fmt.Println(params...)
		case INST_GOTO:
			ic_new, err := vm.getIcForLabel(inst.GotoLabel())
			if err != nil {
				return fmt.Errorf("line %d : %w", inst.LineNumber, err)
			}
			vm.Stack.PushRet(ic + 1)
			ic = ic_new
		case INST_RETURN:
			new_ic, err := vm.Stack.PopRet()
			if err != nil {
				return fmt.Errorf("line %d : %w", inst.LineNumber, err)
			}
			if ic < 0 || ic >= len(vm.instructions) {
				return fmt.Errorf("line %d : ic is out of bounds", inst.LineNumber)
			}
			ic, incrementIc = new_ic, false
		case INST_EXIT:
			ic, incrementIc = math.MaxInt-1, false
		case INST_JUMPEQ:
			var1, err := vm.Stack.Peak(0)
			if err != nil {
				return fmt.Errorf("line %d : %w", inst.LineNumber, err)
			}
			var2, err := vm.Stack.Peak(1)
			if err != nil {
				return fmt.Errorf("line %d : %w", inst.LineNumber, err)
			}
			isEqual, err := var1.CompareEq(var2)
			if err != nil {
				return fmt.Errorf("line %d : %w", inst.LineNumber, err)
			}

			if isEqual {
				ic_new, err := vm.getIcForLabel(inst.JumpEqLabel())
				if err != nil {
					return fmt.Errorf("line %d : %w", inst.LineNumber, err)
				}
				vm.Stack.PushRet(ic + 1)
				ic = ic_new
			}
		case INST_JUMPEQN:
			var1, err := vm.Stack.Peak(0)
			if err != nil {
				return fmt.Errorf("line %d : %w", inst.LineNumber, err)
			}
			var2, err := vm.Stack.Peak(1)
			if err != nil {
				return fmt.Errorf("line %d : %w", inst.LineNumber, err)
			}
			isEqual, err := var1.CompareEqN(var2)
			if err != nil {
				return fmt.Errorf("line %d : %w", inst.LineNumber, err)
			}

			if isEqual {
				ic_new, err := vm.getIcForLabel(inst.JumpEqNLabel())
				if err != nil {
					return fmt.Errorf("line %d : %w", inst.LineNumber, err)
				}
				vm.Stack.PushRet(ic + 1)
				ic = ic_new
			}
		case INST_KVLOAD:
			fileLoc := inst.KvLoadFilePath()
			err := vm.loadKv(fileLoc)
			if err != nil {
				return fmt.Errorf("line %d : %w", inst.LineNumber, err)
			}
		case INST_KVSAVE:
			err := vm.saveKv(vm.kv_filepath)
			if err != nil {
				return fmt.Errorf("line %d : %w", inst.LineNumber, err)
			}
		case INST_KVGET:
			key := inst.KvGet()
			val := vm.kv[key]
			if val == "" {
				return fmt.Errorf("line %d : value for the given key is empty", inst.LineNumber)
			}
			vm.Stack.Push(NewData(DT_STRING, val))
		case INST_KVSET:
			key := inst.KvSet()
			val, err := vm.Stack.Peak(0)
			if err != nil {
				return fmt.Errorf("line %d : %w", inst.LineNumber, err)
			}
			vm.kv[key] = val.GetString()
			vm.Stack.Push(val)
		case INST_KVDELETE:
			key := inst.KvDelete()
			delete(vm.kv, key)
		}
		if incrementIc {
			ic += 1
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
	case INST_JUMPEQ:
		if len(inst) < 2 || inst[1] == "" {
			return fmt.Errorf("line %d : invalid jumpeq label name", counter)
		}
		vm.AddInstruction(NewInstructionJumpEq(counter, inst[1]))
	case INST_JUMPEQN:
		if len(inst) < 2 || inst[1] == "" {
			return fmt.Errorf("line %d : invalid jumpeqn label name", counter)
		}
		vm.AddInstruction(NewInstructionJumpEqN(counter, inst[1]))
	case INST_KVLOAD:
		if len(inst) < 2 {
			break
		}
		params, err := ParseParameters(inst[1])
		if err != nil {
			return fmt.Errorf("line %d: %w", counter, err)
		}
		if len(params) > 1 {
			return fmt.Errorf("line %d: kvload expects only one string param for file path", counter)
		}
		if params[0].Type != DT_STRING {
			return fmt.Errorf("line %d: invalid param type for kvload", counter)
		}
		vm.kv_filepath = params[0].GetString()
		if vm.kv_filepath == "" {
			return fmt.Errorf("line %d: empty param type for kvload", counter)
		}
		vm.AddInstruction(NewInstructionKvLoad(counter, vm.kv_filepath))
	case INST_KVSAVE:
		if len(inst) != 1 {
			return fmt.Errorf("line %d: no parameter is permitted", counter)
		}
		vm.AddInstruction(NewInstructionKvSave(counter))
	case INST_KVGET:
		if len(inst) < 2 {
			break
		}
		params, err := ParseParameters(inst[1])
		if err != nil {
			return fmt.Errorf("line %d: %w", counter, err)
		}
		if len(params) > 1 {
			return fmt.Errorf("line %d: kvget expects only one string param for key", counter)
		}
		if params[0].Type != DT_STRING {
			return fmt.Errorf("line %d: invalid param type for kvget", counter)
		}
		key := params[0].GetString()
		if key == "" {
			return fmt.Errorf("line %d: empty param type for kvget", counter)
		}
		vm.AddInstruction(NewInstructionKvGet(counter, key))
	case INST_KVSET:
		if len(inst) < 2 {
			break
		}
		params, err := ParseParameters(inst[1])
		if err != nil {
			return fmt.Errorf("line %d: %w", counter, err)
		}
		if len(params) > 1 {
			return fmt.Errorf("line %d: kvset expects only one string param for key", counter)
		}
		if params[0].Type != DT_STRING {
			return fmt.Errorf("line %d: invalid param type for kvset", counter)
		}
		key := params[0].GetString()
		if key == "" {
			return fmt.Errorf("line %d: empty param type for kvset", counter)
		}
		vm.AddInstruction(NewInstructionKvSet(counter, key))
	case INST_KVDELETE:
		if len(inst) < 2 {
			break
		}
		params, err := ParseParameters(inst[1])
		if err != nil {
			return fmt.Errorf("line %d: %w", counter, err)
		}
		if len(params) > 1 {
			return fmt.Errorf("line %d: kvdelete expects only one string param for key", counter)
		}
		if params[0].Type != DT_STRING {
			return fmt.Errorf("line %d: invalid param type for kvdelete", counter)
		}
		key := params[0].GetString()
		if key == "" {
			return fmt.Errorf("line %d: empty param type for kvdelete", counter)
		}
		vm.AddInstruction(NewInstructionKvDelete(counter, key))
	}

	return nil
}

func (vm *StackVm) getIcForLabel(label string) (int, error) {
	for i, v := range vm.instructions {
		if v.Type == INST_LABEL && v.LabelName() == label {
			return i, nil
		}
	}
	return math.MaxInt, fmt.Errorf("could not find label: %s", label)
}

func (vm *StackVm) loadKv(filePath string) error {
	vm.kv_filepath = filePath
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			vm.kv = make(map[string]string)
			return nil
		}
		return err
	}
	return json.Unmarshal(bytes, &(vm.kv))
}

func (vm *StackVm) saveKv(filePath string) error {
	bytes, err := json.Marshal(vm.kv)
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, bytes, 0755)
}
