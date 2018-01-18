package myHandler

import (
	"fmt"
	"github.com/Azunyan1111/multilogin/mysql"
	"github.com/Azunyan1111/multilogin/structs"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"net/http"
)

func GetConfirmedNew(c echo.Context) error {
	orm := mysql.GetOrm()
	// サービス存在確認
	serviceUid := c.Param("serviceUid")

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
	if len(userUid) < 6 {
		return c.Render(http.StatusTemporaryRedirect, "userNewForConfirmed.html", nil)
	}

	// サービス情報取得
	var service structs.Service
	orm.Find(&service, "uuid = ?", serviceUid)
	if service.ID == 0 {
		return c.Render(http.StatusBadRequest, "error.html", structs.Error{StatusCode: http.StatusBadRequest,
			Message: "連携するサービス情報が存在しないか、URLが間違っています。"})
	}
	csrf := fmt.Sprintf("%v", c.Get("csrf"))

	var confPage structs.ConfirmedPage
	confPage.Csrf = csrf
	confPage.Service = service
	return c.Render(http.StatusOK, "confirmedNew.html", confPage)
}

func GetConfirmedPost(c echo.Context) error {
	// サービス存在確認
	serviceUid := c.Param("serviceUid")

	// サービス情報取得
	orm := mysql.GetOrm()
	var service structs.Service
	orm.Find(&service, "uuid = ?", serviceUid)
	if service.ID == 0 {
		return c.Render(http.StatusBadRequest, "error.html", structs.Error{StatusCode: http.StatusBadRequest,
			Message: "連携するサービス情報が存在しないか、URLが間違っています。"})
	}

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
	if len(userUid) < 6 {
		return c.Render(http.StatusBadRequest, "error.html", structs.Error{StatusCode: http.StatusBadRequest,
			Message: "連携するにマルチログインにログインしてください"})
	}
	// ユーザーか確認
	var user structs.User
	if orm.First(&user, "uuid = ?", userUid).RowsAffected != 1 {
		return c.Render(http.StatusBadRequest, "error.html", structs.Error{StatusCode: http.StatusBadRequest,
			Message: "サービス管理者は連携することができません。ユーザーでログインしてください。"})
	}

	// 連携情報登録
	var confirmedService structs.ConfirmedService
	confirmedService.ServiceUid = serviceUid
	confirmedService.UserUid = userUid
	orm.NewRecord(&confirmedService)
	orm.Create(&confirmedService)
	// 連携登録確認
	var confirmedServiceCheck structs.ConfirmedService
	orm.Find(&confirmedServiceCheck, "user_uuid = ? and service_uuid = ?", userUid, serviceUid)

	if confirmedService.ID == 0 {
		return c.Render(http.StatusBadRequest, "error.html", structs.Error{StatusCode: http.StatusInternalServerError,
			Message: "連携データを正しく登録することができませんでした。再試行してください。"})
	}
	var ser structs.ServiceMyPage
	ser.Service = service
	ser.Service.CallbackUrl += "?uuid=" + userUid
	return c.Render(http.StatusOK, "confirmedEnd.html", ser)
}

func PostConfirmedDelete(c echo.Context) error {
	// サービス名のuid
	serviceUid := c.Param("serviceUid")
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
	if len(userUid) < 6 {
		return c.Render(http.StatusBadRequest, "error.html", structs.Error{StatusCode: http.StatusBadRequest,
			Message: "連携を解除する前にマルチログインにログインしてください"})
	}

	// サービス情報取得
	orm := mysql.GetOrm()
	var service structs.Service
	orm.Find(&service, "uuid = ?", serviceUid)
	if service.ID == 0 {
		return c.Render(http.StatusBadRequest, "error.html", structs.Error{StatusCode: http.StatusBadRequest,
			Message: "連携するサービス情報が存在しないか、リクエストが不正です。"})
	}
	// 連携解除
	var confirmedService structs.ConfirmedService
	orm.Find(&confirmedService, "user_uuid = ? and service_uuid = ?", userUid, serviceUid)
	if orm.Delete(&confirmedService).RowsAffected != 1 {
		return c.Render(http.StatusInternalServerError, "error.html", structs.Error{StatusCode: http.StatusBadRequest,
			Message: "正常に連携解除する事ができませんでした。"})
	}
	// TODO:ここで削除した趣旨のメッセージを表示したいよね。めんどくさいけど。
	return c.Redirect(http.StatusTemporaryRedirect, "/user/mypage")
}
