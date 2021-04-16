package dig1

import (
	"testing"

	"go.uber.org/dig"
)

func TestDig1(t *testing.T) {
	d := dig.New()
	f1 := func() int {
		return 1
	}
	val1 := 0
	f2 := func(i int) (string, error) {
		val1 += i
		return "ok", nil
	}
	// 顺序无关
	if err := d.Provide(f2); err != nil {
		t.Fatal(err)
	}

	if err := d.Provide(f1); err != nil {
		t.Fatal(err)
	}

	if err := d.Invoke(func(i int) {
		t.Log(i)
		// val1 为0 说明f2没有调用
		t.Log(val1)
	}); err != nil {
		t.Fatal(err)
	}

	if err := d.Invoke(func(s string) {
		t.Log(s)
		// 1
		t.Log(val1)
	}); err != nil {
		t.Fatal(err)
	}

	if err := d.Invoke(func(i int, s string) {
		t.Log("str: ", s)
		t.Log(i)
		// val1
		t.Log(val1)
	}); err != nil {
		t.Fatal(err)
	}
}

func TestDit1(t *testing.T) {
	d := dig.New()
	type out struct {
		Out

		Value []int `group:"val,flatten"`
	}

	provide := func(i []int) {
		if err := d.Provide(func() out {
			return out{Value: i}
		}); err != nil {
			t.Fatal(err)
		}
	}

	provide([]int{1, 2})
	provide([]int{3, 4})

	type in struct {
		In

		Values []int `group:"val"`
	}
	// valcy
	//if err := d.Provide(func(i in) out {
	//	return out{Value: append(i.Values, 7, 8)}
	//}); err != nil {
	//	t.Log(err.Error())
	//	t.Fatal(err)
	//}

	err := d.Invoke(func(i in) {
		t.Log(i.Values)
	})
	if err != nil {
		t.Fatal(err)

	}
}
