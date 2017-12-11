package myHandler

import (
	"fmt"
	"github.com/Azunyan1111/multilogin/model"
	"github.com/labstack/echo"
	"net/http"
	"github.com/Azunyan1111/multilogin/structs"
)

func GetUserNew(c echo.Context) error {
	csrf := fmt.Sprintf("%v", c.Get("csrf"))
	return c.Render(http.StatusOK, "userNew.html", structs.UserNewPage{Csrf: csrf})
}

func PostUserNew(c echo.Context) error {
	csrf := fmt.Sprintf("%v", c.Get("csrf"))
	user := structs.UserNewPage{Csrf: csrf}
	user.User.Email = c.FormValue("InputEmail")
	user.User.UserName = c.FormValue("InputUserName")
	user.User.Image = c.FormValue("InputImage")
	user.User.Age = c.FormValue("InputAge")
	user.User.Birthday = c.FormValue("InputBirthday")
	user.User.Phone = c.FormValue("InputPhone")
	user.User.Address = c.FormValue("InputAddress")

	// 一つでも入ってない場合はまた入力させる
	if val := user.User.Email; val == "" {
		user.Message = "入力されてない箇所があります | 例えば : メールアドレス"
		return c.Render(http.StatusBadRequest, "userNew.html", user)
	}
	if val := user.User.UserName; val == "" {
		user.Message = "入力されてない箇所があります | 例えば : ユーザー名"
		return c.Render(http.StatusBadRequest, "userNew.html", user)
	}
	if err := model.RegisterNewUser(user.User); err != nil{
		user.Message = "サーバーエラーです | 次の内容を管理者に問い合わせてください | " + err.Error()
		return c.Render(http.StatusBadRequest, "userNew.html", user)
	}

	return c.Render(http.StatusTemporaryRedirect, "userNewEnd.html", "")
}