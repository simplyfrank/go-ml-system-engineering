package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", defaultFunc)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func defaultFunc(w http.ResponseWriter, req *http.Request) {
	// Try to get the value from the url
	v := req.FormValue("q")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="post">
		<input type="text" name="q">
		<input type="submit">
	</form>
	</br>
	` + "This was send last time: " +v)

}