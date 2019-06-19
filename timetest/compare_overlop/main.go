package main

import (
	"fmt"
	"time"
)


//(StartDate1 <= EndDate2) and (StartDate2 <= EndDate1)

func main() {
	curTime := time.Now()
	fmt.Println("current time", curTime)
	f1 := "2006-01-02 15:04:05"
	beginTime1, _ := time.Parse(f1, "2019-06-19 03:00:00")
	fmt.Println(beginTime1)
	endTime1, _ := time.Parse(f1, "2019-06-19 03:10:00")
	fmt.Println(endTime1)
	sub1 := beginTime1.Sub(endTime1)
	fmt.Println(sub1)

	// 一样
	beginTime2, _ := time.Parse(f1, "2019-06-19 03:00:00")
	fmt.Println(beginTime2)
	endTime2, _ := time.Parse(f1, "2019-06-19 03:10:00")
	fmt.Println(endTime2)
	sub2 := beginTime1.Sub(endTime2)
	fmt.Println(sub2)

	if beginTime1.Sub(endTime2) <=0 && beginTime2.Sub(endTime1) <= 0 {
		fmt.Println("一样 overlap")
	} else {
		fmt.Println("一样")
	}


	// 包含 1
	beginTime2, _ = time.Parse(f1, "2019-06-19 02:00:00")
	fmt.Println(beginTime2)
	endTime2, _ = time.Parse(f1, "2019-06-19 03:20:00")
	fmt.Println(endTime2)
	sub2 = beginTime1.Sub(endTime2)
	fmt.Println(sub2)
	if beginTime1.Sub(endTime2) <=0 && beginTime2.Sub(endTime1) <= 0 {
		fmt.Println("包含 overlap")
	} else {
		fmt.Println("包含")
	}

	// 左边中间
	beginTime2, _ = time.Parse(f1, "2019-06-19 02:00:00")
	fmt.Println(beginTime2)
	endTime2, _ = time.Parse(f1, "2019-06-19 03:05:00")
	fmt.Println(endTime2)
	sub2 = beginTime1.Sub(endTime2)
	fmt.Println(sub2)
	if beginTime1.Sub(endTime2) <=0 && beginTime2.Sub(endTime1) <= 0 {
		fmt.Println("左边中间 overlap")
	} else {
		fmt.Println("左边中间")
	}

	// 左边左边
	beginTime2, _ = time.Parse(f1, "2019-06-19 02:00:00")
	fmt.Println(beginTime2)
	endTime2, _ = time.Parse(f1, "2019-06-19 03:10:00")
	fmt.Println(endTime2)
	sub2 = beginTime1.Sub(endTime2)
	fmt.Println(sub2)
	if beginTime1.Sub(endTime2) <=0 && beginTime2.Sub(endTime1) <= 0 {
		fmt.Println("左边左边 overlap")
	} else {
		fmt.Println("左边左边")
	}

	// 左边
	beginTime2, _ = time.Parse(f1, "2019-06-19 02:00:00")
	fmt.Println(beginTime2)
	endTime2, _ = time.Parse(f1, "2019-06-19 02:59:00")
	fmt.Println(endTime2)
	sub2 = beginTime1.Sub(endTime2)
	fmt.Println(sub2)
	if beginTime1.Sub(endTime2) <=0 && beginTime2.Sub(endTime1) <= 0 {
		fmt.Println("左边 overlap")
	} else {
		fmt.Println("左边")
	}

	// 中间右边
	beginTime2, _ = time.Parse(f1, "2019-06-19 03:05:00")
	fmt.Println(beginTime2)
	endTime2, _ = time.Parse(f1, "2019-06-19 04:05:00")
	fmt.Println(endTime2)
	sub2 = beginTime1.Sub(endTime2)
	fmt.Println(sub2)
	if beginTime1.Sub(endTime2) <=0 && beginTime2.Sub(endTime1) <= 0 {
		fmt.Println("中间右边 overlap")
	} else {
		fmt.Println("中间右边")
	}

	// 右边右边
	beginTime2, _ = time.Parse(f1, "2019-06-19 03:10:00")
	fmt.Println(beginTime2)
	endTime2, _ = time.Parse(f1, "2019-06-19 04:10:00")
	fmt.Println(endTime2)
	sub2 = beginTime1.Sub(endTime2)
	fmt.Println(sub2)
	if beginTime1.Sub(endTime2) <=0 && beginTime2.Sub(endTime1) <= 0 {
		fmt.Println("右边右边 overlap")
	} else {
		fmt.Println("右边右边")
	}

	// 右边
	beginTime2, _ = time.Parse(f1, "2019-06-19 03:11:00")
	fmt.Println(beginTime2)
	endTime2, _ = time.Parse(f1, "2019-06-19 04:10:00")
	fmt.Println(endTime2)
	sub2 = beginTime1.Sub(endTime2)
	fmt.Println(sub2)
	if beginTime1.Sub(endTime2) <=0 && beginTime2.Sub(endTime1) <= 0 {
		fmt.Println("右边 overlap")
	} else {
		fmt.Println("右边")
	}
}
