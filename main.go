package main

import (
	"github.com/Azunyan1111/multilogin/myHandler"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
	"html/template"
	"io"
	"os"
)

func main() {
	// Echoのインスタンス作る
	e := echo.New()

	logConfig := middleware.LoggerConfig{
		Format: `| ${time_rfc3339} | ${host}${uri} ` +
			`| ${method} | ${status} ` +
			`| ${latency_human}` + "\n",
	}
	// HTML Template Read
	t := &Template{
		templates: template.Must(template.ParseGlob("static/views/*.html")),
	}
	e.Renderer = t

	// 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
	e.Use(middleware.LoggerWithConfig(logConfig))
	e.Use(middleware.Recover())
	e.Use(middleware.CSRF())
	e.Use(middleware.Gzip())
	e.Use(middleware.Secure())
	e.Use(session.Middleware(sessions.NewCookieStore(securecookie.GenerateRandomKey(64))))
	e.Use(customHeader)
	//e.StartAutoTLS(":443")

	e.Static("/static", "static")

	// ルーティング
	e.GET("/", myHandler.HelloWorld)

	e.Start(":" + os.Getenv("PORT"))
}

func customHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Add("Pragma", "no-cache")
		c.Response().Header().Add("Cache-Control", "no-cache, no-store, must-revalidate")
		err := next(c)
		return err
	}
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
