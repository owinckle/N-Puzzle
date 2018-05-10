package main

import (
	"fmt"
	"os/exec"
	"time"
	"strconv"
	"bufio"
	"os"
	"math/rand"
	"bytes"
)

func xd(s string, n int) string {
	if xdlol == false {
		return s
	}

	var buffer bytes.Buffer
	n_1 := n - 1
	l_1 := len(s) - 1
	for i, rune := range s {
		buffer.WriteRune(rune)
		if i%n == n_1 && i != l_1 {
			rand.Seed(time.Now().UnixNano())
			color := "\x1b[37;1m"
			switch rand.Intn(6) {
			case 1:
				color = "\x1b[31;1m"
			case 2:
				color = "\x1b[32;1m"
			case 3:
				color = "\x1b[33;1m"
			case 4:
				color = "\x1b[34;1m"
			case 5:
				color = "\x1b[35;1m"
			case 6:
				color = "\x1b[36;1m"
			}
			buffer.WriteString(color)
		}
	}
	return buffer.String()
}

func prettyPrint(num int) {
	whiteSpaces := 5
	numLen := len(strconv.Itoa(num))
	for whiteSpaces - numLen > 0 {
		fmt.Print(" ")
		whiteSpaces--
	}
}

func printBoard(board [][]int) {
	for row := range board {
		for nb := range board[row] {
			switch board[row][nb] {
			case 0:
				fmt.Print("\x1b[31;1m", board[row][nb], "\x1b[0m", "")
				prettyPrint(board[row][nb])
			default:
				fmt.Print("\x1b[32;1m", board[row][nb], "\x1b[0m", "")
				prettyPrint(board[row][nb])
			}
		}
		fmt.Println()
	}
}

func printSteps(id string, userInput int) {
	idIndex := 1
	for i := len(closed) - 1; i > -1; i-- {
		if idIndex <= len(id) {
			if closed[i].id == id[:idIndex] {
				if idIndex < len(id)+1 {
					if userInput == 2 {
						bufio.NewReader(os.Stdin).ReadBytes('\n')
					}
					time.Sleep(time.Millisecond * 200)
					cmd := "clear"
					lsCmd := exec.Command("bash", "-c", cmd)
					lsOut, _ := lsCmd.Output()
					fmt.Print(string(lsOut))
				}
				fmt.Println(xd("ID: " + closed[i].id, 1))
				fmt.Println(xd("----------", 1))
				printBoard(closed[i].board)
				fmt.Println(xd("----------", 1))
				if userInput == 2 {
					fmt.Println("Steps left : ", closed[0].g - idIndex + 1)
				}
				idIndex++

			}
		}
	}
	fmt.Println(xd("number of moves required :", 1), idIndex-1)
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
