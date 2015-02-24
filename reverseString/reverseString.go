package main

import (
	"github.com/alesr/alesr.go.lib/string/reverseString"
	"bufio"
	"fmt"
	"os"
  "strings"
)

func main() {
	consolereader := bufio.NewReader(os.Stdin)

	fmt.Print("Type a word to reverse: ")

	input, err := consolereader.ReadString('\n') // this will prompt the user for a input

  if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Removing ReadString's delimiters
	s := strings.Fields(input)

	word := strings.ToLower(s[0])

	reverse := reverseString.ReverseString(word)

	fmt.Println(reverse)
}
