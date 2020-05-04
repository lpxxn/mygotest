package main_test

import (
	"os"
	"strings"
	"testing"
	"text/template"
)

func TestTemplate1(t *testing.T) {
	tmpl, err := template.New("test").Parse(`hello {{.name}}!
	obj: {{.}}`)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, map[string]interface{}{
		"name": "world", "age": 18})
	if err != nil {
		panic(err)
	}
}

func TestTemplate2(t *testing.T) {
	tmpl, err := template.New("test").Parse(`hello {{.Name}}!
	obj: {{.}}`)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, struct {
		Name string
		Age  int
	}{Name: "li", Age: 18})
	if err != nil {
		panic(err)
	}
}

func TestTemplate3(t *testing.T) {
	tmpl, err := template.New("test").Parse(`hello:    {{- .Name}}
	age: {{.Age -}}   !!!
	obj:     
	{{- . -}}   end.`)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, struct {
		Name string
		Age  int
	}{Name: "li", Age: 18})
	if err != nil {
		panic(err)
	}
}

func TestTemplate4(t *testing.T) {
	tmpl, err := template.New("test").Parse(`
	name: {{.Name}} 
	{{- if .Name}}
      string .Name true 
	{{else}} 
      string .Name false 
	{{end -}}
	desc: {{.Desc}} 
	{{- if .Desc}}
      string .Desc true 
	{{else}} 
      string .Desc false 
	{{end -}}
	age: {{.Age}} 
	{{- if .Age}}
      number .Age true 
	{{else}} 
	  number .Age true false
	{{end -}}
	isAdmin: {{.IsAdmin}} 
	{{- if .Age}}
      bool .IsAdmin true 
	{{else}} 
	  bool .IsAdmin true false
	{{end}}
	`)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, struct {
		Name    string
		Desc    string
		Age     int
		IsAdmin bool
	}{Name: "", Desc: "xyz", Age: 18, IsAdmin: true})
	if err != nil {
		panic(err)
	}
}

func TestTemplate5(t *testing.T) {
	tmpl, err := template.New("test").Funcs(template.FuncMap{
		"ReplaceAll": func(src string, old, new string) string {
			return strings.ReplaceAll(src, old, new)
		},
	}).Parse(`func replace:  {{ReplaceAll .Name "zhang" "li"}}`)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, struct {
		Name string
		Age  int
	}{Name: "zhang_san zhang_si", Age: 18})
	if err != nil {
		panic(err)
	}
}

func TestTemplate6(t *testing.T) {
	tmpl, err := template.New("test").Parse(`{{printf "name: %s age: %d" .Name .Age}}`)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, struct {
		Name string
		Age  int
	}{Name: "li", Age: 18})
	if err != nil {
		panic(err)
	}
}

func TestTemplate7(t *testing.T) {
	tmpl, err := template.New("test").Parse(`
	{{/* 注释 */}}
	{{define "content"}} hello {{.}} {{end}}
	content: {{template "content" "zhang san"}}`)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		panic(err)
	}
}
