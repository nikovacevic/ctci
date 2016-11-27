package stack

import (
	"fmt"
	"testing"
)

var minTests = []struct {
	push []int
	pop  int
	min  int
	err  error
}{
	{[]int{1, 2, 3}, 1, 1, nil},
	{[]int{5, 2, 3, 2, 1}, 1, 2, nil},
	{[]int{}, 0, 0, fmt.Errorf("Stack is empty")},
}

func TestMin(t *testing.T) {
	for _, tt := range minTests {
		stack := NewIntStack()
		stack.Push(tt.push...)
		for i := 0; i < tt.pop; i++ {
			_, _ = stack.Pop()
		}
		min, err := stack.Min()
		if tt.err != nil {
			if fmt.Sprintf("%v", tt.err) != fmt.Sprintf("%v", err) {
				t.Errorf("%v.Min() expected error \"%v\", actual \"%v\"", stack, tt.err, err)
			}
		} else {
			if min != tt.min {
				t.Errorf("%v.Min() expected \"%v\", actual \"%v\"", stack, tt.min, min)
			}
		}
	}
}

var popAtTests = []struct {
	cap   int
	push  []int
	popAt []int
	exp   *IntSeries
	err   error
}{
	{3, []int{1, 2, 3, 4, 5, 6}, []int{0, 1, 0}, &IntSeries{cap: 3, stacks: []IntStack{IntStack{stack: []int{1}}, IntStack{stack: []int{4, 5}}}}, nil},
	{3, []int{1, 2, 3, 4, 5, 6}, []int{0, 1, 2}, nil, fmt.Errorf("Series has fewer than 2 internal stacks")},
}

func TestPopAt(t *testing.T) {
Tests:
	for _, tt := range popAtTests {
		series := NewIntSeries(tt.cap)
		series.Push(tt.push...)
		var pop = []int{}
		for _, pa := range tt.popAt {
			p, err := series.PopAt(pa)
			if err != nil {
				if tt.err != nil {
					if fmt.Sprintf("%v", tt.err) != fmt.Sprintf("%v", err) {
						t.Errorf("%v.PopAt() expected error: %v, actual error: %v", series, tt.err, err)
					} else {
						continue Tests
					}
				} else {
					t.Errorf("%v.PopAt(%v) error: %v", series, pa, err)
				}
			}
			pop = append(pop, p)
		}
		if !tt.exp.Equals(series) {
			t.Errorf("PopAt() expected: %v, actual: %v", tt.exp, series)
		}

	}
}

var sortTests = []struct {
	push []int
	exp  *IntStack
}{
	{[]int{6, 5, 4, 3, 2, 1}, &IntStack{stack: []int{6, 5, 4, 3, 2, 1}}},
	{[]int{1, 2, 3, 4, 5, 6}, &IntStack{stack: []int{6, 5, 4, 3, 2, 1}}},
	{[]int{2, 3, 5, 6, 4, 1}, &IntStack{stack: []int{6, 5, 4, 3, 2, 1}}},
}

func TestSort(t *testing.T) {
	for _, tt := range sortTests {
		stack := NewIntStack()
		stack.Push(tt.push...)
		stack.Sort()
		if !tt.exp.Equals(stack) {
			t.Errorf("Sort() expected %v, actual %v", tt.exp, stack)
		}
	}
}
