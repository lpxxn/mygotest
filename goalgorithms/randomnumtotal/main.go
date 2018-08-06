package main

import (
	"math/rand"
	"time"
	"fmt"
	"strconv"
	"github.com/shopspring/decimal"
	"strings"
)

func main() {
	for i := 0; i < 1000; i ++ {

		rand.Seed(time.Now().UnixNano())
		//
		//rev := TestRTotalM(3, 0.01, 5)
		//var total float32
		//for _, v := range rev {
		//	fmt.Println(v)
		//	total += v
		//}
		//
		//fmt.Println("Total Money", total)
		//fmt.Println("Total Money", total == 3)
		//fmt.Println("Total Money", strconv.FormatFloat(float64(total), 'f', -1, 64))
		//
		//
		fmt.Println("----------")

		//rev2 := TestRTotalM2(decimal.New(30, 0), decimal.NewFromFloatWithExponent(0.00001, -5), 3)
		//rev2 := TestRTotalM2(decimal.New(30, 0), decimal.NewFromFloatWithExponent(0.0001, -4), 4)

		rev2 := TestRTotalM2(decimal.New(100, 0), decimal.NewFromFloatWithExponent(0.01, -2), 10)
		//rev2 := TestRandomM3(decimal.New(300, 0), decimal.NewFromFloatWithExponent(0.1, -1), 5)
		total2 := decimal.New(0, 0)
		for _, v := range rev2 {
			fmt.Print(v, ",")
			f, _ := v.Float64()
			if f == 0 {
				panic(f)
			}
			total2 = total2.Add(v)
		}
		fmt.Println()
		f, _ := total2.Float64()
		fmt.Println("Total Money", f)
		fmt.Println("Total Money", strconv.FormatFloat(f, 'f', -1, 64))
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


func TestRTotalM2(totalMoney, minMoney decimal.Decimal, nums int64) []decimal.Decimal {
	rev := []decimal.Decimal{}
	var i int64 = 0
	minFloat, _ := minMoney.Float64()
	var precison int32 = 0
	precisionStr := strings.Split(strconv.FormatFloat(minFloat, 'f', -1, 64), ".")
	if len(precisionStr) > 1 {
		precison = int32(len(precisionStr[1]))
	}
	for ;i < nums - 1; i++ {

		remainsNum := decimal.New(nums - i , 0)
		//averageMoney := (totalMoney -  (remainsNum * minMoney)) / remainsNum
		averageMoney := totalMoney.Sub(remainsNum.Mul(minMoney)).Div(remainsNum)
		//fmt.Println("averageMoney :----", averageMoney)
		topMoney := averageMoney.Mul(decimal.NewFromFloat(2))

		rNum := decimal.NewFromFloatWithExponent(float64(rand.Intn(101))/100, -2)

		//individualMoney := rNum.Mul(topMoney).Add(minMoney).Truncate(precison)
		individualMoney := rNum.Mul(topMoney).Add(minMoney).Truncate(precison)

		//individualMoney = float32(int(individualMoney * 1000)) / 1000
		totalMoney = totalMoney.Sub(individualMoney)
		rev = append(rev, individualMoney)
		//fmt.Printf("第 %f 个红包: %s 元， 剩下 %s 元 \n", i, individualMoney.String(), totalMoney.String())
	}

	//rev = append(rev, totalMoney)
	rev = append(rev, totalMoney)
	//fmt.Printf("第 %f 个红包:  %f 元 \n", i, totalMoney)
	return rev
}

func TestRandomM3(totalMoney, minMoney decimal.Decimal, nums int64) []decimal.Decimal {
	minFloat, _ := minMoney.Float64()
	var precison int64 = 1
	precisionStr := strings.Split(strconv.FormatFloat(minFloat, 'f', -1, 64), ".")
	if len(precisionStr) > 1 {
		precison = int64(len(precisionStr[1]))
	}
	// 一共多少份
	allPiece := totalMoney.Div(decimal.NewFromFloat(float64(precison) * 0.1)).IntPart() - nums
	fmt.Println(allPiece)

	rev := make([]decimal.Decimal, nums)
	for ele := range rev {
		rev[ele] = minMoney
	}
	var i int64 = 0
	for ; i < allPiece; i++ {
		r := rand.Intn(int(nums))
		rev[r] = rev[r].Add(minMoney)
	}


	return rev
}
