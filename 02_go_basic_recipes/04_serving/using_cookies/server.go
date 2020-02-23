package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", setHandler)
	http.HandleFunc("/read", readHandler)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)

}

func setHandler(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name: "my-cookie",
		Value: "testing",
	})
	fmt.Fprintf(w, "Cookie written. Check your browsers cache")
	fmt.Fprintln(w, "In Chrome go to dev tools / application / cookies")
}

func readHandler(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, "Your Cookie: ", c)

}
