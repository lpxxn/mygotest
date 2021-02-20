package main

import (
	"fmt"
	"sync"
)

const size int = 1000 * 1000

func main() {
	//slice0 := make([]int, size)
	//fmt.Println("slice0 len: ", len(slice0), " cap :", cap(slice0))
	//doSomeThing(slice0)

	v1 := 1
	v2 := 2
	v3 := 3
	t1 := []*int{&v1, &v2, &v3}
	var wg sync.WaitGroup
	wg.Add(len(t1))
	for i, v := range t1 {
		go func(index int, ov *int) {
			fmt.Println("--------", *v, "-----", *ov)

			fmt.Println(i)
			fmt.Printf("%#v  %d origin %#v \n", *v, i, *t1[i])
			fmt.Printf("param index %d v %#v\n", index, *ov)

			wg.Done()
		}(i, v)
	}
	wg.Wait()
	af := t1[:]
	fmt.Println(af)

	TestAddNil()
}

func doSomeThing(s []int) {
	fmt.Println(len(s))
}

func TestAddNil() {
	arr := []*Element{}
	arr = append(arr, &Element{Name: "abc"})
	fmt.Println(arr, "  ", len(arr))

	arr = append(arr, nil)
	fmt.Println(arr, "  ", len(arr))
}

type Element struct {
	Name string
}

/*
func main() {

	slice0 := make([]int, size)
	fmt.Println("slice0 len: ", len(slice0), " cap :")

	// 创建一个容量和长度均为6的slice
	slice1 := []int{5, 23, 10, 2, 61, 33}

	for i, len := 1, len(slice1); i < len; i++ {
		fmt.Println("index: ", i, " value:", slice1[i])
	}


    var index, value int
	for index, value = range slice1 {
		fmt.Println("index: ", index, &index, " value address : ", &value, " slice1 value address", &slice1[index])

	}

	// 可以忽略我们不关心的元素
	// 只关心value
	for _, value := range slice1 {
		fmt.Println("value ", value)
	}

	// 只关心index, 可以不用 _
	for index := range slice1 {
		fmt.Println("index: ", index)
	}

	// 对slices1进行切片，长度为2容量为3
	slice2 := slice1[1:3:3]
	fmt.Println("cap", cap(slice2))
	fmt.Println("slice2", slice2)

	//修改一个共同指向的元素
	//两个slice的值都会修改
	slice2[0] = 11111
	fmt.Println("slice1", slice1)
	fmt.Println("slice2", slice2)

	// 增加一个元素
	slice2 = append(slice2, 55555)

	fmt.Println("slice1: ", slice1)
	fmt.Println("slice2: ", slice2)

	fmt.Println("\r\n-----")
	slice3 := []int{1, 2, 3}
	fmt.Println("slice2 cap", cap(slice3))

	slice3 = append(slice3, 5)
	fmt.Println("slice2 cap", cap(slice3))

}


*/
