package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index.tmpl", map[string]interface{}{
		"title": "Users",
	})
}