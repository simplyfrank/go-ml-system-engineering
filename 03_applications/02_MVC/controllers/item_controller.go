package controllers

import (
	"LearnGoProject/03_applications/02_MVC/services"
	"encoding/json"
	"net/http"
)
type itemController struct {}

var (
	ItemController itemController
)


func (*itemController) GetItem(w http.ResponseWriter, req *http.Request) {
	//// Check request
	//itemId, err := strconv.ParseUint(req.URL.Query().Get("item_id"), 10, 64)
	//if err != nil {
	//	apiErr := &utils.ApplicationError{
	//		Message:    "item_id must be a number",
	//		StatusCode: http.StatusBadRequest,
	//		Code:       "bad_request",
	//	}
	//
	//	jsonValue, _ := json.Marshal(apiErr)
	//	w.WriteHeader(apiErr.StatusCode)
	//	w.Write(jsonValue)
	//	return
	//}
	itemId := req.URL.Query().Get("item_id")
	item, apiErr := services.ItemService.GetItem(itemId)
	if apiErr != nil {
		// Serialize and send error
		jsonValue, _ := json.Marshal(apiErr)
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonValue)
		return
	}
	// Serialize and return user
	jsonValue, _ := json.Marshal(item)
	w.Write(jsonValue)
}

// Retrieve the full list of registered users
func (*itemController) GetItems(w http.ResponseWriter, req *http.Request) {
	items, apiErr := services.ItemService.GetItems()
	if apiErr != nil {
		// serialize and send error
		jsonValue, _ := json.Marshal(apiErr)
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonValue)
		return
	}

	// Serialize and send items
	jsonValue, _ := json.Marshal(items)
	w.Write(jsonValue)
}
