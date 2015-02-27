package main

import (
	"fmt"
)

// send "x" to channel c
func routineX(c chan string, done chan bool) {
	for {
		select {
		case c <- "x":
		case <-done:
			return
		}
	}
}

// send "o" to channel c
func routineO(c chan string, done chan bool) {
	for {
		select {
		case c <- "o":
		case <-done:
			return
		}
	}
}

func main() {

	strCh := make(chan string) // create a string channel
	doneCh := make(chan bool)

	go routineX(strCh, doneCh)
	go routineO(strCh, doneCh)

	// print time
	for i := 1; i <= 10; i++ {
		go fmt.Printf("%s", <-strCh)
	}
	for k := 1; k <= 10; k++ {
		go fmt.Printf("%s", <-strCh)
	}

	defer fmt.Printf("\n")

}
