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

func (s SAbcDT) String() string {
	return fmt.Sprintf("Name: %s, CreateAt: %#v", s.Name, s.CreateAt)
}

func main() {
	var a string = "afsaf"
	fmt.Println(a)
	arr := AsortObjArray{AsortObj{5}, AsortObj{1}, AsortObj{18}, AsortObj{10}, AsortObj{8}, AsortObj{89}, AsortObj{6}}
	arr.SortDesc()
	fmt.Println(arr)
	arr.LimitEle(5)
	fmt.Println(arr)

	var sabcDTArray []*SAbcDT
	sabcDTArray = append(sabcDTArray, &SAbcDT{
		Name:     "cde",
		CreateAt: time.Now(),
	})

	sabcDTArray = append(sabcDTArray, &SAbcDT{
		Name:     "abcde",
		CreateAt: time.Now().AddDate(-1, 0, 0),
	})

	sabcDTArray = append(sabcDTArray, &SAbcDT{
		Name:     "abcde",
		CreateAt: time.Now().AddDate(-1, 2, 0),
	})

	sabcDTArray = append(sabcDTArray, &SAbcDT{
		Name:     "d",
		CreateAt: time.Now().AddDate(0, 2, 1),
	})

	sabcDTArray = append(sabcDTArray, &SAbcDT{
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
	var err error
	fmt.Printf("%#v\n", err)
	/*
		       [{cde 2020-01-15 12:03:32.50415 +0800 CST m=+0.000191218} {abcde 2019-01-15 12:03:32.50415 +0800 CST} {d 2020-03-16 12:03:32.504211 +0800 CST} {d 2019-08-24 12:03:32.504214 +0800 CST}]
		asc:   [{abcde 2019-01-15 12:03:32.50415 +0800 CST} {d 2019-08-24 12:03:32.504214 +0800 CST} {cde 2020-01-15 12:03:32.50415 +0800 CST m=+0.000191218} {d 2020-03-16 12:03:32.504211 +0800 CST}]
		desc:  [{d 2020-03-16 12:03:32.504211 +0800 CST} {cde 2020-01-15 12:03:32.50415 +0800 CST m=+0.000191218} {d 2019-08-24 12:03:32.504214 +0800 CST} {abcde 2019-01-15 12:03:32.50415 +0800 CST}]

	*/
	//sort.Slice(sabcDTArray, func(i, j int) bool {
	//	return sabcDTArray[i].CreateAt.Before(sabcDTArray[j].CreateAt) && sabcDTArray[i].Name > sabcDTArray[j].Name
	//})
	fmt.Printf("asc: %#v \n", sabcDTArray)
	ModifyName(sabcDTArray)
	fmt.Printf("asc: %#v \n", sabcDTArray)

	SortDtArray(sabcDTArray)
	fmt.Printf("asc: %#v \n", sabcDTArray)
}

func ModifyName(v []*SAbcDT) {
	for _, item := range v {
		item.Name += "aaaaaa"
	}
}

func SortDtArray(sabcDTArray []*SAbcDT) {
	sort.Slice(sabcDTArray, func(i, j int) bool {
		return sabcDTArray[i].CreateAt.Before(sabcDTArray[j].CreateAt)
	})
}
