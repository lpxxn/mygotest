package utils

import (
	"math/rand"
	"time"
)

func init()  {
	rand.Seed(time.Now().UnixNano())
}

var letterRuns = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRuns[rand.Intn(len(letterRuns))]
	}
	return string(b)
}
