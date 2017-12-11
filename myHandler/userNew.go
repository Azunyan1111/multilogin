package myHandler

import (
	"fmt"
	"github.com/Azunyan1111/multilogin/model"
	"github.com/labstack/echo"
	"net/http"
)

func GetUserNew(c echo.Context) error {
	csrf := fmt.Sprintf("%v", c.Get("csrf"))
	return c.Render(http.StatusOK, "userNew.html", model.UserNewPage{Csrf: csrf})
}

func PostUserNew(c echo.Context) error {
	csrf := fmt.Sprintf("%v", c.Get("csrf"))
	user := model.UserNewPage{Csrf: csrf}
	user.User.Email = c.FormValue("InputEmail")
	user.User.UserName = c.FormValue("InputUserName")
	user.User.Image = c.FormValue("InputImage")
	user.User.Age = c.FormValue("InputAge")
	user.User.Birthday = c.FormValue("InputBirthday")
	user.User.Phone = c.FormValue("InputPhone")
	user.User.Address = c.FormValue("InputAddress")

	// 一つでも入ってない場合はまた入力させる
	if val := user.User.Email; val == "" {
		user.Message = "メールアドレス"
		return c.Render(http.StatusBadRequest, "userNew.html", user)
	}
	if val := user.User.UserName; val == "" {
		user.Message = "ユーザー名"
		return c.Render(http.StatusBadRequest, "userNew.html", user)
	}
	if err := model.RegisterNewUser(user.User); err != nil{

	}

	return c.Render(http.StatusTemporaryRedirect, "userNewEnd.html", "")
}

// 非同期の並行処理でインサートする
func UpdateUserData(user model.User){

}