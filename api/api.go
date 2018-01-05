package api

import (
	"github.com/Azunyan1111/multilogin/model"
	"github.com/Azunyan1111/multilogin/mysql"
	"github.com/Azunyan1111/multilogin/structs"
	"github.com/labstack/echo"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// http://localhost:8040/api/sum?value1=1&value2=1
func Sum(c echo.Context) error {
	value1 := c.QueryParam("value1")
	value2 := c.QueryParam("value2")

	var res SumResponse

	valueInt1, err := strconv.Atoi(value1)
	if err != nil {
		res.JsonResponse.StatusCode = http.StatusInternalServerError
		res.JsonResponse.Message = "Error: Request Value1 is not number."
		c.JSON(http.StatusInternalServerError, res)
	}
	valueInt2, err := strconv.Atoi(value2)
	if err != nil {
		res.JsonResponse.StatusCode = http.StatusInternalServerError
		res.JsonResponse.Message = "Error: Request Value2 is not number."
		c.JSON(http.StatusInternalServerError, res)
	}
	value := valueInt1 + valueInt2
	res.JsonResponse.StatusCode = http.StatusOK
	res.JsonResponse.Message = "ok"
	res.Value = value
	return c.JSON(http.StatusOK, res)
}

func apiTemplate(c echo.Context) (structs.Service, structs.User, JsonResponse) {
	orm := mysql.GetOrm()
	var service structs.Service
	var user structs.User
	var jsonResponse JsonResponse

	// パラメーターチェック
	userUid := c.QueryParam("uuid")
	if userUid == "" {
		jsonResponse.StatusCode = http.StatusBadRequest
		jsonResponse.Message = "Error: Request uuid not found. example url '/api/name?uuid=26d2983e-3d5a-421c-bf6f-d4608025e555'"
		return service, user, jsonResponse
	}
	// リクエストからサービス情報を取得する
	if orm.Find(&service, "uuid = ?", GetServiceUid(c.Request().Header.Get("Authorization"))).RowsAffected != 1 {
		jsonResponse.StatusCode = http.StatusBadRequest
		jsonResponse.Message = "Error: Your service is not registered. " +
			"Also, there is no service ID in the 'Authorization' header of the request"
		return service, user, jsonResponse
	}
	// リクエストの署名を確認する
	if IsBadSignature(c.Request().URL, c.Request().Header.Get("Authorization"), service) {
		jsonResponse.StatusCode = http.StatusBadRequest
		jsonResponse.Message = "Error: Request signature is bad."
		return service, user, jsonResponse
	}
	// ユーザーの情報を取得する
	if orm.Find(&user, "uuid = ?", userUid).RowsAffected != 1 {
		jsonResponse.StatusCode = http.StatusBadRequest
		jsonResponse.Message = "Error: The specified user does not exist"
		return service, user, jsonResponse
	}
	// ユーザーが認証しているかを確認する
	var con structs.ConfirmedService
	if orm.Find(&con, "user_uuid = ? and service_uuid = ?", userUid, service.Uid).RowsAffected != 1 {
		jsonResponse.StatusCode = http.StatusBadRequest
		jsonResponse.Message = "Error: This user is not working with your service."
		return service, user, jsonResponse
	}
	return service, user, jsonResponse
}

// TODO:認証を作る。ヘッダーの中身全部打ち込まれるからスプリットしよう
func IsBadSignature(url *url.URL, authorization string, service structs.Service) bool {
	var t string
	var h string
	authorizations := strings.Split(authorization, ",")
	for _, s := range authorizations {
		if strings.HasPrefix(s, "time=") {
			t = strings.Split(s, "time=")[1]
		}
	}
	for _, s := range authorizations {
		if strings.HasPrefix(s, "signature=") {
			h = strings.Split(s, "signature=")[1]
		}
	}
	join := url.String() + "," + t
	if model.CheckHmac(service.Secret, join, h) {
		return false
	} else {
		return true
	}
}

// TODO:authorizationにはアクセストークンが入っているので、アクセストークンからserviceUUIDを取得しよう
func GetServiceUid(authorization string) string {
	var token string
	authorizations := strings.Split(authorization, ",")
	for _, s := range authorizations {
		if strings.HasPrefix(s, "token=") {
			token = s
		}
	}
	if token == "" {
		return ""
	}
	var service structs.Service
	orm := mysql.GetOrm()
	if orm.Find(&service, "token = ?", strings.Split(token, "token=")[1]).RowsAffected != 1 {
		return ""
	}
	return service.Uid
}
