package queue

import "testing"

var stackQueueTests = []struct {
	in  []int
	out []int
	err error
}{
	{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, nil},
}

func TestStackQueue(t *testing.T) {
	for _, tt := range stackQueueTests {
		isq := NewIntStackQueue()
		isq.Add(tt.in...)
		for !isq.IsEmpty() {
			n, err := isq.Remove()
			e := tt.out[0]
			tt.out = tt.out[1:]
			if err != nil {
				t.Errorf("StackQueue unexpected error: %v", err)
			}
			if n != e {
				t.Errorf("Remove() expected %v, actual %v", e, n)
			}
		}
	}
	for _, tt := range stackQueueTests {
		isq := NewIntStackQueue()
		for _, i := range tt.in {
			isq.Add(i)
			n, err := isq.Remove()
			e := tt.out[0]
			tt.out = tt.out[1:]
			if err != nil {
				t.Errorf("StackQueue unexpected error: %v", err)
			}
			if n != e {
				t.Errorf("Remove() expected %v, actual %v", e, n)
			}
		}
	}
}
