package stack

import (
	"errors"
	"strconv"
	"strings"
)

type Stack struct {
	top int // index of the last (top) element in the stack
	arr []int
}

func (s *Stack) Empty() bool {
	return s.top == 0
}

func (s *Stack) Push(el int) {
	if s.top == 0 {
		s.arr = make([]int, 0) // init arr, size 20 at the beginning
	}
	s.arr = append(s.arr, el)
	s.top++
}

func (s *Stack) Pull() (int, error) {
	if s.top == 0 {
		return -1, errors.New("stack is empty")
	}
	el := s.arr[s.top-1]
	s.arr = s.arr[:s.top-1] // remove top element from array
	s.top--
	return el, nil
}

func (s *Stack) String() string {
	var sb strings.Builder
	sb.WriteString("Stack from top to bottom: [")
	for i := s.top - 1; i >= 0; i-- {
		if i == 0 {
			sb.WriteString(strconv.Itoa(s.arr[i]) + "]")
		} else {
			sb.WriteString(strconv.Itoa(s.arr[i]) + ", ")
		}
	}
	sb.WriteString("\n")
	return sb.String()
}
