package handlers

import (
	"html/template"
	"io"
	"net/http"

	"github.com/AslanSN/CurriculumVitae/db"
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

func showInfoDBConstructor(info interface{}, c echo.Context) error {
	result := db.DB.Find(info)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": result.Error.Error(),
		})
	}
	return c.JSON(http.StatusOK, info)
}

/*
!	UNUSED
func (h HomeHandler) CreateHTMLFiles() {
	rootPath := "public"
	name := path.Join(rootPath, "index.html")
	f, err := os.Create(name)
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}

	err = views.Index().Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to write index page: %v", err)
	}
}
*/
