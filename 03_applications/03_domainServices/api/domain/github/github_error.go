package github

type Error struct {
	// define the struct
	Resource string `json:"resource"`
	Code     string `json:"code"`
	Field    string `json:"field"`
	Message  string `json:"message"`
}
type ErrorResponse struct {
	StatusCode       int     `json:"status_code"`
	Message          string  `json:"message"`
	DocumentationUrl string  `json:"documentation_url"`
	Errors           []Error `json:"errors"`
}

func (*Error) Error() string {
	//	Implement error interface
	return "Error interface not yet implemented"
}
