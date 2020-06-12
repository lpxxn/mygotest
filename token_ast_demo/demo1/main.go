package main

import (
	"fmt"
	"go/scanner"
	"go/token"
)

func main() {
	src := []byte(`
// test common
var a bool = true
b := 1 + 2.0
println("hello")
fmt.Println("world")
`)
	fset := token.NewFileSet()
	file := fset.AddFile("hello.go", fset.Base(), len(src))
	var s scanner.Scanner
	s.Init(file, src, nil, scanner.ScanComments)
	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("info pos:%s\t token:%s\t lit:%q\n", fset.Position(pos), tok, lit)
	}

}
