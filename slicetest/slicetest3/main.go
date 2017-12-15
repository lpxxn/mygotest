package main

import (
	"fmt"
)

type RangeInfo struct {
	Begin int64
	End   int64
}


var f = fmt.Println

func main() {
	//data := make([]RangeInfo, 0)
	//data = append(data, RangeInfo{10, 15}, )

	data := []RangeInfo{
		RangeInfo{ 10, 15},
		RangeInfo{17, 19},
		RangeInfo{ 22, 25},
	}

	//for i, v := range data {
	//	fmt.Println("i : ", i, " v: ", v)
	//}


	f("-----", GetValue(10, data))
	f("-----", GetValue(14, data))
	f("-----", GetValue(15, data))
	f("-----", GetValue(18, data))
	f("-----", GetValue(25, data))



}

func GetValue(value int64, data []RangeInfo) int64 {
	value2 := value + 1
	var revValue int64 = -1
	maxIndex := len(data) - 1
	for i, v := range data {
		fmt.Println("index :", i)
		if value >= v.Begin && value <= v.End {
			if value2 <= v.End {
				revValue = value2
				break;
			} else if value2 > v.End {
				if maxIndex == i {
					return -1
				}
				revValue = data[i + 1].Begin
			}
		}
	}

	return revValue

}












