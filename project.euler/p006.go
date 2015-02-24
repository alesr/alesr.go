/*
* Sum square difference
* Problem 6
*
* Find the difference between the sum of the squares of the first
* one hundred natural numbers and the square of the sum.
 */
package main 

import "fmt"

func main() {
	sumOfSq := 0 // Sum of the squares ie. 1^2+2^3...
	sqSum := 0	// Square of the sum ie.  (1+2+...)^3
	for i:= 1; i<=100; i++ {
		sumOfSq += i*i
		sqSum += i
	}
	diff := sqSum*sqSum - sumOfSq
	fmt.Println(diff)
}