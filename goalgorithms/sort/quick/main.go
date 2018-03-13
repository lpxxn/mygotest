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
	arr2 := append([]int{}, arr...)
	arr3 := append([]int{}, arr...)
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
	insertsort(arr2)
	fmt.Println(arr2)

	bubbleSort(arr3)
	fmt.Println(arr3)
}

func insertsort(arr []int) {
	arrLen := len(arr)
	for i := 1; i < arrLen; i++ {
		for j := i - 1; j >= 0; j-- {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
}
/*

	1. 从第一个元素开始，该元素可以认为已经被排序
	2. 取出下一个元素，在已经排序的元素序列中从后向前扫描
	3. 如果该元素（已排序）大于新元素，将该元素移到下一位置
	4. 重复步骤3，直到找到已排序的元素小于或者等于新元素的位置
	5. 将新元素插入到该位置后
	6. 重复步骤2~5

 */

func bubbleSort(arr []int) {
	arrLen := len(arr)
	for i := 0; i < arrLen; i++ {

		for j := i + 1; j < arrLen; j++ {
			if arr[j] < arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
}

/**
原理是临近的数字两两进行比较,按照从小到大或者从大到小的顺序进行交换,
这样一趟过去后,最大或最小的数字被交换到了最后一位,
然后再从头开始进行两两比较交换,直到倒数第二位时结束,其余类似看例子

 */