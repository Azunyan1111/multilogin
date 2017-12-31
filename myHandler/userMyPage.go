package myHandler

import (
	"fmt"
	"github.com/Azunyan1111/multilogin/structs"
	"net/http"
	"github.com/labstack/echo"

	"github.com/labstack/echo-contrib/session"
	"github.com/Azunyan1111/multilogin/mysql"
	"log"
)
// TODO:GetUseMyPageのユーザー取得をormで行う。
func GetUserMyPage(c echo.Context) error {
	orm := mysql.GetOrm()

	// セッション確認
	s, err := session.Get("session", c)
	if err != nil{
		panic(err)
	}
	var userUid string
	if s != nil{
		userUid = fmt.Sprintf("%v", s.Values["uid"])
	}
	if len(userUid) < 6{
		return c.Render(http.StatusBadRequest, "error.html",structs.Error{StatusCode:http.StatusBadRequest,
			Message:"連携するにマルチログインにログインしてください"})
	}

	// ユーザーデータ取得
	var user structs.UserMyPage
	orm.Find(&user.User,"uuid = ?", userUid)
	if user.User.ID == 0 {
		user.Message = "エラー | ユーザーデータが正しく登録できない可能性があります。次の項目を添えて管理者に問い合わせてください | " +  err.Error()
		return c.Render(http.StatusInternalServerError, "userMyPage.html", user)
	}
	// 連携データ取得
	var con []structs.ConfirmedService
	orm.Find(&con,"user_uuid = ?", userUid)
	// 連携しているサービスのデータを入手
	var services []structs.Service
	for _,c := range con{
		var serviced structs.Service
		orm.Find(&serviced,"uuid = ?", c.ServiceUid)
		services = append(services, serviced)
	}
	user.Service = services

	// CSRF
	csrf := fmt.Sprintf("%v", c.Get("csrf"))
	user.Csrf = csrf

	return c.Render(http.StatusOK, "userMyPage.html", user)
}

func PostUserMyPage(c echo.Context) error {
	orm := mysql.GetOrm()

	// セッション確認
	s, err := session.Get("session", c)
	if err != nil{
		panic(err)
	}
	var userUid string
	if s != nil{
		userUid = fmt.Sprintf("%v", s.Values["uid"])
	}
	if len(userUid) < 6{
		return c.Render(http.StatusBadRequest, "error.html",structs.Error{StatusCode:http.StatusBadRequest,
			Message:"連携するにマルチログインにログインしてください"})
	}

	// ユーザーが入力した情報取得s
	var oldUser structs.User
	orm.Find(&oldUser, "uuid = ?", userUid)
	log.Println(userUid)
	var updateUser structs.User
	updateUser = oldUser
	updateUser.UserName = c.FormValue("InputUserName")
	updateUser.Email = c.FormValue("InputEmail")
	updateUser.Image = c.FormValue("InputImage")
	updateUser.Age = c.FormValue("InputAge")
	updateUser.Birthday = c.FormValue("InputBirthday")
	updateUser.Phone = c.FormValue("InputPhone")
	updateUser.Address = c.FormValue("InputAddress")
	log.Println("aaaaaaaaaaaaaaaaaa",oldUser.Age, updateUser.Age)

	if orm.Model(&oldUser).Updates(&updateUser).RowsAffected != 1{
		var user  structs.UserMyPage
		user.Message = "エラー | ユーザーデータが正しく登録できませんでした。"
		return c.Render(http.StatusBadRequest, "userMyPage.html", user)
	}

	return c.Render(http.StatusTemporaryRedirect, "userNewEnd.html", "")
}
