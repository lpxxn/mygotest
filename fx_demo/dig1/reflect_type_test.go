package dig1

import (
	"reflect"
	"testing"
)

type User struct {
	ID     int64  `json:"id" lang:"idLan"`
	Name   string `json:"name" lang:"nameLan"`
	Values []int  `json:"values" lang:"valuesLan"`
}

const (
	LangTag = "lang"
)

func TestStructFile(t *testing.T) {
	user1 := &User{ID: 1, Name: "zhang"}
	nameType := reflect.TypeOf(user1.Name)
	user1Type := reflect.TypeOf(user1)
	nameField, ok := user1Type.Elem().FieldByName("Name")
	t.Log(nameType.String(), " name: ", nameType)
	t.Log(nameField, " ", ok)
	if ok {
		langName := nameField.Tag.Get(LangTag)
		t.Log(langName)
		t.Log("value: ", reflect.ValueOf(nameField).Interface())
		t.Log("value: ", reflect.ValueOf(user1).Elem().FieldByName("Name").String())
		t.Log("field value: ", reflect.Indirect(reflect.ValueOf(user1)).FieldByName("Name"))
		t.Log(getFieldString(user1, "Name"))
	}
	t.Log(getFieldInteger(user1, "ID"))
}

func getFieldString(e *User, field string) string {
	r := reflect.ValueOf(e)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}

func getFieldInteger(e *User, field string) int {
	r := reflect.ValueOf(e)
	f := reflect.Indirect(r).FieldByName(field)
	return int(f.Int())
}
