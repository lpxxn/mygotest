package main

import "fmt"

type T int

func (t *T) Get() T{
	return *t + 1
}

// https://golang.org/ref/spec#Type_declarations
// The type denoted by T is called the receiver base type; it must not be a pointer or interface type
type P T
// type P *T // error when as user receiver
/*
The receiver type must be of the form T or *T where T is a type name.
The type denoted by T is called the receiver base type;
it must not be a pointer or interface type and it must be declared in the same package as the method.
 */
func (p *P) Get() T{
	return T(*p + 2)
}



func main() {
	var i1 = 123
	var v1 T = T(i1)
	var v2 = &v1
	var v3 P= P(v1)
	var v4 *P = &v3
	fmt.Println(v1.Get(), v2.Get(), v3.Get(), v4.Get())


	var u U = U{Name:"U"}
	var u2 U2 = U2(u)
	var u3 U3 = U3(u2)
	var u4 *U3 = &u3
	fmt.Println(u.Get(), u2.Get(), u3.Get(), u4.Get())

	fmt.Println("U3 call u2: ", u3.GetU2_1())
	fmt.Println("U3 call u2 2: ", u3.GetU2_2())
}


type U struct {Name string}

func (U) Get() string {
	return "U"
}

type U2 U

func (U2) Get() string {
	return "U2"
}

type U3 U2

func (U3) Get() string {
	return "U3"
}

func (u3 *U3) GetU2_1() string {
 	return (*U2)(u3).Get()
 	// error  type assertion works for interface only
 	//https://stackoverflow.com/questions/18691927/go-named-type-assertions-and-conversions
 	//return u3.(*U2).Get()
}


func (u3 U3) GetU2_2() string {
	return (U2)(u3).Get()
}



