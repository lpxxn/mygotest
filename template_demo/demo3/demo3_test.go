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
