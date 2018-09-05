package survey

import (
	"github.com/lavpthak/easysurvey/app/model/survey"
	"net/http"
	"github.com/labstack/echo"
)

// GetAll returns all the surveys
func GetAll(c echo.Context) error {
	surveys := survey.GetAll()
	return c.JSON(http.StatusOK, surveys)
}
