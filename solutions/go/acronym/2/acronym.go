package acronym

import (
	"strings"
	"unicode"
)

func Abbreviate(s string) string {
	var result strings.Builder
	atWordStart := true

	for _, r := range s {
		// Treat whitespace and hyphens as word separators.
		if unicode.IsSpace(r) || r == '-' {
			atWordStart = true
			continue
		}

		// Ignore punctuation and symbols.
		if unicode.IsPunct(r) || unicode.IsSymbol(r) {
			continue
		}

		if !atWordStart {
			continue
		}

		result.WriteRune(r)
		atWordStart = false
	}

	return strings.ToUpper(result.String())
}
