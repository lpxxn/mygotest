package twosum

import "testing"

/*
**Example:**
```
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
```
*/

func TestAdd(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	t.Log(myTwoSum(nums, 22))
}
func myTwoSum(arr []int, total int) []int {
	r := map[int]int{}
	for i, item := range arr {
		if idx, ok := r[total-item]; ok {
			return []int{idx, i}
		}
		r[item] = i
	}
	return nil
}
