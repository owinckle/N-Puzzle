package main

func abs(nb int) int {
	if nb < 0 {
		return -nb
	}
	return nb
}

func getDistance(nb int, x int, y int, endState [][]int) int {
	for i := range endState {
		for j := range endState[i] {
			if nb == endState[i][j] {
				dist := abs(i-x) + abs(j-y)
				return dist
			}
		}
	}
	return -1
}

func computeManhattan(board [][]int, endState [][]int) int {
	value := 0
	for i := range board {
		for j := range board[i] {
			value += getDistance(board[i][j], i, j, endState)
		}
	}
	return value
}

func computeHamming(board [][]int, endState [][]int) int {
	value := 0
	for i := range endState {
		for j := range endState[i] {
			if endState[i][j] != board[i][j] {
				value++
			}
		}
	}
	return value
}

func findEndTile(nb int, endState [][]int, x int, y int) int {
	for i := range endState {
		for j := range endState[i] {
			if endState[i][j] == nb {
				if x != i && y != j {
					return 2
				} else if x == i && y == j {
					return 0
				} else {
					return 1
				}
			}
		}
	}
	return 0
}

func computeRowColumn(board [][]int, endState [][]int) int {
	value := 0
	for i := range board {
		for j := range board[i] {
			if board[i][j] != 0 {
				value += findEndTile(board[i][j], endState, i, j)
			}
		}
	}
	return value
}
