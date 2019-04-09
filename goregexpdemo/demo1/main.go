package main

import (
	"fmt"
	"regexp"
	"time"
)

func main() {
	fmt.Println(time.Now().Unix())
	score := 0
	if ok, _ := regexp.MatchString(`[_\-+=*!@#$%^&():;{}[\]|<>,.']`, "sadfasdf+asdfasdf"); ok {
		score++
		fmt.Println("score3:", score)
	}

	if ok, _ := regexp.MatchString(`^ZONE\d*$`, "ZONE12122323"); ok {
		fmt.Println("ZONE12122323 ok")
	}

	if ok, _ := regexp.MatchString(`^ZONE\d*$`, "zone12122323"); ok {
		fmt.Println("zone12122323 ok")
	}

	if ok, _ := regexp.MatchString(`^ZONE\d*$`, "zonezbce12122323"); ok {
		fmt.Println("zonezbce12122323 ok")
	}

	if ok, _ := regexp.MatchString(`^ZONE\d*$`, "ZONEzbce12122323"); ok {
		fmt.Println("ZONEzbce12122323 ok")
	}

	if ok, _ := regexp.MatchString(`^ZONE\d*$`, "ZONE12zbce12122323"); ok {
		fmt.Println("ZONEzbce12122323 ok")
	}

	re := regexp.MustCompile(`^ZONE\d*$`)
	if ok := re.MatchString("ZONE12122323"); ok {
		fmt.Println("MustCompile ZONE12122323 ok")
	}
}
