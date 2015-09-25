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

type projectField struct {
	name, label, inputQuestion, errorMsg, validationMsg string
	program                                             program
}

type program struct {
	setup []string
}

type project struct {
	projectname, hostname, pwd, port, typ projectField
}

var commonSteps = []string{"cd private && echo 'hello world' > hello.txt"}
var yiiSteps = []string{
	"cd private && echo 'hello world' > hello1.txt",
	"cd private && echo 'hello world' > hello2.txt",
	"cd private && echo 'hello world' > hello3.txt",
	"cd private && echo 'hello world' > hello4.txt",
}
var wpSteps = []string{"cd private && echo 'hello world' > hello.txt"}

func main() {

	// Initialization
	project := new(project)

	project.assemblyLine()

	// SSH connection config
	config := &ssh.ClientConfig{
		User: project.projectname.name,
		Auth: []ssh.AuthMethod{
			ssh.Password(project.pwd.name),
		},
	}

	project.typ.program.setup = yiiSteps

	project.connect(config)

}

func (p *project) assemblyLine() {
	// project name
	p.projectname.inputQuestion = "project name: "
	p.projectname.label = "projectname"
	p.projectname.errorMsg = "error getting the project's name: "
	p.projectname.validationMsg = "make sure you type a valid name for your project (3 to 20 characters)."
	ask4Input(&p.projectname)

	// Hostname
	p.hostname.inputQuestion = "hostname: "
	p.hostname.label = "hostname"
	p.hostname.errorMsg = "error getting the project's hostname: "
	p.hostname.validationMsg = "make sure you type a valid hostname for your project. it must contain '.com', '.pt' or '.org', for example.)."
	ask4Input(&p.hostname)

	// Password
	p.pwd.inputQuestion = "password: "
	p.pwd.label = "pwd"
	p.pwd.errorMsg = "error getting the project's password: "
	p.pwd.validationMsg = "type a valid password. It must contain at least 6 digits"
	ask4Input(&p.pwd)

	// Port
	p.port.inputQuestion = "port (default 22): "
	p.port.label = "port"
	p.port.errorMsg = "error getting the project's port"
	p.port.validationMsg = "only digits allowed. min 0, max 999."
	ask4Input(&p.port)

	// Type
	p.typ.inputQuestion = "project type [1]yii [2]wp or gohugo: "
	p.typ.label = "type"
	p.typ.errorMsg = "error getting the project's type"
	p.typ.validationMsg = "pay attention to the options"
	ask4Input(&p.typ)
}

// Takes the assemblyLine's data and mount the prompt for the user.
func ask4Input(field *projectField) {
	consolereader := bufio.NewReader(os.Stdin)
	fmt.Print(field.inputQuestion)
	input, err := consolereader.ReadString('\n')
	checkError(field.errorMsg, err)
	checkInput(field, input)
}

// A simple error checker.
func checkError(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err.Error())
	}
}

func checkInput(field *projectField, input string) {

	var cleanInput string
	var inputLength int

	if len(input) < 3 {
		ask4Input(field)
	}

	cleanInput = strings.Fields(input)[0]
	inputLength = len(cleanInput)

	fmt.Println(inputLength)

	switch field.label {
	case "projectname":

		if inputLength > 20 {
			fmt.Println(field.validationMsg)
			ask4Input(field)
		}

	case "hostname":
		if inputLength <= 5 {
			fmt.Println(field.validationMsg)
			ask4Input(field)
		}
	case "pwd":
		if inputLength <= 6 {
			fmt.Println(field.validationMsg)
			ask4Input(field)
		}
	case "port":

		if inputLength == 0 {
			cleanInput = "22"
		} else if inputLength > 3 {
			fmt.Println(field.validationMsg)
			ask4Input(field)
		}
	case "type":
		if cleanInput != "1" && cleanInput != "2" {
			fmt.Println(field.validationMsg)
			ask4Input(field)
		} else if cleanInput == "1" {
			cleanInput = "Yii"
		} else if cleanInput == "2" {
			cleanInput = "WP"
		}
	}

	field.name = cleanInput
}

func (p *project) connect(config *ssh.ClientConfig) {

	log.Printf("Trying connection...\n")

	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", p.hostname.name, p.port.name), config)
	checkError("Failed to dial: ", err)
	log.Printf("Connection established.\n")

	for step := range p.typ.program.setup {
		p.install(step, conn)
	}

}

func (p *project) install(step int, conn *ssh.Client) {

	session, err := conn.NewSession()
	checkError("Failed to build the session: ", err)

	defer session.Close()

	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf

	log.Printf("Executing command: %s", p.typ.program.setup[step])

	if err := session.Run(p.typ.program.setup[step]); err != nil {
		log.Fatal("Error on command execution", err.Error())
	}
}
