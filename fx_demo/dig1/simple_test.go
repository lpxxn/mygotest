package dig1

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"testing"

	"go.uber.org/dig"
)

func TestSimple1(t *testing.T) {
	type Config struct {
		Prefix string
	}

	c := dig.New()

	err := c.Provide(func() (*Config, error) {
		return &Config{Prefix: "[foo] "}, nil
	})
	if err != nil {
		panic(err)
	}
	err = c.Provide(func(cfg *Config) *log.Logger {
		return log.New(os.Stdout, cfg.Prefix, 0)
	})
	if err != nil {
		panic(err)
	}
	err = c.Invoke(func(l *log.Logger) {
		l.Print("You've been invoked")
	})
	if err != nil {
		panic(err)
	}

	b := &bytes.Buffer{}
	if err := dig.Visualize(c, b); err != nil {
		t.Fatal(err)
	}
	fmt.Println(b.String())
	//t.Log(b.String())

}
