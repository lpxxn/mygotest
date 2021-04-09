package dig1

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStructFile1(t *testing.T) {
	user1 := &User{ID: 1, Name: "zhang"}
	userType := reflect.TypeOf(user1).Elem()

	nameFieldType, _ := userType.FieldByName("Name")
	fmt.Println(nameFieldType.Type)          // string
	fmt.Println(nameFieldType.Type.String()) // string
	fmt.Println(nameFieldType.Type.Name())   // string

	valueFieldType, _ := userType.FieldByName("Values")
	fmt.Println(valueFieldType.Type.String())      // []int
	fmt.Println(valueFieldType.Type.Elem())        // int
	fmt.Println(valueFieldType.Type.Elem().Name()) // int
}

func TestStructFile2(t *testing.T) {
	user1 := &User{ID: 1, Name: "zhang"}
	nameType := reflect.ValueOf(user1.Name)
	t.Log("nameValue: ", nameType.String())

	nameField, ok := nameType.Interface().(reflect.StructField)
	// ok = false
	if ok {
		t.Log(nameField)
	} else {
		t.Log(nameType.Type())
	}
}
