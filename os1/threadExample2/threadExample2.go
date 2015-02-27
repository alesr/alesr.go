package main

import "fmt"

func routineX(c chan string) {
	for i := 0; i <= 100; i++ {
		c <- "X"
	}
}

func routineO(c chan string) {
	for i := 0; i <= 100; i++ {
		c <- "O"
	}
}

func main() {
	strChannel := make(chan string)

	go routineX(strChannel)
	go routineO(strChannel)

	for s := range strChannel {
		fmt.Print(s)
	}
	// defer close(strChannel)
	// defer fmt.Printf("\n")
}
