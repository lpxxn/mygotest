package dig1

import (
	"testing"

	"go.uber.org/dig"
)

func TestIn1(t *testing.T) {
	d := dig.New()

	type Config struct {
		Addr string
		Port int
	}

	type Student struct {
		Name string
		Age  int
	}

	// paramSingle{Type: t}
	f1 := func(c *Config) string {
		t.Log(*c)
		return "OK!"
	}

	f2 := func() (*Config, *Student, error) {
		return &Config{
				Addr: "127.0.0.1",
				Port: 10086,
			}, &Student{
				Name: "Doraemon",
				Age:  10,
			}, nil
	}
	if err := d.Provide(f1); err != nil {
		t.Fatal(err)
	}

	if err := d.Provide(f2); err != nil {
		t.Fatal(err)
	}

	type InParam struct {
		dig.In
		Conf *Config  //  paramObject.Fields -> [0]: paramSingle{Type: t}, [1]: paramSingle{Type: t}
		Stu  *Student
	}
	f3 := func(p InParam) error {
		t.Log(*p.Conf)
		t.Log(*p.Stu)
		return nil
	}
	if err := d.Invoke(f3); err != nil {
		t.Fatal(err)
	}
	t.Log(d.String())

}
