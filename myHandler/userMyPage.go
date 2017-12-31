package myHandler

import (
	"fmt"
	"github.com/Azunyan1111/multilogin/structs"
	"net/http"
	"github.com/labstack/echo"

	"github.com/Azunyan1111/multilogin/mysql"
	"github.com/Azunyan1111/multilogin/model"
	"github.com/labstack/echo-contrib/session"
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
	log.Println(userUid,"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")

	var user structs.UserMyPage
	csrf := fmt.Sprintf("%v", c.Get("csrf"))
	user.Csrf = csrf

	// ユーザーデータ取得
	user.User, err = mysql.SelectUserByUuid(userUid)
	if err != nil {
		user.Message = "エラー | ユーザーデータが正しく登録できない可能性があります。次の項目を添えて管理者に問い合わせてください | " +  err.Error()
		return c.Render(http.StatusInternalServerError, "userMyPage.html", user)
	}
	// 連携データ取得
	var con []structs.ConfirmedService
	orm.Find(&con,"user_uuid = ?", userUid)
	// 連携しているサービスのデータを入手
	var serviceds []structs.Serviced
	for _,c := range con{
		var serviced structs.Serviced
		orm.Find(&serviced,"uuid = ?", c.ServiceUid)
		serviceds = append(serviceds, serviced)
	}
	user.Service = serviceds
	return c.Render(http.StatusOK, "userMyPage.html", user)
}

func PostUserMyPage(c echo.Context) error {
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
	var user structs.UserMyPage
	csrf := fmt.Sprintf("%v", c.Get("csrf"))
	user.Csrf = csrf
	user.User.Uid = userUid
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
	user.User, err = mysql.SelectUserByUuid(userUid)
	if err != nil {
		user.Message = "エラー | ユーザーデータに不自然な値が混入しています。次の項目を添えて管理者に問い合わせてください | " +  err.Error()
		return c.Render(http.StatusBadRequest, "userMyPage.html", user)
	}

	return c.Render(http.StatusTemporaryRedirect, "userNewEnd.html", "")
}
