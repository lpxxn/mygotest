package demo2

import (
	"os"
	"testing"
	"text/template"
)

// a
func TestIfStatement1(t *testing.T) {
	//var b bytes.Buffer
	tmpl, err := template.New("test").Parse(
		`{{if .name}}
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
