package myHandler

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func TestHandler_HelloWorld(t *testing.T) {
	e := echo.New()

	temp := &Template{
		templates: template.Must(template.ParseGlob("../static/views/index.html")),
	}
	e.Renderer = temp

	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, HelloWorld(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Hello World", rec.Body.String())
	}
}
