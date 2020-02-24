/*
Package config
Status: In progress

Target:
Custom implementation of a logging system that can be centrally configured
to log only up to a given level of severity. It offers a number of convenience patterns
to log semantically meaningful situations for later analysis.

It wraps the log package and attempts to provide a more unified interface at the risk
of initially bloating the API.

Open issues known:
*/
package config

var LoggingConfig *loggingConfig


// Define the available log levels in the package
type LogLevel int
type list struct {
	Debug LogLevel
	Info LogLevel
	Warn LogLevel
	Error LogLevel
	Fatal LogLevel
	Off LogLevel
	Race LogLevel
	All LogLevel
}
// Enum for public use
var LogLevels = &list{
	All: 8,
	Debug: 7,
	Info: 6,
	Warn: 5,
	Error: 4,
	Fatal: 3,
	Race: 2,
	Off: 1,
}

type loggingConfig struct {
	logLevel LogLevel
	logString string
}

func init() {
	if LoggingConfig == nil {
		//logger.LogUnsetVariable("Logging", "*loggingConfig", "")
	}
}

// Instantiate and prepopulate a loggingConfig value
// to return to the caller.
func NewLogConfig() *loggingConfig {
	return &loggingConfig{
		logLevel:  LogLevels.Debug,
		logString: "%s:%s",
	}
}


// Provide the setter interface
func (l *loggingConfig) Configure(key string, value interface{}) (ok bool) {
	//switch key {
	//case "LogLevel", "loglevel", "level":
	//	l.SetLogLevel(fmt.Sprintf("%v", value))
	//	ok = true
	//case "Logstring", "LogString", "logstring":
	//	l.SetLogstring(fmt.Sprintf("%v", value))
	//	ok = true
	//default:
	//	// Inserted unknown key. Not setting, and flagging false action
	//	ok = false
	//}
	return true
}

// SetLogLevel sets the logging level for the whole project
func (l *loggingConfig) SetLogLevel(level LogLevel) {
	// Ensure only certain allowed level can be set
	// Set additional settings according with the logger level
	// Set the appropriate level
	l.logLevel = level
}
func (l *loggingConfig) SetLogString(formatstring string) {
	// Check the formatstring accepts the right number of elements
	// Set the format string
	l.logString = formatstring
}

// ---- GETTER Functions ------->
func (l *loggingConfig) LogLevel() LogLevel {
	return l.logLevel
}
func (l *loggingConfig) LogString(numArgs int) string {
	// TODO: Ensure the format string satisfies the right number of arguments
	// Return the logstring
	return l.logString
}
