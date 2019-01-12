package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, errors.New("Strings a and b must be the same length")
	}

	hamDistance := 0

	for i, aVal := range a {
		if string(aVal) != string(b[i]) {
			hamDistance++
		}
	}

	return hamDistance, nil
}
