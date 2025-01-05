package vm

type StackVm struct{}

func NewStackVm() *StackVm {
	return &StackVm{}
}

func (vm *StackVm) Exec(line string) {}
