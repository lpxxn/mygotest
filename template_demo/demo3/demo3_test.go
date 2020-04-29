package demo3

import (
	"html/template"
	"os"
	"testing"
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
