package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func ImpressumHandler(c echo.Context) error {

	return c.Render(http.StatusOK, "impressum.html", map[string]interface{}{
		"name": "Impressum",
		"msg":  "All about Boatswain!",
	})
}
