package main

import "fmt"

func printBoard(board [][]int) {
	for row := range board {
		for nb := range board[row] {
			switch board[row][nb] {
			case 0:
				fmt.Print("\x1b[31;1m", board[row][nb], "\x1b[0m", "		")
			default:
				fmt.Print("\x1b[32;1m", board[row][nb], "\x1b[0m", "		")
			}
		}
		fmt.Println()
	}
}

func printSteps(id string) {
	idIndex := 1
	for i := len(closed) - 1; i > -1; i-- {
		if idIndex <= len(id) {
			if closed[i].id == id[:idIndex] {
				fmt.Println("ID:", closed[i].id)
				printBoard(closed[i].board)
				idIndex++
				if idIndex < len(id)+1 {
					time.Sleep(time.Millisecond * 200)
					cmd := "clear"
					lsCmd := exec.Command("bash", "-c", cmd)
					lsOut, _ := lsCmd.Output()
					fmt.Print(string(lsOut))
				}
			}
		}
	}
	fmt.Println("number of moves required : ", idIndex-1)
}

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
