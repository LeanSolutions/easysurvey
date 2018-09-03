package controller

import (
	"github.com/lavpthak/easysurvey/app/model/survey"
	"github.com/lavpthak/easysurvey/app/route/router"
	"github.com/lavpthak/easysurvey/app/shared/response"
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
