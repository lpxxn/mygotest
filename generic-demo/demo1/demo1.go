package main

import "fmt"

func main() {
	ints := map[string]int64{
		"first":  12,
		"second": 5,
	}

	floats := map[string]float64{
		"first":  1.23,
		"second": 2.45,
	}
	fmt.Printf("gnenric sums: %v and %v\n", SumIntOrFloats[string, int64](ints), SumIntOrFloats[string, float64](floats))
	fmt.Printf("gnenric sums: %v and %v\n", SumIntOrFloats(ints), SumIntOrFloats(floats))

	fmt.Printf("gnenric sums: %v and %v\n", Sum(ints), Sum(floats))
	strs := map[string]string{
		"a": "a",
		"b": "b",
	}
	fmt.Println(Sum(strs))
}

func SumIntOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

type Number interface {
	int64 | float64 | int | string
}

func Sum[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
