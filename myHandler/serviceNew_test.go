package myHandler

import (
	"github.com/Azunyan1111/multilogin/mysql"
	"github.com/Azunyan1111/multilogin/structs"
	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo"
	"github.com/gorilla/sessions"
)

func TestGetServiceNew(t *testing.T) {
	e, req, rec := testTemplateGet("/service/new")
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

	if assert.NoError(t, GetServiceNew(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		doc, _ := goquery.NewDocumentFromReader(rec.Result().Body)
		var text string
		doc.Find("#test_ServiceNew").Each(func(_ int, s *goquery.Selection) {
			text = s.Text()
		})
		assert.Equal(t, "サービス登録", text)
	}
}

func TestPostServiceNew(t *testing.T) {
	mysql.DataBaseInit()
	f := make(url.Values)
	var user structs.Serviced
	user.Email = "test@test.com"
	user.ServiceName = "TestUser114514"
	user.Url = "http://test.com"
	user.CallbackUrl = "http://test.com/callback"

	user.UserName = true
	user.UserEmail = true
	user.UserImage = true
	user.UserAge = true
	user.UserBirthday = true
	user.UserPhone = true
	user.UserAddress = true

	f.Set("InputEmail", user.Email)
	f.Set("InputServiceName", user.ServiceName)
	f.Set("InputUrl", user.Url)
	f.Set("InputCallbackUrl", user.CallbackUrl)

	f.Set("InputUserName", checkString(user.UserEmail))
	f.Set("InputUserEmail", checkString(user.UserName))
	f.Set("InputUserImage", checkString(user.UserImage))
	f.Set("InputUserAge", checkString(user.UserAge))
	f.Set("InputUserBirthday", checkString(user.UserBirthday))
	f.Set("InputUserPhone", checkString(user.UserPhone))
	f.Set("InputUserAddress", checkString(user.UserAddress))

	e, req, rec := testTemplatePost("/service/new", f.Encode())
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

	if assert.NoError(t, PostServiceNew(c)) {
		assert.Equal(t, http.StatusTemporaryRedirect, rec.Code)
		doc, _ := goquery.NewDocumentFromReader(rec.Result().Body)
		var text string
		doc.Find("#test_ServiceNewPost").Each(func(_ int, s *goquery.Selection) {
			text = s.Text()
		})
		assert.Equal(t, "登録完了", text)
	}

	sqlUser := structs.Service{}
	orm := mysql.GetOrm()
	orm.First(&sqlUser,"name = ?", user.ServiceName)


	assert.Equal(t, user.Email, sqlUser.Email)
	assert.Equal(t, user.ServiceName, sqlUser.ServiceName)
	assert.Equal(t, user.Url, sqlUser.Url)
	assert.Equal(t, user.CallbackUrl, sqlUser.CallbackUrl)

	assert.Equal(t, boolToInt(user.UserName), sqlUser.UserName)
	assert.Equal(t, boolToInt(user.UserEmail), sqlUser.UserEmail)
	assert.Equal(t, boolToInt(user.UserImage), sqlUser.UserImage)
	assert.Equal(t, boolToInt(user.UserAge), sqlUser.UserAge)
	assert.Equal(t, boolToInt(user.UserBirthday), sqlUser.UserBirthday)
	assert.Equal(t, boolToInt(user.UserPhone), sqlUser.UserPhone)
	assert.Equal(t, boolToInt(user.UserAddress), sqlUser.UserAddress)

	if err := mysql.DeleteUserByTestService(); err != nil {
		panic(err)
	}
}

func checkString(check bool) string {
	if check {
		return "on"
	} else {
		return ""
	}
}
