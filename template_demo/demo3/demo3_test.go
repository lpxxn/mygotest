package demo3

import (
	"html/template"
	"os"
	"testing"
	"time"
)

// 定义变量
func TestDefineValue(t *testing.T) {
	tmpl, err := template.New("test").Parse(`{{$myName := .name}}
	my name is {{$myName}}
	var Name = "{{$myName}}"
	`)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, map[string]interface{}{"name": "tom"})
	if err != nil {
		panic(err)
	}
}

// 定义方法
func TestFunc1(t *testing.T) {
	tmpl, err := template.New("test").Funcs(template.FuncMap{
		"now": func() string {
			return time.Now().Format("2006-01-02 15:04:05")
		},
		"say": func(name string) string {
			return "hello " + name
		},
	}).Parse(`{{say .name}} 
	time: {{now}}"
	`)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, map[string]interface{}{"name": "tom"})
	if err != nil {
		panic(err)
	}
}

// pipeline
func TestPipline(t *testing.T) {
	// html 内置函数
	// 内置函数 https://golang.org/pkg/text/template/#hdr-Functions
	// 左边的输出作为右边的输入

	tmpl, err := template.New("test").Parse(`{{. | html}} {{. | urlquery}}
	{{with $x := "output"}}{{$x}}{{end}}
    {{with $x := "output"}}{{$x | printf "abc: %s"}}{{end}}
	`)

	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, map[string]interface{}{"name": "中国"})
	if err != nil {
		panic(err)
	}
}

func TestBuildInFunctionCall(t *testing.T) {
	type X struct {
		Y func(a int, b int) int
	}
	x := X{Y: func(a int, b int) int {
		return a + b
	}}
	tmpl, err := template.New("test").Parse(`{{call .X.Y 1 2}}
	`)

	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, map[string]interface{}{"X": x})
	if err != nil {
		panic(err)
	}
}
