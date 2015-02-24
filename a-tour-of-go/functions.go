package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("usage: %s <Type two whole-numbers to add>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	// take arg[1] and convert to int x
	x, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	// take arg[2] and convert to int y
	y, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	// sum the couple calling function add
	sum := add(x, y)
	fmt.Println(sum)
}

// add do you know what
func add(x, y int) int {
	return x + y
}
