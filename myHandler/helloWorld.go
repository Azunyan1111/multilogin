package myHandler

import (
	"github.com/labstack/echo"
	"net/http"
)

func HelloWorld(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", "World")
}
