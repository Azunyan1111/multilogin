package myHandler

import (
	"github.com/Azunyan1111/multilogin/mysql"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo-contrib/session"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"

	"github.com/Azunyan1111/multilogin/structs"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
)

const (
	userUid    = "26d2983e-3d5a-421c-bf6f-d4608025e555"
	serviceUid = "025ad602-7dba-4c08-8226-704b65f2873c"
)

func TestGetConfirmedNew(t *testing.T) {
	mysql.DataBaseInit()

	e, req, rec := TestTemplateGet("/")
	c := e.NewContext(req, rec)

	// URL param
	c.SetParamNames("serviceUid")
	c.SetParamValues("025ad602-7dba-4c08-8226-704b65f2873c")

	// session
	mw := session.Middleware(sessions.NewCookieStore([]byte("secret")))
	h := mw(func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		sess.Values["uid"] = userUid
		sess.Save(c.Request(), c.Response())
		return nil
	})
	h(c)
	ses, _ := session.Get("session", c)
	ses.Values["uid"] = userUid

	if assert.NoError(t, GetConfirmedNew(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		doc, _ := goquery.NewDocumentFromReader(rec.Result().Body)
		assert.Empty(t, doc.Find("#StatusCode").Text())
		assert.Empty(t, doc.Find("#ErrorMessage").Text())
		var text string
		doc.Find("#new > div > div.panel-body > p:nth-child(1)").Each(func(_ int, s *goquery.Selection) {
			text = s.Text()
		})
		assert.Equal(t, "GodServiceと連携してます", text)
	}
}

func TestGetConfirmedPost(t *testing.T) {
	mysql.DataBaseInit()

	e, req, rec := TestTemplateGet("/")

	c := e.NewContext(req, rec)
	// url param
	c.SetParamNames("serviceUid")
	c.SetParamValues(serviceUid)
	// session
	mw := session.Middleware(sessions.NewCookieStore([]byte("secret")))
	h := mw(func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		sess.Values["uid"] = userUid
		sess.Save(c.Request(), c.Response())
		return nil
	})
	h(c)
	ses, _ := session.Get("session", c)
	ses.Values["uid"] = userUid

	orm := mysql.GetOrm()
	var confirmedService structs.ConfirmedService

	if assert.NoError(t, GetConfirmedPost(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		orm.Find(&confirmedService, "user_uuid = ? and service_uuid = ?", userUid, serviceUid)
		assert.Equal(t, userUid, confirmedService.UserUid)
		assert.Equal(t, serviceUid, confirmedService.ServiceUid)

		doc, _ := goquery.NewDocumentFromReader(rec.Result().Body)
		assert.Empty(t, doc.Find("#StatusCode").Text())
		assert.Empty(t, doc.Find("#ErrorMessage").Text())
		assert.Equal(t, "サービス連携が完了しました。サービスに戻ります。", doc.Find("#test_ConfirmedNew").Text())
	}
	//orm.Delete(&confirmedService)
}

func TestGetConfirmedPost2(t *testing.T) {
	mysql.DataBaseInit()
	orm := mysql.GetOrm()

	con := structs.ConfirmedService{UserUid: "daa8123d-be45-4574-88d7-339b145396fc"}

	service := structs.Service{Uid: "025ad602-7dba-4c08-8226-704b65f2873c"}
	user := structs.User{Uid: "26d2983e-3d5a-421c-bf6f-d4608025e555"}

	orm.First(&con)
	orm.First(&service)
	orm.First(&user)

	assert.Equal(t, "daa8123d-be45-4574-88d7-339b145396fc", con.UserUid)
	assert.Equal(t, "GodService", service.ServiceName)
	assert.Equal(t, "Azunyan1111", user.UserName)
}

func TestPostConfirmedDelete(t *testing.T) {
	mysql.DataBaseInit()
	orm := mysql.GetOrm()
	e, req, rec := TestTemplatePost("/confirmed/delete/"+serviceUid, "")
	c := e.NewContext(req, rec)

	c.SetParamNames("serviceUid")
	c.SetParamValues(serviceUid)

	// session
	mw := session.Middleware(sessions.NewCookieStore([]byte("secret")))
	h := mw(func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		sess.Values["uid"] = userUid
		sess.Save(c.Request(), c.Response())
		return nil
	})
	h(c)
	ses, _ := session.Get("session", c)
	ses.Values["uid"] = userUid
	ses.Save(c.Request(), c.Response())

	if assert.NoError(t, PostConfirmedDelete(c)) {
		assert.Equal(t, http.StatusTemporaryRedirect, rec.Code)
		assert.Equal(t, int64(0), orm.Find(&structs.ConfirmedService{},
			"user_uuid = ? and service_uuid = ?", userUid, serviceUid).RowsAffected)
	}
}
