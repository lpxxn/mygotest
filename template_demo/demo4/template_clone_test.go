package demo4

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"
	"testing"
)

func TestBlock1(t *testing.T) {
	const (
		master  = `Names:{{block "list" .}}{{"\n"}}{{range .}}{{println "-" .}}{{end}}{{end}}`
		overlay = `{{define "list"}} {{join . ", "}}{{end}} `
	)
	var (
		funcs     = template.FuncMap{"join": strings.Join}
		guardians = []string{"Gamora", "Groot", "Nebula", "Rocket", "Star-Lord"}
	)
	masterTmpl, err := template.New("master").Funcs(funcs).Parse(master)
	if err != nil {
		log.Fatal(err)
	}

	overlayTmpl, err := template.Must(masterTmpl.Clone()).Parse(overlay)
	if err != nil {
		log.Fatal(err)

	}

	if err := masterTmpl.Execute(os.Stdout, guardians); err != nil {
		log.Fatal(err)
	}

	if err := overlayTmpl.Execute(os.Stdout, guardians); err != nil {
		log.Fatal(err)
	}
}

func TestBlock2(t *testing.T) {
	const (
		overlay = `{{define "list"}} {{join . ", "}}{{end}}`
		//overlay = `Names:{{block "list" .}}{{"\n"}}{{range .}}{{println "-" .}}{{end}}{{end}}`
	)
	var (
		funcs     = template.FuncMap{"join": strings.Join}
		guardians = []string{"Gamora", "Groot", "Nebula", "Rocket", "Star-Lord"}
	)
	masterTmpl, err := template.New("test").Funcs(funcs).Parse(overlay)
	if err != nil {
		log.Fatal(err)
	}
	if err := masterTmpl.Execute(os.Stdout, guardians); err != nil {
		log.Fatal(err)
	}
}

/*

Clone can be used to prepare common templates and use them with variant definitions for other templates by adding the variants after the clone is made.

https://stackoverflow.com/questions/50842389/parsing-multiple-templates-in-go
*/
func TestTemplateClone1(t *testing.T) {
	var base = template.Must(template.New("base").Parse("header\n{{template \"content\"}}\nfooter"))
	var content1 = template.Must(template.Must(base.Clone()).Parse(`{{define "content"}}hello world{{end}}`))
	var content2 = template.Must(template.Must(base.Clone()).Parse(`{{define "content"}}foo bar{{end}}`))

	if err := content1.Execute(os.Stdout, nil); err != nil {
		panic(err)
	}
	fmt.Println()
	fmt.Println()
	if err := content2.Execute(os.Stdout, nil); err != nil {
		panic(err)
	}
	fmt.Println()

}
