package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/crypto/ssh"
)

// A project is made of project fields which has a program on it.
type program struct {
	setup []string
}

type projectField struct {
	name, label, inputQuestion, errorMsg, validationMsg string
	program                                             program
}

type project struct {
	projectname, hostname, pwd, port, typ projectField
}

func main() {

	// Initialization
	project := new(project)

	// Let's build our project!
	project.assemblyLine()

	// SSH connection config
	config := &ssh.ClientConfig{
		User: project.projectname.name,
		Auth: []ssh.AuthMethod{
			ssh.Password(project.pwd.name),
		},
	}

	var yiiSteps = []string{}

	var wpSteps = []string{
		"echo -e '[User]\nname = Pipi, server girl' > .gitconfig",
		"cd ~/www/www/ && git init",
		"cd ~/www/www/ && git add . ",
		"cd ~/www/www/ && git commit -m 'on the beginning was the commit'",
		"cd ~/private/ && mkdir repos && cd repos && mkdir projectname_hub.git && cd projectname_hub.git && git --bare init",
		"cd ~/www/www && git remote add hub ~/private/repos/projectname_hub.git && git push hub master",
		"cd ~/private/repos/projectname_hub.git/hooks && touch post-update",
		"scp post-update-wp " + project.projectname.name + "@" + project.hostname.name + ":/home/" + project.projectname.name + "/private/repos/" + project.projectname.name + "_hub.git/hooks/post-update",
	}

	// Now we need to know which instalation we going to make.
	// And once we get to know it, let's load the setup with
	// the aproppriate set of files and commands.
	if project.typ.name == "Yii" {
		project.typ.program.setup = yiiSteps
	} else {
		project.typ.program.setup = wpSteps
	}

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
	fmt.Print(field.inputQuestion)

	var input string
	_, err := fmt.Scanln(&input)

	if err != nil && err.Error() == "unexpected newline" && field.label != "port" {
		ask4Input(field)
	} else if err != nil && err.Error() == "unexpected newline" {
		input = "22"
		checkInput(field, input)
	} else if err != nil {
		log.Fatal(field.errorMsg, err)
	}

	checkInput(field, input)
}

// A simple error checker.
func checkError(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err.Error())
	}
}

func checkInput(field *projectField, input string) {

	switch inputLength := len(input); field.label {
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
			input = "22"
		} else if inputLength > 3 {
			fmt.Println(field.validationMsg)
			ask4Input(field)
		}
	case "type":
		if input != "1" && input != "2" {
			fmt.Println(field.validationMsg)
			ask4Input(field)
		} else if input == "1" {
			input = "Yii"
		} else if input == "2" {
			input = "WP"
		}
	}
	field.name = input
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
	checkError("Failed to build session: ", err)
	defer session.Close()

	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf

	log.Printf("Executing command: %s", p.typ.program.setup[step])

	if err := session.Run(p.typ.program.setup[step]); err != nil {
		log.Fatal("Error on command execution", err.Error())
	}
}

func readFile(file string) string {
	data, err := ioutil.ReadFile(file)
	checkError("Error on reading file.", err)
	return string(data[:len(data)])
}
