package main

import (
	"fmt"
	"github.com/alesr/lightEnigma"
)

func main() {
	str := "The Secret Life of Walter Mitty"
	enigON := lightEnigma.EnigmaON(str, 150)
	fmt.Println(enigON)
	fmt.Println()
	enigONstr := string(enigON)
	enigOFF := lightEnigma.EnigmaOFF(enigONstr, 150)
	fmt.Println(enigOFF)
}
