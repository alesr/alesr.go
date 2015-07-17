package findAndSumAllMultiples

/*
* Multiples of 3 and 5
* Problem 1
*
* Find the sum of all multiples of 3 or 5 below 1000.
* Expected result: 233168
* More at https://projecteuler.net/problem=1
 */
func FindAndSumAllMultiples(n uint32) uint32 {
	var i, sum uint32
	i = 1
	sum = 0
	for i < n {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
		i++
	}
	return sum
}
