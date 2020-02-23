package main

import (
	"fmt"
	"html/template"
	"net/http"
)


// Define the necessary types
type user struct {
	UserName string
	Password string
	First string
	Last string
}
//
//// manage template setup
//var tpl *template.Template
//// Define generally used types
var dbUsers map[string] user
var dbSessions map[string] string  // Maps username to session status
//
//func init() {
//
//	template.Must(tpl.ParseGlob("./templates/*.gohtml"))
//	//template.Must(tpl.ParseGlob("./templates/*.gohtml"))
//}
var tpl *template.Template
func main() {
	var err error
	tpl, err = template.ParseGlob("./templates/*")
	if err != nil {
		fmt.Println("got an error parsing the templates")
	}

	// Create the handlers
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/signup", signupHandler)
	//http.HandleFunc("/login", loginHandler)
	//http.HandleFunc("/logout", logoutHandler)

	// Instantiate the server
	http.ListenAndServe(":8080", nil)
}

