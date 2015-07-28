package main

import (
	"fmt"
	"os"
	"time"
)

// send "x" to strChannel
func printXs(c chan string) {
	for {
		c <- "x"
	}
}

// send "x" to strChannel
func printOs(c chan string) {
	for {
		c <- "o"
	}
}

// exits the program after 3 seconds
func sleepTime(c chan string) {
	time.Sleep(3 * time.Second)
	close(c)
	fmt.Printf("\n")
	os.Exit(0)
}

func main() {

	strChannel := make(chan string)

	go printXs(strChannel)
	go printOs(strChannel)
	go sleepTime(strChannel)

	// print the channel
	for s := range strChannel {
		fmt.Print(s)
	}
}
