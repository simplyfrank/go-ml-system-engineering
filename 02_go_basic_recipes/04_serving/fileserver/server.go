package main

import (
	"log"
	"net/http"
)

func main() {
	// Instantiate the Handler with a http.Fileserver that expects
	// a RootFilesystem type to be passed, for which we use the http.Dir
	// value and point it at our assets folder
	http.Handle("/", http.FileServer(http.Dir("./assets")))

	// Serving it with the base handler, we are now serving the files
	// stored in the assets folder directly to the user.
	// When we store an index.html file in the root folder, the server
	// will automatically serve only this file, masking access to all others
	// --> Otherwise, the folder will be visibly listed, with http links to
	// all files in that folder
	log.Fatal(http.ListenAndServe(":8080", nil))
}