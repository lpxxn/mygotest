package main

import (
	"testing"
	"encoding/json"
	"github.com/json-iterator/go"
	"time"
	"fmt"
)

func Benchmark_Json(b *testing.B) {
	//p := fmt.Println
	val := []byte(`{"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}`)
	entity := make(map[string]interface{})
	b.StopTimer()

	b.StartTimer()
	for i := 0; i < b.N; i++ {

		//jsoniter.Get(val, "Colors", 0).ToString()
		//jsoniter.Get(val, "Name").ToString()
		jsoniter.Unmarshal(val, &entity)
		entity["time"] = time.Now().Format(time.RFC3339)
		jsoniter.Marshal(entity)
		//jsoniter.Get(val, "Type").ToString()
	}

}

func Benchmark_ByteReplace(b *testing.B) {
	//var rightC = []byte{'}'}
	val := []byte(`{"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}`)
	lenval := len(val)
	//entity := make(map[string]interface{})
	b.StopTimer()

	b.StartTimer()

	for i := 0; i < b.N; i++ {

		newval := []byte(fmt.Sprintf(`,"Time":"%s"}`, time.Now().Format(time.RFC3339)))
		nval := append(val[:lenval -1], newval...)
		_ = jsoniter.Valid(nval)

		//fmt.Println(isValid)
		//bytes.Replace(val, rightC)
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
		entity["time"] = time.Now().Format(time.RFC3339)
		json.Marshal(entity)
	}
}
