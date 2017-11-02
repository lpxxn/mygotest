package main

// TODO import不要用相对路径
// TODO https://golang.org/cmd/go/#hdr-Relative_import_paths
// TODO 相对路径会打破dep的导入规则
// TODO 这个例子只是自己测试相对路径，真实项目中不要这样做
import (
	"fmt"
	// alias name
	lib1 "./mylib1"

	// no alias name
	t "./thrlib"

	// dot alias
	. "./testlib"

	// embed directory
	"./testlib/testlib2"
)

var dt1 lib1.TestData

//var dt2 t.ThadData
func main() {
	// alias name
	fmt.Println("test start")
	lib1.ReadFun()

	dt1 = lib1.TestData{"na"}
	fmt.Println(dt1)
	// no alias name
	de := t.DevicesDb{Name: "test"}
	fmt.Println(de)
	t.DeviceOp()
	td := t.ThadData{Host: "aaaaaHost"}
	fmt.Println(td)
	t.ConnectSql()

	//dt2 = t.ThadData{Host:"test"}
	//print(dt2)
	//thrlib.DeviceOp()

	// dot alias
	PrintLnThing("Hello thr")

	// emabe
	testlib2.ExecMethod()
}

// https://stackoverflow.com/questions/6478962/what-does-the-dot-or-period-in-a-go-import-statement-do
// https://golang.org/ref/spec#Import_declarations
//Import declaration          Local name of Sin
//
//import   "lib/math"         math.Sin
//import M "lib/math"         M.Sin
//import . "lib/math"         Sin
