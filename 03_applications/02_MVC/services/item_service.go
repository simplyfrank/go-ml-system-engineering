package services

import (
	"LearnGoProject/03_applications/02_MVC/domain"
	"LearnGoProject/03_applications/02_MVC/utils"
	"net/http"
)

type itemService struct {}

var (
	ItemService itemService
)

func (*itemService) GetItem(itemId string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "implement me",
		StatusCode: http.StatusInternalServerError,
		Code:       "service_not_implemented",
	}
}
func (*itemService) GetItems() (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "implement me",
		StatusCode: http.StatusInternalServerError,
		Code:       "service_not_implemented",
	}
}
