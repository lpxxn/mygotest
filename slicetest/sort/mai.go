package main

import (
	"sort"
	"fmt"
)

type AsortObj struct {
	V int
}

type AsortObjArray []AsortObj


func (a *AsortObjArray) SortDesc() {
	pData := *a
	sort.SliceStable(pData, func(i, j int) bool {
		return pData[i].V > pData[j].V
	})
}

func (a *AsortObjArray) LimitEle(limit int) {
	pData := *a
	dataLen := len(pData)
	if dataLen > limit {
		pData = pData[:limit]
	}
	*a = pData
}

func main() {
	arr := AsortObjArray{AsortObj{5}, AsortObj{1}, AsortObj{18}, AsortObj{10}, AsortObj{8}, AsortObj{89}, AsortObj{6}}
	arr.SortDesc()
	fmt.Println(arr)
	arr.LimitEle(5)
	fmt.Println(arr)
}
