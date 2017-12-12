package myHandler

import (
	"fmt"
	"github.com/Azunyan1111/multilogin/structs"
	"net/http"
	"github.com/labstack/echo"
	"github.com/ipfans/echo-session"
	"github.com/Azunyan1111/multilogin/mysql"
	"github.com/Azunyan1111/multilogin/model"
)

func GetUserMyPage(c echo.Context) error {
	// セッション確認
	s := session.Default(c)
	var uid string
	if s != nil{
		uid = fmt.Sprintf("%v", s.Get("uid"))
	}
	if uid == "<nil>" || uid == ""{
		// Not Login
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}
	var user structs.UserMyPage
	var err error
	csrf := fmt.Sprintf("%v", c.Get("csrf"))
	user.Csrf = csrf

	// ユーザーデータ取得
	user.User, err = mysql.SelectUserByUuid(uid)
	if err != nil {
		user.Message = "エラー | ユーザーデータが正しく登録できない可能性があります。次の項目を添えて管理者に問い合わせてください | " +  err.Error()
		return c.Render(http.StatusInternalServerError, "userMyPage.html", user)
	}
	// 連携データ取得
	service, err := mysql.SelectConfirmedByUid(uid)
	if err != nil{
		user.Message = "エラー | 連携設定が正しく取得できませんでした。次の項目を添えて管理者に問い合わせてください | " +  err.Error()
		return c.Render(http.StatusInternalServerError, "userMyPage.html", user)
	}
	user.Service = service
	return c.Render(http.StatusOK, "userMyPage.html", user)
}

func PostUserMyPage(c echo.Context) error {
	// セッション確認
	s := session.Default(c)
	var uid string
	if s != nil{
		uid = fmt.Sprintf("%v", s.Get("uid"))
	}
	if uid == "<nil>"{
		// Not Login
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}
	var user structs.UserMyPage
	var err error
	csrf := fmt.Sprintf("%v", c.Get("csrf"))
	user.Csrf = csrf
	user.User.Uid = uid
	user.User.UserName = c.FormValue("InputUserName")
	user.User.Email = c.FormValue("InputEmail")
	user.User.Image = c.FormValue("InputImage")
	user.User.Age = c.FormValue("InputAge")
	user.User.Birthday = c.FormValue("InputBirthday")
	user.User.Phone = c.FormValue("InputPhone")
	user.User.Address = c.FormValue("InputAddress")

	err = model.UpdateNewUser(user.User)
	if err != nil{
		user.Message = "エラー | ユーザーデータが正しく登録できませんでした。次の項目を添えて管理者に問い合わせてください | " +  err.Error()
		return c.Render(http.StatusBadRequest, "userMyPage.html", user)
	}
	user.User, err = mysql.SelectUserByUuid(uid)
	if err != nil {
		user.Message = "エラー | ユーザーデータに不自然な値が混入しています。次の項目を添えて管理者に問い合わせてください | " +  err.Error()
		return c.Render(http.StatusBadRequest, "userMyPage.html", user)
	}

	return c.Render(http.StatusTemporaryRedirect, "userNewEnd.html", "")
}
