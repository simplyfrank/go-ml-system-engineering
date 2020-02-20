package html

import (
	"fmt"
	"html/template"
	"os"
)

// HTMLDifferences highlights some of the differences
// between html/template and text/template
func HTMLDifferences() error {
	t := template.New("html")
	t, err := t.Parse("<h1>Hello! {{.Name}}</h1>n")
	if err != nil { return err }

	// html autoescapes unsafe operations like javascript
	// injections, is contextually aware and will behave differently
	// based on where a variable is rendered
	err = t.Execute(os.Stdout, map[string]string{
		"Name":"<script>alert('Can you see me?')</script>"})
	if err != nil { return err }

	// you can also manually call the escapers
	var input string = "example <example@example.com>"
	fmt.Printf("\n\nManually Escaping inputs '%s':\n\n", input)
	fmt.Println("JSEscaper: \t\t\t",template.JSEscaper(input))
	fmt.Println("HTMLEscaper: \t\t",template.HTMLEscaper(input))
	fmt.Println("URLQueryEscaper: \t",template.URLQueryEscaper(input))

	return nil
}

