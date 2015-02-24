/*
* Multiples of 3 and 5
* Problem 1
*
* Find the sum of all multiples of 3 or 5 below 1000.
*/
package main 

import "fmt"

func main() {
	var i, limit, sum uint32
	i = 1
	limit = 1000
	sum = 0
	for i < limit{
		if i%3==0 || i%5==0 {
			sum += i
		}
		i++
	}
  fmt.Println("Result: ",sum)
}