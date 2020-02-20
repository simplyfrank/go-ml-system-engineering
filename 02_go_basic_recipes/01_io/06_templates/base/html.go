package base

import (
	"os"
	"strings"
	"text/template"
)


const sampleTemplate = `
	{{ .Variable | printf "%#v" }}

	{{ if .Condition }}
	If condition is set this prints
	{{ else }}
	If condition is not set this prints
	{{ end }}
	
	This iterates over an array of items:
	{{ range $idx, $val := .Items }}
	{{ $idx }}: {{$val}}
	{{ end }}

	We can easily import other function like string.Split
	{{ range $idx, $item := split .Words "," }}
	{{$idx}}: {{$item}}
	{{end}}

	Blocks are here for code reuse
	{{block "block_example" .}}
	No Block defined
	{{end}}
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
	t, err:= t.Parse(sampleTemplate)
	if err != nil { return err }

	t2, err := t.Clone()
	if err != nil { return err}

	t2, err = t2.Parse(secondTemplate)
	if err != nil { return err}

	// Write the template to stdout and populate it with data
	if err := t2.Execute(os.Stdout, &data); err != nil {
		return err
	}

	return nil
}

