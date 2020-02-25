package domain

import (
	"LearnGoProject/03_applications/02_MVC/utils"
	"net/http"
)

// Mock user database 
var (
	users = map[uint64] *User{
		123: &User{
			UserID:    123,
			FirstName: "Leonardo",
			LastName:  "DiCaprio",
			Email:     "leonard@cap.io",
		},
		1234: &User{
			1234,
			"Another",
			"Name",
			"another@user.com",
		},
	}
)

func GetUser(userId uint64) (*User, *utils.ApplicationError) {
	if user, ok := users[userId]; ok {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    "user not found in system",
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
func GetUsers() ([]*User, *utils.ApplicationError) {
	var usrList []*User
	for _, user := range users {
		usrList = append(usrList, user)
	}
	if usrList != nil {
		return usrList, nil
	}

	return nil, &utils.ApplicationError{
		Message:    "no registered users in the system",
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
