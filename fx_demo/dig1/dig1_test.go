package dig1

import (
	"testing"

	"go.uber.org/dig"
)

func TestDit1(t *testing.T) {
	d := dig.New()

	type out struct {
		dig.Out

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
		dig.In

		Values []int `group:"val"`
	}

	err := d.Invoke(func(i in) {
		t.Log(i.Values)
	})
	if err != nil {
		t.Fatal(err)
	}
}
