package services

import (
	"LearnGoProject/03_applications/02_MVC/domain"
	"LearnGoProject/03_applications/02_MVC/utils"
)


type userService struct {

}

var (
	UserService userService
)
// Function fulfills request to a user by ID
func(*userService) GetUser(userId uint64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}

func (*userService) GetUsers() ([]*domain.User, *utils.ApplicationError) {
	return domain.GetUsers()
}
