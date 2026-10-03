package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	upper := false

	if len(args) > 0 && args[0] == "--upper" {
		upper = true
		args = args[1:]
	}

	if len(args) == 0 {
		return
	}

	for _, arg := range args {
		n := 0
		valid := true

		for _, c := range arg {
			if c < '0' || c > '9' {
				valid = false
				break
			}
			n = n*10 + int(c-'0')
		}

		if !valid || n < 1 || n > 26 {
			z01.PrintRune(' ')
			continue
		}

		ch := rune('a' + n - 1)
		if upper {
			ch = rune('A' + n - 1)
		}

		z01.PrintRune(ch)
	}

	// ✅ REQUIRED newline
	z01.PrintRune('\n')
}
