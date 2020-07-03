package token_ast_demo

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

const src1 = `package pkgName
import ("a"; "b")
type MyType int
const PI = 3.14
var Length = 1

func main() {}
`

func TestDir(t *testing.T) {

	fSet := token.NewFileSet()
	f, err := parser.ParseFile(fSet, "a.go", src1, parser.AllErrors)
	if err != nil {
		panic(err)
	}
	fmt.Println("package: ", f.Name)
	for _, s := range f.Imports {
		fmt.Println("import: ", s.Path.Value)
	}
	fmt.Println("----------------------")
	for _, d := range f.Decls {
		fmt.Printf("Decl: %T \n", d)
		//fmt.Printf("Decl: %#v \n", d)
	}
	fmt.Println("----------------------")

	for _, v := range f.Decls {
		if s, ok := v.(*ast.GenDecl); ok && s.Tok == token.IMPORT {
			for _, v := range s.Specs {
				fmt.Println("import: ", v.(*ast.ImportSpec).Path.Value)
			}
		}
	}
}

func TestVisitor(t *testing.T) {
	fst := token.NewFileSet()
	f, err := parser.ParseFile(fst, "hello.go", src1, parser.AllErrors)
	if err != nil {
		t.Fatal(err)
	}
	ast.Walk(new(myNodeVisitor), f)
}

type myNodeVisitor struct {
}

/*
type Visitor interface {
	Visit(node Node) (w Visitor)
}
*/
func (p *myNodeVisitor) Visit(n ast.Node) ast.Visitor {
	if x, ok := n.(*ast.Ident); ok {
		fmt.Println("myNodeVisitor.Visit: ", x.Name)
	}
	return p
}
