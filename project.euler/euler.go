package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	// getting action name
	var action string
	action = askAction(action)

	// getting program name
	var programName string
	programName = askProgramName(programName)
	programName = validateProgramName(programName)

	var hasFile bool
	hasFile = checkIfFileExists(programName, action)
	callToAction(action, programName, hasFile)
}

// askAction handler the user input to choose the action
func askAction(action string) string {

	askActionInput := "run, test or create? "
	fmt.Print(askActionInput)
	fmt.Scan(&action)
	action = strings.ToLower(action)
	validAction := action == "run" || action == "test" || action == "create"

	if !validAction {
		fmt.Println("You must choose between run, test or create")
		log.Fatal("Error on action input")
	}
	return action
}

// askProgramName handler the user input to choose the program's name
func askProgramName(programName string) string {

	askProgramNameInput := "program name? "
	fmt.Print(askProgramNameInput)
	fmt.Scan(&programName)

	return programName
}

// check if the program's name is correct
func validateProgramName(s string) string {
	// check if input has a correct length
	if len(s) != 4 {
		fmt.Println("Type the program's name. E.g. \"p003\"")
		log.Fatal("Error on program name input")
	}

	// check if program name starts with letter p
	if strings.ToLower(string(s[0])) != "p" {
		fmt.Println("Program's name must begin with \"p\". E.g. \"p007\"")
		log.Fatal("Error on program name input")
	}

	// check if after letter p, the string is compost of numbers
	var flag uint8
	for i := 1; i < len(s); i++ {
		if _, err := strconv.Atoi(string(s[i])); err == nil {
			flag++
		}
	}
	if flag != 3 {
		fmt.Println("Program's name must begin with \"p\" followed by numbers. E.g. \"p005\"")
		log.Fatal("Error on program name input")
	}
	return s
}

// check if filepath exist
func checkIfFileExists(p, a string) bool {
	if _, err := os.Stat(p); err != nil {
		if os.IsNotExist(err) {
			if a != "create" {
				log.Println("This file does not exists.")
			}
			return false
		}
		log.Print("File exists; processing...")
	}
	return true
}

func askRestart() {
	var restartAction string
	fmt.Print("Start again? (y/n) ")
	fmt.Scan(&restartAction)
	if strings.ToLower(restartAction) != "y" {
		log.Fatal("OK. Bye!")
	}
	main()
}

func callToAction(a, p string, f bool) {
	if !f {
		if a == "create" {
			runAction(a, p)
		} else {
			askRestart()
		}
	} else {
		if a == "run" {
			runAction(a, p)
		} else if a == "test" {
			runAction(a, p)
		} else {
			fmt.Println("File already exist.")
			askRestart()
		}
	}
}

func runAction(a, p string) {
	if a == "run" {
		// go run
		fmt.Println(p)
	} else if a == "test" {
		// go test
		fmt.Println(p)
	} else {
		fmt.Println(p)
		// create program named programName.go
		// write the header
		// create test named programName_test.go
		// write the header
	}
}
