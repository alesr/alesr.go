package main

import (
	"bytes"
	"fmt"
	"github.com/alesr/errorUtil"
	"github.com/alesr/fileUtil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"

	"golang.org/x/crypto/ssh"
)

// A project is made of project fields which has a program on it.
type program struct {
	setup              []string
	postUpdateFilename string
}

type projectField struct {
	name, label, inputQuestion, errorMsg, validationMsg string
	program                                             program
}

type project struct {
	projectname, hostname, pwd, port, typ projectField
}

var postUpdateContent string

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

	// Now we need to know which instalation we going to make.
	// And once we get to know it, let's load the setup with
	// the aproppriate set of files and commands.
	if project.typ.name == "Yii" {

		// Loading common steps into the selected setup
		project.typ.program.setup = []string{}
		project.typ.program.postUpdateFilename = "post-update-yii"

	} else {

		// Loading common steps into the selected setup
		project.typ.program.setup = []string{
			"echo -e '[User]\nname = Pipi, server girl' > .gitconfig",
			"cd ~/www/www/ && git init",
			"cd ~/www/www/ && touch readme.txt && git add . ",
			"cd ~/www/www/ && git commit -m 'on the beginning was the commit'",
			"cd ~/private/ && mkdir repos && cd repos && mkdir " + project.projectname.name + "_hub.git && cd " + project.projectname.name + "_hub.git && git --bare init",
			"cd ~/www/www && git remote add hub ~/private/repos/" + project.projectname.name + "_hub.git && git push hub master",
			"post-update configuration",
			"cd ~/www/www && git remote add hub ~/private/repos/" + project.projectname.name + "_hub.git/hooks && chmod 755 post-update",
			project.projectname.name + ".dev",
			"git clone",
		}
		project.typ.program.postUpdateFilename = "post-update-wp"
	}
	project.connect(config)

	log.Println("Environment configuration done.")
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

	// The port admits empty string as user input. Setting the default value of "22".
	if err != nil && err.Error() == "unexpected newline" && field.label != "port" {
		ask4Input(field)
	} else if err != nil && err.Error() == "unexpected newline" {
		input = "22"
		checkInput(field, input)
	} else if err != nil {
		log.Fatal(field.errorMsg, err)
	}

	// After we've got the input we must check if it's valid.
	checkInput(field, input)
}

// Check invalid parameters on the user input.
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

// Creates a ssh connection between the local machine and the remote server.
func (p *project) connect(config *ssh.ClientConfig) {

	log.Printf("Trying connection...\n")

	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", p.hostname.name, p.port.name), config)
	errorUtil.CheckError("Failed to dial: ", err)
	log.Printf("Connection established.\n")

	session, err := conn.NewSession()
	errorUtil.CheckError("Failed to build session: ", err)
	defer session.Close()

	// Loops over the slice of commands to be executed on the remote.
	for step := range p.typ.program.setup {

		if p.typ.program.setup[step] == "post-update configuration" {
			p.secureCopy(conn)
		} else if p.typ.program.setup[step] == p.projectname.name+".dev" {
			p.makeDirOnLocal(step)
		} else if p.typ.program.setup[step] == "git clone" {
			p.gitOnLocal(step)
		} else {
			p.installOnRemote(step, conn)
		}
	}
}

func (p *project) gitOnLocal(step int) {
	switch p.typ.program.setup[step] {
	case "git clone":
		p.installOnLocal(step)
	}
}

func (p *project) installOnLocal(step int) {
	homeDir := getUserHomeDir()

	if err := os.Chdir(homeDir + string(filepath.Separator) + "sites" + string(filepath.Separator) + p.projectname.name + ".dev/"); err != nil {
		log.Fatal("Failed to change directory.")
	}

	repo := "ssh://" + p.projectname.name + "@" + p.hostname.name + "/home/" + p.projectname.name + "/private/repos/" + p.projectname.name + "_hub.git"

	cmd := exec.Command("git", "clone", repo, ".")
	if err := cmd.Run(); err != nil {
		log.Fatal("Failed to execute git clone: ", err)
	}

}

// Creates a directory on the local machine. Case the directory already exists
// remove the old one and runs the function again.
func (p *project) makeDirOnLocal(step int) {

	log.Println("Creating directory...")

	// Get the user home directory path.
	homeDir := getUserHomeDir()

	// The dir we want to create.
	dir := homeDir + string(filepath.Separator) + "sites" + string(filepath.Separator) + p.typ.program.setup[step]

	// Check if the directory already exists.
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, 0755)
		errorUtil.CheckError("Failed to create directory.", err)
		log.Println(dir + " successfully created.")
	} else {
		log.Println(dir + " already exist.\nRemoving old and creating new...")

		// Remove the old one.
		if err := os.RemoveAll(dir); err != nil {
			log.Fatalf("Error removing %s\n%s", dir, err)
		}
		p.makeDirOnLocal(step)
	}
}

func (p *project) installOnRemote(step int, conn *ssh.Client) {

	// Git and some other programs can send us an unsuccessful exit (< 0)
	// even if the command was successfully executed on the remote shell.
	// On these cases, we want to ignore those errors and move onto the next step.
	ignoredError := "Reason was:  ()"

	// Creates a session over the ssh connection to execute the commands
	session, err := conn.NewSession()
	errorUtil.CheckError("Failed to build session: ", err)
	defer session.Close()

	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf

	log.Printf("Executing command: %s", p.typ.program.setup[step])

	err = session.Run(p.typ.program.setup[step])

	if err != nil && !strings.Contains(err.Error(), ignoredError) {
		log.Printf("Command '%s' failed on execution", p.typ.program.setup[step])
		log.Fatal("Error on command execution: ", err.Error())
	}
}

// Secure Copy a file from local machine to remote host.
func (p *project) secureCopy(conn *ssh.Client) {
	session, err := conn.NewSession()
	errorUtil.CheckError("Failed to build session: ", err)
	defer session.Close()

	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf

	go func() {
		w, _ := session.StdinPipe()
		defer w.Close()
		content := fileUtil.ReadFile(p.typ.program.postUpdateFilename)
		fmt.Fprintln(w, "C0644", len(content), "post-update")
		fmt.Fprint(w, content)
		fmt.Fprint(w, "\x00")
	}()

	if err := session.Run("scp -qrt ~/private/repos/" + p.projectname.name + "_hub.git/hooks"); err != nil {
		log.Fatal("Failed to run SCP: " + err.Error())
	}
}

func getUserHomeDir() string {
	usr, err := user.Current()
	errorUtil.CheckError("Failed to locate user home directory ", err)

	return usr.HomeDir
}
