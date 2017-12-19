package route

import (
	"EasySurvey/app/route/router"
	"net/http"
)

// LoadHTTP loads the HTTP routes
func LoadHTTP() http.Handler {
	return router.Instance()
}
