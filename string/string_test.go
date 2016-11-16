package string

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

// 1.5 One Away
var oneAwayTests = []struct {
	s1       string // input
	s2       string // input
	expected bool
}{
	{"Niko", "Nike", true},
	{"Niko", "Niko", true},
	{"Niko", "Nik", true},
	{"Niko", "Nice", false},
	{"", "a", true},
	{"a", "", true},
}

func TestOneAway(t *testing.T) {
	for _, tt := range oneAwayTests {
		actual := IsOneAway(tt.s1, tt.s2)
		if actual != tt.expected {
			t.Errorf("IsOneAway(\"%s\", \"%s\") expected %v, actual %v", tt.s1, tt.s2, tt.expected, actual)
		}
	}
}

// Compress (1.6)
var compressTests = []struct {
	s   string
	exp string
}{
	{"aaabbbccc", "a3b3c3"},
	{"abc", "abc"},
	{"aabcc", "aabcc"},
	{"aabbcc", "aabbcc"},
	{"aabbbcc", "a2b3c2"},
}

func TestCompress(t *testing.T) {
	for _, tt := range compressTests {
		act := Compress(tt.s)
		if act != tt.exp {
			t.Errorf("Compress(\"%s\") expected %v, actual %v", tt.s, tt.exp, act)
		}
	}
}

// IsRotation (1.9)
var isRotationTests = []struct {
	s1  string
	s2  string
	exp bool
}{
	{"abcde", "cdeab", true},
	{"abcde", "cdeba", false},
	{"aaaba", "baaaa", true},
	{"aaaba", "aaaaa", false},
	{"baaba", "baaba", true},
	{"baaba", "bbaaa", false},
}

func TestIsRotation(t *testing.T) {
	for _, tt := range isRotationTests {
		act := IsRotation(tt.s1, tt.s2)
		if act != tt.exp {
			t.Errorf("IsRotation(\"%s\", \"%s\") expected %v, actual %v", tt.s1, tt.s2, tt.exp, act)
		}
	}
}
