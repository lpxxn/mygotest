package main

import (
	"math/rand"
	"time"
	"fmt"
	"hash/crc32"
)
// consistent hashing
// channel 组， 每一个组里都是buffered channel 3000个
// 这样做一致性哈希后，每个buffered channel 去处理真实的数据。
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
			fmt.Printf("channel index : %d process \n", i)
		}
	}
}


func  main() {

	InitBehaviorCh()

	for i := 0; i < BehCount; i++ {
		go ProcessTask(i)
	}

	//
	dTicker := time.NewTicker(time.Millisecond * 1000)
	for {
		select {
			case <- dTicker.C:
				rand.Seed(time.Now().UnixNano())
				l := rand.Intn(10)
				str := RandStringRunes( l)
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