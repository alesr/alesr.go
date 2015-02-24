/*
* Even Fibonacci numbers
* Problem 2
*
* By considering the terms in the Fibonacci sequence whose values do not exceed four million,
* find the sum of the even-valued terms.
*/
package main 

import "fmt"

func main() {
	var f2,f1,sum uint = 0,1,0
	var fn uint = f2 + f1
	for fn<4000000 {
		if fn%2==0 {
			sum += fn
		}
		f2 = f1
		f1 = fn
		fn = f2 + f1
	}
	fmt.Println("Result: ",sum)	
}
