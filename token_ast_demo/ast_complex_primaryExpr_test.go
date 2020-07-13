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

func TestSliceExprType1(t *testing.T) {
	/*
		切片运算也是在一个主体表达式之后的中括弧中表示，不过切片运算至少有一个冒号分隔符，或者是两个冒号分隔符。切片运算主要包含开始索引、结束索引和最大范围三个部分。下面是x[1:2:3]切片运算的语法树：
		其中X、Low、High、Max分别表示切片运算的主体、开始索引、结束索引和最大范围
	*/
	expr, err := parser.ParseExpr(`x[0:1]`)
	if err != nil {
		t.Fatal(err)
	}
	ast.Print(nil, expr)
	t.Log("----------------------")
	expr, err = parser.ParseExpr(`x[0:1:3]`)
	if err != nil {
		t.Fatal(err)
	}
	ast.Print(nil, expr)
}

func TestAssertionExpr1(t *testing.T) {
	/*
		类型断言是判断一个接口对象是否满足另一个接口、或者接口持有的对象是否是一个确定的非接口类型
	*/
	expr, err := parser.ParseExpr(`x.(y)`)
	if err != nil {
		t.Fatal(err)
	}
	ast.Print(nil, expr)
}
