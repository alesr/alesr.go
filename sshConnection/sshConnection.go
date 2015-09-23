package main

import (
	"log"

	"github.com/pkg/sftp"

	"golang.org/x/crypto/ssh"
)

func main() {
	config := &ssh.ClientConfig{
		User: "ap",
		Auth: []ssh.AuthMethod{
			ssh.Password("9"),
		},
	}

	client, err := ssh.Dial("tcp", "apporg:22", config)

	if err != nil {
		log.Fatal(err)
		return
	}

	session, err := client.NewSession()
	checkError("Failed to create session: ", err)

	defer session.Close()

	sftp, err := sftp.NewClient(client)
	if err != nil {
		log.Fatal(err)
	}
	defer sftp.Close()

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
