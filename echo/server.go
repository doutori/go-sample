package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
)

// Template of Structure
type Template struct {
	templates *template.Template
}

// Member of Structure
type Member struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Render ...
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	e := echo.New()
	e.Static("/assets", "assets")
	e.Renderer = t

	// Routing
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", nil)
	})

	e.GET("/members", func(c echo.Context) error {
		members := []Member{
			{ID: 1, Name: "name1"},
			{ID: 2, Name: "name2"},
		}

		return c.JSON(http.StatusOK, members)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
