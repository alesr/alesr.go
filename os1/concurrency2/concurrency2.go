package main

import (
	"fmt"
)

// send "x" to channel c
func routineX(c chan string) {
	for {
		c <- "x"
	}
}

// send "o" to channel c
func routineO(c chan string) {
	for {
		c <- "o"
	}
}

func main() {

	strCh := make(chan string) // create a string channel

	go routineX(strCh)
	go routineO(strCh)

	// print time
	for i := 1; i <= 10; i++ {
		go fmt.Printf("%s", <-strCh)
	}
	for j := 1; j <= 10; j++ {
		go fmt.Printf("%s", <-strCh)
	}

	fmt.Printf("\n")

}
