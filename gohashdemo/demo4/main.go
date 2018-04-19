package main

import (
	"hash/crc32"
	"fmt"
	"github.com/speps/go-hashids"
	"github.com/mygotest/gohashdemo/demo4/ShortUrlGenerator"
	"strconv"
	"github.com/teris-io/shortid"
)

func main() {
	crc32q := crc32.MakeTable(0xD5828281)
	hashInBytes := crc32.Checksum([]byte("Hello world"), crc32q)
	fmt.Printf("%08x\n", hashInBytes)

	//returnCRC32String = hex.EncodeToString(crc32.new)

	hd := hashids.NewData()
	hd.Salt = "this is my salt"
	hd.MinLength = 30
	h, _ := hashids.NewWithData(hd)
	e, _ := h.Encode([]int{45, 434, 1313, 99})
	fmt.Println(e)
	d, _ := h.DecodeWithError(e)
	fmt.Println(d)


	l_url := "https://www.baidu.com/s?wd=CRC32&rsv_spt=1&rsv_iqid=0x97212db00004f4e6&issp=1&f=8&rsv_bp=0&rsv_idx=2&ie=utf-8&tn=baiduhome_pg&rsv_enter=1&rsv_sug3=1&rsv_sug1=1&rsv_sug7=100&rsv_sug2=0&inputT=6&rsv_sug4=7"
	fmt.Println(ShortUrlGenerator.Transform(l_url))


	var a int64 = 124567894654
	a_16 := strconv.FormatInt(a, 16)
	fmt.Println(a_16)

	//a = 160
	a = 160
	a_16 = strconv.FormatInt(a, 16)
	fmt.Println(a_16)

	sid, _ := shortid.New(1, shortid.DefaultABC, 2342)
	ssid,e1 := sid.Generate()
	fmt.Println(ssid, e1)
}

