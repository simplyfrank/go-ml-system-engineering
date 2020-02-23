package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var tpl *template.Template
var pd pageData
var err error

type pageData struct {
	Title 		string
	CompanyName string
}

func main() {
	// This executes on the start of the package before main
	// and allos us to set package scoped variables
	tpl, err = template.ParseGlob("./templates/*")
	if err != nil { log.Println("Unable to load templates: ", err); os.Exit(-1)}

	pd = pageData{
		Title:       "SkillSmart Evolutions",
		CompanyName: "ABetterStartup Ltd",
	}

	// Set the handler functions
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)

	// Start the server
	fmt.Println("Starting server. Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	// Handle the presentation of the homepage
	w.WriteHeader(200)
	fmt.Println("Generating Home page")
	tpl.ExecuteTemplate(w, "index.gohtml", pd)
}

func aboutHandler(w http.ResponseWriter, req *http.Request) {
	// Handle the presentation fo the about page
	w.WriteHeader(200)
	fmt.Println("Generating About page")
	tpl.ExecuteTemplate(w, "about.gohtml", pd)
}