package main

import (
	"fmt"
	"strconv"
)

type ByteSize uint64

const (
	B  ByteSize = 1 + iota
	KB          = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB

	fnUnmarshalText string = "UnmarshalText"
	maxUint64       uint64 = (1 << 64) - 1
	cutoff          uint64 = maxUint64 / 10
)

func main() {
	fmt.Println(B, KB, MB, TB)
	str := "hello world!"
	bytes := []byte(str)
	var len_bytes = len(bytes)
	fmt.Println(len_bytes, bytes)
	c := float64(len_bytes) / float64(KB)
	svalue := strconv.FormatFloat(c, 'f', 3, 32)
	svalue2 := fmt.Sprintf("%.3f", c)
	f, err := strconv.ParseFloat(svalue, 64)
	fmt.Println(svalue, svalue2, f, err)
	fmt.Println("kilobyte", len_bytes/KB, c)
}
