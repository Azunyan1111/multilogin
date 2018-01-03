package myHandler

import (
	"github.com/Azunyan1111/multilogin/mysql"
	"github.com/Azunyan1111/multilogin/structs"
	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo"
)

func TestGetUserNew(t *testing.T) {
	e, req, rec := TestTemplateGet("/user/new")
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

	if assert.NoError(t, GetUserNew(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		doc, _ := goquery.NewDocumentFromReader(rec.Result().Body)
		var text string
		doc.Find("#test_userNew").Each(func(_ int, s *goquery.Selection) {
			text = s.Text()
		})
		assert.Equal(t, "ユーザー登録", text)
	}
}

func TestPostUserNew(t *testing.T) {
	mysql.DataBaseInit()
	f := make(url.Values)
	var user structs.Usered
	user.Email = "bar@bar.com"
	user.UserName = "TestUser114514"
	user.Image = "https://upload.wikimedia.org/wikipedia/commons/thumb/4/43/Bar-P1030319.jpg/1200px-Bar-P1030319.jpg"
	user.Age = "22"
	user.Birthday = "1985-5-3"
	user.Phone = "080-3749-7392"
	user.Address = "福岡県"
	f.Set("InputEmail", user.Email)
	f.Set("InputUserName", user.UserName)
	f.Set("InputImage", user.Image)
	f.Set("InputAge", user.Age)
	f.Set("InputBirthday", user.Birthday)
	f.Set("InputPhone", user.Phone)
	f.Set("InputAddress", user.Address)

	e, req, rec := TestTemplatePost("/user/new", f.Encode())
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

	if assert.NoError(t, PostUserNew(c)) {
		assert.Equal(t, http.StatusTemporaryRedirect, rec.Code)
		doc, _ := goquery.NewDocumentFromReader(rec.Result().Body)
		var text string
		doc.Find("#test_userNewPost").Each(func(_ int, s *goquery.Selection) {
			text = s.Text()
		})
		assert.Equal(t, "登録完了", text)
	}

	sqlUser := structs.User{}
	orm := mysql.GetOrm()
	orm.First(&sqlUser,"user = ?", user.UserName)

	assert.Equal(t, user.UserName, sqlUser.UserName)
	assert.Equal(t, user.Email, sqlUser.Email)
	assert.Equal(t, user.Image, sqlUser.Image)
	assert.Equal(t, user.Age, sqlUser.Age)
	assert.Equal(t, user.Birthday, sqlUser.Birthday)
	assert.Equal(t, user.Phone, sqlUser.Phone)
	assert.Equal(t, user.Address, sqlUser.Address)
	if err := mysql.DeleteUserByTestUser(); err != nil {
		panic(err)
	}
}
