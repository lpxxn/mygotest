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
			s: reflect.MakeSlice(reflect.SliceOf(t), 0, 10),
		}
}

func (c *Cabinet) Put(val interface{}) {
	if reflect.ValueOf(val).Type() != c.s.Type().Elem() {
		panic(fmt.Sprintf("Put: cannot put a %T into a slice of %s", val, c.s.Type().Elem()))
	}

	c.s = reflect.Append(c.s, reflect.ValueOf(val))
}


func (c *Cabinet) Get(retref interface{}) interface{} {

	//*retref.(*interface{}) = c.s.Index(0)

	//switch d := retref.(type) {
	//case *string:
	//	*d= "abcde"
	//case *float64:
	//	*d = 124
	//}


	//retref = c.s.Index(0)
	v := c.s.Index(0)
	fmt.Println(retref, "   ", v)
	c.s = c.s.Slice(1, c.s.Len())
	return retref
}


func main() {
	f := 3.1415926
	g := 0.0
	c := NewCabinet(reflect.TypeOf(f))
	c.Put(f)
	c.Put(1.2)
	c.Put(0.56)
	v1 := c.Get(&g)
	fmt.Println(v1)
	fmt.Println(v1.(*float64))
	fmt.Println(g)
	fmt.Println(c.s)

}
