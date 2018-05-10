package main

import (
	"fmt"
)

func insert(n *node, mode int) {
	len := len(open)
	i := 0
	switch mode {
	case 1:
		for i < len && n.f > open[i].f {
			i++
		}
		for i < len && n.g <= open[i].g && n.f == open[i].f {
			i++
		}
	case 2:
		for i < len && n.h > open[i].h {
			i++
		}
	case 3:
		for i < len && n.g > open[i].g {
			i++
		}
	}
	open = append(open[:i], append([]*node{n}, open[i:]...)...)
}

func addNeighbours(n *node, endState [][]int, size int, input []int) {
	x := n.coord[0]
	y := n.coord[1]

	if x+1 < size {
		b := copyBoard(n.board, size)
		b[x][y] = b[x+1][y]
		b[x+1][y] = 0
		if !findBoard(b) {
			h := heuristicSwitch(input, b, endState)
			nodeDown := node{
				board: b,
				id:    n.id + "D",
				g:     n.g + 1,
				h:     h,
				f:     h + n.g + 1,
				coord: [2]int{x + 1, y},
			}
			insert(&nodeDown, input[0])
		}
	}

	if x-1 > -1 {
		b := copyBoard(n.board, size)
		b[x][y] = b[x-1][y]
		b[x-1][y] = 0
		if !findBoard(b) {
			h := heuristicSwitch(input, b, endState)
			nodeUp := node{
				board: b,
				id:    n.id + "U",
				g:     n.g + 1,
				h:     h,
				f:     h + n.g + 1,
				coord: [2]int{x - 1, y},
			}
			insert(&nodeUp, input[0])
		}
	}

	if y+1 < size {
		b := copyBoard(n.board, size)
		b[x][y] = b[x][y+1]
		b[x][y+1] = 0
		if !findBoard(b) {
			h := heuristicSwitch(input, b, endState)
			nodeRight := node{
				board: b,
				id:    n.id + "R",
				g:     n.g + 1,
				h:     h,
				f:     h + n.g + 1,
				coord: [2]int{x, y + 1},
			}
			insert(&nodeRight, input[0])
		}
	}

	if y-1 > -1 {
		b := copyBoard(n.board, size)
		b[x][y] = b[x][y-1]
		b[x][y-1] = 0
		if !findBoard(b) {
			h := heuristicSwitch(input, b, endState)
			nodeLeft := node{
				board: b,
				id:    n.id + "L",
				g:     n.g + 1,
				h:     h,
				f:     h + n.g + 1,
				coord: [2]int{x, y - 1},
			}
			insert(&nodeLeft, input[0])
		}
	}
}

func computeAStar(userInput []int, board [][]int, endState [][]int, size int) {
	h := heuristicSwitch(userInput, board, endState)
	firstNode := node{
		board: copyBoard(board, size),
		id:    "0",
		g:     0,
		h:     h,
		f:     h,
		coord: findZero(board),
	}
	open = append(open, &firstNode)
	for open != nil {
		closed = append([]*node{open[0]}, closed...)
		open = append(open[:0], open[1:]...)
		if len(open) > openMax {
			openMax = len(open)
		}
		switch userInput[0] {
		case 3:
			if !findDiff(closed[0].board, endState) {
				if userInput[2] != 1 {
					printSteps(closed[0].id, userInput[2], endState)
					return
				}
				fmt.Println(xd("solution ID : "+closed[0].id, 1))
				fmt.Println(xd("number of moves required : ", 1), closed[0].g+1)
				fmt.Println(xd("complexity in size : ", 1), openMax)
				fmt.Println(xd("complexity in time : ", 1), len(open)+len(closed))
				return
			}
		default:
			if closed[0].h == 0 {
				if userInput[2] != 1 {
					printSteps(closed[0].id, userInput[2], endState)
					return
				}
				fmt.Println(xd("solution ID : "+closed[0].id, 1))
				fmt.Println(xd("number of moves required : ", 1), closed[0].g+1)
				fmt.Println(xd("complexity in size : ", 1), openMax)
				fmt.Println(xd("complexity in time : ", 1), len(open)+len(closed))
				return
			}
		}
		addNeighbours(closed[0], endState, size, userInput)
	}
}
