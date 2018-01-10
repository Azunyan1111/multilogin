package myHandler

import (
	"github.com/Azunyan1111/multilogin/structs"
	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
	"github.com/Azunyan1111/multilogin/mysql"
)

var code string

func TestGetLogin(t *testing.T) {
	mysql.DataBaseInit()
	e, req, rec := TestTemplateGet("/login")
	c := e.NewContext(req, rec)

	// session
	mw := session.Middleware(sessions.NewCookieStore([]byte("secret")))
	h := mw(func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		sess.Values["uid"] = ""
		sess.Save(c.Request(), c.Response())
		return nil
	})
	h(c)

	if assert.NoError(t, GetLogin(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		doc, _ := goquery.NewDocumentFromReader(rec.Result().Body)
		var text string
		doc.Find("#test_userNew").Each(func(_ int, s *goquery.Selection) {
			text = s.Text()
		})
		assert.Equal(t, "ログイン", text)
	}
}

func TestPostLogin(t *testing.T) {
	mysql.DataBaseInit()
	f := make(url.Values)
	var user structs.Usered
	user.Email = "god@god.com"
	f.Set("InputEmail", user.Email)

	e, req, rec := TestTemplatePost("/login", f.Encode())
	c := e.NewContext(req, rec)

	// session
	mw := session.Middleware(sessions.NewCookieStore([]byte("secret")))
	h := mw(func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		sess.Values["uid"] = ""
		sess.Save(c.Request(), c.Response())
		return nil
	})
	h(c)

	if assert.NoError(t, PostLogin(c)) {
		assert.Equal(t, http.StatusTemporaryRedirect, rec.Code)
		doc, _ := goquery.NewDocumentFromReader(rec.Result().Body)
		var text string
		doc.Find("#test_userNew").Each(func(_ int, s *goquery.Selection) {
			text = s.Text()
		})
		assert.Equal(t, "ログインコード入力", text)
	}
}
