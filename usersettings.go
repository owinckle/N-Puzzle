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
	userSettings := make([]int, 3)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print(xd("Select search type:\n1. A* Pathfinding\n2. Greedy Search\n3. Uniform Cost\n\n>> ", 1))
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	switch text {
	case "1\n":
		userSettings[0] = 1
		fmt.Println(xd("Selected search type: A* Pathfinding", 1))
	case "2\n":
		userSettings[0] = 2
		fmt.Println(xd("Selected search type: Greedy Search", 1))
	case "3\n":
		userSettings[0] = 3
		fmt.Println(xd("Selected search type: Uniform Cost", 1))
	default:
		return nil, errors.New(`getUserInput: unknown search type`)
	}

	if userSettings[0] == 3 {
		return userSettings, nil
	}

	reader = bufio.NewReader(os.Stdin)
	fmt.Print(xd("\nSelect a heuristic:\n1. Manhattan\n2. Hamming\n3. Out of Row/Column\n4. Manhattan and Linear Conflict\n\n>> ", 1))

	text, err = reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	switch text {
	case "1\n":
		userSettings[1] = 1
		fmt.Println(xd("Selected heuristic: Manhattan", 1))
	case "2\n":
		userSettings[1] = 2
		fmt.Println(xd("Selected heuristic: Hamming", 1))
	case "3\n":
		userSettings[1] = 3
		fmt.Println(xd("Selected heuristic: Out of Row/Column", 1))
	case "4\n":
		userSettings[1] = 4
		fmt.Println(xd("Selected heuristic: Manhattan and Linear Conflict", 1))
	default:
		return nil, errors.New(`getUserInput: unknown heuristic`)
	}

	reader = bufio.NewReader(os.Stdin)
	fmt.Print(xd("Select display mode:\n1. Data\n2. Step by step\n3. Burst\n\n>> ", 1))

	text, err = reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	switch text {
	case "1\n":
		userSettings[2] = 1
		fmt.Println(xd("Selected display mode: Data", 1))
	case "2\n":
		userSettings[2] = 2
		fmt.Println(xd("Selected display mode: Step by step\nPress 'Enter' to go to the next step", 1))
	case "3\n":
		userSettings[2] = 3
		fmt.Println(xd("Selected display mode: Burst", 1))
	default:
		return nil, errors.New(`getUserInput: unknown display mode`)
	}
	return userSettings, nil
}
