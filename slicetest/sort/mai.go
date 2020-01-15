package main

import (
	"fmt"
	"sort"
	"time"
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

type SAbcDT struct {
	Name     string
	CreateAt time.Time
}

func main() {
	arr := AsortObjArray{AsortObj{5}, AsortObj{1}, AsortObj{18}, AsortObj{10}, AsortObj{8}, AsortObj{89}, AsortObj{6}}
	arr.SortDesc()
	fmt.Println(arr)
	arr.LimitEle(5)
	fmt.Println(arr)

	var sabcDTArray []SAbcDT
	sabcDTArray = append(sabcDTArray, SAbcDT{
		Name:     "cde",
		CreateAt: time.Now(),
	})

	sabcDTArray = append(sabcDTArray, SAbcDT{
		Name:     "abcde",
		CreateAt: time.Now().AddDate(-1, 0, 0),
	})

	sabcDTArray = append(sabcDTArray, SAbcDT{
		Name:     "d",
		CreateAt: time.Now().AddDate(0, 2, 1),
	})

	sabcDTArray = append(sabcDTArray, SAbcDT{
		Name:     "d",
		CreateAt: time.Now().AddDate(0, -5, 9),
	})

	fmt.Println(sabcDTArray)
	sort.Slice(sabcDTArray, func(i, j int) bool {
		return sabcDTArray[i].CreateAt.Before(sabcDTArray[j].CreateAt)
	})
	fmt.Println("asc: ", sabcDTArray)

	sort.Slice(sabcDTArray, func(i, j int) bool {
		return sabcDTArray[i].CreateAt.After(sabcDTArray[j].CreateAt)
	})
	fmt.Println("desc: ", sabcDTArray)
}
