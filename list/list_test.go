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

func TestKthToLasts(t *testing.T) {
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
