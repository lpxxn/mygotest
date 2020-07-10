package token_ast_demo

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestFunc1(t *testing.T) {
	src := `package t
type Person struct {
	name string
	desc string
	age  int
}
func (p *Person) Name() string {
	return p.name
}
func (p *Person) SetInfo(name, desc string, age int) (bool, error) {
	p.name = name
	p.desc = desc
	p.age = age
	return true, nil
}
func NewPerson() *Person {
	return &Person{}
}
`
	fst := token.NewFileSet()
	f, err := parser.ParseFile(fst, "a.go", src, parser.AllErrors)
	if err != nil {
		t.Fatal(err)
	}
	for _, decl := range f.Decls {
		t.Logf("decl type: %T\n", decl)
		if f, ok := decl.(*ast.FuncDecl); ok {
			t.Logf("recv: %#v, name: %#v, type: %#v, body %#v\n", f.Recv, f.Name, f.Type, f.Body)
			ast.Print(nil, f)
		} else {
			ast.Print(nil, decl)
		}
	}
}
