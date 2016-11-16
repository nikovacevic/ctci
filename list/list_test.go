package list

import "testing"

// 2.1 RemoveDuplicates
var removeDuplicatesTests = []struct {
	in  List
	exp List
}{
	{*NewList(1, 2, 3), *NewList(1, 2, 3)},
	{*NewList(1, 2, 3, 3), *NewList(1, 2, 3)},
	{*NewList(1, 1, 2, 1, 2, 3, 1, 3, 2), *NewList(1, 2, 3)},
}

func TestRemoveDuplicates(t *testing.T) {
	for _, tt := range removeDuplicatesTests {
		act := RemoveDuplicates(tt.in)
		if !Equal(act, tt.exp) {
			t.Errorf("RemoveDuplicates() expected %v, actual %v", &tt.exp, &act)
		}
	}
}
