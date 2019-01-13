/*Package luhn contains functions for verifying
that a number is valid per the Luhn algorithm
https://en.wikipedia.org/wiki/Luhn_algorithm */
package luhn

import (
	"strconv"
	"strings"
)

// Valid checks if a given number is valid per
// the Luhn algorithm
func Valid(toCheck string) bool {

	toCheck = strings.Replace(toCheck, " ", "", -1)

	if len(toCheck) < 2 {
		return false
	}

	var toSum []int
	for i := 0; i < len(toCheck); i++ {
		num, err := strconv.Atoi(toCheck[i : i+1])
		if err != nil {
			return false
		}

		if (len(toCheck)-i)%2 == 0 {
			toSum = append(toSum, doubleRestrictNine(num))
		} else {
			toSum = append(toSum, num)
		}
	}

	if sum(toSum)%10 == 0 {
		return true
	}
	return false
}

func doubleRestrictNine(num int) int {
	num = num * 2
	if num > 9 {
		return num - 9
	}
	return num
}

func sum(toSum []int) int {
	sum := 0
	for _, n := range toSum {
		sum += n
	}

	return sum
}
