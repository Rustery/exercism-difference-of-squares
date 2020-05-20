// Package diffsquares is Exercism.io exercise
package diffsquares

// SquareOfSum — get square of the sum of the first N natural numbers
func SquareOfSum(n int) int {
	sum := ((n * n) + n) / 2
	return sum * sum
}

// SumOfSquares — get sum of the squares of the first N natural numbers
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference — get the difference between the square of the sum of the first N natural numbers
// and the sum of the squares of the first N natural numbers
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
