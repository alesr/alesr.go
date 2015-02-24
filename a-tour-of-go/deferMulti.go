package main

import "fmt"

func main() {
	fmt.Println("counting")

	// deferred calls are executed in last-in-first-out order
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
