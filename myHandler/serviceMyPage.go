package myHandler

import (
	"fmt"
	"github.com/Azunyan1111/multilogin/structs"
	"net/http"
	"github.com/labstack/echo"

	"github.com/labstack/echo-contrib/session"
	"github.com/Azunyan1111/multilogin/mysql"
)

func GetServiceMyPage(c echo.Context) error {
	orm := mysql.GetOrm()

	// セッション確認
	s, err := session.Get("session", c)
	if err != nil{
		panic(err)
	}
	var serviceUid string
	if s != nil{
		serviceUid = fmt.Sprintf("%v", s.Values["uid"])
	}
	if len(serviceUid) < 6{
		return c.Render(http.StatusBadRequest, "error.html",structs.Error{StatusCode:http.StatusBadRequest,
			Message:"連携するにマルチログインにログインしてください"})
	}
	// サービス情報取得
	var serviceMyPage structs.ServiceMyPage
	rows := orm.Find(&serviceMyPage.Service,"uuid = ?", serviceUid).RowsAffected
	if rows != 1{
		return c.Render(http.StatusBadRequest, "error.html",structs.Error{StatusCode:http.StatusBadRequest,
			Message:"サービス管理者としてログインしていません。サービスの管理を行うには再度ログインしてください。"})
	}
	// CSRF
	csrf := fmt.Sprintf("%v", c.Get("csrf"))
	serviceMyPage.Csrf = csrf
	return c.Render(http.StatusOK, "serviceMyPage.html", serviceMyPage)
}

func PostServiceMyPage(c echo.Context) error {
	orm := mysql.GetOrm()
	// セッション確認
	s, err := session.Get("session", c)
	if err != nil{
		panic(err)
	}
	var serviceUid string
	if s != nil{
		serviceUid = fmt.Sprintf("%v", s.Values["uid"])
	}
	if len(serviceUid) < 6{
		return c.Render(http.StatusBadRequest, "error.html",structs.Error{StatusCode:http.StatusBadRequest,
			Message:"連携するにマルチログインにログインしてください"})
	}
	// サービス情報取得
	var serviceMyPage structs.ServiceMyPage
	rows := orm.Find(&serviceMyPage.Service,"uuid = ?", serviceUid).RowsAffected
	if rows != 1{
		return c.Render(http.StatusBadRequest, "error.html",structs.Error{StatusCode:http.StatusBadRequest,
			Message:"サービス管理者としてログインしていません。サービスの管理を行うには再度ログインしてください。"})
	}

	// ユーザーが入力した情報取得
	var service structs.Service
	orm.Find(&service, "uuid = ?", serviceUid)
	// サービス情報
	service.Email = c.FormValue("InputEmail")
	service.ServiceName = c.FormValue("InputServiceName")
	service.Url = c.FormValue("InputUrl")
	service.CallbackUrl = c.FormValue("InputCallbackUrl")
	// 権限
	service.UserName = checkbox2(c.FormValue("InputUserName"))
	service.UserEmail = checkbox2(c.FormValue("InputUserEmail"))
	service.UserImage = checkbox2(c.FormValue("InputUserImage"))
	service.UserAge = checkbox2(c.FormValue("InputUserAge"))
	service.UserBirthday = checkbox2(c.FormValue("InputUserBirthday"))
	service.UserPhone = checkbox2(c.FormValue("InputUserPhone"))
	service.UserAddress = checkbox2(c.FormValue("InputUserAddress"))
	row := orm.Save(&service)

	if row.RowsAffected != 1{
		var user  structs.UserMyPage
		user.Message = "エラー | ユーザーデータが正しく登録できませんでした。"
		return c.Render(http.StatusBadRequest, "userMyPage.html", user)
	}

	return c.Render(http.StatusTemporaryRedirect, "serviceNewEnd.html", "")
}

func checkbox2(check string) int {
	if check == "on" {
		return 1
	} else {
		return 0
	}
}
