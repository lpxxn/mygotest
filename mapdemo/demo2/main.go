package main

import (
	"time"
	"fmt"
)

func main() {

	m1 := map[string]string{"name": "li"}
	t1 := fmt.Sprint(time.Now().Unix())
	fmt.Println(t1)
	t2 := t1[:6]
	fmt.Println(t2)
	tmap1(m1)
	fmt.Printf("%+v", m1)
	fmt.Printf("%#v", m1)

}

func tmap1(m map[string]string) {
	m["add"] = "abcde"
}