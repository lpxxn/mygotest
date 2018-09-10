package utils

import (
	"math/rand"
	"time"
)

func Random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max - min + 1) + min
}

func ReplaceAtIndex(in string, r string, index int) string {
	return in[:index] + r + in[index+1:]
}
