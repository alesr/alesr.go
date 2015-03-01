package main

import (
	"bufio"
	"fmt"
	"github.com/alesr/lightEnigma"
	"os"
)

// in case you don't know
func startUsage() {
	fmt.Println("Usage: type encrypt or decrypt.")
	fmt.Println("Example: ./enigma encrypt")
}

// check and output error
func checkErr(msg string, err error) {
	if err != nil {
		fmt.Println(msg, err)
		os.Exit(1)
	}
}

func runEncryption() {

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter text: ")
	// text, _ := reader.ReadString('\n')
	// fmt.Println(text)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the file's name that you want to encrypt: ")
	filename, err := reader.ReadString('\n')
	checkErr("Something wrong with your input. Check if it contains the .text and exists on the directory you are running the program.", err)

	fmt.Print("\nType a integer between 1 and 255 (and do not forget it): ")
	key, err := reader.ReadInt('\n')
	checkErr("Enter a number between 1 and 255.", err)

	fmt.Print("\nEnter the file name for the output file: ")
	output, err := reader.ReadString('\n')
	checkErr("Type the name you want for the output file.", err)

	run(filename, output, key)

}

// starts the magic
func run(filename string, outputName string, key uint8) {
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
		line := scan.Text()
		content += lightEnigma.EnigmaON(line, 1) + "\n" // tcharann!
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
	if len(os.Args) < 1 { // check the console input
		startUsage()
		os.Exit(1)
	} else {
		option := os.Args[1]
		fmt.Println(option)
		if option == "encrypt" {
			runEncryption()
		} else if option == "decrypt " {
			//runDecryption()
		} else {
			startUsage()
		}
	}
}
