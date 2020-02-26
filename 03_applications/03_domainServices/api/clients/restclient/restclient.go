package restclient

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Post takes performs a rest call with method "POST" to the url, accepting a body and headers
func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {
	jsonBytes, err := json.Marshal(body); if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	request.Header = headers

	client := http.Client{}
	return client.Do(request)
}