/*
Check provides functionality for error handling and checking
It integrates with SkillSmart/go/logger for handling the logging of error incidents.

Its behaviour can be customized through the SkillSmart/go/config package at project level.

Status: Under development
*/

package check

import (
	"LearnGoProject/03_applications/01_codeOrg/modules/config"
	"LearnGoProject/03_applications/01_codeOrg/modules/logger"
	"database/sql"
)

// ------- Handle ERROR CHECKS ------------->
func LogNotNil(err error, level config.LogLevel, callFunc string, task string) {
	if err != nil {
		logger.LogError(err, level, callFunc, task)
	}
}
func RaiseNotNil(err error, callFunc string, task string) {
	if err != nil {
		// Set logging level to fatal
		logger.LogFatal(err, callFunc, task)
	}
}
func HandleNotNil(err error) {

}

// --------- Handle connection checks ----------->
func DBConnect(db *sql.DB, err error) *sql.DB {
	callFunc, task := "db.NewDB", "establish DB connection"
	logger.LogError(err, config.LogLevels.Error, callFunc, task)
	return db
}


// --------- Handle Value expectations ---------->

// Ensures that value gv is of an expected type et.
// Failing the check results in logging and a selected response type.

