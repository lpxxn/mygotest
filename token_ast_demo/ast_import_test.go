package token_ast_demo

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestImport(t *testing.T) {
	src := `package foo
	import "pkg-a"
	import pab_b_v2 "pkg-b-v2"
	import . "pkg-c"
	import _ "pkg-d"
`
	fst := token.NewFileSet()
	f, err := parser.ParseFile(fst, "hello.go", src, parser.AllErrors)
	if err != nil {
		t.Fatal(err)
	}
	for _, s := range f.Imports {
		t.Logf("import name: %s, path: %#v", s.Name, s.Path)
	}
}
