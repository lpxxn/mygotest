package token_ast_demo

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)
// stmt expr
// 是 statement 和expression 也就是 语句 和 表达式
// 10 语句块和语句

func TestBlockFunc1(t *testing.T) {
	src := `package pkg_a
func main() {
}
`
	fst := token.NewFileSet()
	f, err := parser.ParseFile(fst, "a.go", src, parser.AllErrors)
	if err != nil {
		t.Fatal(err)
	}
	//ast.Print(nil, f)
	ast.Print(nil, f.Decls[0].(*ast.FuncDecl))
}

func TestBlockFunc2(t *testing.T) {
	src := `package pkg_a
func main() {
	{
	}
	{
	}
}
`
	fst := token.NewFileSet()
	f, err := parser.ParseFile(fst, "a.go", src, parser.AllErrors)
	if err != nil {
		t.Fatal(err)
	}
	//ast.Print(nil, f)
	ast.Print(nil, f.Decls[0].(*ast.FuncDecl))
}
