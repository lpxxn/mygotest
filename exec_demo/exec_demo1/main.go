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
/*

tell application "iTerm"
	activate

	set W to current window
	if W = missing value then set W to create window with default profile
	tell W's current session
		split vertically with default profile
		split vertically with default profile
	end tell
	set T to W's current tab
	write T's session 1 text "cd ~/Desktop"
	write T's session 2 text "cd ~/Downloads"
	write T's session 3 text "cd ~/Documents"
end tell

 */
