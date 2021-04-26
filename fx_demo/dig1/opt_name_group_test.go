package dig1

import (
	"testing"

	"go.uber.org/dig"
)

func TestOpt(t *testing.T) {
	type Student struct {
		Name string
		Age  int
	}

	NewUser := func(name string, age int) func() *Student {
		return func() *Student {
			return &Student{name, age}
		}
	}
	container := dig.New()

	if err := container.Provide(NewUser("tom", 3), dig.Name("s1")); err != nil {
		t.Fatal(err)
	}
	if err := container.Provide(NewUser("jerry", 1), dig.Name("s2")); err != nil {
		t.Fatal(err)
	}
	type UserParams struct {
		dig.In

		User1 *Student `name:"s1"`
		User2 *Student `name:"s2"`
	}
	PrintInfo := func(params UserParams) error {
		t.Log("Student 1")
		t.Logf("Name: %s", params.User1.Name)
		t.Logf("Age: %d", params.User1.Age)

		t.Log("Student 2")
		t.Logf("Name: %s", params.User2.Name)
		t.Logf("Age: %d", params.User2.Age)
		return nil
	}
	if err := container.Invoke(PrintInfo); err != nil {
		t.Fatal(err)
	}
}

func TestOptGroup(t *testing.T) {
	type Student struct {
		Name string
		Age  int
	}
	NewUser := func(name string, age int) func() *Student {
		return func() *Student {
			return &Student{name, age}
		}
	}
	container := dig.New()

	if err := container.Provide(NewUser("tom", 3), dig.Group("stu")); err != nil {
		t.Fatal(err)
	}
	if err := container.Provide(NewUser("jerry", 1), dig.Group("stu")); err != nil {
		t.Fatal(err)
	}

	type UserParams struct {
		dig.In

		StudentList []*Student `group:"stu"`
	}

	Info := func(params UserParams) error {
		for _, u := range params.StudentList {
			t.Log(u.Name, u.Age)
		}

		return nil
	}

	if err := container.Invoke(Info); err != nil {
		t.Fatal(err)
	}
}
