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
	}
	return 0
}

func computeSwitch(userInput []int, board [][]int, endState [][]int, size int) {
	switch userInput[0] {
	case 1:
		computeAStar(userInput[1], board, endState, size)
	case 2:
		fmt.Println("[WIP] compute greedy search")
	}
}
