package main

import (
	"api/utils"
	"fmt"
	"log"

	"golang.org/x/crypto/ssh"
	// "fyne.io/fyne/v2/app"
	// "fyne.io/fyne/v2/widget"
	// "golang.org/x/crypto/ssh/knownhosts"
)

func main() {
	username, password, ip := utils.GetCredentials()
	if username == "" && password == "" && ip == "" {
		return
	}

	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", ip+":22", config)
	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}

	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}
	if err := session.RequestPty("linux", 80, 40, modes); err != nil {
		log.Fatal("request for pseudo terminal failed: ", err)
	}
	// myApp := app.New()
	// myWindow := myApp.NewWindow("GoUploadToServer")

	// content := widget.NewButton("click me", func() {

	// })

	// myWindow.SetContent(content)
	// myWindow.ShowAndRun()

}
