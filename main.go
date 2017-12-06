package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
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

	// 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
	e.Use(middleware.LoggerWithConfig(logConfig))
	e.Use(middleware.Recover())
	e.Use(middleware.CSRF())
	e.Use(middleware.Gzip())
	e.Use(customHeader)

	// ルーティング
	e.GET("/", func(e echo.Context) error {
		return e.JSON(http.StatusOK,"Hello World")
	})

	e.Start(":" + os.Getenv("PORT"))
}

func customHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Add("X-Frame-Options", "SAMEORIGIN")
		c.Response().Header().Add("X-XSS-Protection", "1")
		c.Response().Header().Add("X-Content-Type-Options", "nosniff")
		c.Response().Header().Add("Pragma", "no-cache")
		c.Response().Header().Add("Cache-Control", "no-cache, no-store, must-revalidate")
		err := next(c)
		return err
	}
}