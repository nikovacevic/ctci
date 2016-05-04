package strings

import "testing"

// 1.1 Is Unique
var isUniqueTests = []struct {
	str      string // input
	expected bool
}{
	{"run", true},
	{"strava", false},
	{"I am Niko.", false},
	{"ABC abc", true},
	{"", true},
}

func TestIsUnique(t *testing.T) {
	for _, tt := range isUniqueTests {
		actual := IsUnique(tt.str)
		if actual != tt.expected {
			t.Errorf("IsUnique(\"%s\") expected %v, actual %v", tt.str, tt.expected, actual)
		}
	}
}

// 1.2 Check Permutation
var isPermutationTests = []struct {
	s1       string // input
	s2       string // input
	expected bool
}{
	{"run", "fun", false},
	{"strava", "vastar", true},
	{"", "", true},
	{"racecar", "racecar", true},
	{"race car", "ra cecar", true},
	{"race car", "ra cecar ", false},
}

func TestIsPermutation(t *testing.T) {
	for _, tt := range isPermutationTests {
		actual := IsPermutation(tt.s1, tt.s2)
		if actual != tt.expected {
			t.Errorf("IsPermutation(\"%s\", \"%s\") expected %v, actual %v", tt.s1, tt.s2, tt.expected, actual)
		}
	}
}
