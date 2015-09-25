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

// DEBUG ALERT --> Uma instalação contemém métodos relacionados aos comandos.
// e.g. adicionar comando, executar commando, ...
// Yii é um tipo de Programa que contém commandos e implementa a interface
// Installer.
// Wp é outro tipo de Programa que contém commandos e implementa a interface
// Installer.
// Quaisquer novos sotfwares a serem instalados por este programas deverão ser
// do tipo Programa e implementar a interface Installer.
type ProjectField struct {
	name, label, inputQuestion, errorMsg, validationMsg string
}

type Program struct {
	ProjectField
	commands []string
}

type Project struct {
	projectname, hostname, pwd, port ProjectField
	typ                              Program
}

type Installer interface {
	connect()
}

func main() {

	// Initialization
	project := new(Project)

	project.assemblyLine()

	// SSH connection config
	config := &ssh.ClientConfig{
		User: project.projectname.name,
		Auth: []ssh.AuthMethod{
			ssh.Password(project.pwd.name),
		},
	}

	// commands := []string{
	// 	"cd private && echo 'hello world' > hello.txt",
	// 	}
	// project.connect(commands, config)

}

func (p *Project) assemblyLine() {

	// Project name
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
func ask4Input(field *ProjectField) {
	consolereader := bufio.NewReader(os.Stdin)
	fmt.Print(field.inputQuestion)
	input, err := consolereader.ReadString('\n')
	checkError(field.errorMsg, err)
	checkInput(field, input)
}

// Right on hand error checker.
func checkError(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err.Error())
	}
}

func checkInput(field *ProjectField, input string) {

	var cleanInput string
	var inputLength int

	if len(input) > 1 {
		cleanInput = strings.Fields(input)[0]
		inputLength = len(cleanInput)
	}

	switch field.label {
	case "projectname":

		if inputLength < 3 || inputLength > 20 {
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

func (p *Project) connect(commands []string, config *ssh.ClientConfig) {

	log.Printf("Trying connection...\n")

	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", p.hostname.name, p.port.name), config)
	checkError("Failed to dial: ", err)
	log.Printf("Connection established.\n")

	session, err := conn.NewSession()
	checkError("Failed to build the session: ", err)

	defer session.Close()

	log.Printf("Session created.\n")

	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf

	// log.Printf("Preparing to execute command: %s", commands[0])
	// if err := session.Run(commands[0]); err != nil {
	// 	log.Fatal("Error on command execution", err.Error())
	// }

	log.Printf("Preparing to execute command: %s", p.typ.command[0])
}
