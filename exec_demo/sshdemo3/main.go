package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

// https://github.com/nanobox-io/golang-ssh/blob/master/client.go
var (
	Bastion    = "5.8.19.4:22"
	Target     = "xxx.xxx.xxx.xxx:22"
	BastionPem = "/Users/l/.ssh/cw1-.pem"
	DestPem    = "/Users/me/.ssh/xxx-dest.pem"
	Timeout    = 30 * time.Second
)

func main() {
	bastionSSH, err := BastionConfig()
	if err != nil {
		panic(err)
	}
	conn, err := ssh.Dial("tcp", Bastion, bastionSSH)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fd := int(os.Stdin.Fd())
	//state, err := terminal.MakeRaw(fd)
	//if err != nil {
	//	log.Fatal(err)
	//}
	termWidth, termHeight, err := terminal.GetSize(fd)
	if err != nil {
		termWidth = 80
		termHeight = 24
	}

	fmt.Println(termWidth, termHeight)
	// Create a session
	session, err := conn.NewSession()
	//defer terminal.Restore(fd, state)

	defer session.Close()
	// 成功
	//session.Stdout = os.Stdout
	//session.Stderr = os.Stderr
	//session.Run("ls /; pwd;")

	// 成功
	// Set up terminal modes
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}
	// Request pseudo terminal
	//if err := session.RequestPty("xterm", termWidth, termHeight, modes); err != nil {
	// vt100  VT220 xterm  xterm-256color ansi
	// echo $TERM  echo $SHELL echo $0
	if err := session.RequestPty("xterm-256color", termWidth, termHeight, modes); err != nil {
		log.Fatal("request for pseudo terminal failed: ", err)
	}
	// Start remote shell
	if err := session.Shell(); err != nil {
		log.Fatal("failed to start shell: ", err)
	}
	session.Run("ls /; pwd;")
	if err := session.Wait(); err != nil {
		panic(err)
	}
}

func BastionConfig() (*ssh.ClientConfig, error) {
	pemBytes, err := ioutil.ReadFile(BastionPem)
	if err != nil {
		log.Fatal(err)
	}
	signer, err := ssh.ParsePrivateKey(pemBytes)
	if err != nil {
		log.Fatalf("parse key failed:%v", err)
	}
	config := &ssh.ClientConfig{
		User:    "ec2-user",
		Auth:    []ssh.AuthMethod{ssh.PublicKeys(signer)},
		Timeout: Timeout,
		//HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		//	return nil
		//},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	return config, err
}

func DestConfig() (*ssh.ClientConfig, error) {
	pemBytes, err := ioutil.ReadFile(DestPem)
	if err != nil {
		log.Fatal(err)
	}
	signer, err := ssh.ParsePrivateKey(pemBytes)
	if err != nil {
		log.Fatalf("parse key failed:%v", err)
	}
	config := &ssh.ClientConfig{
		User:    "ec2-user",
		Auth:    []ssh.AuthMethod{ssh.PublicKeys(signer)},
		Timeout: Timeout,
	}
	return config, err
}
func Connect() {
	config, _ := BastionConfig()
	bClient, err := ssh.Dial("tcp", Bastion, config)
	if err != nil {
		log.Fatal("dial bastion error:", err)
	}
	log.Println("dial bastion ok...")
	// Dial a connection to the service host, from the bastion
	conn, err := bClient.Dial("tcp", Target)
	if err != nil {
		log.Fatal("dial target error", err)
	}
	targetConfig, _ := DestConfig()
	ncc, chans, reqs, err := ssh.NewClientConn(conn, Target, targetConfig)
	if err != nil {
		log.Fatal("new target conn error:", err)
	}
	log.Printf("target conn[%s] ok\n", Target)

	targetClient := ssh.NewClient(ncc, chans, reqs)
	if err != nil {
		log.Fatalf("target ssh error:%v", err)
	}

	session, err := targetClient.NewSession()

	if err != nil {
		log.Fatalf("session failed:%v", err)
	}
	defer session.Close()
	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	err = session.Run("hostname")
	if err != nil {
		log.Fatalf("Run failed:%v", err)
	}
	log.Printf(">%s", stdoutBuf)

}
