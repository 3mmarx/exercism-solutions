package pangram

import "strings"

func IsPangram(input string) bool {
	lower := strings.ToLower(input)

	for ch := 'a'; ch <= 'z'; ch++ {
		if !strings.ContainsRune(lower, ch) {
			return false
		}
	}

	return true
}
