package main

import (
	"github.com/Azunyan1111/multilogin/myHandler"
	"github.com/Azunyan1111/multilogin/mysql"
	"github.com/labstack/echo"

	"github.com/Azunyan1111/multilogin/api"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
	"html/template"
	"io"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	// Echoのインスタンス作る
	e := echo.New()
	if err := mysql.DataBaseInit(); err != nil {
		panic(err)
	}

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
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:csrf",
	}))
	e.Use(middleware.Gzip())
	e.Use(middleware.Secure())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("key") /*securecookie.GenerateRandomKey(64)*/)))

	e.Use(customHeader)
	//e.StartAutoTLS(":443")

	e.Static("/static", "static")

	// ルーティング
	e.GET("/", myHandler.HelloWorld)
	// ログイン・ログアウト
	e.GET("/logout", myHandler.Logout)
	e.GET("/login", myHandler.GetLogin)
	e.POST("/login", myHandler.PostLogin)
	e.POST("/login/code", myHandler.PostLoginCode)
	// ユーザー
	e.GET("/user/new", myHandler.GetUserNew)
	e.POST("/user/new", myHandler.PostUserNew)
	e.GET("/user/mypage", myHandler.GetUserMyPage)
	e.POST("/user/mypage", myHandler.PostUserMyPage)
	// サービス
	e.GET("/service/new", myHandler.GetServiceNew)
	e.POST("/service/new", myHandler.PostServiceNew)
	e.GET("/service/mypage", myHandler.GetServiceMyPage)
	e.POST("/service/mypage", myHandler.PostServiceMyPage)

	// 連携
	e.GET("/confirmed/:serviceUid", myHandler.GetConfirmedNew)
	e.POST("/confirmed/:serviceUid", myHandler.GetConfirmedPost)
	e.POST("/confirmed/delete/:serviceUid", myHandler.PostConfirmedDelete)

	// API
	e.GET("/api/sum", api.Sum)
	e.GET("/api/user/name", api.GetName)
	e.GET("/api/user/image", api.GetImage)
	e.GET("/api/user/age", api.GetAge)
	e.GET("/api/user/birthday", api.GetBirthday)
	e.GET("/api/user/email", api.GetEmail)
	e.GET("/api/user/phone", api.Sum)
	e.GET("/api/user/address", api.Sum)

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
