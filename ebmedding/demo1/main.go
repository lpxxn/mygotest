package main

import (
	"fmt"
)

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

// Dog 的Animal是 值类型 所以 值类型的 Dog 不会修改 Animal的 EatFood的值
func (cat Cat) CatEat(v string) {
	cat.EatFood += v
}


// Dog 的Animal是 指针 所以 值类型的 Dog 也会修改 Animal的 EatFood的值
func (dog Dog) DogEat(v string) {
	dog.EatFood += v
}
func main() {
	cat := &Cat{Animal{Name: "cat"}}
	ObjEat(cat, "fish")
	fmt.Println(cat)

	dog := &Dog{&Animal{Name: "dog"}}
	ObjEat(dog, "bone")
	fmt.Println(dog.Animal)

	fmt.Println("----- Begin Value:")
	cat2 := Cat{Animal{Name:"cat2"}}
	cat2.Eat("aaa")
	// error
	//ObjEat(cat2, "cat2eee")
	ObjEat(&cat2, "cat2aaa")
	fmt.Println(cat2.Animal)

	cat2.CatEat(" cat eat")
	fmt.Println("cat eat : ",cat2.Animal)


	dog2 := Dog{&Animal{Name: "dog2"}}
	dog2.Eat("bbbb")
	ObjEat(dog2, "cccc")
	fmt.Println(dog2.Animal)
	dog2.DogEat(" dog eat")
	fmt.Println("dog eat: ", dog2.Animal)
}
