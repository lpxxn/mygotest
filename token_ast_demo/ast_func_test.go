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

func TestBasicList1(t *testing.T) {
	expr, _ := parser.ParseExpr(`"9527"`)
	ast.Print(nil, expr)

	expr, _ = parser.ParseExpr(`9527`)
	ast.Print(nil, expr)

	expr, _ = parser.ParseExpr(`952.7`)
	ast.Print(nil, expr)

	expr, _ = parser.ParseExpr(`true`)
	ast.Print(nil, expr)
	/*
	   0  *ast.Ident {
	   1  .  NamePos: 1
	   2  .  Name: "true"
	   3  .  Obj: *ast.Object {
	   4  .  .  Kind: bad
	   5  .  .  Name: ""
	   6  .  }
	   7  }
	*/

	lit := &ast.BasicLit{
		Kind:  token.FLOAT,
		Value: "1",
	}
	ast.Print(nil, lit)

	ast.Print(nil, ast.NewIdent(`false`))
	/*
		Bad表示未知的类型，其它的分别对应Go语言中包、常量、类型、变量、函数和标号等语法结构。而对于标识符中更具体的类型（比如是整数还是布尔类型）则是由ast.Object的其它成员描述。
	*/

}

func TestBasicList2(t *testing.T) {
	expr, err := parser.ParseExpr(`a == true`)
	if err != nil {
		t.Fatal(err)
	}
	ast.Print(nil, expr)

	fset := token.NewFileSet()
	const src = `package pkgname
var a int = 1
var length bool = true
`
	f, err := parser.ParseFile(fset, "hello.go", src, parser.AllErrors)
	if err != nil {
		t.Fatal(err)
	}
	ast.Print(nil, f)
}
