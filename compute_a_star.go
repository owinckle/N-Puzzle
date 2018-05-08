package main

import (
	"fmt"
	"os/exec"
	"time"
)

func addNeighbours(n *node, endState [][]int, size int, input int) {
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
			open = append(open, &nodeDown)
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
			open = append(open, &nodeUp)
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
			open = append(open, &nodeRight)
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
			open = append(open, &nodeLeft)
		}
	}
}

func computeAStar(input int, board [][]int, endState [][]int, size int) {
	h := heuristicSwitch(input, board, endState)
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
		i := findNodeIndex()
		closed = append([]*node{open[i]}, closed...)
		open = append(open[:i], open[i+1:]...)
		if closed[0].h == 0 {
			printSteps(closed[0].id)
			return
		}
		addNeighbours(closed[0], endState, size, input)
	}
}
