package vm

import "math"

type InstructionType string

const (
	INST_PUSH     InstructionType = "push"
	INST_POP      InstructionType = "pop"
	INST_LABEL    InstructionType = "label"
	INST_PUSHADDR InstructionType = "pushaddr"
	INST_GOTO     InstructionType = "goto"
	INST_RETURN   InstructionType = "return"
	INST_PRINT    InstructionType = "print"
	INST_EXIR     InstructionType = "exit"
)

type Instruction struct {
	Type       InstructionType
	LineNumber int
	Data       any
}

func (inst Instruction) PushParam() []Data {
	if inst.Type == INST_PUSH {
		return inst.Data.([]Data)
	}
	return nil
}

func (inst Instruction) PopParam() int {
	if inst.Type == INST_POP {
		return inst.Data.(int)
	}
	return 1
}

func (inst Instruction) PrintParam() int {
	if inst.Type == INST_PRINT {
		return inst.Data.(int)
	}
	return 1
}

func (inst Instruction) LabelIc() int {
	if inst.Type == INST_LABEL {
		return inst.Data.(int)
	}
	return math.MaxInt
}

func NewInstructionPush(lineNo int, params []Data) Instruction {
	return Instruction{
		Type:       INST_PUSH,
		LineNumber: lineNo,
		Data:       params,
	}
}

func NewInstructionPop(lineNo int, count int) Instruction {
	return Instruction{
		Type:       INST_POP,
		LineNumber: lineNo,
		Data:       count,
	}
}

func NewInstructionPrint(lineNo int, count int) Instruction {
	return Instruction{
		Type:       INST_PRINT,
		LineNumber: lineNo,
		Data:       count,
	}
}

func NewInstructionLabel(lineNo, ic int) Instruction {
	return Instruction{
		Type:       INST_LABEL,
		LineNumber: lineNo,
		Data:       ic,
	}
}
