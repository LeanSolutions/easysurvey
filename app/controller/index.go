package controller

import (
	"EasySurvey/app/route/router"
	"EasySurvey/app/shared/response"
	"net/http"
)

func init() {
	// Main page
	router.Get("/", Index)
}

// Index is the root url
func Index(w http.ResponseWriter, r *http.Request) {
	response.Send(w, http.StatusOK, "Welcome to Easy Survey", 0, nil)
}
