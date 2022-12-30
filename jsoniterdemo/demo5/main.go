package main

import (
	"fmt"
	j "github.com/tidwall/gjson"
	sj "github.com/tidwall/sjson"
)

//var OriginJson =

func main() {
	myJson := MyJson{}
	myJson.OriginJson = `{"name":"tom","age":18,"hobby":["basketball","football"],"address":{"city":"beijing","street":"xizhimen"}}`
	fmt.Println("------开始遍历--------")

	newVal := "new"
	myJson.ParseJSON(newVal)
}

type MyJson struct {
	OriginJson string
	rev        []string
}

func (m *MyJson) ParseJSON(newVal string) {
	m.rev = make([]string, 0)
	keyPath := ""
	j.Parse(m.OriginJson).ForEach(func(key, value j.Result) bool {
		keyPath = key.String()
		m.parseJSON(key, value, keyPath, newVal)
		return true
	})
}

func (m *MyJson) GetResult() []string {
	return m.rev
}

func (m *MyJson) parseJSON(key, value j.Result, keyPath string, newVal string) {
	if value.IsObject() {
		value.ForEach(func(key, value j.Result) bool {
			m.parseJSON(key, value, fmt.Sprintf("%s.%s", keyPath, key), newVal)
			return true
		})
	} else if value.IsArray() {
		for idx, name := range value.Array() {
			tmpKeyPath := fmt.Sprintf("%s.%d", keyPath, idx)
			if name.IsObject() {
				name.ForEach(func(key, value j.Result) bool {
					m.parseJSON(key, value, fmt.Sprintf("%s.%s", tmpKeyPath, key), newVal)
					return true
				})
			} else {
				v, _ := sj.Set(m.OriginJson, tmpKeyPath, name.String()+newVal)
				m.rev = append(m.rev, v)
			}
		}
	} else {
		v, _ := sj.Set(m.OriginJson, keyPath, value.String()+newVal)
		m.rev = append(m.rev, v)
	}
}
