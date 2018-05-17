package main

import (
	"fmt"
)

type Person struct {
	name string
	age int
}

func main() {

	p := make(People)
	p["HM"] = Person{"Hank McNamara", 39}
//	p["HM"].age = p["HM"].age + 1 // error




	t := p["HM"]
	t.age = t.age + 1
	p["HM"] = t
	// or use point
	// p2 := make(map[string]*Person)

	fmt.Printf("age: %d\n", p["HM"].age)



	s1 := make([]Person, 0)
	s1 = append(s1, Person{name: "li"})
	// moidfy field of element
	s1[0].age = 111
	fmt.Println(s1)
	t.age = 111

	fmt.Println(s1)

}


type People map[string]Person
/*



 */