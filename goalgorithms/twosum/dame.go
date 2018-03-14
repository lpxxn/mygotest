package main

import "fmt"

func main() {
	// two Sum
	nums := []int {1, 2, 5, 18, 21}
	revArr := twoSum(nums, 20)
	fmt.Println(revArr)

}

func twoSum(sum []int, target int) []int {
	mp := make(map[int]int)

	for i, num := range sum {
		if idx, ok := mp[target - num]; ok {
			return []int {idx, i}
		}
		mp[num] = i;
	}

	return nil
}