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
	var prev *Node
	for v, value := range values {
		if v == 0 {
			head := &Node{value, nil, nil, &list}
			list.Head = head
			list.Length = 1
			prev = head
		} else {
			node := &Node{value, prev, nil, &list}
			(*prev).Next = node
			prev = node
			list.Length++
		}
	}
	return &list
}

// Equal returns true if l1 and l2 have identical elements in the same order
// and returns false otherwise
func Equal(list1, list2 List) bool {
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

// Remove removes a
func (l *List) Remove(n *Node) error {
	if (*n).List.Head != l.Head {
		return fmt.Errorf("List %v does not contain node %v", l, n)
	}
	(*n.Prev).Next = n.Next
	if (*n).Next != nil {
		(*n.Next).Prev = n.Prev
	}
	(*l).Length--
	return nil
}

// String prints a List in readable format
func (l *List) String() string {
	str := "["
	for n := (*l).Head; n != nil; n = (*n).Next {
		str += strconv.Itoa(n.Value)
		if (*n).Next != nil {
			str += ", "
		}
	}
	str += "]"
	return str
}

// RemoveDuplicates removes duplicates from an un-sorted linked list
func RemoveDuplicates(l List) List {
	m := make(map[int]bool)
	for n := l.Head; n != nil; n = n.Next {
		if m[n.Value] {
			l.Remove(n)
		}
		m[n.Value] = true
	}
	return l
}
