package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func getUserInput() ([]int, error) {
	cmd := "clear"
	lsCmd := exec.Command("bash", "-c", cmd)
	lsOut, _ := lsCmd.Output()
	fmt.Print(string(lsOut))
	userSettings := make([]int, 2)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Select search type:\n1. A* Pathfinding\n2. Greedy Search\n\n>> ")

	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	switch text {
	case "1\n":
		userSettings[0] = 1
		fmt.Println("Selected search type: A* Pathfinding")
	case "2\n":
		userSettings[0] = 2
		fmt.Println("Selected search type: Greedy Search")
	default:
		return nil, errors.New(`getUserInput: unknown search type`)
	}

	reader = bufio.NewReader(os.Stdin)
	fmt.Print("\nSelect a heuristic:\n1. Manhattan\n2. Hamming\n3. Out of Row/Column\n4. Manhattan and Linear Conflict\n5. All\n\n>> ")

	text, err = reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	switch text {
	case "1\n":
		userSettings[1] = 1
		fmt.Println("Selected heuristic: Manhattan")
	case "2\n":
		userSettings[1] = 2
		fmt.Println("Selected heuristic: Hamming")
	case "3\n":
		userSettings[1] = 3
		fmt.Println("Selected heuristic: Out of Row/Column")
	case "4\n":
		userSettings[1] = 4
		fmt.Println("Selected heuristic: Manhattan and Linear Conflict")
	default:
		return nil, errors.New(`getUserInput: unknown heuristic`)
	}
	return userSettings, nil
}
