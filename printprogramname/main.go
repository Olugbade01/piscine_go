package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[0]

	// find last '/'
	start := 0
	for i := 0; i < len(name); i++ {
		if name[i] == '/' {
			start = i + 1
		}
	}

	for _, c := range name[start:] {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
