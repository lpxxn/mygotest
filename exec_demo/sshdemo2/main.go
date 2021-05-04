package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"golang.org/x/crypto/ssh"
)

// github.com/kevinburke/ssh_config
func main() {

	var hostKey ssh.PublicKey

	key, err := ioutil.ReadFile("/Users/li/.ssh/id_rsa")
	if err != nil {
		fmt.Println(err)
		return
	}

	singer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		fmt.Println(err)
		return
	}

	config := &ssh.ClientConfig{
		User:            "root",
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(singer)},
		HostKeyCallback: ssh.FixedHostKey(hostKey),
		Timeout:         30 * time.Second,
	}

	client, err := ssh.Dial("tcp", "172.16.0.178:22", config)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer session.Close()

	var exitcode int

	output, err := session.CombinedOutput(`echo $PATH`)
	if err != nil {

		if ins, ok := err.(*ssh.ExitError); ok {
			exitcode = ins.ExitStatus()
		} else {
			exitcode = ins.ExitStatus()
		}
	}
	fmt.Println(string(output), exitcode)
}
