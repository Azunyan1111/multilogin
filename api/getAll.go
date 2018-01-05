package api

import (
	"github.com/labstack/echo"
	"net/http"
)

// http://localhost:8040/api/user/all?uuid=26d2983e-3d5a-421c-bf6f-d4608025e555
func GetAll(c echo.Context) error {
	// response struct
	var res AllResponse

	service, user, response := apiTemplate(c)
	if response.StatusCode != 0 {
		c.JSON(response.StatusCode, AllResponse{JsonResponse: response})
	}

	// 権限を確認する
	if service.UserName == 1 {
		res.User.UserName = user.UserName
	}
	if service.UserEmail == 1 {
		res.User.Email = user.Email
	}
	if service.UserImage == 1 {
		res.User.Image = user.Image
	}
	if service.UserAge == 1 {
		res.User.Age = user.Age
	}
	if service.UserBirthday == 1 {
		res.User.Birthday = user.Birthday
	}
	if service.UserPhone == 1 {
		res.User.Phone = user.Phone
	}
	if service.UserAddress == 1 {
		res.User.Address = user.Address
	}

	// 権限があるのでユーザーの情報を返す
	res.JsonResponse.StatusCode = http.StatusOK
	res.JsonResponse.Message = "ok"
	res.User.Uid = user.Uid
	return c.JSON(http.StatusOK, res)
}
