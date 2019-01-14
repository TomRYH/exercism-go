package grains

import "errors"

// Square calculates the amount of grains on the given
// square by raising 2 to the  squares index minus 2
func Square(square int) (uint64, error) {

	if !(square > 0 && square < 65) {
		return 0, errors.New("number not in range of board")
	}

	if square == 1 {
		return 1, nil
	}

	return 2 << uint(square-2), nil
}

// Total calculates how many total grains by working out
// the amount in final square then doubling and subtracting 1
func Total() uint64 {
	var x uint64
	x = 2 << 62
	return (x * 2) - 1
}
