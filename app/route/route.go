package route

import (
	"EasySurvey/app/route/router"
	"net/http"
	"EasySurvey/app/route/cors"
)

// LoadHTTP loads the HTTP routes
func LoadHTTP() http.Handler {
	var instance = router.Instance()

	return cors.Handler(instance) 
}
