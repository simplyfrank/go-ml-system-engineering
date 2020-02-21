package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", baseHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func baseHandler(w http.ResponseWriter, req *http.Request) {

	fmt.Println(req.Method)
	switch req.Method {
		// Requesting the page for the first time will send a GET
		// request, which we can now handle here in a switch statement
		// on the req.Method value which is passed down from the server
		// to our function
		case "GET":
			fmt.Println("Checking GET")
			// Prepare the reponse
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			io.WriteString(w, `
				<form method="POST" enctype="multipart/form-data">
					<input type="file" name="fileupload">
					<input type="submit">
				</form>
				</br>
			`)

		// Submitting from the form will send a POST request
		// to the server sending the file as a named variable, which we
		// can extract from the request object easily.
		case "POST":
			fmt.Println("Checking POST")
			// open the file
			f, _, err := req.FormFile("fileupload")
			if err != nil { http.Error(w, err.Error(), http.StatusInternalServerError)}
			defer f.Close()

			// read the file
			bs, err := ioutil.ReadAll(f)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			// For confirmation we are just writing the file back to the caller
			io.WriteString(w, string(bs))
		}
}
