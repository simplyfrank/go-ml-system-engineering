package base

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

// Create TEamplate will create a template file that contains data
func CreateTemplate(path string, data string) error {
	return ioutil.WriteFile(path, []byte(data), os.FileMode(0755))
}

// Init Templates sets up templates from a directory
func InitTemplates() error {
	tempdir, err := ioutil.TempDir("", "temp")
	if err != nil { return err }
	defer os.RemoveAll(tempdir)

	err = CreateTemplate(filepath.Join(tempdir, "t1.tmpl"),
		`Template 1! {{ .Var1 }}
	{{ block "template2" .}}  {{ end }}
	{{ block "template3" .}} {{ end }}
	`)

	if err != nil {
		return err
	}

	err = CreateTemplate(filepath.Join(tempdir, "t2.tmpl"),
		`
	{{ define "templat2"}}Template 2! {{ .Var2 }}{{end}}
	`)
	if err != nil {
		return err
	}

	err = CreateTemplate(filepath.Join(tempdir, "t3.tmpl"),
		`
	{{ define "template3"}}Template 3! {{ .Var3 }}{{end}}
	`)
	if err != nil {
		return err
	}

	pattern := filepath.Join(tempdir, "*.tmpl")

	// Parse glob will combine all the files that match
	// glob and combine them into a single template
	tmpl, err := template.ParseGlob(pattern)
	if err != nil { return err }

	// Execute can also work with a map instead of a struct
	tmpl.Execute(os.Stdout, map[string]string{
		"Var1":"Variable 1 printed here...",
		"Var2":"Variable 2 is printed here",
		"Var3":"Variable 3 is printed here",
	})

	return nil
}