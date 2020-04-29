package demo2

import (
	"os"
	"testing"
	"text/template"
)

// if else
func TestIfStatement1(t *testing.T) {
	//var b bytes.Buffer
	// 当.name 为bool类型的时候，则为true表示执行，当.name 为string类型的时候，则非空表示执行。
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
				假设我们需要逻辑判断，比如与或、大小不等于等判断的时候，我们需要一些内置的模板函数来做这些工作，目前常用的一些内置模板函数有：

		not 非
		{{if not .condition}}
		{{end}}
		and 与
		{{if and .condition1 .condition2}}
		{{end}}
		or 或
		{{if or .condition1 .condition2}}
		{{end}}
		eq 等于
		{{if eq .var1 .var2}}
		{{end}}
		ne 不等于
		{{if ne .var1 .var2}}
		{{end}}
		lt 小于 (less than)
		{{if lt .var1 .var2}}
		{{end}}
		le 小于等于
		{{if le .var1 .var2}}
		{{end}}
		gt 大于
		{{if gt .var1 .var2}}
		{{end}}
		ge 大于等于
		{{if ge .var1 .var2}}
		{{end}}
	*/
}

func Test_Range0(t *testing.T) {
	tmpl, err := template.New("test").Parse(`
	{{- range .name}}
 	   {{.}}
	{{- end}}

	index value 方式：
	{{- range $idx, $value := .name}}
	   id: {{$idx}}, value: {{$value}}
	{{- end}}
	结构体内部字段	
	{{- range .Num}}
	   {{.N}}
	{{- end}}
`)
	if err != nil {
		panic(err)
	}
	type Numbers struct{ N int }
	err = tmpl.Execute(os.Stdout, map[string]interface{}{"name": []string{"tom", "jack", "angle"},
		"Num": []Numbers{{1}, {2}, {3}}})
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
	tmpl, err := template.New("test").Parse(`{{- range .name}}
		{{.}}
	{{- else}}
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
	/*
		However, to aid in formatting template source code, if an action's left delimiter (by default "{{") is followed immediately by a minus sign and ASCII space character ("{{- "), all trailing white space is trimmed from the immediately preceding text. Similarly, if the right delimiter ("}}") is preceded by a space and minus sign (" -}}"), all leading white space is trimmed from the immediately following text. In these trim markers, the ASCII space must be present; "{{-3}}" parses as an action containing the number -3.
	*/
}
