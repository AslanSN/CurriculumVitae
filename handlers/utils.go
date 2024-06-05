package handlers

import (
	"html/template"
	"io"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string,
	data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplate() *Templates {
	println(template.ParseGlob("../views/index.html"))
	return &Templates{
		templates: template.Must(template.ParseGlob("../views/*.html")),
	}
}

func show(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, data)
}
