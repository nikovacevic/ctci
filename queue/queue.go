package queue

import (
	"fmt"
	"sync"

	"github.com/nikovacevic/ctci/stack"
)

// Queue defines behavior of a stack data structure
type Queue interface {
	Add(interface{})
	Remove() (interface{}, error)
	Peek() (interface{}, error)
	IsEmpty() bool
}

// IntQueue implements a thread-safe Queue of type int
type IntQueue struct {
	lock  sync.Mutex
	queue []int
}

// IntStackQueue (3.4) implements a queue of ints using IntStacks
type IntStackQueue struct {
	lock sync.Mutex
	in   *stack.IntStack
	out  *stack.IntStack
}

// NewIntQueue creates and returns a reference to a new instance of an IntQueue
func NewIntQueue() *IntQueue {
	return &IntQueue{queue: []int{}}
}

// Add takes one or more ints, adding each to the IntQueue in the given
// order; e.g. NewIntQueue().Add(1, 2, 3) causes 1 to be the head element.
func (q *IntQueue) Add(n ...int) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.queue = append(q.queue, n...)
}

// Remove removes and returns the head element in the IntQueue.
func (q *IntQueue) Remove() (int, error) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.IsEmpty() {
		return 0, fmt.Errorf("Queue is empty")
	}
	n := q.queue[0]
	q.queue = q.queue[1:]
	return n, nil
}

// Peek returns the value of the head element, but does not remove it.
func (q *IntQueue) Peek() (int, error) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.IsEmpty() {
		return 0, fmt.Errorf("Queue is empty")
	}
	return q.queue[0], nil
}

// IsEmpty returns true if the IntQueue has no elements.
func (q *IntQueue) IsEmpty() bool {
	return len(q.queue) == 0
}

// NewIntStackQueue creates and returns reference to a new IntStackQueue
func NewIntStackQueue() *IntStackQueue {
	return &IntStackQueue{
		in:  stack.NewIntStack(),
		out: stack.NewIntStack(),
	}
}

// Add takes one or more ints, adding each to the IntStackQueue in the given
// order; e.g. NewIntStackQueue().Add(1, 2, 3) causes 1 to be the head element.
func (q *IntStackQueue) Add(n ...int) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.in.Push(n...)
}

// Remove removes and returns the head element in the IntStackQueue.
func (q *IntStackQueue) Remove() (int, error) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.IsEmpty() {
		return 0, fmt.Errorf("Queue is empty")
	}
	if !q.out.IsEmpty() {
		return q.out.Pop()
	}
	if err := q.transfer(); err != nil {
		return 0, err
	}
	return q.out.Pop()
}

// Peek returns the value of the head element, but does not remove it.
func (q *IntStackQueue) Peek() (int, error) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.IsEmpty() {
		return 0, fmt.Errorf("Queue is empty")
	}
	if !q.out.IsEmpty() {
		return q.out.Peek()
	}
	if err := q.transfer(); err != nil {
		return 0, err
	}
	return q.out.Peek()
}

// transfer moves all elements from in to out, reversing their order
func (q *IntStackQueue) transfer() error {
	for !q.in.IsEmpty() {
		n, err := q.in.Pop()
		if err != nil {
			return err
		}
		q.out.Push(n)
	}
	return nil
}

// IsEmpty returns true if the IntStackQueue has no elements.
func (q *IntStackQueue) IsEmpty() bool {
	return q.in.IsEmpty() && q.out.IsEmpty()
}
