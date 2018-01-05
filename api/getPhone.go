package api

import (
	"github.com/labstack/echo"
	"net/http"
)

// http://localhost:8040/api/user/phone?uuid=26d2983e-3d5a-421c-bf6f-d4608025e555
func GetPhone(c echo.Context) error {
	var res PhoneResponse

	service, user, response := apiTemplate(c)
	if response.StatusCode != 0{
		c.JSON(response.StatusCode,PhoneResponse{JsonResponse:response})
	}

	if service.UserPhone != 1{
		res.JsonResponse.StatusCode = http.StatusBadRequest
		res.JsonResponse.Message = "Error: Your service does not have GetPhone authority."
		return c.JSON(http.StatusBadRequest,res)
	}

	res.JsonResponse.StatusCode = http.StatusOK
	res.JsonResponse.Message = "ok"
	res.Phone = user.Phone
	return c.JSON(http.StatusOK,res)
}
