package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	bodyStr := `{"a" = "aaaa"}`
	re, err := http.NewRequest(http.MethodPost, "http://www.baidu.com", strings.NewReader(bodyStr))
	re.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		panic(err)
	}
	fmt.Println(re.ParseForm())
	fmt.Println(re.Form)
}