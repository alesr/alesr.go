package main

import (
	"fmt"
	"log"
	"bytes"
	//"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func main() {
	config := &ssh.ClientConfig{
		User: "appanc",
		Auth: []ssh.AuthMethod{
			ssh.Password("91827744"),
		},
	}

	command := "cd private && touch hello.txt"
	conn := connect(command, "appanc.org", "22", config)

	fmt.Println(conn)

}


func connect(command string, hostname string, port string, config *ssh.ClientConfig) string {
	log.Printf("Trying connection...\n")

	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", hostname, port), config)
	checkError("Failed to dial: ", err)

	log.Printf("Connection established.\n")

	session, err := conn.NewSession()
	defer session.Close()

	log.Printf("Session created.\n")

  var stdoutBuf bytes.Buffer
  session.Stdout = &stdoutBuf
  if err := session.Run(command); err != nil {
		log.Fatal("Error on command execution", err.Error())
	}

	return fmt.Sprint("%s -> %s", hostname, stdoutBuf.String())
}





func checkError(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err.Error())
	}
}
