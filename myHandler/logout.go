package myHandler

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/labstack/echo-contrib/session"
)

func Logout(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Values["uid"] = ""
	sess.Save(c.Request(), c.Response())
	return c.Render(http.StatusOK, "logout.html", "正常にログアウトしました。")
}
