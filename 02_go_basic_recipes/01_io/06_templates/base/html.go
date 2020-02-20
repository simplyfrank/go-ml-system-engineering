package base

import (
	"os"
	"strings"
	"text/template"
)


const sampleTemplate = `
	This template demonstrates printing a {{ .Variable | printf "%#v""}}  
	{{if condition}}
	If condition is set, we'll print this
	{{else}}
	We will print this instead if condition is not set

	Next we'll iterate over an array of strings:
	{{ range $index, $item := split .Words ","}}
	{{$index}}: {{$item}}
	{{end}}

	Blocks are a way to embed templates ino one another
 	{{block "block_example" .}}
	No block defined!
	{{end}}


	{{/*
	This is a way to insert
	a multi line comment in a template
	*/}}
`

const secondTemplate = `
	{{ define "block_example" }}
	{{.OtherVariable}}
	{{end}}
`


func RunTemplate() error {
	data := struct{
		Condition bool
		Variable string
		Items []string
		Words string
		OtherVariable string
	}{
		Condition: true,
		Variable: "variable",
		Items: []string{"item1", "item2", "item3","item4"},
		Words: "another_item1,another_item2,another_item3,another_item4",
		OtherVariable: "I'am defined in a second template that is nested here.",
	}

	funcmap := template.FuncMap{
		"split": strings.Split,
	}

	// these can also be chained
	t := template.New("example")
	t = t.Funcs(funcmap)

	// We coudl use Must instead to panic on error
	// template.Must(t.Parse
	template.Must(t.Parse(sampleTemplate))

	// Write the template to stdout and populate it with data
	if err := t.Execute(os.Stdout, &data); err != nil {
		return err
	}

	return nil
}

