package main

import (
	"bufio"
	"errors"
	"io/ioutil"
	"log"
	"os"
)

var xdlol bool = false

func main() {
	if len(os.Args) != 2 {
		if len(os.Args) != 3 {
			log.Fatal(`Usage: ./puzzle <demo_file> -xd`)
		}
		if len(os.Args) == 3 && os.Args[2] == "-xd" {
			xdlol = true
		} else if len(os.Args) == 3 && os.Args[2] != "-xd" {
			log.Fatal(`Usage: ./puzzle <demo_file> -xd`)
		}
	}

	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(file)
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}

	board, err := parse(content)
	if err != nil {
		log.Fatal(err)
	}

	endState := generateTable(len(board))
	if !isSolvable(board, endState) {
		log.Fatal(errors.New(`unsolvable puzzle`))
	}

	userInput, err := getUserInput()
	if err != nil {
		log.Fatal(err)
	}

	computeAStar(userInput, board, endState, len(board))
	return
}
