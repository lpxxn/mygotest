package main

import "fmt"

func main() {
	arr1 := []int {1, 3, 5, 7, 2, 6}
	index := partitionArray(arr1, 4)
	fmt.Println("index: ", index, "  arr: ", arr1)
}

func partitionArray(nums []int, k int) int {
	start := 0
	end := len(nums) - 1
	for end - start >= 0 {
		if nums[start] >= k {
			if nums[end] < k {
				nums[start], nums[end] = nums[end], nums[start]
			} else {
				end--
			}
		} else {
			start++
		}
	}

	return start
}


/*

问题描述：
给出一个整数数nums和一个整数k。划分数组（即移动数组nums中的元素），使得：

所有小于k的元素移到左边
所有大于等于k的元素移到右边
返回数组划分的位置，即数组中第一个位置i，满足nums[i]大于等于k。

样例：给出数组nums = [3, 2, 2, 1]和 k=2，返回 1。

挑战：要求在原地使用O(n)的时间复杂度来划分数组。


 */