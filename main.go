package main

import (
	"bufio"
	"errors"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		log.Fatal(`Usage: ./puzzle <demo_file>`)
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

	computeSwitch(userInput, board, endState, len(board))
	return
}
