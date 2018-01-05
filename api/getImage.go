package api

import (
	"github.com/labstack/echo"
	"net/http"
)

// http://localhost:8040/api/user/image?uuid=26d2983e-3d5a-421c-bf6f-d4608025e555
func GetImage(c echo.Context) error {
	// response struct
	var res ImageResponse

	service, user, response := apiTemplate(c)
	if response.StatusCode != 0{
		c.JSON(response.StatusCode,NameResponse{JsonResponse:response})
	}

	// 権限を確認する
	if service.UserImage != 1{
		res.JsonResponse.StatusCode = http.StatusBadRequest
		res.JsonResponse.Message = "Error: Your service does not have GetImage authority."
		return c.JSON(http.StatusBadRequest,res)
	}

	// 権限があるのでユーザーの情報を返す
	res.JsonResponse.StatusCode = http.StatusOK
	res.JsonResponse.Message = "ok"
	res.Image = user.Image
	return c.JSON(http.StatusOK,res)
}