package myHandler

import (
	"github.com/stretchr/testify/assert"
	"github.com/PuerkitoBio/goquery"
	"testing"
	"github.com/labstack/echo-contrib/session"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"net/http"
	"github.com/Azunyan1111/multilogin/structs"
	"net/url"
)

var code string

func TestGetLogin(t *testing.T) {
	e, req, rec := testTemplateGet("/login")
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
	f := make(url.Values)
	var user structs.Usered
	user.Email = "god@god.com"
	f.Set("InputEmail", user.Email)

	e, req, rec := testTemplatePost("/login", f.Encode())
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
