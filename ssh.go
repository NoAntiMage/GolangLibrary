package main

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/ssh"
)

func main() {
	sshHost := "localhost"
	sshUser := "root"
	sshPassword := "psword"
	sshType := "password"
	sshPort := 22

	config := &ssh.ClientConfig{
		Timeout:         time.Second * 3,
		User:            sshUser,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	if sshType == "password" {
		config.Auth = []ssh.AuthMethod{ssh.Password(sshPassword)}
	} else {
		return
	}

	addr := fmt.Sprintf("%s:%d", sshHost, sshPort)
	sshClient, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		log.Fatal("failed to ssh", err)
	}
	defer sshClient.Close()

	session, err := sshClient.NewSession()
	if err != nil {
		log.Fatal("failed to session ", err)
	}

	combo, err := session.CombinedOutput("ls")
	if err != nil {
		log.Fatal("remote cmd failed", err)
	}
	log.Println("cmd output: ", string(combo))
}
