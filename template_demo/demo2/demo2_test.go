package demo2

import (
	"os"
	"testing"
	"text/template"
)

// if else
func TestIfStatement1(t *testing.T) {
	//var b bytes.Buffer
	tmpl, err := template.New("test").Parse(`
		{{if .name}}
		a
		{{else}}
		b 	
		{{end}}`)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, map[string]interface{}{"name": "world"})
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		panic(err)
	}
	//t.Log(b.String())
	/*
		{{if pipeline}} T1 {{else}} T0 {{end}}
		If the value of the pipeline is empty, T0 is executed;
		otherwise, T1 is executed. Dot is unaffected.
	*/
}

func Test_Range0(t *testing.T) {
	tmpl, err := template.New("test").Parse(`
	{{range .name}}
    {{.}}
    {{- end}}
`)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, map[string]interface{}{"name": []string{"tom", "jack", "angle"}})
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		panic(err)
	}
}

// 多排显示
func Test_Range1(t *testing.T) {
	tmpl, err := template.New("test").Parse(`{{range .name}}
		{{.}}
		{{else}}
		b 	
		{{end}}`)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, map[string]interface{}{"name": []string{"tom", "jack", "angle"}})
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		panic(err)
	}
}

// 放在一排显示
func Test_Range2(t *testing.T) {
	tmpl, err := template.New("test").Parse(`
		{{range .name}}{{.}} {{else}}
		b 	
		{{end}}`)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, map[string]interface{}{"name": []string{"tom", "jack", "angle"}})
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		panic(err)
	}
}

// {{-  -}} 去掉空格{{ 左边 或者}} 右边，所有的空格
func Test_RemoeSpace(t *testing.T) {
	tmpl, err := template.New("test").Parse(`
	--
	a {{ 23 }} < {{ 45 }} b
	--
	`)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		panic(err)
	}

	tmpl, err = template.New("test").Parse(`
	--
	a  {{ 23 -}} < {{- 45 }}  b
	--
	`)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		panic(err)
	}

	tmpl, err = template.New("test").Parse(`
	--
	a  {{- 23 -}} < {{- 45 -}}  b
	--
	`)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		panic(err)
	}
	tmpl, err = template.New("test").Parse(`
	--
	{{- 23 -}} < {{- 45 -}}
	--
	`)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		panic(err)
	}
}
