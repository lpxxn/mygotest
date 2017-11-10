package main

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "World", "A name to say hello to.")

var spanish bool


func init(){
	flag.BoolVar(&spanish, "spanish", false, "Use Spanish language")
	flag.BoolVar(&spanish, "s", false, "Use Spanish language")

}
// go build main.go
// main.exe -s=true -name=abcdef  //bool类型必须用=号 name 可以用等也可以不用
// main.exe -s=false -name=abcdef

const (
	Header string = "Header"
)
const (
	H1 = 1 << iota
	H2
	H3
)

func main() {
	flag.Parse()

	if spanish == true {
		fmt.Printf("Hola %s \n", *name)
	} else {
		fmt.Printf("Hello %s! \n", *name)
	}
	fmt.Println(H1, H2, H3)
}


//Command line flag syntax:
//
//-flag
//-flag=x
//-flag x  // non-boolean flags only

// The last form is not permitted for boolean flags because the meaning of the command