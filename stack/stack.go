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

// Series (3.3) implements a Stack composed of Stacks which cannot grow past
// a certain capacity. Provides the additional ability to pop from a sub-stack
// by stack index.
type Series interface {
	Stack
	Cap() int
	PopAt(int) (interface{}, error)
}

// IntStack implements a thread-safe Stack of type int
type IntStack struct {
	lock     sync.Mutex
	stack    []int
	minstack []int
}

// IntSeries implements a thread-safe Series of type int
type IntSeries struct {
	lock   sync.Mutex
	cap    int
	stacks []IntStack
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

// String returns the string representation of IntStack
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

// NewIntSeries returns a new IntSeries with each IntStack of capacity cap
func NewIntSeries(cap int) *IntSeries {
	return &IntSeries{
		cap: cap,
		stacks: []IntStack{
			*NewIntStack(),
		},
	}
}

// Cap returns the capacity of the IntStack stack constituting the IntSeries
func (s *IntSeries) Cap() int {
	return s.cap
}

// Push adds nums to the top of the IntSeries
func (s *IntSeries) Push(nums ...int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	for _, n := range nums {
		if len(s.stacks) == 0 || len(s.stacks[len(s.stacks)-1].stack) >= s.Cap() {
			s.stacks = append(s.stacks, *NewIntStack())
		}
		s.stacks[len(s.stacks)-1].Push(n)
	}
}

// Pop removes values from the top of the IntSeries
func (s *IntSeries) Pop() (int, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.stacks) == 0 {
		return 0, fmt.Errorf("Series is empty")
	}
	n, err := s.stacks[len(s.stacks)-1].Pop()
	if err != nil {
		return 0, err
	}
	if s.stacks[len(s.stacks)-1].IsEmpty() {
		s.stacks = s.stacks[:len(s.stacks)-1]
	}
	return n, nil
}

// PopAt (3.3) removes values from the top of the IntStack, designated by the
// given index within IntSeries
func (s *IntSeries) PopAt(i int) (int, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.stacks) <= i {
		return 0, fmt.Errorf("Series has fewer than %v internal stacks", i)
	}
	if len(s.stacks[i].stack) == 0 {
		return 0, fmt.Errorf("Series stack %v is empty", i)
	}
	n, err := s.stacks[i].Pop()
	if err != nil {
		return 0, err
	}
	if s.stacks[i].IsEmpty() {
		s.stacks = append(s.stacks[:i], s.stacks[i+1:]...)
	}
	return n, nil
}

// IsEmpty returns true if the IntSeries has no elements
func (s *IntSeries) IsEmpty() bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return len(s.stacks) == 0
}

// Equals returns true if the values and structure of s are equal to that of t
func (s *IntSeries) Equals(t *IntSeries) bool {
	if len(s.stacks) != len(t.stacks) {
		return false
	}
	for si := 0; si < len(s.stacks); si++ {
		if len(s.stacks[si].stack) != len(t.stacks[si].stack) {
			return false
		}
		for i := 0; i < len(s.stacks[si].stack); i++ {
			if s.stacks[si].stack[i] != t.stacks[si].stack[i] {
				return false
			}
		}
	}
	return true
}

// String returns the string representation of IntSeries
func (s *IntSeries) String() string {
	str := "{"
	for i := range s.stacks {
		str += fmt.Sprintf("%v", &s.stacks[i])
	}
	str += "}"
	return str
}
