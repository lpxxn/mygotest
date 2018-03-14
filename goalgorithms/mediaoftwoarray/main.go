package main

import "fmt"

func main() {
	n1 := []int {1, 2}
	n2 := []int {3, 4}

	fm := findMedianSortedArrays
	rev :=fm(n1, n2)
	fmt.Println(rev)

	fmt.Println(fm([]int {3, 4, 5}, []int {2, 7}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1 := nums1
	n2 := nums2

	pa := 0  // 当前数组1的位值
	pb := 0  // 当前数组2的位值

	k := 0	// k 位置 //找到的第k小的数

	n1len := len(n1)
	n2len := len(n2)

	first := 0
	second := 0

	// odd  or even
	isEven := (n1len + n2len) % 2 == 0

	max := (n1len + n2len) / 2

	for {
		if k > max {
			break
		}

		min := 0

		// 长度不一样的
		if pa < n1len && pb < n2len {
			item1 := n1[pa]
			item2 := n2[pb]

			if item1 < item2 {
				min = item1
				pa++
			} else {
				min = item2
				pb++
			}

		// 其中有一个长度走完了后
		} else if pa < n1len {
			//
			min = n1[pa]
			pa++
		} else {
			min = n2[pb]
			pb++
		}

		// 奇偶情况
		if isEven {
			fmt.Println("Even")
			if (k + 1)  == max {
				first = min
				fmt.Println("in for first: ", first)
			}

			if k == max {
				second = min
				fmt.Println("in for second: ", second)
			}
		} else {
			fmt.Println("Odd")
			if k == max {
				first = min
				second = 0
			}
		}
		k++
		fmt.Println("k: ", k)
	}

	if isEven {
		return (float64(first) + float64((second))) / 2
	} else {
		return float64(first)
	}


}



/*

nums1 = [1, 3]
nums2 = [2]

The median is 2.0



nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5


两个有序数组的中位数，可以转换为求第k小的数
奇数是找到第k小的就可以了
偶数是找到k 和 k + 1 在取平均值



 */