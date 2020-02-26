package github

// Model the domain API request structure
type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	Private     bool   `json:"private"`
	HasIssues   bool   `json:"has_issues"`
	HasProjects bool   `json:"has_projects"`
	HasWiki     bool   `json:"has_wiki"`
}

// Models the response from the API we receive --> Test first and model from actual response wherever possible
type CreateRepoResponse struct {
	Id          int64           `json:"id"`
	Name        string          `json:"name"`
	FullName    string          `json:"full_name"`
	Owner       RepoOwner       `json:"owner"`
	Permissions RepoPermissions `json:"permissions"`
}

// Separate nested structs so we can change and compose them easily.
// DO NOT directly nest them into the response and request structs
type RepoOwner struct {
	Id int64 `json:"id"`
	Login string `json:"login"`
	Url string `json:url`
	HtmlUrl string `json:"html_url"`
}

// Same deal here. Keep it separate for ease of maintenance
type RepoPermissions struct {
	IsAdmin bool `json:"admin"`
	HasPull bool `json:"pull"`
	HasPush bool `json:"push"`
}
