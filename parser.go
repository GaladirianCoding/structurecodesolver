package main

import (
	"errors"
	"os"
	"strings"
)

func getArguments() (string, error) {
	args := os.Args[1:]

	if len(args) < 1 {
		return "", errors.New("Missing puzzle argument." +
			"\nPlease enter the puzzle as a comma-deliniated list of character pairs.")
	}

	return args[0], nil
}

func parseArguments(puzzleString string) (puzzle, error) {
	tuples := strings.Split(strings.ToUpper(puzzleString), ",")

	returnSlice := puzzle{}
	for _, tuple := range tuples {
		if len(tuple) < 2 {
			return nil,
				errors.New("puzzle rune \"" + tuple + "\" was too short. Please ensure each rune is two characters long")
		}
		returnSlice = append(returnSlice, puzzleRune{rune(tuple[0]), rune(tuple[1])})
	}

	return returnSlice, nil
}
