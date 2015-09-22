package main

import (
	"io"
	"log"
	"os"
	"sftp"

	"golang.org/x/crypto/ssh"
)

func main() {
	config := &ssh.ClientConfig{
		User: "a",
		Auth: []ssh.AuthMethod{
			ssh.Password("9"),
		},
	}

	client, err := ssh.Dial("tcp", "app:22", config)

	if err != nil {
		log.Fatal(err)
		return
	}

	session, err := client.NewSession()
	checkError("Failed to create session: ", err)

	stdin, err := session.StdinPipe()
	if err != nil {
		log.Fatal("Unable to setup stdin for session: ", err)
	}
	go io.Copy(stdin, os.Stdin)

	stdout, err := session.StdoutPipe()
	if err != nil {
		log.Fatal("Unable to setup stdout for session: ", err)
	}
	go io.Copy(os.Stdout, stdout)

	stderr, err := session.StderrPipe()
	if err != nil {
		log.Fatal("Unable to setup stderr for session: ", err)
	}
	go io.Copy(os.Stderr, stderr)

	defer session.Close()

	// leave your mark
	f, err := sftp.Create("hello.txt")
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte("Hello world!")); err != nil {
		log.Fatal(err)
	}
}

func checkError(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err.Error())
	}
}
