package myHandler

import (
	"fmt"
	"github.com/Azunyan1111/multilogin/model"
	"github.com/Azunyan1111/multilogin/structs"
	"github.com/labstack/echo"
	"net/http"
)

func GetServiceNew(c echo.Context) error {
	csrf := fmt.Sprintf("%v", c.Get("csrf"))
	return c.Render(http.StatusOK, "serviceNew.html", structs.ServiceNewPage{Csrf: csrf})
}

func PostServiceNew(c echo.Context) error {
	csrf := fmt.Sprintf("%v", c.Get("csrf"))
	user := structs.ServiceNewPage{Csrf: csrf}
	// サービス情報
	user.Service.Email = c.FormValue("InputEmail")
	user.Service.ServiceName = c.FormValue("InputServiceName")
	user.Service.Url = c.FormValue("InputUrl")
	user.Service.CallbackUrl = c.FormValue("InputCallbackUrl")

	// 権限
	user.Service.UserName = checkbox(c.FormValue("InputUserName"))
	user.Service.UserEmail = checkbox(c.FormValue("InputUserEmail"))
	user.Service.UserImage = checkbox(c.FormValue("InputUserImage"))
	user.Service.UserAge = checkbox(c.FormValue("InputUserAge"))
	user.Service.UserBirthday = checkbox(c.FormValue("InputUserBirthday"))
	user.Service.UserPhone = checkbox(c.FormValue("InputUserPhone"))
	user.Service.UserAddress = checkbox(c.FormValue("InputUserAddress"))

	// 一つでも入ってない場合はまた入力させる
	if val := user.Service.Email; val == "" {
		user.Message = "入力されてない箇所があります | 例えば : メールアドレス"
		return c.Render(http.StatusBadRequest, "serviceNew.html", user)
	}
	if val := user.Service.ServiceName; val == "" {
		user.Message = "入力されてない箇所があります | 例えば : ユーザー名"
		return c.Render(http.StatusBadRequest, "serviceNew.html", user)
	}
	if err := model.RegisterNewService(user.Service); err != nil {
		user.Message = "サーバーエラーです | 次の内容を管理者に問い合わせてください | " + err.Error()
		return c.Render(http.StatusBadRequest, "serviceNew.html", user)
	}

	return c.Render(http.StatusTemporaryRedirect, "serviceNewEnd.html", "")
}

func checkbox(check string) bool {
	if check == "on" {
		return true
	} else {
		return false
	}
}
