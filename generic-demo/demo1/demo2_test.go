package main_test

import (
	"testing"
)

/*
~ 在这里应该可以理解为 泛类型 ，即所有以int64为基础类型的类型都能够被约束。

我们来举个例子：现在我们声明一个以 int64 为基础类型，取名为testInt
*/
type Number interface {
	~int64 | float64 | string
}

type testInt int64

func Sum[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func TestGenInt1(t *testing.T) {
	ints := map[string]testInt{
		"first":  34,
		"second": 12,
	}

	t.Logf("Generic Sum: %v\n", Sum(ints))
}
