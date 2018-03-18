package main

import "fmt"

type TList struct {
	value int
	Next *TList
}

func  hasLoop(list *TList) (*TList, bool, int) {

	if list == nil {
		return nil, false, -1
	}
	fast := list

	slow := list


	length := 1
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return slow, true, length
		}

		length++
	}
	return nil, false, -1
}

func main() {
	list1 := &TList{value:1}
	list2 := &TList{value:2}
	list3 := &TList{value:3}
	list4 := &TList{value:4}

	list1.Next = list2
	list2.Next = list1
	list3.Next = list4

	flist, ok, length := hasLoop(list1)
	fmt.Println(flist, ok, length)

}
