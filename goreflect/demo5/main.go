package main

import (
	"fmt"
	"reflect"
)

func main() {
	s2 := []int64{1, 2, 3}
	isEqual := ElmInSlice(int64(1), s2)
	fmt.Println("is equal :", isEqual)

	s1 := []int{1, 2, 3}
	isEqual = ElmInSlice(int(1), s1)
	fmt.Println("is equal :", isEqual)
}

func ElmInSlice(id interface{}, slice interface{}) bool {
	sliceType := reflect.TypeOf(slice)
	fmt.Println(sliceType.Kind())
	if sliceType.Kind() != reflect.Slice {
		return false
	}
	sliceValues := reflect.ValueOf(slice)
	idValue := reflect.ValueOf(id)
	for n := 0; n < sliceValues.Len(); n++ {
		if reflect.DeepEqual(sliceValues.Index(n).Interface(), idValue.Interface()) {
			return true
		}
	}
	return false
}
