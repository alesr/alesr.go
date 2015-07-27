package primeFactor

// PrimeFactor - Lists all prime factors of a positive integer.
// Does not include repeated factors.
func PrimeFactor(n int) []int {
	i := 2
	var factors []int = make([]int, 0)

	for i <= n {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
			i++
		}
		i++
	}
	return factors
}
