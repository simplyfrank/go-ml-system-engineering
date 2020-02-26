package github_provider

import (
	"LearnGoProject/03_applications/03_domainServices/api/clients/restclient"
	"LearnGoProject/03_applications/03_domainServices/api/domain/github"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	headerAuthorization       = "Authorization"
	headerAuthorizationFormat = "token %s"

	urlCreateRepo = "https://api.gitub.com/user/repos"
)

func getAuthorizationHeader(accessToken string) http.Header {
	h := http.Header{}
	h.Add(headerAuthorization, fmt.Sprintf(headerAuthorizationFormat, accessToken))
	return h
}

func CreateRepo(accessToken string, req github.CreateRepoRequest) (*github.CreateRepoResponse, *github.ErrorResponse) {
	// Create request header
	header := getAuthorizationHeader(accessToken)

	// Send request to API
	response, err := restclient.Post(urlCreateRepo, req, header)
	if err != nil {
		return nil, &github.ErrorResponse{
			StatusCode:       http.StatusInternalServerError,
			Message:          err.Error(),
		}
	}
	// Read response from request body
	bytes, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		return nil, &github.ErrorResponse{
			StatusCode: http.StatusInternalServerError, Message: err.Error(),}
	}

	// Check if error is present
	if response.StatusCode > 299 {
		// We have an error encountered
		var errResponse github.ErrorResponse
		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &github.ErrorResponse{
				StatusCode: http.StatusInternalServerError,	Message: "invalid crateRepoError response body",
			}
		}
		// Return errorResponse object
		return nil, &errResponse
	}

	// Create CreateRepoResponse object from request body
	var createResponse github.CreateRepoResponse
	if err := json.Unmarshal(bytes, &createResponse); err != nil {
		log.Printf("unable to parse createResponse from response body")
		return nil, &github.ErrorResponse{StatusCode: http.StatusInternalServerError, Message: "invalid createRepo response body"}
	}
	// Return success
	return &createResponse, nil

}
