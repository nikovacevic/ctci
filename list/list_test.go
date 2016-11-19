package list

import (
	"fmt"
	"testing"
)

var equalTests = []struct {
	a   *List
	b   *List
	exp bool
}{
	{NewList(), NewList(), true},
	{NewList(1), NewList(1), true},
	{NewList(1), NewList(2), false},
	{NewList(1, 2, 3), NewList(1, 2, 3), true},
	{NewList(1, 2, 3), NewList(1, 3, 2), false},
	{NewList(1, 2, 3), NewList(1, 2, 3, 4), false},
}

func TestEqual(t *testing.T) {
	for _, tt := range equalTests {
		act := Equal(tt.a, tt.b)
		if act != tt.exp {
			t.Errorf("Equal(%v, %v) expected %v, actual %v", tt.a, tt.b, tt.exp, act)
		}
	}
}

// 2.1 RemoveDuplicates
var removeDuplicatesTests = []struct {
	in  *List
	exp *List
}{
	{NewList(1, 2, 3), NewList(1, 2, 3)},
	{NewList(1, 2, 3, 3), NewList(1, 2, 3)},
	{NewList(1, 1, 2, 1, 2, 3, 1, 3, 2), NewList(1, 2, 3)},
}

func TestRemoveDuplicates(t *testing.T) {
	for _, tt := range removeDuplicatesTests {
		tt.in.RemoveDuplicates()
		if !Equal(tt.in, tt.exp) {
			t.Errorf("RemoveDuplicates() expected %v, actual %v", &tt.exp, &tt.in)
		}
	}
}

// 2.2 Kth to Last
var kthToLastTests = []struct {
	list *List
	k    int
	exp  int
	err  error
}{
	{NewList(1, 2, 3), -1, 0, fmt.Errorf("Index k out of bounds")},
	{NewList(1, 2, 3), 0, 3, nil},
	{NewList(1, 2, 3), 1, 2, nil},
	{NewList(1, 2, 3), 2, 1, nil},
	{NewList(1, 2, 3), 3, 0, fmt.Errorf("Index k out of bounds")},
}

func TestKthToLast(t *testing.T) {
	for _, tt := range kthToLastTests {
		act, err := tt.list.KthToLast(tt.k)
		if tt.err != nil {
			if err == nil {
				t.Errorf("%v.KthToLast(%v) expected (%v, %v), actual (%v, %v)", tt.list, tt.k, tt.exp, tt.err, act, err)
			}
		} else {
			if tt.exp != act.Value {
				t.Errorf("%v.KthToLast(%v) expected (%v, %v), actual (%v, %v)", tt.list, tt.k, tt.exp, tt.err, act, err)
			}
		}
	}
}

// 2.4 Partition
var partitionTests = []struct {
	list *List
	p    int
	exp  *List
}{
	{NewList(3, 1, 2), 3, NewList(2, 1, 3)},
	{NewList(1, 2, 3, 4, 5), 3, NewList(2, 1, 3, 4, 5)},
	{NewList(1, 2, 1, 9), 10, NewList(9, 1, 2, 1)},
	{NewList(1, 2, 1, 9), 0, NewList(1, 2, 1, 9)},
}

func TestPartition(t *testing.T) {
	for _, tt := range partitionTests {
		act := tt.list.Partition(tt.p)
		if !Equal(tt.exp, act) {
			t.Errorf("%v.Partition(%v) expected %v, actual %v", tt.list, tt.p, tt.exp, act)
		}
	}
}

// 2.5 SumLists
var sumListsTests = []struct {
	a   *List
	b   *List
	exp int
}{
	{NewList(0), NewList(0), 0},
	{NewList(1), NewList(2), 3},
	{NewList(3, 1, 2), NewList(2, 1, 3), 525},
	{NewList(8, 3, 1, 2), NewList(2, 1, 3), 2450},
	{NewList(9, 9, 9, 9, 9, 9, 9), NewList(9, 9, 9, 9), 10009998},
}

func TestSumLists(t *testing.T) {
	for _, tt := range sumListsTests {
		act := SumLists(tt.a, tt.b)
		if tt.exp != act {
			t.Errorf("SumLists(%v, %v) expected %v, actual %v", tt.a, tt.b, tt.exp, act)
		}
	}
}
