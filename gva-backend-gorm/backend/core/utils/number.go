package utils

import "golang.org/x/exp/constraints"

// IsEven returns true if the given number is even, false otherwise
func IsEven[T constraints.Integer](n T) bool {
	return n%2 == 0
}

// IsOdd returns true if the given number is odd, false otherwise
func IsOdd[T constraints.Integer](n T) bool {
	return n%2 != 0
}

// Sum returns the sum of the given numbers
func Sum[T constraints.Integer](numbers []T) T {
	var total T
	for _, number := range numbers {
		total += number
	}
	return total
}
