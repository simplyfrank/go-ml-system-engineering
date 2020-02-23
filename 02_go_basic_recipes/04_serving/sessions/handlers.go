package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, req *http.Request) {
	//u := getUser(w, req)
	fmt.Print("Rendeing template index....")
	if err := tpl.ExecuteTemplate(w, "index", struct{}{}); err != nil {
		fmt.Print("Can not find template")
	}
}

func bar(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	if !alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}