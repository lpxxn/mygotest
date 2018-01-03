package main

import "fmt"

type MyObj interface {
	Eat(v string)
}


type Animal struct {
	Name string
	EatFood string
}

func (a *Animal) Eat(v string) {
	a.EatFood = v
}


type Cat struct {
	Animal	"Animal"
}

type Dog struct {
	*Animal
}

func ObjEat(obj MyObj, v string) {
	obj.Eat(v)
}

func main() {
	cat := &Cat{Animal{Name: "cat"}}
	ObjEat(cat, "fish")
	fmt.Println(cat)

	dog := &Dog{&Animal{Name: "dog"}}
	ObjEat(dog, "bone")
	fmt.Println(dog.Animal)


	cat2 := Cat{Animal{Name:"cat2"}}
	cat2.Eat("aaa")
	// error
	//ObjEat(cat2, "cat2eee")
	ObjEat(&cat2, "cat2aaa")
	fmt.Println(cat2.Animal)

	dog2 := Dog{&Animal{Name: "dog2"}}
	dog2.Eat("bbbb")
	ObjEat(dog2, "cccc")
	fmt.Println(dog2.Animal)
}
