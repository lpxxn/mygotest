package main

import (
	"fmt"

	"github.com/mygotest/exec_demo"
)

func main() {
	err := exec_demo.RunCmd("go build /Users/lipeng/go/src/github.com/mygotest/exec_demo/bad_test")
	fmt.Printf("err: %#v\n", err)

	err = exec_demo.RunCmd(`
osascript -e '
tell application "iTerm"
  activate
  tell current window to set tb to create tab with default profile
  tell current session of current window to write text "cafe_sandbox_test1"  
end tell
'
`)
	fmt.Printf("err: %#v\n", err)

	err = exec_demo.RunCmd(`
osascript -e '
tell application "iTerm2"
    set newWindow to (create window with default profile)
    tell current session of newWindow
        write text "cafe_sandbox_test1"
    end tell
end tell
'
`)
	fmt.Printf("err: %#v\n", err)
}
