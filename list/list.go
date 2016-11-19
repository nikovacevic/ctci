package list

import (
	"fmt"
	"math"
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
func (n *Node) InsertAfter(a *Node) error {
	l := a.List
	next := a.Next
	n.Prev = a
	n.Next = next
	a.Next = n
	l.Length++
	return nil
}

// InsertBefore inserts a Node into a List before the given Node
func (n *Node) InsertBefore(b *Node) error {
	l := b.List
	var prev *Node
	if b == l.Head {
		l.Head = n
		prev = nil
	} else {
		prev = b.Prev
	}
	n.Prev = prev
	n.Next = b
	b.Prev = n
	l.Length++
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

// KthToLast (2.2) returns the kth to last Node in the List. Returns an error if k is
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

// Partition (2.4) arranges the nodes in the given List around the given value,
// p, such that all Nodes with value less than p come before those with value
// greater than or equal to p.
func (l *List) Partition(p int) *List {
	if l.Length < 2 {
		return l
	}
	pl := NewList(l.Head.Value)
	b := pl.Head
	for a := l.Head.Next; a != nil; a = a.Next {
		if a.Value < p {
			i := &Node{a.Value, nil, nil, pl}
			i.InsertBefore(pl.Head)
		} else {
			i := &Node{a.Value, nil, nil, pl}
			i.InsertAfter(b)
			b = i
		}
	}
	return pl
}

// SumLists (2.5) accepts two base-10 natural numbers represented by ones, tens,
// hundreds, ... place as a linked list, returning the sum of the two.
func SumLists(a, b *List) int {
	if b.Length > a.Length {
		a, b = b, a
	}
	na := a.Head
	nb := b.Head
	s := 0 // sum
	p := 0 // power of 10
	c := 0 // carry
	for nb != nil {
		val := na.Value + nb.Value + c
		c = 0
		if val > 10 {
			val -= 10
			c = 1
		}
		s += val * int(math.Pow10(p))
		na = na.Next
		nb = nb.Next
		p++
	}
	for na != nil {
		val := na.Value + c
		c = 0
		if val > 10 {
			val -= 10
			c = 1
		}
		s += val * int(math.Pow10(p))
		na = na.Next
		p++
	}
	return s
}

// IsPalindrome (2.6) checks if a List is a palindrome
func (l *List) IsPalindrome() bool {
	if l.Length < 2 {
		return true
	}
	a, b := l.Head, l.Head
	for b.Next != nil {
		b = b.Next
	}
	for ia, ib := 0, l.Length-1; ia < ib; ia, ib = ia+1, ib-1 {
		if a.Value != b.Value {
			return false
		}
		a = a.Next
		b = b.Prev
	}
	return true
}
