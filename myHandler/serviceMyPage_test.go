package myHandler

import (
	"testing"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"github.com/Azunyan1111/multilogin/structs"
	"net/url"
	"time"
	"strconv"
	"github.com/Azunyan1111/multilogin/mysql"
	"math/rand"
)


func TestGetServiceMyPage(t *testing.T) {

	e, req, rec := testTemplateGet("/service/mypage")
	c := e.NewContext(req, rec)
	// session
	mw := session.Middleware(sessions.NewCookieStore([]byte("secret")))
	h := mw(func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		sess.Values["uid"] = serviceUid
		sess.Save(c.Request(), c.Response())
		return nil
	})
	h(c)

	if assert.NoError(t, GetServiceMyPage(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		doc, _ := goquery.NewDocumentFromReader(rec.Result().Body)
		email, _ := doc.Find("#new > div:nth-child(1) > div.panel-body >" +
			" form > div:nth-child(1) > div > input").Attr("value")
		assert.Equal(t,"god@god.com",email)
		assert.Equal(t, "サービス登録情報更新", doc.Find("#TestServiceMyPage").Text())
	}
}

func TestPostServiceMyPage(t *testing.T) {
	orm := mysql.GetOrm()
	var testService structs.Service
	orm.Find(&testService, "uuid = ?", serviceUid)

	// アップデートする項目をセット
	f := make(url.Values)
	rand.Seed(time.Now().UnixNano())
	randomNumber := strconv.Itoa(rand.Intn(100))
	curl := "http://bar.com/callback" +randomNumber

	// 既存の情報
	f.Set("InputEmail", testService.Email)
	f.Set("InputServiceName", testService.ServiceName)
	f.Set("InputUrl", testService.Url)
	f.Set("InputCallbackUrl", curl)

	f.Set("InputUserName", "")//boolToString(testService.UserName))
	f.Set("InputUserEmail", "on")
	f.Set("InputUserImage", boolToString(testService.UserImage))
	f.Set("InputUserAge", boolToString(testService.UserAge))
	f.Set("InputUserBirthday", boolToString(testService.UserBirthday))
	f.Set("InputUserPhone", boolToString(testService.UserPhone))
	f.Set("InputUserAddress", boolToString(testService.UserAddress))


	e, req, rec := testTemplatePost("/service/mypage",f.Encode())
	c := e.NewContext(req, rec)
	// session
	mw := session.Middleware(sessions.NewCookieStore([]byte("secret")))
	h := mw(func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		sess.Values["uid"] = serviceUid
		sess.Save(c.Request(), c.Response())
		return nil
	})
	h(c)


	if assert.NoError(t, PostServiceMyPage(c)) {
		assert.Equal(t, http.StatusTemporaryRedirect, rec.Code)
		doc, _ := goquery.NewDocumentFromReader(rec.Result().Body)

		var service structs.Service
		orm.Find(&service,"uuid = ?", serviceUid)
		assert.Equal(t,curl, service.CallbackUrl)
		assert.Equal(t,0, service.UserName)
		assert.Equal(t,1 , service.UserEmail)
		assert.Equal(t, "登録完了", doc.Find("#test_ServiceNewPost").Text())
		var old structs.Service
		old = service
		old.UserName = 1
		old.UserEmail = 1
		old.UserImage = 1
		old.UserAge = 1
		old.UserBirthday = 1
		old.UserPhone = 1
		old.UserAddress = 1
		assert.Equal(t,int64(1),orm.Model(&service).Updates(&old).RowsAffected)
	}

}

func boolToString(b int)(string){
	if b == 1{
		return "on"
	}else {
		return ""
	}
}

func boolToInt(b bool)(int){
	if b{
		return 1
	}else {
		return 0
	}
}
