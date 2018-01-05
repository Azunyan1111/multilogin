package myHandler

import (
	"fmt"
	"github.com/Azunyan1111/multilogin/structs"
	"github.com/labstack/echo"
	"net/http"

	"github.com/Azunyan1111/multilogin/mysql"
	"github.com/labstack/echo-contrib/session"
)

func GetServiceMyPage(c echo.Context) error {
	orm := mysql.GetOrm()

	// セッション確認
	s, err := session.Get("session", c)
	if err != nil {
		panic(err)
	}
	var serviceUid string
	if s != nil {
		serviceUid = fmt.Sprintf("%v", s.Values["uid"])
	}
	if len(serviceUid) < 6 {
		return c.Render(http.StatusBadRequest, "error.html", structs.Error{StatusCode: http.StatusBadRequest,
			Message: "連携するにマルチログインにログインしてください"})
	}
	// サービス情報取得
	var serviceMyPage structs.ServiceMyPage
	rows := orm.Find(&serviceMyPage.Service, "uuid = ?", serviceUid).RowsAffected
	if rows != 1 {
		return c.Render(http.StatusBadRequest, "error.html", structs.Error{StatusCode: http.StatusBadRequest,
			Message: "サービス管理者としてログインしていません。サービスの管理を行うには再度ログインしてください。"})
	}

	// 連携ユーザー情報取得
	serviceMyPage.UserResponse = getUser(serviceMyPage.Service)
	// CSRF
	csrf := fmt.Sprintf("%v", c.Get("csrf"))
	serviceMyPage.Csrf = csrf
	return c.Render(http.StatusOK, "serviceMyPage.html", serviceMyPage)
}

func PostServiceMyPage(c echo.Context) error {
	orm := mysql.GetOrm()
	// セッション確認
	s, err := session.Get("session", c)
	if err != nil {
		panic(err)
	}
	var serviceUid string
	if s != nil {
		serviceUid = fmt.Sprintf("%v", s.Values["uid"])
	}
	if len(serviceUid) < 6 {
		return c.Render(http.StatusBadRequest, "error.html", structs.Error{StatusCode: http.StatusBadRequest,
			Message: "連携するにマルチログインにログインしてください"})
	}
	// サービス情報取得
	var serviceMyPage structs.ServiceMyPage
	rows := orm.Find(&serviceMyPage.Service, "uuid = ?", serviceUid).RowsAffected
	if rows != 1 {
		return c.Render(http.StatusBadRequest, "error.html", structs.Error{StatusCode: http.StatusBadRequest,
			Message: "サービス管理者としてログインしていません。サービスの管理を行うには再度ログインしてください。"})
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

	if row.RowsAffected != 1 {
		var user structs.UserMyPage
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

func getUser(service structs.Service) []structs.UserResponse {
	var users []structs.UserResponse
	var con []structs.ConfirmedService
	orm := mysql.GetOrm()
	orm.Find(&con, "service_uuid = ?", service.Uid).Limit(20)
	for _, c := range con {
		var res structs.UserResponse
		var user structs.User
		orm.Find(&user, "uuid = ?", c.UserUid)
		// 権限を確認する
		if service.UserName == 1 {
			res.UserName = user.UserName
		}
		if service.UserEmail == 1 {
			res.Email = user.Email
		}
		if service.UserImage == 1 {
			res.Image = user.Image
		}
		if service.UserAge == 1 {
			res.Age = user.Age
		}
		if service.UserBirthday == 1 {
			res.Birthday = user.Birthday
		}
		if service.UserPhone == 1 {
			res.Phone = user.Phone
		}
		if service.UserAddress == 1 {
			res.Address = user.Address
		}
		res.Uid = user.Uid
		if user.Uid == "" {
			continue
		}
		users = append(users, res)
	}
	return users
}
