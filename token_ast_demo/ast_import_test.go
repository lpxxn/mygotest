package token_ast_demo

import (
	"go/ast"
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

/*
Go语⾔中通过type关键字声明类型：⼀种是声明新的类型，另⼀种是为已
有的类型创建⼀个别名
*/
func TestAlia(t *testing.T) {
	// 新类型
	type MyInt1 int
	// 别名
	type MyInt2 = int

	var i int = 1
	var a1 MyInt1 = 11
	var a2 MyInt2 = 22
	// 错
	//if a1 == i {}
	if a1 == MyInt1(i) {
	}
	if a2 == i {
	}
	src := `package fool
	type MyInt1 int
	type Int2 = int
`
	fst := token.NewFileSet()
	f, err := parser.ParseFile(fst, "a.go", src, parser.AllErrors)
	if err != nil {
		t.Fatal(err)
	}
	for _, decl := range f.Decls {
		t.Logf("Decl: type: %T, v: %#v\n", decl, decl)
		if v, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range v.Specs {
				t.Logf("spec : %#v\n", spec)
				if v.Tok == token.TYPE {
					if ts, ok := spec.(*ast.TypeSpec); ok {
						t.Logf("typsSpec name: %#v type: %#v asign: %#v \n", ts.Name, ts.Type, ts.Assign)
					}
					/*
						其中最重要的是 TypeSpec.Name 成员，表示新声明类型的名字或者是已
						有类型的别名。⽽ TypeSpec.Assign 成员对应 = 符号的位置，如果该
						成员表示的位置有效，则表示这是为已有类型定义⼀个别名（⽽不是定义
						新的类型）
					*/
				}
			}
		}
	}
}

func TestConst(t *testing.T) {
	src := `package foo
const Pi = 3.14
const E float64 = 2.71828
const a1, a2 = 1, "aaa"
`
	fst := token.NewFileSet()
	f, err := parser.ParseFile(fst, "a.go", src, parser.AllErrors)
	if err != nil {
		t.Fatal(err)
	}
	for _, decl := range f.Decls {
		if v, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range v.Specs {
				t.Logf("spec : %#v\n", spec)
				if vs, ok := spec.(*ast.ValueSpec); ok {
					t.Logf("ValueSpec name: %#v type: %#v values: %#v \n", vs.Names, vs.Type, vs.Values)
					ast.Print(nil, vs)
				}
			}
		}
	}
}

func TestVar(t *testing.T) {
	src := `package foo
var Pi = 3.14
var E float64 = 2.71828
var a1, a2 = 1, "aaa"
`
	fst := token.NewFileSet()
	f, err := parser.ParseFile(fst, "a.go", src, parser.AllErrors)
	if err != nil {
		t.Fatal(err)
	}
	for _, decl := range f.Decls {
		if v, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range v.Specs {
				t.Logf("spec : %#v\n", spec)
				if vs, ok := spec.(*ast.ValueSpec); ok {
					t.Logf("ValueSpec name: %#v type: %#v values: %#v \n", vs.Names, vs.Type, vs.Values)
					ast.Print(nil, vs)
				}
			}
		}
	}
}

func TestVarGroup(t *testing.T) {
	src := `package foo
const Pi = 3.14
var (
	a1 = 1
	a2 = "aaa"
)
`
	fst := token.NewFileSet()
	f, err := parser.ParseFile(fst, "a.go", src, parser.AllErrors)
	if err != nil {
		t.Fatal(err)
	}
	// len(f.Decls) = 2 就2个元素
	for _, decl := range f.Decls {
		if v, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range v.Specs {
				t.Logf("spec : %#v\n", spec)
				if vs, ok := spec.(*ast.ValueSpec); ok {
					t.Logf("ValueSpec name: %#v type: %#v values: %#v \n", vs.Names, vs.Type, vs.Values)
					ast.Print(nil, vs)
				}
			}
		}
	}
}
