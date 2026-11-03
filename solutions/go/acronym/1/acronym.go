package acronym

import (
	"strings"
	"unicode"
)

func Abbreviate(s string) string {
	var res strings.Builder
	normalized := normalizePunctuation(s)

	for word := range strings.FieldsSeq(normalized) {
		res.WriteByte(word[0])
	}

	return strings.ToUpper(res.String())
}

func normalizePunctuation(s string) string {
	return strings.Map(func(r rune) rune {
		if r == '-' {
			return ' '
		}

		if unicode.IsPunct(r) || unicode.IsSymbol(r) {
			return -1
		}

		return r
	}, s)
}
