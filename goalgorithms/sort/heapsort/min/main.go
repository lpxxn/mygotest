package main

import (
	"fmt"
)
/*
https://www.cnblogs.com/MOBIN/p/5374217.html
https://billjh.github.io/blog/2017/heap-sort/
https://www.cnblogs.com/chengxiao/p/6129630.html
 */
func HeapSort(a []int) {
	alen := len(a) - 1

	for i := alen/2 - 1; i >= 0; i-- {
		adjustHeap(a, i, alen)
	}

	for ;alen >= 0; alen-- {
		a[0], a[alen] = a[alen], a[0]
		adjustHeap(a, 0, alen)
	}
}

func adjustHeap(a []int, i, length int) {
	// 判断当前的父节点有没有左节点
	left := leftChild(i)
	for left < length {

		right := left + 1 // 右节点
		j := left	// j 指向左节点
		// 选出较小的节点
		if right < length && a[left] > a[right] {
			j++
		}

		// 如果父节点比最大的孩子节点大，则交换
		if a[i] > a[j] {
			a[i], a[j] = a[j], a[i]

		} else {
			// 比孩子的节点都大则跳出
			break
		}
		i = j
		left = leftChild(i)
	}
}

// 左节点
func leftChild(i int) int {
	return 2 * i + 1
}

func rightChild(i int) int  {
	return 2 * i + 2
}

func main() {
	//a := []int {20,50,20,40,70,10,80,30,60}
	a := []int {1,5,6,3,2,4}
	//
	HeapSort(a)
	fmt.Println(a)
}
