package main

import (
	"log"

	"fmt"

	"github.com/pkg/sftp"

	"golang.org/x/crypto/ssh"
)

func main() {
	config := &ssh.ClientConfig{
		User: "appanc",
		Auth: []ssh.AuthMethod{
			ssh.Password("91827744"),
		},
	}

	client, err := ssh.Dial("tcp", "appanc.org:22", config)
	checkError("Failed to dial: ", err)

	sftp, err := sftp.NewClient(client)
	checkError("Failure over sftp", err)
	defer sftp.Close()

	session, err := client.NewSession()
	checkError("Failed to create session: ", err)
	defer session.Close()

	d, err := sftp.ReadDir("www/www")
	checkError("Error on reading directory.", err)

	fmt.Println(len(d))
	fmt.Println(d[0].Name())

	// var b bytes.Buffer
	// session.Stdout = &b
	// if err := session.Run("/usr/bin/whoami"); err != nil {
  //   panic("Failed to run: " + err.Error())
	// }
	// fmt.Println(b.String())


}


func checkError(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err.Error())
	}
}
