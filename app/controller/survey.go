package controller

import (
	"EasySurvey/app/model/survey"
	"EasySurvey/app/route/router"
	"EasySurvey/app/shared/response"
	"net/http"
)

func init() {
	router.Get("/surveys", SurveyGetAll)
}

// SurveyGetAll returns all the surveys
func SurveyGetAll(w http.ResponseWriter, r *http.Request) {
	surveys := survey.GetAll()
	response.Send(w, http.StatusOK, "Gets all surveys", len(surveys), surveys)
}
