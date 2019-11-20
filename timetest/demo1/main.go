package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
 	var td time.Duration
 	fmt.Println(td)
	t1 := time.Now().UnixNano()
	fmt.Println(t1)
	t2 := time.Now().Unix()
	fmt.Println(t2)
	fmt.Println(makeTimestamp())
	fmt.Println(RandomString("e", 15))
}
func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func RandomString(prefix string, n int) string {
	var letterRunes = []rune("1234567890")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return prefix + string(b)
}
