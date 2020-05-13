// Package diffsquares is Exercism.io exercise
package diffsquares

import "math"

// SquareOfSum — get square of the sum of the first N natural numbers
func SquareOfSum(n int) int{
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	return int(math.Pow(float64(sum), 2))
}

// SumOfSquares — get sum of the squares of the first N natural numbers
func SumOfSquares(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += int(math.Pow(float64(i), 2))
	}
	return sum
}

// Difference — get the difference between the square of the sum of the first N natural numbers
// and the sum of the squares of the first N natural numbers
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}