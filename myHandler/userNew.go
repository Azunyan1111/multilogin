package myHandler

import (
	"fmt"
	"github.com/Azunyan1111/multilogin/model"
	"github.com/Azunyan1111/multilogin/structs"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"net/http"
)

func GetUserNew(c echo.Context) error {
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
			Message: "すでにログインしています。動作が不安定な場合はブラウザのクッキーを削除してください。"})
	}

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
	uid, err := model.RegisterNewUser(user.User)
	if err != nil {
		user.Message = "サーバーエラーです | 次の内容を管理者に問い合わせてください | " + err.Error()
		return c.Render(http.StatusInternalServerError, "userNew.html", user)
	}

	// セッションに自分のuuidをつけて返す
	// セッション確認
	s, err := session.Get("session", c)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.html", structs.Error{StatusCode: http.StatusInternalServerError,
			Message: "サーバーの処理に異常がありました。エラーコード:1001"})
	}
	s.Values["uid"] = uid
	s.Save(c.Request(), c.Response().Writer)

	return c.Render(http.StatusTemporaryRedirect, "userNewEnd.html", "")
}
