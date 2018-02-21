package main

import (
	"fmt"
)

type TestS struct {
	status_code int
	headers     map[string]string
	body        string
	msg         *string
	intValue    *int
}

func (t TestS) validResponse() bool {
	if t.status_code < 300 {
		return true
	}
	return false
}

func (t TestS) Status() int {
	return t.status_code
}

// 不会被修改
func (t TestS) updateStatusV(new_status int) {
	t.status_code = new_status
}

func (t *TestS) updateStatusP(new_status int) {
	t.status_code = new_status
}

// 会被修改, 因为map底层是指针
func (t TestS) updateMapV(header, value string) {
	t.headers[header] = value
}

func (t *TestS) updateMapP(header, value string) {
	t.headers[header] = value
}

// 因为msg 是指针所有都会变
func (t TestS) updateMsgV(msg string) {
	*t.msg = msg
}

// 因为msg 是指针所有都会变
func (t *TestS) updateMsgP(msg string) {
	*t.msg = msg
}

// 因为 intValue是指针，所有都会变
func (t TestS) updateIntV(value int) {
	*t.intValue = value
}

// 因为 intValue是指针，所有都会变
func (t *TestS) updateIntP(value int) {
	*t.intValue = value
}

func main() {
	value1 := TestS{headers: make(map[string]string), intValue: new(int), msg: new(string)}
	value1.updateStatusP(123)
	fmt.Println(value1.status_code)

	value1.updateMapP("li", "peng")
	fmt.Println(value1.headers)

	value1.updateStatusV(3333)
	value1.updateMapV("li", "san")
	fmt.Println(value1.status_code)
	fmt.Println(value1.headers)

	value1.updateMsgP("li-msg-point")
	fmt.Println("msg P : ", *value1.msg)
	value1.updateMsgV("li-msg")
	fmt.Println("msg V : ", *value1.msg)

	value1.updateIntP(111111111111)
	fmt.Println("intValue P: ", *value1.intValue)
	value1.updateIntV(222222222222)
	fmt.Println("intValue V: ", *value1.intValue)
	fmt.Println("------------------")

	value2 := &TestS{headers: make(map[string]string), msg:new(string)}
	value2.updateStatusP(1111)
	fmt.Println(value2.status_code)

	value2.updateStatusV(2222)
	fmt.Println(value2.status_code)

	point1 := &TestS{headers: make(map[string]string)}
	point1.updateStatusP(33333)
	fmt.Println(point1.status_code)
	// 指针变量调用也没有发生变化
	point1.updateStatusV(44444)
	fmt.Println(point1.status_code)


	value2.updateMsgP("li-msg-point")
	fmt.Println("msg2 P : ",*value2.msg)
	value2.updateMsgV("li-msg")
	fmt.Println("msg2 V : ", *value2.msg)


	var a *int = new(int)
	*a = 123123123
	fmt.Println(*a)

	valuP := TestS{}
	// *valuP.intValue = 123 // error
	valuP.intValue = new(int)
	*valuP.intValue = 11
	fmt.Println(*valuP.intValue)
	// valuP.headers["aa"] = "aa"// error
	valuP.headers = make(map[string]string)
	valuP.headers["aa"] = "aa"

}
