package algorithms

import (
	"errors"
)

type Stack struct {
	Elements []int
}

func (s *Stack) IsEmpty() bool {
	return len(s.Elements) == 0
}

func (s *Stack) Push(values int) {
	s.Elements = append(s.Elements, values)
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return -1, errors.New("the stack is empty, no elements to pop.")
	}
	// return the value that was popped
	poppedValue := s.Elements[len(s.Elements)-1]
	s.Elements = s.Elements[:len(s.Elements)-1]

	return poppedValue, nil
}

func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return -1, errors.New("the stack is empty, no elements to pop.")
	}
	topStackValue := s.Elements[len(s.Elements)-1]

	return topStackValue, nil
}
