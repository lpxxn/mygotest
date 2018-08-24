package main

import (
	"fmt"
	"reflect"
)

type NInterTest interface {
	GetName() string
}

type TestNilI struct {
	Name string
	
}

func (t *TestNilI) GetName() string {
	return t.Name
}

func NilTest(t *TestNilI) bool {
	if t == nil {
		fmt.Println("NilTest t is nil")

		return true
	}
	fmt.Println("NilTest not nil ")
	return false
}

func NilInterface(t interface{}) {
	if t == nil {
		fmt.Println("NilInterface interface t is nil")

	} else {
		fmt.Println("NilInterface interface t is not nil")
	}
	fmt.Printf("NilInterface interface reflect is nil: %v \n", reflect.ValueOf(t).IsNil())
}

func NilInterface2(t NInterTest) {
	if t == nil {
		fmt.Println("NilInterface2 NInterTest t is nil")

	} else {
		fmt.Println("NilInterface2 NInterTest t is not nil")
	}

	fmt.Printf("NilInterface2 NInterTest reflect is nil: %v \n", reflect.ValueOf(t).IsNil())
}


func main() {
	t1 := &TestNilI{}
	if t1 == nil {
		fmt.Println("t1 is nil")
	}
	NilTest(t1)
	fmt.Println(t1)

	t2 := new(TestNilI)
	fmt.Println(t2)
	NilTest(t2)
	if t2 == nil {
		fmt.Println("t2 is nil")
	}

	t2 = nil
	fmt.Println(t2)

	if t2 == nil {
		fmt.Println("t2 is nil")
	}
	NilTest(t2)
	NilInterface(t2)
	NilInterface2(t2)

	rev1 := ReturnPointerValue()
	fmt.Printf("ReturnPointerValue is nil: %v \n", rev1 == nil)
	fmt.Printf("ReturnInterface is nil: %v \n", ReturnInterface() == nil)

	fmt.Printf("ReturnInterface2 is nil: %v \n", ReturnInterface2() == nil)  // false
	fmt.Printf("ReturnInterface3 is nil: %v \n", ReturnInterface3() == nil)  // false
}

func ReturnPointerValue() *TestNilI {
	return nil
}

func ReturnInterface() NInterTest {
	return nil
}

func ReturnInterface2() NInterTest {
	rev := ReturnPointerValue()
	return rev
}

func ReturnInterface3() NInterTest {
	var rev   *TestNilI
	return rev
}
