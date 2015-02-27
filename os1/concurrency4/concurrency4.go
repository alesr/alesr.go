package main

import (
	"fmt"
	"sync"
)

// routineX prints "x" 10 times
func routineX(wgX *sync.WaitGroup) {
	for i := 1; i <= 10; i++ {
		fmt.Print("x")
	}
	wgX.Done() // flag the routine's end
}

// routineO prints "o" 10 times
func routineO(wgO *sync.WaitGroup) {
	for i := 1; i <= 10; i++ {
		fmt.Print("o")
	}
	wgO.Done() // flag the routine's end
}

func main() {

	// sync.WaitGroup will wait the routine's end to finish the program
	var wgX sync.WaitGroup
	var wgO sync.WaitGroup

	wgX.Add(1) // flag the routineX's start
	go routineX(&wgX)

	wgO.Add(1) // flag the routineO's start
	go routineO(&wgO)

	// you shall not pass until wg.Done
	wgX.Wait()
	wgO.Wait()

	fmt.Println()
}