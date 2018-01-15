package myHandler

import (
	"fmt"
	"github.com/Azunyan1111/multilogin/model"
	"github.com/Azunyan1111/multilogin/mysql"
	"github.com/Azunyan1111/multilogin/redis"
	"github.com/Azunyan1111/multilogin/structs"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

func GetLogin(c echo.Context) error {
	// セッション確認
	s, err := session.Get("session", c)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.html", structs.Error{StatusCode: http.StatusInternalServerError,
			Message: "サーバーの処理に異常がありました。エラーコード:1001"})
	}
	var userUid string
	if s != nil {
		userUid = fmt.Sprintf("%v", s.Values["uid"])
	}
	if len(userUid) > 6 {
		return c.Render(http.StatusBadRequest, "error.html", structs.Error{StatusCode: http.StatusBadRequest,
			Message: "ログイン済みです。別のアカウントにログインする場合はログアウトしてください。"})
	}

	csrf := fmt.Sprintf("%v", c.Get("csrf"))
	return c.Render(http.StatusOK, "login.html", structs.LoginPage{Csrf: csrf, Email: ""})
}

func PostLogin(c echo.Context) error {
	orm := mysql.GetOrm()
	csrf := fmt.Sprintf("%v", c.Get("csrf"))

	// セッション確認
	s, err := session.Get("session", c)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.html", structs.Error{StatusCode: http.StatusInternalServerError,
			Message: "サーバーの処理に異常がありました。エラーコード:1001"})
	}
	var userUid string
	if s != nil {
		userUid = fmt.Sprintf("%v", s.Values["uid"])
	}
	if len(userUid) > 6 {
		return c.Render(http.StatusBadRequest, "error.html", structs.Error{StatusCode: http.StatusBadRequest,
			Message: "ログイン済みです。別のアカウントにログインする場合はログアウトしてください。"})
	}

	// ユーザー登録確認
	email := c.FormValue("InputEmail")
	var user structs.User
	var service structs.Service
	uid := ""
	if orm.Find(&user, "email = ?", email).RowsAffected == 1 {
		uid = user.Uid
	}
	if orm.Find(&service, "email = ?", email).RowsAffected == 1 {
		uid = service.Uid
	}
	if uid == "" {
		return c.Render(http.StatusBadRequest, "login.html", structs.LoginPage{Csrf: csrf, Email: email,
			Message: "アカウントが存在しないかデータベースに不具合があります。お問い合わせください。"})
	}
	var loginCodePage structs.LoginCodePage
	loginCodePage.Csrf = csrf

	// ログインコード登録
	code := strconv.Itoa(rand.Intn(9)) + strconv.Itoa(rand.Intn(9)) + strconv.Itoa(rand.Intn(9)) +
		strconv.Itoa(rand.Intn(9)) + strconv.Itoa(rand.Intn(9)) + strconv.Itoa(rand.Intn(9))
	log.Println(code)
	redis.Set(code, uid)
	model.SendMail(email, "ログインコードお知らせします。", "ログインコードは「"+code+"」です。数字四桁になのます。")
	return c.Render(http.StatusTemporaryRedirect, "loginCode.html", loginCodePage)
}

func PostLoginCode(c echo.Context) error {
	csrf := fmt.Sprintf("%v", c.Get("csrf"))

	// セッション確認
	s, err := session.Get("session", c)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.html", structs.Error{StatusCode: http.StatusInternalServerError,
			Message: "サーバーの処理に異常がありました。エラーコード:1001"})
	}
	var userUid string
	if s != nil {
		userUid = fmt.Sprintf("%v", s.Values["uid"])
	}
	if len(userUid) > 6 {
		return c.Render(http.StatusBadRequest, "error.html", structs.Error{StatusCode: http.StatusBadRequest,
			Message: "ログイン済みです。別のアカウントにログインする場合はログアウトしてください。"})
	}

	// ユーザー登録確認
	code := c.FormValue("InputCode")
	uid := redis.Get(code)
	if uid == "" {
		var loginCodePage structs.LoginCodePage
		loginCodePage.Csrf = csrf
		loginCodePage.Message = "認証コードが間違っています"
		return c.Render(http.StatusTemporaryRedirect, "loginCode.html", loginCodePage)
	}
	// セッションに自分のuuidをつけて返す
	s.Values["uid"] = uid
	s.Save(c.Request(), c.Response().Writer)

	return c.Render(http.StatusTemporaryRedirect, "loginCodeEnd.html", nil)
}
