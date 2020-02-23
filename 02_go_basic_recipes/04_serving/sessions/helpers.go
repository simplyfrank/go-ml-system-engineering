package main

import "net/http"

// Deifne the handler functions
func getUser(w http.ResponseWriter, req *http.Request) user {
	// Check if a cookie is set
	c, err := req.Cookie("session")
	if err != nil {
		sID := "lakjlkejaer2e23423"
		c = &http.Cookie{
			Name: "session",
			Value: sID,
		}
	}
	http.SetCookie(w, c)

	// If the user exists, retrieve him
	var u user
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}
	return u
}

//func logoutHandler(w http.ResponseWriter, req *http.Request) {
//
//}

// Helper functions
func alreadyLoggedIn(req *http.Request)  bool {
	c, err := req.Cookie("seession")
	if err != nil {
		return false
	}
	un := dbSessions[c.Value]
	_, ok := dbUsers[un]
	return ok
}
