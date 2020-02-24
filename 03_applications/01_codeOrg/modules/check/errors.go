package check

import (
	"LearnGoProject/03_applications/01_codeOrg/modules/config"
	"LearnGoProject/03_applications/01_codeOrg/modules/logger"
)

// --- Error Handling ----------->
func Check(err error, level config.LogLevel, callFunc string, task string) {
	// Set defaults (and try to extract from metadata)
	// Run the check
	if err != nil {
		// Log the error according to set logger level
		logger.LogError(err, level, callFunc, task)
	}
}
