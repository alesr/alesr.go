package lineReverse

import (
	"bufio"
	"log"
	"os"

	"github.com/alesr/alesr.go/reverseString"
)

// LineReverse invert the content of all strings in a sentence
// The file must be in the program directory
func LineReverse(filename string, outputName string) {
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

// check and output error
func checkErr(msg string, err error) {
	if err != nil {
		log.Println(msg, err)
		os.Exit(1)
	}
}

func main() {
	LineReverse("input.txt", "output.txt")
}
