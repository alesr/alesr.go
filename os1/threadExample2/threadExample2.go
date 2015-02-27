package main

import (
	"fmt"
	"os"
)

func routineX(x chan string) {
	for i := 0; i <= 1000; i++ {
		fmt.Print("x")
	}
	//close(x)
}

func routineO(x chan string) {
	for i := 0; i <= 1000; i++ {
		fmt.Print("o")
	}
	//close(o)
}

func main() {

	xCh := make(chan string)
	//oCh := make(chan string)

	go routineX(xCh)
	routineO(xCh)

	defer os.Exit(0)
}
