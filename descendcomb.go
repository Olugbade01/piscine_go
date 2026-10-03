package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for first := 99; first >= 1; first-- {
		for second := first - 1; second >= 0; second-- {
			z01.PrintRune(rune('0' + first/10))
			z01.PrintRune(rune('0' + first%10))
			z01.PrintRune(' ')
			z01.PrintRune(rune('0' + second/10))
			z01.PrintRune(rune('0' + second%10))

			if !(first == 1 && second == 0) {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}
