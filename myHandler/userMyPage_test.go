package myHandler

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/echo-contrib/session"
	"net/http"
	"github.com/labstack/echo"
	"github.com/gorilla/sessions"
	"net/url"
	"math/rand"
	"github.com/Azunyan1111/multilogin/structs"
	"github.com/Azunyan1111/multilogin/mysql"
	"time"
	"strconv"
)


func TestGetUserMyPage(t *testing.T) {

	e, req, rec := testTemplateGet("/user/mypage")
	c := e.NewContext(req, rec)
	// session
	mw := session.Middleware(sessions.NewCookieStore([]byte("secret")))
	h := mw(func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		sess.Values["uid"] = userUid
		sess.Save(c.Request(), c.Response())
		return nil
	})
	h(c)

	if assert.NoError(t, GetUserMyPage(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		doc, _ := goquery.NewDocumentFromReader(rec.Result().Body)
		email, _ := doc.Find("#new > div:nth-child(1) > div.panel-body >" +
			" form > div:nth-child(1) > div > input").Attr("value")
		assert.Equal(t,"azunyan1111@azunyan.me",email)
		assert.Equal(t, "登録情報更新", doc.Find("#test_UserMyPage").Text())
	}
}

func TestPostUserMyPage(t *testing.T) {
	orm := mysql.GetOrm()

	var testUser structs.User
	orm.Find(&testUser, "uuid = ?", userUid)


	// アップデートする項目をセット
	f := make(url.Values)
	rand.Seed(time.Now().UnixNano())
	age := strconv.Itoa(rand.Intn(100))
	f.Set("InputAge", age)

	// 既存の情報
	f.Set("InputUserName", testUser.UserName)
	f.Set("InputEmail", testUser.Email)
	f.Set("InputImage", testUser.Image)
	//f.Set("InputAge", testUser.Age)
	f.Set("InputBirthday", testUser.Birthday)
	f.Set("InputPhone", testUser.Phone)
	f.Set("InputAddress", testUser.Address)


	e, req, rec := testTemplatePost("/user/mypage",f.Encode())
	c := e.NewContext(req, rec)
	// session
	mw := session.Middleware(sessions.NewCookieStore([]byte("secret")))
	h := mw(func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		sess.Values["uid"] = userUid
		sess.Save(c.Request(), c.Response())
		return nil
	})
	h(c)

	if assert.NoError(t, PostUserMyPage(c)) {
		assert.Equal(t, http.StatusTemporaryRedirect, rec.Code)
		doc, _ := goquery.NewDocumentFromReader(rec.Result().Body)

		var user structs.User
		orm.Find(&user,"uuid = ?", userUid)
		assert.Equal(t,age,user.Age)
		assert.Equal(t, "登録完了", doc.Find("#test_userNewPost").Text())
	}
}