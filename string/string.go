package string

import "strconv"

// IsUnique (1.1)
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

// IsPermutation (1.2)
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

// URLify (1.3)
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

// IsPalindromePermutation (1.4)
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
			}
			oddExists = true
		}
	}

	return true
}

// IsOneAway (1.5)
// O(len(max(s1, s2)))
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

// Compress (1.6)
// O(len(s))
func Compress(s string) string {
	if len(s) < 3 {
		return s
	}
	out := ""    // compressed output string
	comp := 0    // compression score
	curr := s[0] // current Unicode code point
	count := 1   // consecutive count of curr
	for i := 1; i < len(s); i++ {
		if s[i] == curr {
			count++
			continue
		}
		// New Unicode code point
		out += string(curr) + strconv.Itoa(count)
		comp += count - 2
		count = 1
		curr = s[i]
	}
	out += string(curr) + strconv.Itoa(count)
	comp += count - 2
	if comp > 0 {
		// Compression succeeded
		return out
	}
	return s
}

// 1.9 String Rotation
// Assume you have a functino isSubstring, which checks if one word is a
// substring of another. Given two strings, s1 and s2, check if s2 is a rotation
// of s2 using only one call to isSubstring.
