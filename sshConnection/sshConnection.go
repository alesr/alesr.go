package main

import (
	"fmt"
	"log"
	"bytes"
	"os"
	"bufio"
	//"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func main() {
	config := &ssh.ClientConfig{
		User: "a",
		Auth: []ssh.AuthMethod{
			ssh.Password("91"),
		},
	}

	ask4Input()

	command := "cd private && touch hello.txt"
	connect(command, "app@org", "22", config)

}


func connect(command string, hostname string, port string, config *ssh.ClientConfig) string {
	log.Printf("Trying connection...\n")

	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", hostname, port), config)
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

	return fmt.Sprint("%s -> %s", hostname, stdoutBuf.String())
}


func ask4Input() string {

	consolereader := bufio.NewReader(os.Stdin)

	fmt.Println("Project name: ")
	projectname, err := consolereader.ReadString('\n')
	checkError("Error getting the project name:", err)

	return projectname
}


func checkError(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err.Error())
	}
}
