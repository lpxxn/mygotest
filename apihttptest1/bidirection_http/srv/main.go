package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	//bodyStr := `{"a" = "aaaa"}`
	bodyStr := `a=aaaa&c=bdee`
	re, err := http.NewRequest(http.MethodPost, "http://www.baidu.com?wudsa=dasjkdjaskfj", bytes.NewReader([]byte(bodyStr)))
	re.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		panic(err)
	}
	fmt.Println(re.ParseForm())
	fmt.Println(re.Form)
}
