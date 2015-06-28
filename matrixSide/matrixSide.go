// Workout the left and right sides of a matrix n*n
package main

import "fmt"

func main() {
	//fmt.Print(isLeftSide(8, 4))
	fmt.Print(isRightSide(10, 4))
}

// Print the left collumn
func leftSide(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(i * n)
	}
}

// Print the right collumn
func rightSide(n int) {
	for i := 1; i <= n; i++ {
		fmt.Println((n * i) - 1)
	}
}

// Return true if i is on the left collumn
func isLeftSide(i, n int) bool {
	flag := false
	for j := 0; j < n; j++ {
		if n*j == i {
			flag = true
			break
		}
	}
	return flag
}

// Return true if i is on the right collumn
func isRightSide(i, n int) bool {
	flag := false
	for j := 1; j <= n; j++ {
		if (n*j)-1 == i {
			flag = true
			break
		}
	}
	return flag
}
