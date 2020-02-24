package logger

import (
	"LearnGoProject/03_applications/01_codeOrg/modules/config"
	logpk "log"
)
// ---- Logging Functionalities -------->
// log switches on the set config.LogLevel to ensure the right type of logging procedure
// is used
// TODO: Continue developing the individual logging rules
func log(logFunc func(), level config.LogLevel) {
	switch config.LoggingConfig.LogLevel(){
	case config.LogLevels.All:
		// execute the logging statement

	case config.LogLevels.Debug:
		if level <= config.LogLevels.Debug {
			// Accept all log levels up to Debug
			logFunc()
		}
	case config.LogLevels.Info:
		if level <= config.LogLevels.Info {
			// Accept all log levels up to Info
			logFunc()
		}
	case config.LogLevels.Warn:
		if level <= config.LogLevels.Warn {
			// Accept all log levels up to Warn
			logFunc()
		}
	case config.LogLevels.Error:
		if level <= config.LogLevels.Error {
			// Accept all log levels up to Error
			logFunc()
		}
	case config.LogLevels.Fatal:
		if level <= config.LogLevels.Fatal {
			// Accept all log levels up to Fatal
			logFunc()
		}
	case config.LogLevels.Off:
		if level <= config.LogLevels.Off {
			return
			// This doesn't allow any logging
		}
	case config.LogLevels.Race:
		if level <= config.LogLevels.Race{
			// Accept all log levels up to Race
			logFunc()
		}
	}
}

// ----------- CONVENIENCE LOGGERS --------->>>>
// Functions are high level abstractions to quickyl set up a meaningful logging system

// LogUnsetValue logs an event where a value that is required to have been set was found uninitiated.
func LogUnsetValue(value interface{}, fn interface{}, msg string ) {
	// define the logging function as an anonymous function to be passed
	logFn := func() {
		logpk.Printf("Value %v:: Within function %s found to be uninitialized:: %s", value, fn, msg)
	}
	// Pass the function along with its level of severity
	log(logFn, config.LogLevels.Warn)
}

// ----------- CUSTOM LOGGING ------------>>>>
// Following functions enable a customized logging behaviour at a low level. Internally
// used to implement convenience loggers in the package. Do not alter their API !!

func LogError(err error, level config.LogLevel, callFunc string, task string) {
	// define the logging function as an anonymous function to be passed
	logFn := func() {
		logpk.Printf("%s:: Unable to %s:: %s", callFunc, task, err)
	}
	// Pass the function along with its level of severity
	log(logFn, config.LogLevels.Error)
}

func LogFatal(err error, callFunc string, task string) {
	logFn := func() {
		logpk.Fatalf("%s:: Unable to %s:: %s", callFunc, task, err)
	}
	log(logFn, config.LogLevels.Fatal)
}

