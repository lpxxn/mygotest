package main

import (
	"fmt"

	"strconv"
)

type TokeType uint16

const (
	KeyWord TokeType = iota
	IDENTIFIER

)

type IFather interface {

	GetName() string
}

type Father struct {
	KeyType TokeType
	Name string
}

func (f Father) GetName() string {
	return f.Name + "father"
}

type Father2 struct {
	Father
	Age int
}

type Father3 struct {
	Father
	Name string
}

func (f Father3) GetName() string {
	return f.Name + "f3"
}

type Father4 struct {
	IFather
	Sex bool
}

func (f Father4) GetName() string {
	return f.IFather.GetName() + strconv.FormatBool(f.Sex)
}

func main() {

	f2 := Father2{Father: Father{Name: "li", KeyType:IDENTIFIER}, Age: 28}
	fmt.Println(f2.GetName(), f2.Name)

	f3 := Father3{Father: Father{Name:"li", KeyType:KeyWord}, Name: "peng"}
	fmt.Println(f3.GetName(), f3.Name, f3.Father.Name)


	f4 := Father4{IFather: Father2{Father: Father{Name:"peng", KeyType:KeyWord}, Age: 29}, Sex: true}

	fmt.Println(f4.GetName())

}
