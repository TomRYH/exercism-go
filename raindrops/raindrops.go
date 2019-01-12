package raindrops

import "strconv"

var conversions = []struct {
	num int
	out string
}{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
}

// Convert takes an integer and converts it
// to a string based on its  factors
func Convert(i int) string {
	retStr := ""

	for _, conv := range conversions {
		if i%conv.num == 0 {
			retStr += conv.out
		}

	}

	if retStr == "" {
		return strconv.Itoa(i)
	}

	return retStr
}
