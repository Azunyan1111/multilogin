package myHandler

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/echo-contrib/session"
	"net/http"
	"github.com/labstack/echo"
	"github.com/gorilla/sessions"
)

func TestGetUserMyPage(t *testing.T) {

	e, req, rec := testTemplateGet("/user/mypage")
	c := e.NewContext(req, rec)
	// session
	mw := session.Middleware(sessions.NewCookieStore([]byte("secret")))
	h := mw(func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		sess.Values["uid"] = "26d2983e-3d5a-421c-bf6f-d4608025e555"
		sess.Save(c.Request(), c.Response())
		return nil
	})
	h(c)

	if assert.NoError(t, GetUserMyPage(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		doc, _ := goquery.NewDocumentFromReader(rec.Result().Body)
		email := doc.Find("#new > div:nth-child(1) > div.panel-body > form > div:nth-child(1) > div > input").Text()
		assert.Equal(t,"azunyan1111@azunyan.me",email)
		assert.Equal(t, "登録情報更新", doc.Find("#test_UserMyPage").Text())
	}
}