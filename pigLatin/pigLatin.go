package main

import (
	"bufio"
	"fmt"
	"github.com/alesr/alesr.go.lib/pigLatin"
	"os"
	"strings"
)

func main() {
	consolereader := bufio.NewReader(os.Stdin)

	fmt.Print("To Pig Latin: ")

	input, err := consolereader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Removing ReadString's delimiters
	s := strings.Fields(input)

	var word string
	for i := range s {
		word += pigLatin.PigLatinTranslator(s[i])
		word += " "
	}

	fmt.Println(word)
}
