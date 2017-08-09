package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("publish crm..")

	execCmd("./test.bat")
}

func execCmd(path string, args ...string) error{
	fmt.Printf("Running: %q %q\n", path, strings.Join(args, ","))
	cmd := exec.Command("test.bat", args...)
	bs, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error", err)
		return err
	}
	fmt.Printf("Output: %s", bs)
	return nil
}