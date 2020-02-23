package main

import (
	"fmt"
	"github.com/google/uuid"
	"html/template"
	"net/http"
)

type user struct {
	UserName string
	Password []byte
	First string
	Last string
}

var tpl *template.Template
var dbUsers = map[string]user{}
var dbSessions = map[uuid.UUID]string{}


func init() {
	tpl = template.Must(template.ParseGlob("./templates/*"))
}

func main() {
	// Set the handler routes
	http.HandleFunc("/", baseHandler)
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	// Set up the server
	fmt.Println("Starting the webserver:\n\nListening on Port localhost:8080")
	http.ListenAndServe(":8080", nil)
}

