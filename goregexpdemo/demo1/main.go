package main

import (
	"regexp"
	"fmt"
	"time"
)


func main() {
	fmt.Println(time.Now().Unix())
	score := 0
	if ok, _ := regexp.MatchString(`[_\-+=*!@#$%^&():;{}[\]|<>,.']`, "sadfasdf+asdfasdf"); ok{
		score++
		fmt.Println("score3:",score)
	}




}
