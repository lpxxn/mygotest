package dig1

import (
	"reflect"
	"testing"
)

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func TestStructFile(t *testing.T) {
	user1 := &User{ID: 1, Name: "zhang"}
	nameType := reflect.TypeOf(user1.Name)
	user1Type := reflect.TypeOf(user1)
	nameField, ok := user1Type.Elem().FieldByName(nameType.Name())
	t.Log(nameType.String(), " name: ", nameType)
	t.Log(nameField, " ", ok)
}
