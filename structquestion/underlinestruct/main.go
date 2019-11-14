package main

//This technique enforces keyed fields when declaring a struct.
//For example, the struct:
type SomeType struct {
	F1 string
	F2 bool
	_  struct{}
}

type SomeType2 struct {
	F1 string
	F2 bool
}

//One reason for doing this is to allow additional fields to be added to the struct in the future without breaking existing code.
func main() {
	// ALLOWED:
	_ = SomeType{F1: "f"}
	// compile error
	//underT1Error := SomeType{"abc", true}

	// ok
	_ = SomeType2{"faf", false}
}
