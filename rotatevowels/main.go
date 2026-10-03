package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isVowel(c rune) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
		c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		z01.PrintRune('\n')
		return
	}

	var vowels []rune

	// collect vowels
	for _, arg := range args {
		for _, c := range arg {
			if isVowel(c) {
				vowels = append(vowels, c)
			}
		}
	}

	// if no vowels, print original args
	if len(vowels) == 0 {
		for i, arg := range args {
			for _, c := range arg {
				z01.PrintRune(c)
			}
			if i != len(args)-1 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
		return
	}

	// reverse vowels and replace
	k := len(vowels) - 1

	for i, arg := range args {
		for _, c := range arg {
			if isVowel(c) {
				z01.PrintRune(vowels[k])
				k--
			} else {
				z01.PrintRune(c)
			}
		}
		if i != len(args)-1 {
			z01.PrintRune(' ')
		}
	}

	z01.PrintRune('\n')
}
