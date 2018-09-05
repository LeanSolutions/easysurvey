package routes

import (
	"github.com/labstack/echo"
	"net/http"
)


// Index is the root url
func Index(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to Easy Survey!")
}
