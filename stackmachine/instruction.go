package stackmachine

type InstructionType string

const (
	INST_PUSH     InstructionType = "push"
	INST_POP      InstructionType = "pop"
	INST_LABEL    InstructionType = "label"
	INST_GOTO     InstructionType = "goto"
	INST_RETURN   InstructionType = "return"
	INST_PRINT    InstructionType = "print"
	INST_EXIT     InstructionType = "exit"
	INST_JUMPEQ   InstructionType = "jumpeq"
	INST_JUMPEQN  InstructionType = "jumpeqn"
	INST_KVLOAD   InstructionType = "kvload"
	INST_KVSAVE   InstructionType = "kvsave"
	INST_KVGET    InstructionType = "kvget"
	INST_KVSET    InstructionType = "kvset"
	INST_KVDELETE InstructionType = "kvdelete"
	INST_GHCHECK  InstructionType = "ghcheck"
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

func (inst Instruction) LabelName() string {
	if inst.Type == INST_LABEL {
		return inst.Data.(string)
	}
	return ""
}

func (inst Instruction) GotoLabel() string {
	if inst.Type == INST_GOTO {
		return inst.Data.(string)
	}
	return ""
}

func (inst Instruction) JumpEqLabel() string {
	if inst.Type == INST_JUMPEQ {
		return inst.Data.(string)
	}
	return ""
}

func (inst Instruction) JumpEqNLabel() string {
	if inst.Type == INST_JUMPEQN {
		return inst.Data.(string)
	}
	return ""
}

func (inst Instruction) KvLoadFilePath() string {
	if inst.Type == INST_KVLOAD {
		return inst.Data.(string)
	}
	return ""
}

func (inst Instruction) KvGet() string {
	if inst.Type == INST_KVGET {
		return inst.Data.(string)
	}
	return ""
}

func (inst Instruction) KvSet() string {
	if inst.Type == INST_KVSET {
		return inst.Data.(string)
	}
	return ""
}

func (inst Instruction) KvDelete() string {
	if inst.Type == INST_KVDELETE {
		return inst.Data.(string)
	}
	return ""
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

func NewInstructionLabel(lineNo int, name string) Instruction {
	return Instruction{
		Type:       INST_LABEL,
		LineNumber: lineNo,
		Data:       name,
	}
}

func NewInstructionGoto(lineNo int, label string) Instruction {
	return Instruction{
		Type:       INST_GOTO,
		LineNumber: lineNo,
		Data:       label,
	}
}

func NewInstructionReturn(lineNo int) Instruction {
	return Instruction{
		Type:       INST_RETURN,
		LineNumber: lineNo,
	}
}

func NewInstructionExit(lineNo int) Instruction {
	return Instruction{
		Type:       INST_EXIT,
		LineNumber: lineNo,
	}
}

func NewInstructionJumpEq(lineNo int, label string) Instruction {
	return Instruction{
		Type:       INST_JUMPEQ,
		LineNumber: lineNo,
		Data:       label,
	}
}

func NewInstructionJumpEqN(lineNo int, label string) Instruction {
	return Instruction{
		Type:       INST_JUMPEQN,
		LineNumber: lineNo,
		Data:       label,
	}
}

func NewInstructionKvLoad(lineNo int, filePath string) Instruction {
	return Instruction{
		Type:       INST_KVLOAD,
		LineNumber: lineNo,
		Data:       filePath,
	}
}

func NewInstructionKvSave(lineNo int) Instruction {
	return Instruction{
		Type:       INST_KVSAVE,
		LineNumber: lineNo,
	}
}

func NewInstructionKvGet(lineNo int, key string) Instruction {
	return Instruction{
		Type:       INST_KVGET,
		LineNumber: lineNo,
		Data:       key,
	}
}

func NewInstructionKvSet(lineNo int, key string) Instruction {
	return Instruction{
		Type:       INST_KVSET,
		LineNumber: lineNo,
		Data:       key,
	}
}

func NewInstructionKvDelete(lineNo int, key string) Instruction {
	return Instruction{
		Type:       INST_KVDELETE,
		LineNumber: lineNo,
		Data:       key,
	}
}

func NewInstructionGhCheck(lineNo int, key string) Instruction {
	return Instruction{
		Type:       INST_GHCHECK,
		LineNumber: lineNo,
		Data:       key,
	}
}
