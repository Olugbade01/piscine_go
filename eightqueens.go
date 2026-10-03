package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	var board [8]int
	solve(board, 0)
}

func solve(board [8]int, col int) {
	if col == 8 {
		printSolution(board)
		return
	}

	for row := 1; row <= 8; row++ {
		if isSafe(board, col, row) {
			board[col] = row
			solve(board, col+1)
		}
	}
}

func isSafe(board [8]int, col int, row int) bool {
	for i := 0; i < col; i++ {
		// same row
		if board[i] == row {
			return false
		}
		// diagonal check
		if abs(board[i]-row) == col-i {
			return false
		}
	}
	return true
}

func printSolution(board [8]int) {
	for i := 0; i < 8; i++ {
		z01.PrintRune(rune(board[i] + '0'))
	}
	z01.PrintRune('\n')
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
