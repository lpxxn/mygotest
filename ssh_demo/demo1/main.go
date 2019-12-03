package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/ssh"
)

func main() {

	homeDir := os.Getenv("HOME")
	user := "user"
	key, err := ioutil.ReadFile(homeDir + "/.ssh/a.pem")
	if err != nil {
		panic(err)
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		panic(err)
	}
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			// Add in password check here for moar security.
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", "15.2.3.4:22", config)
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		panic(err)
	}

	defer session.Close()
	combo, err := session.CombinedOutput("whoami; ls -al;pwd")
	if err != nil {
		panic("远程执行cmd 失败" + err.Error())
	}
	fmt.Println("命令输出:", string(combo))

	//cmd := os.Args[1]
	//hosts := os.Args[2:]
	//
	//results := make(chan string, 10)
	//timeout := time.After(10 * time.Second)
	//
	//port := os.Getenv("PORT")
	//if len(port) == 0 {
	//	port = "22"
	//}
	//
	//config := &ssh.ClientConfig{
	//	User: os.Getenv("USER"),
	//	Auth: []ssh.ClientAuth{makeKeyring()},
	//}
	//
	//for _, hostname := range hosts {
	//	go func(hostname string, port string) {
	//		results <- executeCmd(cmd, hostname, port, config)
	//	}(hostname, port)
	//}
	//
	//for i := 0; i < len(hosts); i++ {
	//	select {
	//	case res := <-results:
	//		fmt.Print(res)
	//	case <-timeout:
	//		fmt.Println("Timed out!")
	//		return
	//	}
	//}
}

//func executeCmd(command, hostname string, port string, config *ssh.ClientConfig) string {
//	conn, _ := ssh.Dial("tcp", fmt.Sprintf("%s:%s", hostname, port), config)
//	session, _ := conn.NewSession()
//	defer session.Close()
//
//	var stdoutBuf bytes.Buffer
//	session.Stdout = &stdoutBuf
//	session.Run(command)
//
//	return fmt.Sprintf("%s -> %s", hostname, stdoutBuf.String())
//}
//
//// Use either .ssh/id_rsa or .ssh/id_dsa keys to sign in.
//type SignerContainer struct {
//	signers []ssh.Signer
//}
//
//func (t *SignerContainer) Key(i int) (key ssh.PublicKey, err error) {
//	if i >= len(t.signers) {
//		return
//	}
//	key = t.signers[i].PublicKey()
//	return
//}
//
//func (t *SignerContainer) Sign(i int, rand io.Reader, data []byte) (sig []byte, err error) {
//	if i >= len(t.signers) {
//		return
//	}
//	sig, err = t.signers[i].Sign(rand, data)
//	return
//}
//
//func makeSigner(keyname string) (signer ssh.Signer, err error) {
//	fp, err := os.Open(keyname)
//	if err != nil {
//		return
//	}
//	defer fp.Close()
//
//	buf, _ := ioutil.ReadAll(fp)
//	signer, _ = ssh.ParsePrivateKey(buf)
//	return
//}
//
//func makeKeyring() ssh.ClientAuth {
//	signers := []ssh.Signer{}
//	keys := []string{os.Getenv("HOME") + "/.ssh/id_rsa", os.Getenv("HOME") + "/.ssh/id_dsa"}
//
//	for _, keyname := range keys {
//		signer, err := makeSigner(keyname)
//		if err == nil {
//			signers = append(signers, signer)
//		}
//	}
//
//	return ssh.ClientAuthKeyring(&SignerContainer{signers})
//}
