package dig1

import (
	"bytes"
	"fmt"
	"testing"

	"go.uber.org/dig"
)

func TestOut1(t *testing.T) {
	d := dig.New()

	type Student struct {
		Name string
		Age  int
	}
	type out struct {
		dig.Out

		StudentList []*Student `group:"students,flatten"`
	}

	provide := func(i []int) {
		if err := d.Provide(func() out {
			rev := out{StudentList: []*Student{}}
			for _, item := range i {
				rev.StudentList = append(rev.StudentList, &Student{Name: fmt.Sprintf("name: %d", item), Age: item})
			}
			return rev
		}); err != nil {
			t.Fatal(err)
		}
	}

	provide([]int{1, 2})
	provide([]int{3, 4})

	type in struct {
		dig.In

		Values []*Student `group:"students"`
	}

	err := d.Invoke(func(i in) {
		t.Log(i.Values)
	})
	if err != nil {
		t.Fatal(err)

	}

	b := &bytes.Buffer{}
	if err := dig.Visualize(d, b); err != nil {
		t.Fatal(err)
	}
	t.Log(b.String())
}
