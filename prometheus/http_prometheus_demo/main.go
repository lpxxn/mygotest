package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	fmt.Println("begnin")
	f1 := func() {
		var i int64 = 0
		for {
			<- time.After(time.Second * 5)
			arr := make([]int64, 100, 100)
			arr = append(arr, i)
			arrStr := make([]string, 50, 50)
			arrStr = append(arrStr, "aaaaaaaaaaaaaa")
			str := `{"name": "li"}`
			arrStr = append(arrStr, str)
			m := make(map[string]interface{})
			json.Unmarshal([]byte(str), &m)
			json.Marshal(m)
			json.Marshal(arrStr)
			i += 1
			fmt.Println(fmt.Sprintf("current value %d", i))
		}
	}
	go f1()
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":5112", nil); err != nil {
		panic(err)
	}
	fmt.Println("end")
}
