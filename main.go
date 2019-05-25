package main

import "fmt"

func main() {
	puzzleString, err := getArguments()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	data, err := parseArguments(puzzleString)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	shuffle(data)
	result, err := solvePuzzle(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
