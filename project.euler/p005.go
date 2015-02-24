/*
* Smallest multiple
* Problem 5
*
* What is the smallest positive number that is evenly
* divisible by all of the numbers from 1 to 20?
 */
package main 

import "fmt"
/*
* Too much pain to solve using math.
* So, let's be brute.
 */
func main() {
	flag := false
	k := 1
	for flag != true {
		i:=1
		for i<=20 {
			if k%i==0{
				flag = true
			} else {
				flag = false
				break
			}
			i++
		}
		k++
	}
	fmt.Println(k-1)
}