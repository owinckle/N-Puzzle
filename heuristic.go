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

func indexOf(nb int, endState [][]int) [2]int {
	for i := range endState {
		for j := range endState[i] {
			if endState[i][j] == nb {
				return [2]int{i, j}
			}
		}
	}
	return [2]int{0, 0}
}

func checkColumn(x int, y int, board [][]int, endState [][]int) int {
	value := 0
	realPos := indexOf(board[x][y], endState)
	checkPos := indexOf(board[x][y], endState)

	if realPos[1] < y && x == realPos[0] {
		for i := 0; i < y; i++ {
			checkPos = indexOf(board[x][i], endState)
			if checkPos[1] >= y && checkPos[0] == x && board[x][i] != 0 {
				value++
			}
		}
	}

	if realPos[1] > y && x == realPos[0] {
		for i := 0; i > y; i-- {
			checkPos = indexOf(board[i][y], endState)
			if checkPos[1] <= y && checkPos[0] == x && board[i][y] != 0 {
				value++
			}
		}
	}
	return value
}

func checkRow(x int, y int, board [][]int, endState [][]int) int {
	value := 0
	realPos := indexOf(board[x][y], endState)
	checkPos := indexOf(board[x][y], endState)

	if realPos[0] < x && y == realPos[1] {
		for i := 0; i < x; i++ {
			checkPos = indexOf(board[x][i], endState)
			if checkPos[0] >= x && checkPos[1] == y && board[x][i] != 0 {
				value++
			}
		}
	}

	if realPos[0] > x && y == realPos[1] {
		for i := 0; i > x; i-- {
			checkPos = indexOf(board[i][y], endState)
			if checkPos[0] <= x && checkPos[1] == y && board[i][y] != 0 {
				value++
			}
		}
	}
	return value
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

func computeLinearConflict(board [][]int, endState [][]int) int {
	value := 0
	for i := range board {
		for j := range board[i] {
			if board[i][j] != 0 {
				value += checkColumn(i, j, board, endState)
				value += checkRow(i, j, board, endState)
			}
		}
	}
	return value * 2
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

func heuristicSwitch(input []int, board [][]int, endState [][]int) int {
	if input[0] == 3 {
		return 0
	}
	switch input[1] {
	case 1:
		return computeManhattan(board, endState)
	case 2:
		return computeHamming(board, endState)
	case 3:
		return computeRowColumn(board, endState)
	case 4:
		return computeManhattan(board, endState) + computeLinearConflict(board, endState)
	}
	return 0
}
