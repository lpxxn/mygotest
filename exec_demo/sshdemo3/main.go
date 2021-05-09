package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

// https://github.com/nanobox-io/golang-ssh/blob/master/client.go
var (
	Bastion    = "5.3.1.4:22"
	Target     = "xxx.xxx.xxx.xxx:22"
	BastionPem = "/Users/l/.ssh/c1.pem"
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
	state, err := terminal.MakeRaw(fd)
	if err != nil {
		return
	}
	defer terminal.Restore(fd, state)

	termWidth, termHeight, err := terminal.GetSize(fd)
	if err != nil {
		panic(err)
	}

	fmt.Println(termWidth, termHeight)

	// Create a session
	session, err := conn.NewSession()
	if err != nil {
		panic(err)
	}

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
	term := os.Getenv("TERM")
	if err := session.RequestPty(term, termWidth, termHeight, modes); err != nil {
		log.Fatal("request for pseudo terminal failed: ", err)
	}
	// Terminal resize goroutine.
	winch := syscall.Signal(0x1c)
	signalchan := make(chan os.Signal, 1)
	signal.Notify(signalchan, winch)
	go func() {
		for {
			s := <-signalchan
			switch s {
			case winch:
				fd := int(os.Stdout.Fd())
				width, height, err := terminal.GetSize(fd)
				if err != nil {
					session.WindowChange(height, width)
				}
			}
		}
	}()
	if err := session.Shell(); err != nil {
		log.Fatal("failed to start shell: ", err)
	}

	go SendKeepAlive(session)
	// Start remote shell

	if err := session.Wait(); err != nil {
		panic(err)
	}
}

func SendKeepAlive(session *ssh.Session) {
	// keep alive interval (default 30 sec)
	interval := 30

	// keep alive max (default 5)
	max := 5

	// keep alive counter
	i := 0
	for {
		// Send keep alive packet
		_, err := session.SendRequest("keepalive", true, nil)
		// _, _, err := c.Client.SendRequest("keepalive", true, nil)
		if err == nil {
			i = 0
		} else {
			i += 1
		}

		// check counter
		if max <= i {
			session.Close()
			return
		}

		// sleep
		time.Sleep(time.Duration(interval) * time.Second)
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
