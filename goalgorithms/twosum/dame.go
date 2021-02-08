package main

import (
	"fmt"
	"strconv"
)

func main() {
	// two Sum
	nums := []int{1, 2, 5, 18, 21}
	revArr := twoSum(nums, 20)
	fmt.Println(revArr)




	// 123
	n1 := NewNum(523)
	fmt.Println(n1.String())
	fmt.Println(n1.Value())

	n2 := NewNum(987)
	fmt.Println(n2.String())
	fmt.Println(n2.Value())
	dist := add(n1, n2)
	fmt.Println(dist.String())
	fmt.Println(dist.Value())
}

func twoSum(sum []int, target int) []int {
	mp := make(map[int]int)

	for i, num := range sum {
		if idx, ok := mp[target-num]; ok {
			return []int{idx, i}
		}
		mp[num] = i
	}

	return nil
}

type Num struct {
	Val  int
	Next *Num
}

func (n Num) String() string {
	rev := strconv.Itoa(n.Val)
	if n.Next != nil {
		return fmt.Sprintf("%d -> %s", n.Val, n.Next.String())
	}
	return rev
}

func (n Num) Value() int {
	if n.Next != nil {
		v, _ := strconv.Atoi(fmt.Sprintf("%d%d", n.Next.Value(), n.Val))
		return v
	}
	return n.Val
}

func add(n1, n2 *Num) *Num {
	rev := &Num{}
	next := rev
	carry := 0
	for n1 != nil || n2 != nil {
		v1, v2 := 0, 0
		if n1 != nil {
			v1 = n1.Val
			n1 = n1.Next
		}
		if n2 != nil {
			v2 = n2.Val
			n2 = n2.Next
		}
		sum := v1 + v2 + carry
		next.Next = &Num{Val: sum % 10}
		carry = sum / 10
		next = next.Next
	}
	if carry > 0 {
		next.Next = &Num{Val: carry}
	}
	return rev.Next
}

func NewNum(v int) *Num {
	rev := &Num{}
	next := rev
	for v > 0 {
		next.Next = &Num{Val: v % 10}
		next = next.Next
		v /= 10
	}
	return rev.Next
}
