package mylib1

import (
	"../testlib"
	"../testlib/testlib2"
	"bytes"
	"fmt"
)

type TestData struct {
	Name string `json:"name"`
}

type Person struct {
	testlib2.User
	Desc string
}

type tinh struct {
	T1 string
	T2 string
}

func (ti *tinh) TString() string {
	var buffer bytes.Buffer
	buffer.WriteString(ti.T2)
	buffer.WriteString(" hi ")
	buffer.WriteString(ti.T1)
	return buffer.String()
	//	return ti.T2 + "hi " + ti.T1

}

type Athinh struct {
	tinh
	T3 string
}

func ReadFun() (string, error) {

	a := Athinh{T3: "abcde"}
	a.T1 = "111:"
	a.T2 = "bbbb"
	fmt.Println(a.TString())
	fmt.Println("return string and err")

	testlib.PrintLnThing("test in readutil.go ")

	p := Person{Desc: "hello"}
	p.Name = "li"

	fmt.Println(p)

	var u testlib2.User = testlib2.User{Name: "hi"}
	fmt.Println(u)
	return "hello ", nil
}
