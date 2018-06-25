package main

import (
	"math/rand"
	"time"
	"fmt"
	"hash/crc32"
)

var BehaviorCh []chan string
const BehCount int = 32

func InitBehaviorCh() {
	for i := 0; i < BehCount; i++ {
		BehaviorCh = append(BehaviorCh, make(chan string, 3000))
	}
}


func ProcessTask(i int) {
	for {
		value, ok := <- BehaviorCh[i]
		fmt.Printf("index : %d , value: %s \n", i, value)
		if ok {
			fmt.Printf("index : %d ok \n", i)
		}
	}
}


func  main() {

	InitBehaviorCh()

	for i := 0; i < BehCount; i++ {
		go ProcessTask(i)
	}

	//
	dTicker := time.NewTicker(time.Millisecond * 100)
	for {
		select {
			case <- dTicker.C:
				rand.Seed(time.Now().UnixNano())
				l := rand.Intn(10)
				str := RandCode("c", l)
				// consistent hash modular get the index
				BehaviorCh[GetCrc(str) % uint32(BehCount)] <- str

		}
	}

}


// GetCrc ...
func GetCrc(key string) uint32 {
	if len(key) < 64 {
		var scratch [64]byte
		copy(scratch[:], key)
		return crc32.ChecksumIEEE(scratch[:len(key)])
	}
	return crc32.ChecksumIEEE([]byte(key))
}

// RandCode ...
func RandCode(ty string, l int) string {
	model := map[string]string{
		"d":   "0123456789",
		"c":   "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"s":   "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"mix": "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^",
	}

	return GetRandomString(model[ty], l)
}

// GetRandomString ...
func GetRandomString(str string, l int) string {
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}