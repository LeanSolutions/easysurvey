package controller

import (
	"EasySurvey/app/route/router"
	"net/http"
)

func init() {
	router.get("/surveys", SurveyGetAll)
}

func SurveyGetAll(w *http.ResponseWriter, r *http.Request) {

}
