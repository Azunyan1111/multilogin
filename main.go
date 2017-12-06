package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"os"
	"github.com/labstack/echo-contrib/session"
	"github.com/gorilla/sessions"
	"github.com/gorilla/securecookie"
	"log"
	"io"
	"html/template"
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
		templates: template.Must(template.ParseGlob("public/views/*.html")),
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
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK,"Hello World")
	})

	e.Start(":" + os.Getenv("PORT2"))
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