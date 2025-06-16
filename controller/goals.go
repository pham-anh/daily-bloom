package controller

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

func NewTemplateRenderer(pattern string) *TemplateRenderer {
	return &TemplateRenderer{
		templates: template.Must(template.ParseGlob(pattern)),
	}
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func GoalsAdd(c echo.Context) error {
	return c.Render(http.StatusOK, "goals_add", nil)
}
