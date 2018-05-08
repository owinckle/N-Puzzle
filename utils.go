package main

func copyBoard(board [][]int, size int) [][]int {
	copy := make([][]int, size)
	for i := 0; i < size; i++ {
		copy[i] = make([]int, size)
	}
	for i := range board {
		for j := range board[i] {
			copy[i][j] = board[i][j]
		}
	}
	return copy
}

func findDiff(board [][]int, cmp [][]int) bool {
	for i := range board {
		for j := range board[i] {
			if board[i][j] != cmp[i][j] {
				return true
			}
		}
	}
	return false
}

func findBoard(board [][]int) bool {
	for i := range closed {
		if !findDiff(board, closed[i].board) {
			return true
		}
	}
	return false
}

func findZero(board [][]int) [2]int {
	x := 0
	for i := range board {
		y := 0
		for j := range board[i] {
			if board[i][j] == 0 {
				return ([2]int{x, y})
			}
			y++
		}
		x++
	}
	return [2]int{}
}
