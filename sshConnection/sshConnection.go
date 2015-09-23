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

type project struct {
	name     string
	hostname string
	pwd      string
	port     string
	typ      string
}

func main() {

	p := assemblyLine()

	// Initialization
	project := &project{
		name:     p[0],
		hostname: p[1],
		pwd:      p[2],
		port:     p[3],
		typ:      p[4],
	}

	// SSH connection config
	config := &ssh.ClientConfig{
		User: project.name,
		Auth: []ssh.AuthMethod{
			ssh.Password(project.pwd),
		},
	}

	command := "cd private && touch hello.txt"
	project.connect(command, config)

}

func (p *project) connect(command string, config *ssh.ClientConfig) {

	log.Printf("Trying connection...\n")

	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", p.hostname, p.port), config)
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
func assemblyLine() []string {
	name := ask4Input("project name: ", "error getting the project's name: ")
	hostname := ask4Input("hostname: ", "error getting the project's hostname: ")
	password := ask4Input("Password: ", "error getting the project's password: ")
	port := ask4Input("port (default 22): ", "error getting the project's port: ")
	typ := ask4Input("Project type [1]yii [2]wp or gohugo: ", "error getting the project's type: ")

	return []string{name, hostname, password, port, typ}
}

// Remove delimitator added by ReadString
func removeDelimitator(s string) string {
	return strings.Fields(s)[0]
}

//
func checkError(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err.Error())
	}
}

// Takes the assemblyLine's data and mount the prompt for the user
func ask4Input(question, errorMsg string) string {
	consolereader := bufio.NewReader(os.Stdin)
	fmt.Print(question)
	s, err := consolereader.ReadString('\n')
	checkError(errorMsg, err)

	return removeDelimitator(s)
}
