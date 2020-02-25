package services

import (
	"LearnGoProject/03_applications/02_MVC/domain"
	"LearnGoProject/03_applications/02_MVC/utils"
)


// Function fulfills request to a user by ID
func GetUser(userId uint64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}

func GetUsers() ([]*domain.User, *utils.ApplicationError) {
	return domain.GetUsers()
}