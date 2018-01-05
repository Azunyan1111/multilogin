package api

import (
	"github.com/labstack/echo"
	"net/http"
)

// TODO:コメントアウト
// http://localhost:8040/api/user/address?uuid=26d2983e-3d5a-421c-bf6f-d4608025e555
// TODO:メソッド名
func GetAddress(c echo.Context) error {
	// TODO:構造体
	var res AddressResponse

	service, user, response := apiTemplate(c)
	if response.StatusCode != 0{
		// TODO:構造体
		c.JSON(response.StatusCode,AddressResponse{JsonResponse:response})
	}

	// TODO:権限
	if service.UserAddress != 1{
		res.JsonResponse.StatusCode = http.StatusBadRequest
		// TODO:メッセージ
		res.JsonResponse.Message = "Error: Your service does not have GetAddress authority."
		return c.JSON(http.StatusBadRequest,res)
	}

	res.JsonResponse.StatusCode = http.StatusOK
	res.JsonResponse.Message = "ok"
	// TODO:ユーザー情報
	res.Address = user.Address
	return c.JSON(http.StatusOK,res)
}

