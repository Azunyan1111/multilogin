package myHandler

import (
	"fmt"
	"github.com/Azunyan1111/multilogin/structs"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"net/http"
	"github.com/Azunyan1111/multilogin/mysql"
	"log"
)

func GetConfirmedNew(c echo.Context) error {
	// サービス存在確認
	serviceUid := c.Param("serviceUid")

	// セッション確認
	s, err := session.Get("session", c)
	if err != nil{
		panic(err)
	}
	var uid string
	if s != nil{
		uid = fmt.Sprintf("%v", s.Values["uid"])
	}
	log.Println(uid)
	if len(uid) < 6{
		return c.Render(http.StatusBadRequest, "error.html",structs.Error{StatusCode:http.StatusBadRequest,
		Message:"連携するにマルチログインにログインしてください"})
	}

	// サービス情報取得
	orm := mysql.GetOrm()
	var service structs.Service
	orm.LogMode(true)
	orm.Find(&service, "uuid = ?",serviceUid)
	log.Println(service)
	if service.ID == 0{
		return c.Render(http.StatusBadRequest, "error.html",structs.Error{StatusCode:http.StatusBadRequest,
		Message:"連携するサービス情報が存在しないか、URLが間違っています。"})
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
	if !mysql.IsServiceByUuid(serviceUid){
		log.Println("Not Serivce")
		return c.Render(http.StatusBadRequest, "error.html", structs.Error{StatusCode:http.StatusBadRequest, Message:"Serviced not found"})
	}
	return c.Render(http.StatusBadRequest, "error.html", structs.Error{StatusCode:http.StatusBadRequest, Message:"Serviced not found"})

	/*/
	// セッション確認
	s := session.Default(c)
	var uid string
	if s != nil{
		log.Println("nil session")
		uid = fmt.Sprintf("%v", s.Get("uid"))
	}
	if len(uid) < 5{
		return c.Render(http.StatusBadRequest, "error.html", structs.Error{StatusCode:http.StatusBadRequest, Message:"Not Login. Please login."})
	}

	// 連携情報登録
	err := mysql.InsertConfirmedByUidAndUid(uid,serviceUid)
	if err != nil{
		return c.Render(http.StatusBadRequest, "error.html", structs.Error{StatusCode:http.StatusInternalServerError, Message:"Insert Not success." + err.Error()})
	}
	// サービスのコールバックURLへ飛ばす
	return c.Render(http.StatusOK, "confirmedNew.html", nil)
	//*/
}
