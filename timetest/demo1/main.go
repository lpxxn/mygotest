package main

import (
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"go.planetmeican.com/kiwi/toolbox/hashids"
	"hash/fnv"
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

	hd := hashids.NewData()
	hd.Salt = "THIS IS MY SALT"
	hd.Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	hd.MinLength = 10
	h, _ := hashids.NewWithData(hd)
	e, _ := h.EncodeInt64([]int64{88347612751293423, 38778162789239})
	fmt.Println("encode", e)
	d, _ := h.DecodeWithError(e)
	fmt.Println(d)
	decode(e)
}
func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func decode(s string) (uint32, error) {
	a, e := base64.StdEncoding.DecodeString(s)
	if e != nil {
		return 0, e
	}
	return binary.LittleEndian.Uint32(append(a, 0)), nil
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
