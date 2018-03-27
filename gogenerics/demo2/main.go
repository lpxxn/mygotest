package main

import (
	"reflect"
	"fmt"
)

type Cabinet struct {
	s reflect.Value
}

func NewCabinet(t reflect.Type) *Cabinet {
	return &Cabinet{
		s: reflect.MakeSlice(reflect.SliceOf(t), 0, 10), // cap is arbitrary, we need to pass one here
	}
}


func (c *Cabinet) Put(val interface{}) {

	if reflect.ValueOf(val).Type() != c.s.Type().Elem() {
		panic(fmt.Sprintf("Put: cannot put a %T into a slice of %s", val, c.s.Type().Elem()))
	}

	c.s = reflect.Append(c.s, reflect.ValueOf(val))
}

func (c *Cabinet) Get(retref interface{}) {
	retref = c.s.Index(0)
	c.s = c.s.Slice(1, c.s.Len())
}

func reflectExample() {
	f := 3.14152
	g := 0.0
	c := NewCabinet(reflect.TypeOf(f))
	c.Put(f)
	c.Get(&g)
	fmt.Printf("reflectExample: %f (%T)\n", g, g)
}

func main() {
	reflectExample()
}