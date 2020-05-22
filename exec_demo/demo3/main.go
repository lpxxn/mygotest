package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

// ok
func main() {
	//cmd := exec.Command("ssh", "cafetest1dev", "ls; whoami; cd baseinfo; pwd;")
	cmd := exec.Command("ssh", "cafe2dev", "whoami;sudo systemctl restart userrelations;")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out.String())
}

// https://stackoverflow.com/questions/37679939/how-do-i-execute-a-command-on-a-remote-machine-in-a-golang-cli
