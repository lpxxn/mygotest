package main

import (
	"fmt"
)

func main() {
	userList := []string{"a", "b", "c", "d", "e", "f", "g"}
	const defaultCount = 4
	lenUserList := len(userList)
	fmt.Println("len:", lenUserList)
	for i := 0; i < lenUserList; i += defaultCount {
		max := i + defaultCount
		if max > lenUserList {
			max = lenUserList
		}
		items := userList[i:max]
		var param [defaultCount]string
		copy(param[:], items)
		fmt.Printf("values: %+v, len of param: %d\n", param, len(param))
		for i, v := range param {
			if v == "" {
				fmt.Print("is empty ")
			}
			fmt.Println("i: ", i, " v:", v)
		}
	}
}
