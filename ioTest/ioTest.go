package main

import (
	"bufio"
	"fmt"
	"github.com/alesr/alesr.go.lib/string/reverseString"
	"os"
  "log"
)

var outputName string

// in case you don't know
func usage() {
  log.Println("Usage: enter input.txt and output.txt (to be created)")
  log.Println("eg. ./ioTest myFile.txt myNewFile.txt")
}

// check and output error
func checkErr(msg string, err error) {
	if err != nil {
		fmt.Println(msg, err)
		os.Exit(1)
	}
}

// starts the magic
func run(filename string) {
	// open file
	inputFile, err := os.Open(filename)
	checkErr("Not possible to open this file.\n", err)
	defer inputFile.Close()
	loadFile(inputFile)
}

// read file and store data on a temp string
func loadFile(inputFile *os.File) {
	var content string
	scan := bufio.NewScanner(inputFile)
	for scan.Scan() {
		l := scan.Text()
		content += reverseString.ReverseString(l) + "\n"
	}
	reWrite(content)
}

// write data on a brand new shiny file
func reWrite(content string) *os.File {
	// create file
	outputFile, err := os.Create(outputName)
	checkErr("Some nasty error to create the file.\n", err)
	// it's time to write
	w := bufio.NewWriter(outputFile)
	w.WriteString(content)
	w.Flush()

	return outputFile
}

func main() {
  if len(os.Args) < 3 { // check the console input
    usage()
    os.Exit(1)
  } else {
    outputName := os.Args(3)
    run(os.Args[1])
  }
}
