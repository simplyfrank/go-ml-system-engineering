package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", baseFunc)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// We are now logging all requests done to localhost:8080/q?....
// and print them back out to the caller
func baseFunc(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	fmt.Fprintf(w, `
	The posted request on q reads: %s
`, v)
}

// Go and visit: http://localhost:8080/?q=TestingTheWorld
// Change up the value of q, and try other combinations of variable
// names and assigned values