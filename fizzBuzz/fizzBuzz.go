package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open(os.Args[1])
	checkError(err)
	readData(input)
	defer input.Close()
	os.Exit(0)
}

func readData(input *os.File) {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		// the blue print
		serie := strings.Split(scanner.Text(), " ")
		// first divisor
		x, err := strconv.Atoi(serie[0])
		checkError(err)
		// second divisor
		y, err := strconv.Atoi(serie[1])
		checkError(err)
		// limit to range over
		limit, err := strconv.Atoi(serie[2])
		checkError(err)
		// magic
		fizzBuzz(x, y, limit)
	}
}

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func fizzBuzz(x, y, limit int) {
	for i := 1; i <= limit; i++ {
		switch {
		case i%(x*y) == 0:
			fmt.Print("FB")
		case i%(x) == 0:
			fmt.Print("F")
		case i%(y) == 0:
			fmt.Print("B")
		default:
			fmt.Print(i)
		}
		if i < limit {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
