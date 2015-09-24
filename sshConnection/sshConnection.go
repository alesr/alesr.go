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
	label , errorMsg, validationMsg string
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

	project.name.errorMsg = "error getting the project's name: "
	project.hostname.errorMsg = "error getting the project's hostname: "
	project.pwd.errorMsg = "error getting the project's password: "
	project.port.errorMsg = "error getting the project's port"
	project.typ.errorMsg = "error getting the project's type"

	project.name.validationMsg = "make sure you type a valid name for your project."
	project.hostname.validationMsg = "make sure you type a valid hostname for your project."
	project.pwd.validationMsg = "type a valid password. It must contain at least 6 digits"
	project.port.validationMsg = "type a port between 1 and 999"
	project.typ.validationMsg = "pay attention to the options"


	// SSH connection config
	config := &ssh.ClientConfig {
		User: project.name.label,
		Auth: []ssh.AuthMethod{
			ssh.Password(project.pwd.label),
		},
	}

	command := "cd private && echo 'hello world' > hello.txt"
	project.connect(command, config)

}

func (p *Project) connect(command string, config *ssh.ClientConfig) {

	log.Printf("Trying connection...\n")

	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", p.hostname.label, p.port.label), config)
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


	// Project name
	p.name.label = ask4Input("project name: ", p.name.errorMsg)
	valid := checkInput(p.name.label, "name")
	p.check4ValidInput(valid, "name")

	// Hostname
	p.hostname.label = ask4Input("hostname: ", p.hostname.errorMsg)
	valid = checkInput(p.hostname.label, "hostname")
	p.check4ValidInput(valid, "hostname")

	// Password
	p.pwd.label = ask4Input("Password: ", p.pwd.errorMsg)
	valid = checkInput(p.hostname.label, "pwd")
	p.check4ValidInput(valid, "pwd")

	// Port
	p.port.label = ask4Input("port (default 22): ", p.port.errorMsg)
	valid = checkInput(p.hostname.label, "port")
	p.check4ValidInput(valid, "port")

	// Type
	p.typ.label = ask4Input("project type [1]yii [2]wp or gohugo: ", p.typ.errorMsg)
	valid = checkInput(p.hostname.label, "type")
	p.check4ValidInput(valid, "type")
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
func checkInput(input, kind string) bool {
	switch kind {
		case "name":
			if len(input) <= 2 {
				return false
			}
		case "hostname":
			if len(input) <= 5 {
				return false
			}
		case "pwd":
			if len(input) <= 6 {
				return false
			}
		case "port":
			if len(input) == 0 {
				return false
			}
		case "type":
			if len(input) == 0 || len(input) > 1{
				return false
			}
	}
	return true

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
// fmt.Sprintf("%s:%s", p.hostname.label, p.port.label)
func (p *Project) check4ValidInput(flag bool, kind string)  {
	if !flag {
		fmt.Println(p.name.validationMsg)
		p.name.label = ask4Input("project name: ", p.name.errorMsg)
	}
}
