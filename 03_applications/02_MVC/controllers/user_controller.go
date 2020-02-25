package controllers

import (
	"LearnGoProject/03_applications/02_MVC/services"
	"LearnGoProject/03_applications/02_MVC/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

// GetUser retrieves a single user by userId
func GetUser(w http.ResponseWriter, req *http.Request) {
	// Check request
	userId, err := strconv.ParseUint(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		jsonValue, _ := json.Marshal(apiErr)
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonValue)
		return
	}

	user, apiErr := services.GetUser(userId)
	if apiErr != nil {
		// Serialize and send error
		jsonValue, _ := json.Marshal(apiErr)
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonValue)
		return
	}
	// Serialize and return user
	jsonValue, _ := json.Marshal(user)
	w.Write(jsonValue)
}

// Retrieve the full list of registered users
func GetUsers(w http.ResponseWriter, req *http.Request) {
	users, apiErr := services.GetUsers()
	if apiErr != nil {
		// serialize and send error
		jsonValue, _ := json.Marshal(apiErr)
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonValue)
		return
	}

	// Serialize and send users
	jsonValue, _ := json.Marshal(users)
	w.Write(jsonValue)
}
