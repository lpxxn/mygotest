package main

import (
	"reflect"
	"fmt"
)

type Cat struct {
	Name string
	Age int
}

func (c Cat) GetName() string {
	return c.Name
}

func (c Cat)GetAge() int {
	return c.Age
}

func main() {

	var cp *Cat = new(Cat)
	cp.Name = "little cat"

	t := reflect.TypeOf(cp)

	for i := 0; i < t.NumMethod(); i++ {
		fmt.Println(t.Method(i).Name)
	}


	var ib IB = Cat{Name:"tom", Age:23}
	fmt.Println(ib.GetAge())
	fmt.Println(ib.GetName())

}

/*
The method set of any other type T consists of all methods declared with receiver type T. The method set of the corresponding pointer type *T is the set of all methods declared with receiver *T or T (that is, it also contains the method set of T).

So if you have a method with Cat as the receiver type, that method is also part of the method set of *Cat. So *Cat will already have that method, attempting to declare "another" one with the same name and *Cat as the receiver type will be a duplicate.


 */

 type IA interface {
 	GetName() string
 }

 type IB interface {
 	IA
 	GetAge() int
 }
