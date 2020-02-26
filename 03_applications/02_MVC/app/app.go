package app

import (
	"LearnGoProject/03_applications/02_MVC/controllers"
	"net/http"
)

func StartApp() {
	// Set the handlers
	http.HandleFunc("/user", controllers.UserController.GetUser)
	http.HandleFunc("/users", controllers.UserController.GetUsers)

	http.HandleFunc("/item", controllers.ItemController.GetItem)
	http.HandleFunc("/items", controllers.ItemController.GetItems)
	// Start the server
	http.ListenAndServe(":8080", nil)
}
