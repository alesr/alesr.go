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
func run(filename string, outputName string) {
	// open file
	inputFile, err := os.Open(filename)
	checkErr("Not possible to open this file.\n", err)
	defer inputFile.Close()
	reWrite(loadFile(inputFile), outputName)
}

// read file and store data on a temp string
func loadFile(inputFile *os.File) string {
	var content string
	scan := bufio.NewScanner(inputFile)
	for scan.Scan() {
		l := scan.Text()
		content += reverseString.ReverseString(l) + "\n"
	}
	return content
}

// write data on a brand new shiny file
func reWrite(content string, outputName string) *os.File {
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
    run(os.Args[1], os.Args[2])
  }
}
