package main

import (
  "fmt"
  "os"
  "path/filepath"
  "github.com/alesr/alesr.go.lib/string/reverseString"
  "strings"
)

func main() {
  if len(os.Args) != 2 {
    fmt.Printf("usage: %s <type a word to check if is a palindrome>\n",
    filepath.Base(os.Args[0]))
    os.Exit(1)
  }

  input := os.Args[1]
  word := strings.ToLower(input)

  if word == reverseString.ReverseString(word) {
    fmt.Println("General Hans Landa says: Oooh, that's a bingo!")
  } else {
    fmt.Println("Not this time.")
  }
}
