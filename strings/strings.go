package strings

// 1.1 Is Unique
// O(len(s))
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
// O(len(s1) + len(s2)) : N = len(s1), M = len(s2)
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
// O(len(s))
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

// 1.4 Palindrome Permutation
// O(len(s))
func IsPalindromePermutation(s string) bool {
	// map of character frequency
	charMap := make(map[byte]int)
	// s as array of bytes
	str := []byte(s)

	for _, v := range str {
		_, ok := charMap[v]
		if !ok {
			charMap[v] = 0
		}
		charMap[v]++
	}

	// if character exists w/ odd frequency, flag as true
	oddExists := false
	for _, c := range charMap {
		if c%2 == 1 {
			if oddExists {
				// more than one odd frequency
				return false
			} else {
				oddExists = true
			}
		}
	}

	return true
}

// 1.5 One Away
// O()
func IsOneAway(s1, s2 string) bool {
	if s1 == s2 {
		// same string
		return true
	} else if len(s1) == len(s2) {
		// compare substitutions
		return isOneSubstitutionAway(s1, s2)
	} else if len(s1) == len(s2)+1 {
		// compare removals
		return isOneRemovalAway(s1, s2)
	} else if len(s1) == len(s2)-1 {
		// compare insertions
		return isOneRemovalAway(s2, s1)
	} else {
		return false
	}
}

func isOneSubstitutionAway(s1, s2 string) bool {
	str1 := []byte(s1)
	str2 := []byte(s2)

	diffCount := 0
	for i, c := range str1 {
		if c != str2[i] {
			diffCount++
		}
		if diffCount > 1 {
			return false
		}
	}

	return true
}

func isOneRemovalAway(s1, s2 string) bool {
	str1 := []byte(s1)
	str2 := []byte(s2)

	removalCount := 0
	for i, c := range str1 {
		if len(str2) <= i || c != str2[i] {
			removalCount++
		}
		if removalCount > 1 {
			return false
		}
	}

	return true
}
