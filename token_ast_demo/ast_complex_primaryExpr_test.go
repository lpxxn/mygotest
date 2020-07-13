package token_ast_demo

import (
	"go/ast"
	"go/parser"
	"testing"
)

// 9 复合表达式

func TestConvertType1(t *testing.T) {
	expr, err := parser.ParseExpr(`int(x)`)
	if err != nil {
		t.Fatal(err)
	}
	ast.Print(nil, expr)
	/*
		转型操作居然是用ast.CallExpr表示，这说明在语法树中转型和函数调用的结构是完全一样的。这是因为在语法树解析阶段，解析器并不知道int(x)中的int是一个类型还是一个函数，因此也无法知晓这是一个转型操作还是一个函数调用。
	*/
}

func TestSelectorType1(t *testing.T) {
	expr, err := parser.ParseExpr(`x.y`)
	if err != nil {
		t.Fatal(err)
	}
	ast.Print(nil, expr)
}


func TestIndexType1(t *testing.T) {
	expr, err := parser.ParseExpr(`x[y]`)
	if err != nil {
		t.Fatal(err)
	}
	ast.Print(nil, expr)
}
