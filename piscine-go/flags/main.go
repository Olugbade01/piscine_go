package flags
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
}

func println(s string) {
	printStr(s)
	z01.PrintRune('\n')
}

func printHelp() {
	println("--insert")
	println("  -i\t Insert the string given to the --insert flag into the string argument.")
	println("--order")
	println("  -o\t Order the string argument.")
	println("--help")
	println("  -h\t Show this help message.")
}

func sortString(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		for j := i + 1; j < len(runes); j++ {
			if runes[i] > runes[j] {
				runes[i], runes[j] = runes[j], runes[i]
			}
		}
	}
	return string(runes)
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		printHelp()
		return
	}

	insert := ""
	order := false
	mainStr := ""
	for _, arg := range args {
		if arg == "--help" || arg == "-h" {
			printHelp()
			return
		} else if len(arg) > 9 && arg[:9] == "--insert=" {
			insert = arg[9:]
		} else if len(arg) > 3 && arg[:3] == "-i=" {
			insert = arg[3:]
		} else if arg == "--order" || arg == "-o" {
			order = true
		} else {
			mainStr = arg
		}
	}

	result := mainStr + insert

	if order {
		result = sortString(result)
	}

	println(result)
}
