package main

import (
	"fmt"
	//"sync/atomic"
	"log"
	"strconv"
)

type People interface {
	Speak(string) string
}

type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

type User struct {
}
type MyUser1 User
type MyUser2 = User

func (i MyUser1) m1() {
	fmt.Println("MyUser1.m1")
}
func (i User) m2() {
	fmt.Println("User.m2")
}

type student struct {
	Name string
	Age  int
}

func (s *student) String() string {
	return fmt.Sprintln(s.Age, s.Name)
}

// 错
func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	// 错
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	fmt.Printf("%#v\n", m)
}

// 这样也可以
func pase_student2() {
	m := make(map[string]*student)
	stus := []*student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = stu
	}
	fmt.Printf("%#v\n", m)
}

const (
	x = iota
	y
	z = "zz"
	k
	p = iota
	g
)

func main() {
	pase_student()
	pase_student2()
	fmt.Println(x, y, z, k, p, g)

	a := []int{1, 2, 3, 3}
	b := a[:2:2]
	c := a[2:]
	b = append(b, 55)
	fmt.Println(b, "   ", c, "       ", a)

	v, mb, t, err := strconv.UnquoteChar(`\"Fran & Freddie's Diner\"`, '"')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("value:", string(v))
	fmt.Println("multibyte:", mb)
	fmt.Println("tail:", t)
	strA := "c"
	switch strA {
	case "c":
		println("c")
		fallthrough
	case "b", "a":
		println("a")
	}
}

func GetValue(m map[int]string, id int) (string, bool) {
	if _, exist := m[id]; exist {
		return "存在数据", true
	}
	return "", false
}

func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc3(i int) (t int) {
	//defer func() {
	//	t += i
	//}()
	t = i + 1
	return 26
}

func getType(in interface{}) {
	switch in.(type) {
	case People:
		fmt.Println("people")
	default:
		fmt.Printf("%T", in)
	}
}
