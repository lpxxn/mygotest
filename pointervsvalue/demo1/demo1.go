package main

import "fmt"

type TestS struct {
	status_code int
	headers map[string]string
	body string
}

func (t TestS) validResponse() bool {
	if t.status_code < 300 {
		return true
	}
	return false;
}

func (t TestS) Status() int {
	return t.status_code
}
// 不会被修改
func (t TestS) updateStatusV(new_status int)  {
	t.status_code = new_status
}

func (t *TestS) updateStatusP(new_status int) {
	t.status_code = new_status
}

// 会被个性map底层是指针
func (t TestS) updateMapV(header, value string) {
	t.headers[header] = value
}

func (t *TestS) updateMapP(header, value string) {
	t.headers[header] = value
}


func main() {
	value1 := TestS{headers:make(map[string]string)}
	value1.updateStatusP(123)
	fmt.Println(value1.status_code)

	value1.updateMapP("li", "peng")
	fmt.Println(value1.headers)

	value1.updateStatusV(3333)
	value1.updateMapV("li", "san")
	fmt.Println(value1.status_code)
	fmt.Println(value1.headers)

	fmt.Println("------------------")

	value2 := &TestS{headers:make(map[string]string)}
	value2.updateStatusP(1111)
	fmt.Println(value2.status_code)

	value2.updateStatusV(2222)
	fmt.Println(value2.status_code)


	point1 := &TestS{headers:make(map[string]string)}
	point1.updateStatusP(33333)
	fmt.Println(point1.status_code)
	// 指针变量调用也没有发生变化
	point1.updateStatusV(44444)
	fmt.Println(point1.status_code)

}
