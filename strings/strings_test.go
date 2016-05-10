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

// 1.3 URLify
var urlifyTests = []struct {
	s        string // input
	expected string
}{
	{"Niko", "Niko"},
	{"I run far", "I%20run%20far"},
}

func TestURLify(t *testing.T) {
	for _, tt := range urlifyTests {
		actual := URLify(tt.s)
		if actual != tt.expected {
			t.Errorf("URLify(\"%s\") expected %v, actual %v", tt.s, tt.expected, actual)
		}
	}
}

// 1.4 Palindrome Permutation
var palindromeTests = []struct {
	s        string // input
	expected bool
}{
	{"Niko", false},
	{"I run far", false},
	{"racecar", true},
	{"rceraca", true},
}

func TestPalindrome(t *testing.T) {
	for _, tt := range palindromeTests {
		actual := IsPalindromePermutation(tt.s)
		if actual != tt.expected {
			t.Errorf("IsPalindromePermutation(\"%s\") expected %v, actual %v", tt.s, tt.expected, actual)
		}
	}
}
