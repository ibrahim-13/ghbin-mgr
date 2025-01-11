package stackmachine

import (
	"errors"
	"fmt"
)

type Stack struct {
	stack        []Data
	return_stack []int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(d Data) {
	s.stack = append(s.stack, d)
}

func (s *Stack) Peak(num int) (Data, error) {
	if len(s.stack) < 1 {
		return Data{}, errors.New("stack empty")
	}
	if num < 0 {
		return Data{}, errors.New("peak number can not be negetive")
	}
	return s.stack[len(s.stack)-num-1], nil
}

func (s *Stack) Pop() (Data, error) {
	if len(s.stack) < 1 {
		return Data{}, errors.New("stack empty")
	}
	ret := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return ret, nil
}

func (s *Stack) PushRet(r int) {
	s.return_stack = append(s.return_stack, r)
}

func (s *Stack) PopRet() (int, error) {
	if len(s.return_stack) < 1 {
		return -1, errors.New("return stack empty")
	}
	ret := s.return_stack[len(s.return_stack)-1]
	s.return_stack = s.return_stack[:len(s.return_stack)-1]
	return ret, nil
}

func (s *Stack) PrintStack() {
	for i, v := range s.stack {
		fmt.Println(fmt.Sprintf("%d:", i), v.data)
	}
}
