package queue

import (
	"fmt"
	"sync"
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
	q.lock.Lock()
	defer q.lock.Unlock()
	return len(q.queue) == 0
}
