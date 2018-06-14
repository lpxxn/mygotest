package main

import (
	"math/rand"
	"time"
	"fmt"
	"math"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	//
	rev := TestRTotalM(3, 0.01, 5)
	var total float32
	for _, v := range rev {
		fmt.Println(v)
		total += v
	}
	//fmt.Println("Total Money", total)
	fmt.Printf("Total Money %f", total)

}


func TestRTotalM(totalMoney, minMoney float32, nums int) []float32 {
	rev := []float32{}
	i := 0
	for ;i < nums; i++ {
		remainsNum := float32(nums - (i - 1))
		averageMoney := (totalMoney -  (remainsNum * minMoney)) / remainsNum
		fmt.Println("averageMoney :----", averageMoney)
		topMoney := averageMoney * 2;

		rNum := float32(rand.Intn(100) + 1)/100
		//rNum := rand.Float32()

		fmt.Println("random rNum: ", rNum)

		individualMoney := rNum * topMoney + minMoney
		individualMoney = float32(math.Floor(float64(individualMoney * 1000)) / 1000)
		totalMoney -= individualMoney
		rev = append(rev, individualMoney)
		fmt.Printf("第 %f 个红包: %f 元， 剩下 %f 元 \n", i, individualMoney, totalMoney)

	}

	rev = append(rev, float32(math.Round(float64(totalMoney * 1000)) / 1000))
	fmt.Printf("第 %f 个红包:  %f 元 \n", i, totalMoney)
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
