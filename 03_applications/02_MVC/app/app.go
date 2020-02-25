package app

import (
	"LearnGoProject/03_applications/02_MVC/controllers"
	"net/http"
)

func StartApp() {
	// Set the handlers
	http.HandleFunc("/user", controllers.GetUser)
	http.HandleFunc("/users", controllers.GetUsers)
	// Start the server
	http.ListenAndServe(":8080", nil)
}
