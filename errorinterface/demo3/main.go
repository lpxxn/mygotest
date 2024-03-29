package main

import "fmt"

type MyError struct {
	Code string
}

func (e *MyError) Error() string {
	return e.Code
}

func TestErr1() *MyError {
	return nil
}

func TestErr2() *MyError {
	return TestErr1()
}


type TestF struct {

}

func (t *TestF) Err1() (string, *MyError) {
	return "", nil
}

func (t *TestF) Err2() (string, *MyError) {
	return t.Err1()
}

func TestNormalError() error {
	return nil
}


func main() {
	err1 := TestErr1()
	if err1 != nil {
		fmt.Println("err1 is not nil", err1)
	}
	err2 := TestErr2()
	if err2 != nil {
		fmt.Println("err2 is not nil", err2)
	}

	t1 := TestF{}
	_, err3 := t1.Err1()
	if err3 != nil {
		fmt.Println("err3 is not nil", err3)
	}

	// 这时的err4是error 接口类型
	err4 := TestNormalError()
	// 再给err4赋值，err4是error接口，指向MyError的空类型
	s, err4 := t1.Err2()
	if err4 != nil {
		fmt.Println("err4 is not nil", err4)
	}
	fmt.Println(s)
	//下面 类型错误，err3已经是 MyError
	//err3, str := TestEStr()
	//fmt.Println(err3, str)
}

func TestEStr() (string, string) {
	return "a", "b"
}
