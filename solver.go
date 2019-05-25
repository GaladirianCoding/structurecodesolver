package main

import "errors"

func solvePuzzle(puzzleToSolve puzzle) (string, error) {
	for i, pair := range puzzleToSolve {
		newList := remove(append(puzzleToSolve[:0:0], puzzleToSolve...), i)
		solution, err := solvePuzzleRecursive(newList, string(pair[:]))
		if err == nil {
			return solution, err
		}
	}
	return "", errors.New(failedMessage)
}

func solvePuzzleRecursive(puzzleToSolve puzzle, solution string) (string, error) {
	if len(puzzleToSolve) == 0 {
		return solution, nil
	}
	for i, pair := range puzzleToSolve {
		if pair[0] == rune(solution[len(solution)-1]) {
			solution += string(pair[1])
			newList := remove(append(puzzleToSolve[:0:0], puzzleToSolve...), i)
			result, err := solvePuzzleRecursive(newList, solution)
			if err == nil {
				return result, err
			}
			solution = solution[:len(solution)-1]
		}
	}
	return "", errors.New(failedMessage)
}
