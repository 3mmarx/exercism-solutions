// Package leap should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package leap

// IsLeapYear accepts a year and returns weather it's a leap year or not.
func IsLeapYear(year int) bool {
	if year%4 != 0 {
		return false
	}

	// # A year must be divisible by 4, except century years, which must also be divisible by 400.

	if year%100 == 0 {
		return year%400 == 0
	}

	return true
}
