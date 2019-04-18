package main

import (
	"fmt"
	"time"
)

func makeTimestamp(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

func main() {
	t := time.Now()
	javaTime := makeTimestamp(t)
	fmt.Println("currentTime : ", t)
	fmt.Println("javaTime    : ", javaTime)
	pT := time.Unix(0, javaTime*int64(time.Millisecond))
	fmt.Println("pt          : ", pT)

	//return pData[i].ParseCreateTime().After(pData[j].ParseCreateTime())
}
