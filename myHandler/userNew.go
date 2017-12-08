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
	user.NewUser.Email = c.FormValue("InputEmail")
	user.NewUser.UserName = c.FormValue("InputUserName")
	user.NewUser.Image = c.FormValue("InputImage")
	user.NewUser.Age = c.FormValue("InputAge")
	user.NewUser.Birthday = c.FormValue("InputBirthday")
	user.NewUser.Phone = c.FormValue("InputPhone")
	user.NewUser.Address = c.FormValue("InputAddress")

	// 一つでも入ってない場合
	if val := user.NewUser.Email; val == "" {
		user.Message = "メールアドレス"
		return c.Render(http.StatusBadRequest, "userNew.html", user)
	}
	if val := user.NewUser.UserName; val == "" {
		user.Message = "ユーザー名"
		return c.Render(http.StatusBadRequest, "userNew.html", user)
	}
	if val := user.NewUser.Image; val == "" {
		user.Message = "サムネイル"
		return c.Render(http.StatusBadRequest, "userNew.html", user)
	}
	if val := user.NewUser.Age; val == "" {
		user.Message = "年齢"
		return c.Render(http.StatusBadRequest, "userNew.html", user)
	}
	if val := user.NewUser.Birthday; val == "" {
		user.Message = "生年月日"
		return c.Render(http.StatusBadRequest, "userNew.html", user)
	}
	if val := user.NewUser.Phone; val == "" {
		user.Message = "電話番号"
		return c.Render(http.StatusBadRequest, "userNew.html", user)
	}
	if val := user.NewUser.Address; val == "" {
		user.Message = "住所"
		return c.Render(http.StatusBadRequest, "userNew.html", user)
	}

	return c.Render(http.StatusTemporaryRedirect, "userNewEnd.html", "")
}
