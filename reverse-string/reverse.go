// Package reverse contains a string reversal function
package reverse

// String reverses the given string
func String(s string) string {
	retStr := ""
	for _, l := range s {
		retStr = string(l) + retStr
	}

	return retStr
}
