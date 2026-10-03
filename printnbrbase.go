package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	// Validate base
	if len(base) < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	for i := 0; i < len(base); i++ {
		if base[i] == '+' || base[i] == '-' {
			z01.PrintRune('N')
			z01.PrintRune('V')
			return
		}
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
	}

	// Handle negative numbers
	if nbr < 0 {
		z01.PrintRune('-')
		if -nbr < 0 {
			// MinInt64: can't negate safely, peel off last digit in negative space
			lastDigit := -(nbr % len(base))
			nbr = -(nbr / len(base))
			PrintNbrBase(nbr, base)
			z01.PrintRune(rune(base[lastDigit]))
			return
		}
		nbr = -nbr
	}

	baseLen := len(base)

	// Recursive conversion
	if nbr >= baseLen {
		PrintNbrBase(nbr/baseLen, base)
	}
	z01.PrintRune(rune(base[nbr%baseLen]))
}
