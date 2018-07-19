package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
)

// Template ...
type Template struct {
	templates *template.Template
}

// Render ...
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// Sample Handler ...
func Sample(c echo.Context) error {
	return c.Render(http.StatusOK, "sample", "Sampleでぇす!!!!")
}

func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	e := echo.New()
	e.Renderer = t
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/sample", Sample)

	e.Logger.Fatal(e.Start(":8080"))
}
