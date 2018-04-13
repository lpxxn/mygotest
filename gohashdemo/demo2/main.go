package main

import (
	"github.com/teris-io/shortid"
	"fmt"
)

func main() {
	sid, err := shortid.New(1, shortid.DefaultABC, 2342)
	if err != nil {
		panic(err)
	}

	// then either:
	fmt.Println(sid.Generate())
	fmt.Println(sid.Generate())


	// or:
	shortid.SetDefault(sid)
	// followed by:
	fmt.Println(shortid.Generate())
	fmt.Println(shortid.Generate())
}
