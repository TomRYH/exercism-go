// Package scrabble contains functions for calculating
// the scrabble score of words
package scrabble

import (
	"strings"
)

var tiles = []struct {
	letters string
	score   int
}{
	{"AEIOULNRST", 1},
	{"DG", 2},
	{"BCMP", 3},
	{"FHVWY", 4},
	{"K", 5},
	{"JX", 8},
	{"QZ", 10},
}

// Score calculates the scrabble score of a word
func Score(word string) int {

	score := 0
	word = strings.ToUpper(word)

	for _, letter := range word {
		score += lookup(string(letter))
	}

	return score
}

// lookup the score of a given word
func lookup(s string) int {
	for _, tile := range tiles {
		if strings.Contains(tile.letters, s) {
			return tile.score
		}
	}
	return 0
}
