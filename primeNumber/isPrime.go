package primes

import "math"

// IsPrime checks if a given n number is prime returnig a boolean
func IsPrime(n int) bool {
	// Considering primes are greater than 1
	if n < 2 {
		return false
	}

	// For even, only true if equal to two.
	if n%2 == 0 {
		return n == 2
	}

	root := int(math.Sqrt(float64(n)))
	// Tries to divide n for all odd numbers from 3 to n
	for i := 3; i <= root; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
