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
