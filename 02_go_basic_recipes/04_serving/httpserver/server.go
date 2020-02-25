package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// To define the Handler functions to use with the http.HandlerFunc method
// we have to accept a http.ResponseWriter and a request object
func responseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	Host: %s
	Methos: %s
	RemoteAddr: %s
	Response: %v
	URL: %s
`, r.Host, r.Method, r.RemoteAddr, r.Response, r.URL)
}

func nameHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `
	<h1>Welcome to the show</h1>
	<p>This page is a welcome page for you alone {{.}}</p>
`)
}

func main() {
	// We are defining the routes we want to publish
	// Each route gets assigned a handler function
	http.HandleFunc("/response/", responseHandler )
	http.HandleFunc("/:name/", nameHandler)

	// WE can use the shorthand way to set a server
	//logger.Fatal(http.ListenAndServe(":8080", nil))

	// Or create a custom server value
	s := &http.Server{
		Addr:              ":8080",
		Handler:           nil,
		TLSConfig:         nil,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 0,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       0,
		MaxHeaderBytes:    1 << 20,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
	// Now we are calling the Listen and Serve loop directly on our
	// custom server instance
	log.Fatal(s.ListenAndServe())
}
