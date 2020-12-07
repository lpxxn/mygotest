package main

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func TestPerm(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for _, value := range rand.Perm(4) {
		fmt.Println(value)
	}
}

func TestIntSliceToString(t *testing.T) {
	delim := ","
	a := []int64{112132312, 25649684321654, 3154854, 5351546654658}
	str := strings.Replace(fmt.Sprint(a), " ", delim, -1)
	fmt.Println(str)
}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
	//return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
	//return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}
