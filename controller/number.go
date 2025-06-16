package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Number(c echo.Context) error {
	if c.Request().Method == http.MethodGet {
		return c.Render(http.StatusOK, "number.tmpl", map[string]interface{}{
			"title": "Number",
		})
	}
	// POST process (to be implemented)
	return nil
}
