package config

var Error *errorConfig

type errorConfig struct {
}

// Provide a general function to set values
func (e *errorConfig) Configure(key string, value interface{}) (ok bool) {
	// Switch allows to route to the appropriate function to handle
	// the pre-set operations
	return true
}

func init() {
	// Instantiate the necessary variables for the module here
	// Set default error handling configuration
}





