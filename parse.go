package main

import (
	"errors"
	"strconv"
	"strings"
)

func parse(content []byte) ([][]int, error) {

	slicedContent := strings.Split(string(content), "\n")
	for i, slice := range slicedContent {
		slicedContent[i] = strings.Trim(strings.Split(slice, "#")[0], " ")
	}

	maxLen := len(slicedContent)
	for i := 0; i < maxLen; i++ {
		if slicedContent[i] == "" {
			slicedContent = append(slicedContent[:i], slicedContent[i+1:]...)
			maxLen--
			i--
		}
	}

	size, err := strconv.Atoi(slicedContent[0])
	if err != nil {
		return nil, err
	}

	if len(slicedContent) != size+1 {
		return nil, errors.New(`parse: too much (or too little) shit`)
	}

	board := make([][]int, size)
	for i := range board {
		board[i] = make([]int, size)
		slice := strings.Fields(slicedContent[i+1])
		if len(slice) != size {
			return nil, errors.New(`parse: row "` + slicedContent[i+1] + `": size mismatch`)
		}
		for j := range board[i] {
			board[i][j], err = strconv.Atoi(slice[j])
			if err != nil {
				return nil, err
			}
		}
	}

	checker := make([]int, size*size)
	for i := range board {
		for j := range board[i] {
			if board[i][j] > (size*size-1) || board[i][j] < 0 {
				return nil, errors.New(`parse: number value invalid`)
			}
			switch checker[board[i][j]] {
			case 0:
				checker[board[i][j]] = 1
			default:
				return nil, errors.New(`parse: multi-occurence of number value`)
			}
		}
	}

	return board, nil
}
