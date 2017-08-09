package utilities

import (
	"fmt"
	"strings"
	"os/exec"
)

func ExecBatCmd(path string, args ...string) error{
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
