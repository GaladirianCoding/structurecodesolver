package main

import (
	"math/rand"
	"time"
)

//Remove returns a puzzle with puzzleRune at index i removed
func remove(slice puzzle, s int) puzzle {
	return append(slice[:s], slice[s+1:]...)
}

//Shuffle shuffles a puzzle in place.
func shuffle(vals puzzle) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for n := len(vals); n > 0; n-- {
		randIndex := r.Intn(n)
		vals[n-1], vals[randIndex] = vals[randIndex], vals[n-1]
	}
}
