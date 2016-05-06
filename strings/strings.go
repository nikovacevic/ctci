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
// O()
func URLify(s string) string {
	str := []byte(s)

	// Count the spaces
	count := 0
	for _, v := range str {
		if v == byte(' ') {
			count++
		}
	}
	newLen := len(str) + count*2
	newStr := make([]byte, newLen)

	// Scan the string backwards
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == byte(' ') {
			newStr[newLen-1] = '0'
			newStr[newLen-2] = '2'
			newStr[newLen-3] = '%'
			newLen -= 3
		} else {
			newStr[newLen-1] = str[i]
			newLen--
		}
	}

	return string(newStr)
}
