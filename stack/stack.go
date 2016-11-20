package stack

import "sync"

// Stack defines behavior of a stack data structure
type Stack interface {
	Pop() interface{}
	Push(interface{})
	Peek() interface{}
	IsEmpty() bool
}

// IntStack implements a thread-safe Stack of type int
type IntStack struct {
	lock  sync.Mutex
	stack []int
}

// NewIntStack creates and returns a reference to a new instance of an IntStack
func NewIntStack() *IntStack {
	return &IntStack{stack: []int{}}
}

// Pop removes and returns the top element in the IntStack.
func (s *IntStack) Pop() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	n := s.stack[len(s.stack)-1]
	s.stack = s.stack[0 : len(s.stack)-1]
	return n
}

// Push takes one or more ints, pushing each onto the IntStack in the given
// order; e.g. Push(1, 2, 3) causes 3 to be the top element.
func (s *IntStack) Push(n ...int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.stack = append(s.stack, n...)
}

// Peek returns the value of the top element, but does not remove it.
func (s *IntStack) Peek() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.stack[len(s.stack)-1]
}

// IsEmpty returns true if the IntStack has no elements.
func (s *IntStack) IsEmpty() bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return len(s.stack) == 0
}
