package main

import (
	"bytes"
	"encoding/binary"
)

func main() {
	str1 := "abcdef"
	var writeBuf bytes.Buffer
	if err := binary.Write(&writeBuf, binary.BigEndian, []byte(str1)); err != nil {
		panic(err)
	}

	writeBuf.Reset()
	var vI32 int32 = 123456789
	if err := binary.Write(&writeBuf, binary.BigEndian, vI32); err != nil {
		panic(err)
	}
	print(writeBuf.Bytes())

	writeBuf.Reset()
	var vI64 int64 = 123
	if err := binary.Write(&writeBuf, binary.BigEndian, vI64); err != nil {
		panic(err)
	}

	writeBuf.Reset()
	str1 = "汉字"
	if err := binary.Write(&writeBuf, binary.BigEndian, []byte(str1)); err != nil {
		panic(err)
	}
}
