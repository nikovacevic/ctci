package list

import (
	"fmt"
	"strconv"
)

// List implements a linked list of type int
type List struct {
	Head   *Node
	Length int
}

// Node implements a node in a doubly-linked list
type Node struct {
	Value int
	Prev  *Node
	Next  *Node
	List  *List
}

// NewList creates a new List from a variadic set of integers
func NewList(values ...int) *List {
	list := List{}
	if len(values) == 0 {
		return &list
	}

	head := &Node{values[0], nil, nil, &list}
	list.Head = head
	list.Length = 1
	prev := head

	values = values[1:]
	for _, value := range values {
		node := &Node{value, prev, nil, &list}
		(*prev).Next = node
		prev = node
		list.Length++
	}

	return &list
}

// Equal returns true if l1 and l2 have identical elements in the same order
// and returns false otherwise
func Equal(list1, list2 *List) bool {
	if list1.Length != list2.Length {
		return false
	}
	for n1, n2 := list1.Head, list2.Head; n1 != nil; n1, n2 = n1.Next, n2.Next {
		if n1.Value != n2.Value {
			return false
		}
	}
	return true
}

// Remove a Node from a List
func Remove(n *Node) error {
	l := (*n).List
	(*n.Prev).Next = n.Next
	if (*n).Next != nil {
		(*n.Next).Prev = n.Prev
	}
	(*l).Length--
	return nil
}

// InsertAfter inserts a Node into a List after the given Node
func (n *Node) InsertAfter(after *Node) error {
	l := (*after).List
	next := (*after).Next
	(*after).Next = n
	(*n).Next = next
	(*l).Length++
	return nil
}

// InsertBefore inserts a Node into a List before the given Node
func (n *Node) InsertBefore(before *Node) error {
	l := (*before).List
	if before == (*l).Head {
		(*l).Head = n
		(*n).Next = before
		(*l).Length++
		return nil
	}
	prev := (*before).Prev
	(*before).Prev = n
	(*n).Prev = prev
	(*l).Length++
	return nil
}

// String prints a List in readable format
func (l List) String() string {
	str := "["
	for n := l.Head; n != nil; n = (*n).Next {
		str += strconv.Itoa(n.Value)
		if (*n).Next != nil {
			str += ", "
		}
	}
	str += "]"
	return str
}

// String prints a Node in readable format
func (n Node) String() string {
	return fmt.Sprintf("{%v}", n.Value)
}

// RemoveDuplicates (2.1) removes duplicates from an un-sorted linked list
func (l *List) RemoveDuplicates() {
	m := make(map[int]bool)
	for n := (*l).Head; n != nil; n = (*n).Next {
		if m[(*n).Value] {
			Remove(n)
		}
		m[n.Value] = true
	}
}

// KthToLast returns the kth to last Node in the List. Returns an error if k is
// out of bounds. (Does not leverage advantages of doubly-linked list because
// the prompt calls for a singly-linked list)
func (l *List) KthToLast(k int) (*Node, error) {
	if k < 0 {
		return nil, fmt.Errorf("Index k out of bounds")
	}
	b := l.Head
	for j := 0; j < k; j++ {
		b = b.Next
		if b == nil {
			return nil, fmt.Errorf("Index k out of bounds")
		}
	}
	a := l.Head
	for b.Next != nil {
		a = a.Next
		b = b.Next
	}
	return a, nil
}
