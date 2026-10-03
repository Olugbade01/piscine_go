package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	var count [10]int

	// extract digits
	for n > 0 {
		count[n%10]++
		n /= 10
	}

	// print digits in ascending order
	for i := 0; i < 10; i++ {
		for count[i] > 0 {
			z01.PrintRune(rune(i + '0'))
			count[i]--
		}
	}
}
