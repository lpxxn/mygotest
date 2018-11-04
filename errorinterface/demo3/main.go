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

	err4 := TestNormalError()

	s, err4 := t1.Err2()
	if err4 != nil {
		fmt.Println("err4 is not nil", err4)
	}
	fmt.Println(s)
}
