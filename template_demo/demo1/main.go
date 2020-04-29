package main

import (
	"os"
	"text/template"
)

func main() {
	comment1()
	test1()
	test2()
	test3()
	test4Functions()
	fieldInStruct1()
}

func comment1() {
	tmpl, err := template.New("test").Parse(`
	hello {{/* a comment */}} 
	{{"\"output\""}} A string constant
	{{printf "%q" "haha"}} 	A function call.
	{{"abcdef" | printf "%q"}}

	A raw string constant.
`)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		panic(err)
	}
}

// Renders the root element
func test1() {
	tmpl, err := template.New("test").Parse("hello {{.}}! \n")
	if err != nil {
		panic(err)
	}
	//err = tmpl.Execute(os.Stdout, map[string]interface{}{"name": "world"})
	err = tmpl.Execute(os.Stdout, map[string]interface{}{"name": "world", "a": "bcdef"})
	// output hello my name is A!
	if err != nil {
		panic(err)
	}
}

// Renders the root element
func test2() {
	tmpl, err := template.New("test").Parse("hello {{.}}! \n")
	if err != nil {
		panic(err)
	}
	//err = tmpl.Execute(os.Stdout, map[string]interface{}{"name": "world"})
	err = tmpl.Execute(os.Stdout, "my name is A")
	// output hello map[a:bcdef name:world]!
	if err != nil {
		panic(err)
	}
}

// Renders the root element
func test3() {
	v := struct {
		Values []string
	}{
		[]string{"a", "b", "c"},
	}
	tmpl, err := template.New("test").Parse(`
--------------------	
{{- range .Values}}
"{{.}}"
{{- end}}
--------------------
`)
	if err != nil {
		panic(err)
	}
	//err = tmpl.Execute(os.Stdout, map[string]interface{}{"name": "world"})
	err = tmpl.Execute(os.Stdout, v)
	// output hello map[a:bcdef name:world]!
	if err != nil {
		panic(err)
	}
}

// function
func test4Functions() {
	tmpl, err := template.New("test").Funcs(template.FuncMap{
		"name": func() string {
			return "A"
		},
	}).Parse(`
test4Functions {{name}}!
`)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		panic(err)
	}
}

// Renders the -field in a nested element
func fieldInStruct1() {
	user := struct {
		Name string
	}{
		"jack",
	}
	tmpl, err := template.New("test").Parse("hello {{.Name}}! \n")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, user)
	// output hello map[a:bcdef name:world]!
	if err != nil {
		panic(err)
	}
}
