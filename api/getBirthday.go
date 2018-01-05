package api

import (
	"github.com/labstack/echo"
	"net/http"
)

// http://localhost:8040/api/user/birthday?uuid=26d2983e-3d5a-421c-bf6f-d4608025e555
func GetBirthday(c echo.Context) error {
	var res BirthdayResponse

	service, user, response := apiTemplate(c)
	if response.StatusCode != 0{
		c.JSON(response.StatusCode,BirthdayResponse{JsonResponse:response})
	}

	if service.UserBirthday != 1{
		res.JsonResponse.StatusCode = http.StatusBadRequest
		res.JsonResponse.Message = "Error: Your service does not have GetBirthday authority."
		return c.JSON(http.StatusBadRequest,res)
	}

	res.JsonResponse.StatusCode = http.StatusOK
	res.JsonResponse.Message = "ok"
	res.Birthday = user.Birthday
	return c.JSON(http.StatusOK,res)
}
