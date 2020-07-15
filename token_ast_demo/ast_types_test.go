package token_ast_demo

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

/*
11
主流的编译器前端遵循词法解析、语法解析、语义解析等流程，然后才是基于中间表示的层层优化并最终产生目标代码。在得到抽象的语法树之后就表示完成了语法解析的工作。
不过在进行中间优化或代码生成之前还需要对抽象语法树进行语义分析。语义分析需要更深层次理解代码的语义，比如两个变量相加是否合法，外层作用域有多个同名的变量时如何选择等。
本章简单讨论go/types包的用法，展示如果通过该包实现语法树的类型检查功能。
*/

/*
虽然Go语言是基于包和目录来组织代码，但是Go语言在语法树解析阶段并不关心包之间的依赖关系。这是因为在语法树解析阶段并不对代码本身做语义检测，因此很多语法正确但是语义错误的代码也可以生成语法树。
*/
func TestAstTreeTypes1(t *testing.T) {
	src := `package a
func main() {
	var _ = "a" + 1
}
`
	f := func(src string) {
		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, "a.go", src, parser.AllErrors)
		if err != nil {
			t.Fatal(err)
		}
		ast.Print(fset, f)
		//ast.Print(fset, f)
		pkg, err := new(types.Config).Check("a.go", fset, []*ast.File{f}, nil)
		if err != nil {
			t.Error(err)
		}
		t.Log(pkg)
	}
	f(src)
	src = `package a
func main() {
	var _ = 2 + 1
}
`
	f(src)
}

/*
func (conf *Config) Check(path string, fset *token.FileSet, files []*ast.File, info *Info) (*Package, error)

第一个参数表示要检查包的路径，
第二个参数表示全部的文件集合（用于将语法树中元素的位置信息解析为文件名和行列号），
第三个参数是该包中所有文件对应的语法树，
最后一个参数可用于存储检查过程中产生的分析结果。如果成功该方法返回一个types.Package对象，表示当前包的信息。
*/

func TestAstTreeMultiPkg1(t *testing.T) {
	src := `package a
import "math"
func main() {
	var _ = "a" + math.Pi
}
`
	f := func(src string) {
		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, "a.go", src, parser.AllErrors)
		if err != nil {
			t.Fatal(err)
		}
		ast.Print(fset, f)
		//ast.Print(fset, f)
		conf := &types.Config{
			Importer: importer.Default(),
		}
		pkg, err := conf.Check("a.go", fset, []*ast.File{f}, nil)
		if err != nil {
			t.Error(err)
		}
		t.Log(pkg)
	}
	f(src)
	src = `package a
import "math"

func main() {
	var _ = 2 + math.Pi
}
`
	f(src)
}
