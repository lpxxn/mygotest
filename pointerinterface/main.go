package main

import (
	"fmt"
)

type ITestBase interface {
	GetName() string
	GetAge() int
}

type MyTestC struct {
	name string
	age int
}
//  method declared with a pointer receiver of type MyTestC
//  you can call the method directly from  pointers only
func (t *MyTestC) GetName() string {
	return t.name
}


// method declared with a value receiver of type MyTestC
// you can call the method directly from values and pointers
func (t MyTestC) GetAge() int {
	return t.age
}

// the package to test that interface call the method
func main() {
	// interface values are represented as two-word pair giving a pointer to information about the type stored
	// in the interface and a pointer to the asociated data
	// https://stackoverflow.com/questions/13511203/why-cant-i-assign-a-struct-to-an-interface
	// that is why Interface and not *Interface is the correct type to hold a poin ter to a struct implementing Interface

	var it1 ITestBase = &MyTestC{name:"peng", age: 1}

	fmt.Println(it1.GetName())
	fmt.Println(it1.GetAge())

	// you can do this
	var it *ITestBase = new(ITestBase)
	var tc = &MyTestC{"li", 2}
	*it = tc;
	//it.GetName() //error
	name := (*it).GetName()
	fmt.Println(name)
	fmt.Println((*it).GetAge())


	// error GetName method has pointer receiver
//	var it3 ITestBase = MyTestC{name: "lili", age:3}
	// this work
	var it3 MyTestC = MyTestC{"lili", 3}
	fmt.Println(it3.GetName())
	fmt.Println(it3.GetAge())
}