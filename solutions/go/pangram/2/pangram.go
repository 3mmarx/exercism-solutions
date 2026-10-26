package pangram

func IsPangram(input string) bool {
	seen := make(map[rune]struct{})

	for _, ch := range input {
		if ch >= 'A' && ch <= 'Z' {
			ch += 'a' - 'A'
		}

		if ch >= 'a' && ch <= 'z' {
			seen[ch] = struct{}{}
		}
	}

	return len(seen) == 26
}
