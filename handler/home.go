package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

//HomeHandler can be called as a handler
func HomeHandler(c echo.Context) error {

	return c.Render(http.StatusOK, "impressum.html", map[string]interface{}{
		"name": "Home",
		"msg":  "this is the Home of the Website!",
	})
}
