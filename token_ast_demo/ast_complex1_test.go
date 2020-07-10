package token_ast_demo

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

// https://github.com/chai2010/go-ast-book/blob/master/ch7/readme.md
func TestType1(t *testing.T) {
	const src = `package foo
type Int1 int
type Int2 pkgA.int
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "a.go", src, parser.AllErrors)
	if err != nil {
		t.Fatal(err)
	}
	for _, decl := range f.Decls {
		//ast.Print(nil, decl.(*ast.GenDecl).Specs[0])
		ast.Print(nil, decl.(*ast.GenDecl).Specs)
	}
	// 而Int2的Type定义对应的时候*ast.SelectorExpr表示是其它包的命名类型
}

func TestPointer1(t *testing.T) {
	const src = `package foo
type Int1 *int
type Int2 *pkgA.int
type Int1 **int
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "a.go", src, parser.AllErrors)
	if err != nil {
		t.Fatal(err)
	}
	for _, decl := range f.Decls {
		//ast.Print(nil, decl.(*ast.GenDecl).Specs[0])
		ast.Print(nil, decl.(*ast.GenDecl).Specs)
	}
}

func TestArray1(t *testing.T) {
	const src = `package foo
type a [11]int
type IntArrayArray [1][2]int
type b []string
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "a.go", src, parser.AllErrors)
	if err != nil {
		t.Fatal(err)
	}
	for _, decl := range f.Decls {
		//ast.Print(nil, decl.(*ast.GenDecl).Specs[0])
		ast.Print(nil, decl.(*ast.GenDecl).Specs)
	}
}

/*
type ArrayType struct {
	Lbrack token.Pos // position of "["
	Len    Expr      // Ellipsis node for [...]T array types, nil for slice types
	Elt    Expr      // element type
}
*/

func TestStruct1(t *testing.T) {
	const src = `package foo
type MyStruct struct {
	a, b int "int value"
	string
	c int32
}
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "a.go", src, parser.AllErrors)
	if err != nil {
		t.Fatal(err)
	}
	for _, decl := range f.Decls {
		//ast.Print(nil, decl.(*ast.GenDecl).Specs[0])
		ast.Print(nil, decl.(*ast.GenDecl).Specs)
	}
}

func TestStruct2(t *testing.T) {
	const src = `package foo
type B struct {
	b int
}
type MyStruct struct {
	a, b int "int value"
	string
	c int32
	bs B
	dd pagD.A
	ee E
}
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "a.go", src, parser.AllErrors)
	if err != nil {
		t.Fatal(err)
	}
	for _, decl := range f.Decls {
		//ast.Print(nil, decl.(*ast.GenDecl).Specs[0])
		ast.Print(nil, decl.(*ast.GenDecl).Specs)
	}
}

func TestMap1(t *testing.T) {
	const src = `package foo
type IntStringMap map[int]string
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "a.go", src, parser.AllErrors)
	if err != nil {
		t.Fatal(err)
	}
	for _, decl := range f.Decls {
		//ast.Print(nil, decl.(*ast.GenDecl).Specs[0])
		ast.Print(nil, decl.(*ast.GenDecl).Specs)
	}
}

func TestChan1(t *testing.T) {
	const src = `package foo
type c chan int
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "a.go", src, parser.AllErrors)
	if err != nil {
		t.Fatal(err)
	}
	for _, decl := range f.Decls {
		//ast.Print(nil, decl.(*ast.GenDecl).Specs[0])
		ast.Print(nil, decl.(*ast.GenDecl).Specs)
	}
	/*
	    12  .  .  Type: *ast.ChanType {
	    13  .  .  .  Begin: 20
	    14  .  .  .  Arrow: 0
	    15  .  .  .  Dir: 3
	    16  .  .  .  Value: *ast.Ident {
	    17  .  .  .  .  NamePos: 25
	    18  .  .  .  .  Name: "int"
	    19  .  .  .  }
	    20  .  .  }
	其中ast.ChanType.Dir值是3，也就是SEND|RECV比特位组合
	 */
}

func TestFuncType1(t *testing.T) {
	const src = `package foo
type FuncType func(a, b int) bool
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "a.go", src, parser.AllErrors)
	if err != nil {
		t.Fatal(err)
	}
	for _, decl := range f.Decls {
		//ast.Print(nil, decl.(*ast.GenDecl).Specs[0])
		ast.Print(nil, decl.(*ast.GenDecl).Specs)
	}
}

func TestInterfaceType1(t *testing.T) {
	const src = `package foo
type IntReader interface {
	Read() int
}
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "a.go", src, parser.AllErrors)
	if err != nil {
		t.Fatal(err)
	}
	for _, decl := range f.Decls {
		//ast.Print(nil, decl.(*ast.GenDecl).Specs[0])
		ast.Print(nil, decl.(*ast.GenDecl).Specs)
	}
}

