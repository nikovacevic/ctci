package strings

// 1.1 Is Unique
// O(N) : N = len(s)
func IsUnique(s string) bool {
	runes := make(map[rune]bool)

	for _, r := range s {
		_, found := runes[r]
		if found {
			return false
		}
		runes[r] = true
	}

	return true
}

// 1.2 Check Permutation
// O(N+M) : N = len(s1), M = len(s2)
func IsPermutation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	// Map rune to instance count
	runes := make(map[rune]int)

	for _, r := range s1 {
		runes[r]++
	}

	for _, r := range s2 {
		count, ok := runes[r]
		if !ok || count == 0 {
			return false
		}
		runes[r]--
	}

	return true
}

// 1.3 URLify
