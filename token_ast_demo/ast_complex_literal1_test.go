package token_ast_demo

import (
	"go/ast"
	"go/parser"
	"testing"
)

// 函数面值
func TestFuncLiteral1(t *testing.T) {
	// FunctionLit   = "func" Signature FunctionBody .
	// 同样是由func关键字开始，后面是函数签名（输入参数和返回值）和函数体。函数面值和函数声明的最大差别是没有函数名字。
	// 该函数面值没有输入参数和返回值，同时函数体也没有任何语句，而且没有涉及上下文的变量引用，可以说是最简单的函数面值。因为面值也是一种表达式，因此可以用表达式的方式解析其语法树：
	expr, err := parser.ParseExpr("func() {}")
	if err != nil {
		t.Fatal(err)
	}

	ast.Print(nil, expr)
	expr, err = parser.ParseExpr("func (a string, b int) bool { return true }")
	if err != nil {
		t.Fatal(err)
	}

	ast.Print(nil, expr)
}
