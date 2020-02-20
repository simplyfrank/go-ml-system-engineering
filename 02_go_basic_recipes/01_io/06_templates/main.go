package main

import (
	basetempl "github.com/skillsmart/go/02_basic_recipes/06_templates/base"
	htmltempl "github.com/skillsmart/go/02_basic_recipes/06_templates/html"
)
func main() {

	// Populate a single template with a set
	// dictionary of variables
	if err := basetempl.RunTemplate(); err != nil {
		panic(err)
	}

	// Generate templates from common folder and
	// display in a combined template
	if err := basetempl.InitTemplates(); err != nil {
		panic(err)
	}

	// Generate the html template and display
	if err := htmltempl.HTMLDifferences(); err != nil {
		panic(err)
	}

}