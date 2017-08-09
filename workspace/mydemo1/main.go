package main

import (
	"fmt"
	// alias name
	lib1 "./mylib1"

	// no alias name
	"./thrlib"

	// dot alias
	. "./testlib"

	// embed directory
	"./testlib/testlib2"
)
var dt1 lib1.TestData
func main() {
	// alias name
	fmt.Println("test start")
	lib1.ReadFun()

	// no alias name
	var thrdDt thrlib.ThrdData = thrlib.ThrdData{Host:"test"}
	print(thrdDt)
	thrlib.DeviceOp()

	// dot alias
	PrintLnThing("Hello thr")

	// emabe
	testlib2.ExecMethod()
}

// https://stackoverflow.com/questions/6478962/what-does-the-dot-or-period-in-a-go-import-statement-do
//Import declaration          Local name of Sin
//
//import   "lib/math"         math.Sin
//import M "lib/math"         M.Sin
//import . "lib/math"         Sin