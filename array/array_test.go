package array

import (
	"testing"
)

// 1.7 Rotate Matrix
var rotateMatrixTests = []struct {
	in  [][]int
	exp [][]int
}{
	{[][]int{[]int{1}}, [][]int{[]int{1}}},
	{[][]int{[]int{1, 2}, []int{4, 3}}, [][]int{[]int{4, 1}, []int{3, 2}}},
	{[][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}, [][]int{[]int{7, 4, 1}, []int{8, 5, 2}, []int{9, 6, 3}}},
}

func TestRotateMatrix(t *testing.T) {
	for _, tt := range rotateMatrixTests {
		act := RotateMatrix(tt.in)
		if !matricesAreEqual(act, tt.exp) {
			t.Errorf("RotateMatrix(\"%v\") expected %v, actual %v", tt.in, tt.exp, act)
		}
	}
}

func matricesAreEqual(a, b [][]int) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !slicesAreEqual(a[i], b[i]) {
			return false
		}
	}
	return true
}

func slicesAreEqual(a, b []int) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
