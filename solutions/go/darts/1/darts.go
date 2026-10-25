package darts

func Score(x, y float64) int {
	pos := x*x + y*y

	// outside target
	if pos > 100 {
		return 0
	}

	// outer circle
	if pos > 25 {
		return 1
	}

	// middle circle
	if pos > 1 {
		return 5
	}

	// inner circle
	return 10
}
