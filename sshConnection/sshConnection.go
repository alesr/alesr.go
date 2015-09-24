package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/ssh"
)

type ProjectField struct {
	fieldName , fieldError string
}

type Project struct {
	name      ProjectField
	hostname  ProjectField
	pwd       ProjectField
	port      ProjectField
	typ       ProjectField
}

func main() {

	// Initialization
	project := new(Project)

	project.assemblyLine()

	project.name.fieldError = "error getting the project's name: "
	project.hostname.fieldError = "error getting the project's hostname: "
	project.pwd.fieldError = "error getting the project's password: "
	project.port.fieldError = "error getting the project's port"
	project.typ.fieldError = "error getting the project's type"

	// SSH connection config
	config := &ssh.ClientConfig {
		User: project.name.fieldName,
		Auth: []ssh.AuthMethod{
			ssh.Password(project.pwd.fieldName),
		},
	}

	command := "cd private && echo 'hello world' > hello.txt"
	project.connect(command, config)

}

func (p *Project) connect(command string, config *ssh.ClientConfig) {

	log.Printf("Trying connection...\n")

	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", p.hostname.fieldName, p.port.fieldName), config)
	checkError("Failed to dial: ", err)
	log.Printf("Connection established.\n")

	session, err := conn.NewSession()
	checkError("Failed to build the session: ", err)

	defer session.Close()

	log.Printf("Session created.\n")

	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf

	if err := session.Run(command); err != nil {
		log.Fatal("Error on command execution", err.Error())
	}

}

// Assembly questions and error messages for user's input
func (p *Project) assemblyLine() {


	p.name.fieldName = ask4Input("project name: ", p.name.fieldError)

	message, valid := checkInput(p.name.fieldName, "name")

	if !valid {
		fmt.Println(message)
		p.name.fieldName = ask4Input("project name: ", p.name.fieldError)
	}

	// hostnameError := "error getting the project's hostname: "
	p.hostname.fieldName = ask4Input("hostname: ", p.hostname.fieldError)

	message, valid = checkInput(p.hostname.fieldName, "hostname")

	if !valid {
		fmt.Println(message)
		p.hostname.fieldName = ask4Input("hostname: ", p.hostname.fieldError)
	}

	p.pwd.fieldName = ask4Input("Password: ", p.pwd.fieldError)
	p.port.fieldName = ask4Input("port (default 22): ", p.port.fieldError)
	p.typ.fieldName = ask4Input("project type [1]yii [2]wp or gohugo: ", p.typ.fieldError)
}

// Takes the assemblyLine's data and mount the prompt for the user.
func ask4Input(question, errorMsg string) string {
	consolereader := bufio.NewReader(os.Stdin)
	fmt.Print(question)
	s, err := consolereader.ReadString('\n')
	checkError(errorMsg, err)

	return removeDelimitator(s)
}

// Remove delimitator added by ReadString.
func removeDelimitator(s string) string {
	return strings.Fields(s)[0]
}

// Right on hand error checker.
func checkError(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err.Error())
	}
}


// WAITING FOR REFACTORING
func checkInput(input, kind string) (string, bool) {
	switch kind {

		case "name":
			if len(input) <= 2 {
				return "make sure you type a valid name for your project.", false
			}

		case "hostname":
			if len(input) <= 5 {
				return "make sure you type a valid hostname for your project.", false
			}

			if hasDigits(input) {
				return "the hostname must not contain digits", false
			}

			// Em construção. Terminar validação dos restantes inputs...

	}

	return "", true

}

// Iterates over a string to check if it has digits.
func hasDigits(input string) bool {
	for _, value := range input {
		if value >= '0' && value <= '9' {
			return true
		}
	}
	return false
}
