package main

import (
	"fmt"
)

// send "x" to channel strCh1
func routineX(x chan string) {
	for {
		x <- "x"
	}
}

// send "o" to channel strCh2
func routineO(o chan string) {
	for {
		o <- "o"
	}
}

func main() {

	strChX := make(chan string) // create a string channel to receive "x"
	strChO := make(chan string) // create a string channel to receive "o"

	// calling troops
	go routineX(strChX)
	go routineO(strChO)

	// printing both channels
	for j := 0; j <= 1000; j++ {
		go fmt.Printf("%s", <-strChX)
		go fmt.Printf("%s", <-strChO)
	}

	// closing channels
	close(strChX)
	close(strChO)

	fmt.Printf("\n")
}
