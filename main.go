package main

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

func main() {
	pass := "T$2z-artek"
	host := "172.16.0.1"
	user := "root"

	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{ssh.Password(pass)},
	}
	sshConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey()

	client, err := ssh.Dial("tcp", host, sshConfig)
	if err != nil {
		panic(err)
	}

	session, err := client.NewSession()
	if err != nil {
		client.Close()
		panic(err)
	}

	out, err := session.CombinedOutput("system-view")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
	client.Close()
}
