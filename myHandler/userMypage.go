package myHandler

import (
	"fmt"
	"github.com/Azunyan1111/multilogin/structs"
	"net/http"
	"github.com/labstack/echo"
	"github.com/ipfans/echo-session"
	"log"
)

func GetUserMyPage(c echo.Context) error {
	// セッション確認
	s := session.Default(c)
	var uid string
	// Test is not run
	if s != nil{
		uid = fmt.Sprintf("%v", s.Get("uid"))
	}
	log.Println(uid)
	if uid == "<nil>"{
		// Not Login
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}
	csrf := fmt.Sprintf("%v", c.Get("csrf"))
	return c.Render(http.StatusOK, "userNew.html", structs.UserNewPage{Csrf: csrf})
}
