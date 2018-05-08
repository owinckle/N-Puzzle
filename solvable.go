package main

func inversions(b [][]int, size int) int {
	row := make([]int, size*size)
	for i := range b {
		for j := range b[i] {
			row[(i*size)+j] = b[i][j]
		}
	}

	inversionsNumber := 0
	for i := range row {
		for j := 0; j < i; j++ {
			if row[j] > row[i] && row[i] != 0 && row[j] != 0 {
				inversionsNumber++
			}
		}
	}

	return inversionsNumber
}

func isSolvable(board [][]int, endState [][]int) bool {
	size := len(board)

	startInversions := inversions(board, size)
	endInversions := inversions(endState, size)

	if size % 2 == 0 {
		startInversions += findZero(board)[0]
		endInversions += findZero(endState)[0]
	}
	return startInversions % 2 == endInversions % 2
}
