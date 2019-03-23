package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

//ImpressumHandler can be called as a handler
func ImpressumHandler(c echo.Context) error {

	return c.Render(http.StatusOK, "pagetemplate.html", map[string]interface{}{
		"name": "Impressum",
		"msg":  "All about this site!",
	})
}
