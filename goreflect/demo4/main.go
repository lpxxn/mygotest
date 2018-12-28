package main

import (
	"fmt"
	"reflect"
)

type R1 struct {
	Name string
}

type MySlice []R1


func (m *MySlice) SetReflect() {
	reflectSlice(m)
}

func (m *MySlice) SetReflect2() {
	reflectSlice2(m)
}

func reflectSlice(m interface{}) {
	sliceValue := reflect.ValueOf(m).Elem()
	for i := 0; i < sliceValue.Len(); i++ {
		v := sliceValue.Index(i)

		fmt.Println("file number", v.NumField(), " v type :", v.Type().String())
		//revValue := v.Interface().(R1)

		pv := reflect.New(v.Type())
		pv.Elem().Set(v)
		v.Field(0).SetString("cccccc")
		reflectStructInterface(pv.Interface())
	}
}



func reflectSlice2(m interface{}) {
	sliceValue := reflect.ValueOf(m).Elem()
	for i := 0; i < sliceValue.Len(); i++ {
		v := sliceValue.Index(i)

		fmt.Println("file number", v.NumField(), " v type :", v.Type().String())
		//revValue := v.Interface().(R1)

		pv := reflect.New(v.Type())
		pv.Elem().Set(v)
		v.Field(0).SetString("cccccc")
		reflectStructInterface(pv.Interface())
		v.Set(pv.Elem())
	}
}

func reflectStructInterface(param interface{}) {
	elType := reflect.TypeOf(param)
	eleValue := reflect.ValueOf(param)

	if elType.Kind() != reflect.Ptr {
		fmt.Println("not pointer")
		return
	}
	if eleValue.Elem().Kind() != reflect.Struct {
		fmt.Println("not Struct")
		return
	}

	fmt.Println("file number", eleValue.Elem().NumField())
	eleValue.Elem().Field(0).SetString("sfsdfsdf")
}


func main() {
	m1 := MySlice{
		{Name: "li"},
		{Name: "peng"},
	}
	m1.SetReflect()
	fmt.Println(m1)
	m1.SetReflect2()
	fmt.Println(m1)
}