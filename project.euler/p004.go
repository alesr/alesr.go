/*
* Largest palindrome product
* Problem 4
*
* Find the largest palindrome made from the product of two 3-digit numbers.
 */
package main

import (
	"fmt"
	"strconv"

	"github.com/alesr/alesr.go.lib/reverseStr"
)

func main() {
	var maxValue int
	// Decrementing i, j from 999,999 to find the largest product and test if it's a palindrome.
	for i := 999; i > 1; i-- {
		for j := 999; j > 1; j-- {
			product := i * j
			strProd := strconv.Itoa(product) // Cast int into string.
			revStr := reverseStr.ReverseStr(strProd)
			if strProd == revStr && product > maxValue { // Check equality and certifies the is the greater product.
				maxValue = product
				break
			}
		}
	}
	fmt.Println("Result: ", maxValue)
}
