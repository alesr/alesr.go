package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	var result float64
	for {
		z = z - ((z*z - x) / 2 * z)
		if z-math.Sqrt(2) < 0.01 {
			result = z
			break
		}
	}
	return result

}

func main() {
	fmt.Println(Sqrt(2))
}
