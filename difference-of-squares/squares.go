/*Package diffsquares contains functions for finding
the difference between the sum of squares and sqaure of sum
for a a given number*/
package diffsquares

// SquareOfSum finds the square of the sum
// of natural numbers bound by 0 and n (inclusive)
func SquareOfSum(num int) int {

	squareOfSum := num * (num + 1) / 2
	squareOfSum = sqr(squareOfSum)

	return squareOfSum
}

// SumOfSquares finds the sum of all squared
// natural numbers bound by 0 and n (inclusive)
func SumOfSquares(num int) int {
	sumOfSquares := 0

	for num != 0 {
		sumOfSquares += sqr(num)
		num--
	}

	return sumOfSquares
}

// Difference finds the difference between the the sum of squares and square of sum
func Difference(num int) int {
	return SquareOfSum(num) - SumOfSquares(num)
}

func sqr(num int) int {
	return num * num
}
