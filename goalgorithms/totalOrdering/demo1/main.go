package main

import "fmt"

var totalNum int = 0

func perm(arr []int, k, m int) {
	//i := 0
	//fmt.Println("k: ", k, " m:", m)
	if k > m {
		//fmt.Println("arr begin")
		for i := 0; i <= m; i++ {
			fmt.Print(arr[i])
		}
		//fmt.Println("\narr end")
		fmt.Println()
		totalNum++;
	} else {
		for i := k; i <= m; i++ {
			//fmt.Println("order: ", arr)
			//fmt.Printf("i: %d k: %d \n", i , k)
			arr[i], arr[k] = arr[k], arr[i]
			//fmt.Println("change: 1",  arr)
			perm(arr, k + 1, m)
			arr[i], arr[k] = arr[k], arr[i]
			//fmt.Println("end change: 2",  arr, "--------")
		}
	}
}

func main() {
	arr := []int {1, 2, 3}
	perm(arr, 0, len(arr) -1 )
	fmt.Println("total: ", totalNum)
}

/*
核心思想是交换，具体来说，对于一个长度为n的串，要得到其所有排列，我们可以这样做：

1.把当前位上的元素依次与其后的所有元素进行交换

2.对下一位做相同处理，直到当前位是最后一位为止，输出序列

[需要注意的一点：我们的思想是“交换”，也就是直接对原数据进行修改，那么在交换之后一定还要再换回来，否则我们的原数据就发生变化了，肯定会出错]

如果觉得上面的解释还是很难懂的话，那么记住这句话：核心思想就是让你后面的所有人都和你交换一遍（而你是一个指针，从前向后按位移动...

*/
