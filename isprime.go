package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	if nb == 2 {
		return true
	}
	if nb%2 == 0 {
		return false
	}
	return checkPrime(nb, 3)
}

func checkPrime(nb int, divisor int) bool {
	if divisor*divisor > nb {
		return true
	}
	if nb%divisor == 0 {
		return false
	}
	return checkPrime(nb, divisor+2)
}
