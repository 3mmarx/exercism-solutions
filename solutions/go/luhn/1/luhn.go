package luhn

import (
	"strings"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	if len(id) <= 1 {
		return false
	}

	for _, l := range id {
		if l < '0' || l > '9' {
			return false
		}
	}

	position := 0
	sum := 0
	for i := len(id) - 1; i >= 0; i-- {
		digit := int(id[i] - '0')

		if position%2 == 1 {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		position++
	}

	return sum%10 == 0
}
