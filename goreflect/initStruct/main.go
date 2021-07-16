package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Config struct {
	Name string
	Meta struct {
		Desc       string
		Properties map[string]string
		Users      []string
	}
}

func initializeStruct(t reflect.Type, v reflect.Value) {
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		ft := t.Field(i)
		switch ft.Type.Kind() {
		case reflect.Map:
			f.Set(reflect.MakeMap(ft.Type))
		case reflect.Slice:
			f.Set(reflect.MakeSlice(ft.Type, 0, 0))
		case reflect.Chan:
			f.Set(reflect.MakeChan(ft.Type, 0))
		case reflect.Struct:
			initializeStruct(ft.Type, f)
		case reflect.Ptr:
			fv := reflect.New(ft.Type.Elem())
			initializeStruct(ft.Type.Elem(), fv.Elem())
			f.Set(fv)
		default:
		}
	}
}

func main() {
	t := reflect.TypeOf(Config{})
	v := reflect.New(t)
	initializeStruct(t, v.Elem())
	c := v.Interface().(*Config)
	fmt.Println(*c)
	cJson, err := json.Marshal(c)
	fmt.Println(string(cJson), err)
	c.Meta.Properties["color"] = "red"          // map was already made!
	c.Meta.Users = append(c.Meta.Users, "srid") // so was the slice.
	fmt.Println(v.Interface())
}
