package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

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
	userList = userList[:len(userList)-1]
	fmt.Println(userList, " cap", cap(userList), " len", len(userList))
	userList = append(userList, "cc")
	fmt.Println(userList)


	max, min := 10, 1
	// remove slice element
	for lenU := len(userList) -1 ; lenU >= 0; lenU-- {
		c := rand.Intn(max - min + 1) + min
		if c > 4 {
			userList = append(userList[:lenU], userList[lenU + 1:]...)
		}
	}
	fmt.Println(userList)

	for i := 0; i < 10; i++ {
		fmt.Println(testRandom())
	}
}
func testRandom() int {
	max, min := 10, 1
	return rand.Intn(max - min + 1) + min
}
