package logger

import "net/http"

// Logging functionality to handle web related event
/*
- Enable smart extraction of information directly from the request object
- Enable easy configuration of logging behavior on request data from outside the code
- Enable optimized log exports
*/

// The WebLogger struct enables to set configuration for a given
// logging function. Each function can ensure the logger
type WebLogger struct {
	// Define all relevant metrics to log
	// Some values can be flags
	LogIP bool // Extract and log the ip of the requestor
	LogLoc bool // Extract and log the location of the request
	LogCookie bool // Log the state of all registered cookies
	LogMethod bool // Log the method of the request
	LogEncrypted bool // Encrypt information before writing to log
}
type FileLogger struct {
	// TODO: Fill in attributes
}

type DBLogger struct {

}

type APILogger struct {

}

func LogRequest(req *http.Request, *WebLogger) {

}