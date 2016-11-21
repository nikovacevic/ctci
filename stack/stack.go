package stack

import (
	"fmt"
	"sync"
)

const maxInt = 1<<31 - 1

// Stack defines behavior of a stack data structure
type Stack interface {
	Push(interface{})
	Pop() (interface{}, error)
	Peek() (interface{}, error)
	IsEmpty() bool
}

// IntStack implements a thread-safe Stack of type int
type IntStack struct {
	lock     sync.Mutex
	stack    []int
	minstack []int
}

// NewIntStack creates and returns a reference to a new instance of an IntStack
func NewIntStack() *IntStack {
	return &IntStack{
		stack:    make([]int, 0),
		minstack: []int{maxInt},
	}
}

// Push takes one or more ints, pushing each onto the IntStack in the given
// order; e.g. Push(1, 2, 3) causes 3 to be the top element.
func (s *IntStack) Push(nums ...int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	for _, n := range nums {
		s.stack = append(s.stack, n)
		s.minstack = append(s.minstack, s.minstack[len(s.minstack)-1])
		if n < s.minstack[len(s.minstack)-1] {
			s.minstack[len(s.minstack)-1] = n
		}
	}
}

// Pop removes and returns the top element in the IntStack.
func (s *IntStack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("Stack is empty")
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	i := len(s.stack) - 1
	n := s.stack[i]
	s.stack = s.stack[0:i]
	s.minstack = s.minstack[0 : len(s.minstack)-1]
	return n, nil
}

// Peek returns the value of the top element, but does not remove it.
func (s *IntStack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("Stack is empty")
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.stack[len(s.stack)-1], nil
}

// IsEmpty returns true if the IntStack has no elements.
func (s *IntStack) IsEmpty() bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return len(s.stack) == 0
}

func (s *IntStack) String() string {
	return fmt.Sprintf("%v", s.stack)
}

// Min (2.2) returns the values of the minimum element
func (s *IntStack) Min() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("Stack is empty")
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.minstack[len(s.minstack)-1], nil
}
