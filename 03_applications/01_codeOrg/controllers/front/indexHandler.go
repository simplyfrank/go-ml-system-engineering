package front

import (
	"LearnGoProject/03_applications/01_codeOrg/controllers"
	"LearnGoProject/03_applications/01_codeOrg/modules/logger"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// Read from context

	// Do the switch
	switch r.Method {
	case "GET":
		logger.LogRequest(r)
		controllers.Tpl.Execute(w, "index.gohtml")
	default:
		http.Error(w, "Page not found.", 404)
	}
}