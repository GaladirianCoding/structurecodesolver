package main

//failedMessage is the message that is returned if the puzzle has no solution
const failedMessage = "No solution"

//puzzleRune is a type representing each two charcater pair in the puzzle
type puzzleRune = [2]rune

//puzzle is a type representing a slice of puzzleRunes
type puzzle = []puzzleRune
