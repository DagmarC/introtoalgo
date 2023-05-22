package stack

import (
	"errors"
	"fmt"
	"strings"
)

type Stack[T fmt.Stringer] struct {
	top int // index of the last (top) element in the stack
	arr []T
}

func (s *Stack[T]) Empty() bool {
	return s.top == 0
}

func (s *Stack[T]) Push(el T) {
	if s.top == 0 {
		s.arr = make([]T, 0) // init arr
	}
	s.arr = append(s.arr, el)
	s.top++
}

func (s *Stack[T]) Pull(empty T) (T, error) {

	if s.top == 0 {
		return empty, errors.New("stack is empty")
	}
	el := s.arr[s.top-1]
	s.arr = s.arr[:s.top-1] // remove top element from array
	s.top--
	return el, nil
}

func (s *Stack[T]) Top(empty T) (T, error) {
	if s.top == 0 {
		return empty, errors.New("stack is empty")
	}
	return s.arr[s.top-1], nil
}

func (s *Stack[T]) String() string {
	var sb strings.Builder
	sb.WriteString("Stack from top to bottom: [")
	for i := s.top - 1; i >= 0; i-- {
		if i == 0 {
			sb.WriteString(s.arr[i].String() + "]")
		} else {
			sb.WriteString(s.arr[i].String() + ", ")
		}
	}
	sb.WriteString("\n")
	return sb.String()
}
