package renderers

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

// Template is used as template renderer
type Template struct {
	templates *template.Template
}

// Render executes a template to be used when serving HTML pages
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// Register registers a template renderer service to be used in echo server
func Register(server *echo.Echo) {
	// // All templates must be parsed first
	// tmpl := template.Must(template.ParseGlob("files/pages/templates/*.html"))

	// The last to be parsed has to be the files not being used as templates anywhere
	tmpl := template.Must(template.ParseGlob("files/pages/*.html"))

	t := &Template{
		templates: tmpl,
	}

	server.Renderer = t
}
