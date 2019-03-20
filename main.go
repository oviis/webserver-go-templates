package main

import (
	"errors"
	"io"
	"net/http"
	"text/template"

	// High minimalist GO web framework https://echo.labstack.com/
	"github.com/labstack/echo"

	"github.com/oviis/webserver-go-templates/handler"
)

//TemplateRegistry Define the template registry struct
type TemplateRegistry struct {
	templates map[string]*template.Template
}

//Render Implement e.Renderer interface
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}
	return tmpl.ExecuteTemplate(w, "base.html", data)
}

func main() {
	e := echo.New()
	e.GET("/hello-world", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// curl http://localhost:1323/json
	e.GET("/json", func(c echo.Context) error {
		return c.JSONBlob(
			http.StatusOK,
			[]byte(`{ "id": "1", "msg": "Hello Ovi!" }`),
		)
	})

	e.GET("/", func(c echo.Context) error {
		return c.HTML(
			http.StatusOK,
			"<h1>Hello, Ovi, have a good day!<h1><br><strong>Hello, World!</strong>",
		)
	})

	// Instantiate a template registry with an array of template set
	// Ref: https://gist.github.com/rand99/808e6e9702c00ce64803d94abff65678
	templates := make(map[string]*template.Template)
	templates["home.html"] = template.Must(template.ParseFiles("views/home.html", "views/base.html"))
	templates["impressum.html"] = template.Must(template.ParseFiles("views/impressum.html", "views/base.html"))
	e.Renderer = &TemplateRegistry{
		templates: templates,
	}

	// Route => handler
	e.GET("/home", handler.HomeHandler)
	e.GET("/impressum", handler.ImpressumHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
