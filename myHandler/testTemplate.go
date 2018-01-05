package myHandler

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func TestTemplateGet(target string) (e *echo.Echo, req *http.Request, rec *httptest.ResponseRecorder) {
	e = echo.New()

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	temp := &Template{
		templates: template.Must(template.ParseGlob("../static/views/*.html")),
	}
	e.Renderer = temp

	req = httptest.NewRequest(echo.GET, target, nil)
	rec = httptest.NewRecorder()

	return e, req, rec
}
func TestTemplatePost(target string, json string) (e *echo.Echo, req *http.Request, rec *httptest.ResponseRecorder) {
	e = echo.New()

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	temp := &Template{
		templates: template.Must(template.ParseGlob("../static/views/*.html")),
	}
	e.Renderer = temp

	req = httptest.NewRequest(echo.POST, target, strings.NewReader(json))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec = httptest.NewRecorder()

	return e, req, rec
}
