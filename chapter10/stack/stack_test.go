package stack

import (
	"strconv"
	"testing"
)
type MyInt int

func (mi MyInt) String() string {
	return strconv.Itoa(int(mi))
}

func TestStackEmpty(t *testing.T) {
	s := new(Stack[MyInt])
	if !s.Empty() {
		t.Errorf("expexted empty stack, got %v\n", s.arr)
	}
}

func TestStackPush(t *testing.T) {
	expTop := 2
	s := new(Stack[MyInt])
	s.Push(MyInt(1))
	s.Push(MyInt(3))
	s.Push(MyInt(2))
	if s.top != 3 && s.arr[s.top-1] == 2 {
		t.Errorf("expexted top element to be %d at %d got index s.top=%d, el %d\n", expTop, expTop, s.top, s.arr[s.top-1])
	}
}

func TestStackPull(t *testing.T) {
	expTop := 1
	s := new(Stack[MyInt])
	s.Push(1)
	s.Push(2)
	s.Pull(-1) // -1 will be returned, used because of the Generics
	if s.top != 1 && s.arr[s.top-1] != 1 {
		t.Errorf("expexted top element to be %d at %d got index s.top=%d, el %d\n", expTop, expTop, s.top, s.arr[s.top-1])
	}
}

func TestStackPullMore(t *testing.T) {
	s := new(Stack[MyInt])
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	for !s.Empty() {
		s.Pull(-1)
	}
	if s.top != 0 || !s.Empty() {
		t.Errorf("expected empty stack, got %v\n", s.arr)
	}
}
