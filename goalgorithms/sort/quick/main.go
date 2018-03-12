package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func partition(a []int, l, r int) int {
	fmt.Println("old a: ", a)
	p := a[r] //
	storeIndex := l
	for i := l; i < r; i++ {
		if a[i] < p {
			a[storeIndex], a[i] = a[i], a[storeIndex]
			storeIndex++
		}
	}
	fmt.Println("a1 :", a)
	a[r], a[storeIndex] = a[storeIndex], a[r]
	fmt.Println("a2 :", a)
	return storeIndex
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func QuickSort(a []int, left, right int) {
	if right > left {

		pivotIndex := partition(a, left, right)
		fmt.Println("pivot : ", pivotIndex)
		QuickSort(a, left, pivotIndex-1)
		QuickSort(a, pivotIndex+1, right)

	}
}

func main() {
	arr := []int{6, 7, 8, 5, 2, 1, 9, 5, 4}
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
