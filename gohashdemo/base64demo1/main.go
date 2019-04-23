package main

import(
	"encoding/base64"
	"errors"
	"fmt"
)

func main() {
	str := "122341243-123asdfasdf-124123"
	encodedStr := base64.StdEncoding.EncodeToString([]byte(str))
	fmt.Println("Encoded: ", encodedStr)

	raw, err := base64.StdEncoding.DecodeString(encodedStr)
	if err != nil {
		panic(err)
	}
	if string(raw) != str{
		panic(errors.New(string(raw)))
	}
	fmt.Println(string(raw))
}
