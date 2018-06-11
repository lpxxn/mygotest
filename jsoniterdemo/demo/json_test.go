package main

import (
	"testing"
	"github.com/json-iterator/go"
	"encoding/json"
)

func Benchmark_Json(b *testing.B) {
	//p := fmt.Println
	val := []byte(`{"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}`)

	b.StopTimer()

	b.StartTimer()
	for i := 0; i < b.N; i++ {

		//jsoniter.Get(val, "Colors", 0).ToString()
		jsoniter.Get(val, "Name").ToString()
	}

}

func Benchmark_GoJson(b *testing.B) {
	//p := fmt.Println
	val := []byte(`{"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}`)
	entity := make(map[string]interface{})
	b.StopTimer()

	b.StartTimer()



	for i := 0; i < b.N; i++ {
		json.Unmarshal(val, &entity)
		//_ = entity["Name"]
	}
}
