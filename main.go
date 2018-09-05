package main

import (
	"github.com/lavpthak/easysurvey/app/routes"
	"github.com/lavpthak/easysurvey/app/routes/survey"
	"github.com/labstack/echo"
)


func main() {
	e := echo.New()

	e.GET("/", routes.Index)
	e.GET("/survey/all", survey.GetAll)

	e.Logger.Fatal(e.Start(":8081"))
}
