package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

// Output the number of days for a selected month
func main() {

	consolereader := bufio.NewReader(os.Stdin)

	fmt.Print("Month: ")

	input, err := consolereader.ReadString('\n') // this will prompt the user for input
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Removing ReadString's delimiters
	s := strings.Fields(input)
	month := s[0]

	month = strings.ToLower(month)

	var x float64

	switch month {
	case "january", "janeiro", "jan", "01", "1":
		x = 1
	case "february", "fevereiro", "feb", "02", "2":
		x = 2
	case "march", "março", "mar", "03", "3":
		x = 3
	case "april", "abril", "apr", "04", "4":
		x = 4
	case "may", "maio", "05", "5":
		x = 5
	case "june", "junho", "06", "6":
		x = 6
	case "july", "julho", "07", "7":
		x = 7
	case "august", "agosto", "aug", "08", "8":
		x = 8
	case "september", "setembro", "stp", "09", "9":
		x = 9
	case "october", "outubro", "oct", "10":
		x = 10
	case "november", "novembro", "nov", "11":
		x = 11
	case "december", "dezembro", "dec", "12":
		x = 12
	default:
		fmt.Println("Enter a valid month name")
		os.Exit(1)
	}

	var days = 28 + byte((x+math.Floor(x/8)))%2 + 2%byte(x) + byte(2*math.Floor(1/x))

	fmt.Println(days)
}
