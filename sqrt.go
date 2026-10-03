package piscine

func Sqrt(nb int) int {
	if nb <= 0 {
		return 0
	}
	return sqrtHelper(nb, 1)
}

func sqrtHelper(nb int, guess int) int {
	if guess*guess == nb {
		return guess
	}
	if guess*guess > nb {
		return 0
	}
	return sqrtHelper(nb, guess+1)
}
