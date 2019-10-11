package utils

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func RandomInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func RandomSleep(min, max int) {
	sleepT := RandomInt(min, max)
	if sleepT > 0 {
		sleepTotal := time.Second * time.Duration(sleepT)
		time.Sleep(sleepTotal)

	}
}