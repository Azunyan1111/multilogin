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

	// 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	e.GET("/", func(e echo.Context) error {
		return e.JSON(http.StatusOK,"Hello World")
	})

	// サーバー起動
	e.Start(":" + os.Getenv("PORT"))    //ポート番号指定してね
}