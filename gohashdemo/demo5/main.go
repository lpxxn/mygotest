package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"errors"
)

func main() {
	hasher := md5.New()
	strs := `https://www.google.co.kr/search?ei=CBnQWuNzkPb_BJz2goAJ&q=mongodb+auto+increment+integer&oq=mongodb+auto+int&gs_l=psy-ab.3.1.0i30k1j0i8i30k1.58876.59812.0.63611.5.5.0.0.0.0.357.621测试一二三.2-1j1.2.0....0...1.1.64.psy-ab..3.2.621...0j0i67k1j0i131k1j0i7i30k1.0.OiUWokhME1w`
	hasher.Write([]byte(strs))
	hash_str := hex.EncodeToString(hasher.Sum(nil))
	fmt.Println(hash_str)


	fmt.Println(len(letterRunes))
	//rand.Seed(time.Now().UnixNano())
	//rand.Seed(1000000)
	//fmt.Println(RandStringRunes(8))

	var count int64 = 10000000
	map_data := make(map[string]struct{})
	var i int64
	for i = 0;count > i; i++ {
		//fmt.Println(i)
		rand.Seed(i)
		//rand.Seed(time.Now().UnixNano())
		temp_v := RandStringRunes(1)
		if _, ok := map_data[temp_v]; ok {
			fmt.Println("-------have key: ", temp_v)
			fmt.Println(len(map_data))
			panic(errors.New(temp_v))
		} else {
			map_data[temp_v] = struct{}{}
		}
	}
	fmt.Println(len(map_data))
}


var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}