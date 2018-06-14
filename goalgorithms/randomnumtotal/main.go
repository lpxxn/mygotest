package main

import (
	"math/rand"
	"time"
	"fmt"
	"math"
	"strconv"
)

func main() {
	for i := 0; i < 10; i ++ {

		rand.Seed(time.Now().UnixNano())
		//
		rev := TestRTotalM(3, 0.01, 5)
		var total float32
		for _, v := range rev {
			fmt.Println(v)
			total += v
		}

		fmt.Println("Total Money", total)
		fmt.Println("Total Money", total == 3)
		fmt.Println("Total Money", strconv.FormatFloat(float64(total), 'f', -1, 64))
	}
}


func TestRTotalM(totalMoney, minMoney float32, nums int) []float32 {
	rev := []float32{}
	i := 0
	for ;i < nums; i++ {
		remainsNum := float32(nums - (i - 1))
		averageMoney := (totalMoney -  (remainsNum * minMoney)) / remainsNum
		//fmt.Println("averageMoney :----", averageMoney)
		topMoney := averageMoney * 2;

		rNum := float32(rand.Intn(100) + 1)/100
		//rNum := rand.Float32()

		//fmt.Println("random rNum: ", rNum)

		individualMoney := rNum * topMoney + minMoney
		individualMoney = float32(int(individualMoney * 1000)) / 1000
		totalMoney = totalMoney - individualMoney
		rev = append(rev, individualMoney)
		//fmt.Printf("第 %f 个红包: %f 元， 剩下 %f 元 \n", i, individualMoney, totalMoney)

	}

	//rev = append(rev, totalMoney)
	rev = append(rev, totalMoney)
	//fmt.Printf("第 %f 个红包:  %f 元 \n", i, totalMoney)
	return rev
}


func TestRTotalM2(totalMoney, minMoney float32, nums int) []float32 {
	rev := []float32{}
	i := 0
	for ;i < nums; i++ {
		remainsNum := float32(nums - (i - 1))
		max := totalMoney/(remainsNum) * 2
		fmt.Println("max :----", max)

		individualMoney := rand.Float32() * max
		if individualMoney <= minMoney {
			individualMoney = minMoney
		} else {
			individualMoney = float32(math.Floor(float64(individualMoney * 100)) / 100)
		}

		totalMoney -= individualMoney
		rev = append(rev, individualMoney)
		fmt.Printf("第 %f 个红包: %f 元， 剩下 %f 元 \n", i, individualMoney, totalMoney)

	}

	rev = append(rev, float32(math.Floor(float64(totalMoney * 100)) / 100))
	fmt.Printf("第 %f 个红包: 剩下 %f 元 \n", i, totalMoney)
	return rev
}
