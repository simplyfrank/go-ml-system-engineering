package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"net/http"
	"net/url"
	"github.com/google/uuid"
)

var tmpData struct {
	active bool
	sessID string
	username string
	passwordHash []byte
}

// ----------------- HANDLERS ---------------------
func baseHandler(w http.ResponseWriter, req *http.Request) {

	us, ok := getUserFromSession(req)

	var dat interface{}
	if ok {
		dat = struct{
				User user
				Users map[string]user
				Req http.Request
			}{
			us,
			dbUsers,
			*req,
			}
		}
	tpl.ExecuteTemplate(w, "index.gohtml", dat)
}

func signupHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	switch req.Method {
	case "GET":
		if isLoggedIn(req) {
			// User already exists and is logged in. Redirect
			http.Redirect(w, req, "/", http.StatusSeeOther)
		}
		// User does not yet exist. Display login form
		w.Header().Set("Status", "200")
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		tpl.ExecuteTemplate(w, "signup.gohtml", nil)

	case "POST":
		log.Print("POST /signup")
		// Get information from form
		req.ParseForm()
		log.Printf("Parsed form output as: %v", req.PostForm)
		// Pass information explicitly to register function
		currUser, ok := registerUser(w, req.PostForm)
		log.Printf("User Registration returned with : %v\nstatus: %v", currUser, ok)
		if !ok {
			log.Print("User registration failed. Redirecting")
			//http.Redirect(w, req, "/signup", http.StatusSeeOther)
		}

		// Associate user with new session
		sID := logIn(w, currUser.UserName)
		log.Printf("Successfully registered a new session:\nSession ID:%v\nUser:%v", sID, currUser)
		http.Redirect(w, req, "/signup/success", 302)
	}
}

func loginResultHandler(w http.ResponseWriter, req *http.Request) {
	// Ensure only authenticated users can reach the page
	requireAuth(w, req)

	// Display the result of the registration metrics
	log.Printf("Preparing success page after login -> Getting user")
	currUser, ok := getUserFromSession(req)
	if !ok {
		log.Print("Unable to retrieve current user from session")
	}
	sID := getSession(req)
	if !ok {log.Print("Unable to get session from cookie")}

	_, _ = currUser, sID
	// Display the results on a sepearate success page for confirmation
	log.Print("Executing template success.gohtml")
	tpl.ExecuteTemplate(w, "success.gohtml", struct{
		session string
		user user
	}{
		sID.String(),
		currUser,
	})
}


// loginHandler manages the login of registered users to the system
func loginHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		// Check if already logged and if so redirect
		if isLoggedIn(req) {
			http.Redirect(w, req, "/", http.StatusSeeOther)
		}
		// Present the login form
		tpl.ExecuteTemplate(w, "login.gohtml", nil)
	case "POST":
		// Login in the user
		uName := req.FormValue("username")
		pWrd := req.FormValue("password")

		if isValidUser(uName, pWrd) {
			sessID := logIn(w, uName)
			io.WriteString(w, "You have been successfully logged in with session: "+sessID.String())
		}

		// if not validated
		http.Redirect(w, req, "/login", http.StatusUnauthorized)
	}
}

// logOut Handler manages the logout of a user from the system
func logoutHandler(w http.ResponseWriter, req *http.Request) {
	// Log out the user from the system
	if logOut(w, req) {
		io.WriteString(w, "You have been successfully logged out")
	} else {
		io.WriteString(w, "We were unable to log you out. Please try again")
	}
}

// --------------- Helper Functions ------------------------------

// logOut removes the session from the db and deletes
// the session cookie for the user.
func logOut(w http.ResponseWriter, req *http.Request) bool {
	// Check if logged in
	if isLoggedIn(req) {
		// remove session from dbSessions
		sessId := getSession(req)
		delete(dbSessions, sessId)
		// Remove the cookie from the user
		http.SetCookie(w, &http.Cookie{
			Name:       "session",
			Value:      "",
			MaxAge:     0,
		})
		return true
	}
	return false
}

// registerUser manages the safe addition of a new userAccount
// with the system
func registerUser(w http.ResponseWriter, formvals url.Values) (user, bool) {
	log.Printf("Start registering user, passing: %v", formvals)
	ut := user{
		formvals.Get("username"),
		[]byte(formvals.Get("password")),
		formvals.Get("firstname"),
		formvals.Get("lastname"),
	}
	// Register user with the system
	log.Printf("Registered users as: %v", ut)
	dbUsers[ut.UserName] = ut

	// Login the user for a new session
	log.Printf("Logging in user: %v", ut)
	sID := logIn(w, ut.UserName)
	log.Printf("Logged in user to session: %v", sID)
	return ut, true
}

// isLoggedIn checks if there is an active current session
// for a given user and returns a status
func isLoggedIn(req *http.Request) bool {
	sessID := getSession(req)
	// Check if the sessionID is stored in the database
	_, ok := dbSessions[sessID]
	return ok
}

// getSession retrieves the sessionID from the current request
func getSession(req *http.Request) uuid.UUID {
	c, err := req.Cookie("session")
	if err != nil {
		panic(err)
	}
	// Check if session can be parsed
	sID, err := uuid.Parse(c.Value)
	if err != nil {
		return uuid.Nil
	}
	return sID
}

// getUser attempts to retrieve the user stored with the active
// session and return it to the call. In case there is no active
// session, a session cookie will be set.
func getUserFromSession(req *http.Request) (user, bool){
	// check if there is a cookie
	log.Println("getUserFromSession:: getting cookie ")
	c, err := req.Cookie("session")
	if err != nil {
		log.Println("getUserFromSession:: unable to retrieve cookie")
		return user{}, false
	}

	// if the user exists, retrieve, else redirect
	var u user
	log.Printf("getUserFromSession:: retrieving user: currently set to %v", c.Value)
	sID, err := uuid.Parse(c.Value)
	if err != nil {
		// Mark the user as non-existent
		return u, false
	}
	if un, ok := dbSessions[sID]; ok {
		log.Printf("getUserFromSession:: user found. retrieving: %v", un)
		u = dbUsers[un]
		return u, true
	}
	return u, false
}

func isValidUser(username string, password string) bool {
	// Check if user is registered
	user, ok := dbUsers[username]
	if !ok {
		return false
	}

	// if registered check validity of password
	return validatePassword(user.Password, password)
}
// logIn generates a new session for a given username
// and stores it into the dbSessions map
func logIn(w http.ResponseWriter, username string) uuid.UUID {
	// Generate a new session id
	sID, err := uuid.NewUUID()
	if err != nil { log.Print("Unable to generate UUID")}
	// Check if the user is registered
	dbSessions[sID] = username

	// Set the cookie on the user
	http.SetCookie(w, &http.Cookie{
		Name:       "session",
		Value:      sID.String(),
	})
	log.Println("Successfully set session cookie")
	return sID
}

// encryptPW encryptes a given password and returns it as an
// encrypted bytestring for storage with the user type
func encryptPW(password string) ([]byte, bool) {
	encPW, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return nil, false
	}
	return encPW, true
}
func validatePassword(hash []byte, plain string) bool {
	if err := bcrypt.CompareHashAndPassword(hash, []byte(plain)); err != nil {
		return false
	}
	return true
}

func requireAuth(w http.ResponseWriter, req *http.Request) {
	// If user not authenticated, redirect to signup
	log.Println("requireAuth:: Checking authentication status")
	if !isLoggedIn(req) {
		log.Print("requireAuth:: Request blocked: Not Authenticated. Redirect to /login")
		http.Redirect(w, req, "/signup", 302)
	}
	log.Print("User was authenticated and was granted access")
}
