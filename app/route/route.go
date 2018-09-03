package route

import (
	"github.com/lavpthak/easysurvey/app/route/router"
	"net/http"
	"github.com/lavpthak/easysurvey/app/route/cors"
)

// LoadHTTP loads the HTTP routes
func LoadHTTP() http.Handler {
	var instance = router.Instance()

	return cors.Handler(instance) 
}
