package main

import (
	"fmt"
)

func heuristicSwitch(input int, board [][]int, endState [][]int) int {
	switch input {
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

func computeSwitch(userInput []int, board [][]int, endState [][]int, size int, showInterface bool) {
	switch userInput[0] {
	case 1:
		computeAStar(userInput[1], board, endState, size, showInterface)
	case 2:
		fmt.Println("[WIP] compute greedy search")
	}
}
