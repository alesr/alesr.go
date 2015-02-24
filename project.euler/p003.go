/*
* Largest prime factor
* Problem 3
*
* The prime factors of 13195 are 5, 7, 13 and 29.
* What is the largest prime factor of the number 600851475143?
*/
package main 

import ("fmt";
"github.com/alesr/alesr.go.lib/primes/primeFactor")

func main() {
	f := primeFactor.PrimeFactor(600851475143)
	fmt.Println("Result: ",f[len(f)-1:])
}