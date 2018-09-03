package controller

import (
	"github.com/lavpthak/easysurvey/app/shared/response"
	"net/http"
	"github.com/lavpthak/easysurvey/app/route/router"
)

func init() {
	router.Get("/", Index)
}

// Index is the root url
func Index(w http.ResponseWriter, r *http.Request) {
	response.Send(w, http.StatusOK, "Welcome to Easy Survey", 0, nil)
}
