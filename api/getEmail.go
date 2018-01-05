package api

import (
	"github.com/labstack/echo"
	"net/http"
)

// http://localhost:8040/api/user/email?uuid=26d2983e-3d5a-421c-bf6f-d4608025e555
func GetEmail(c echo.Context) error {
	var res EmailResponse

	service, user, response := apiTemplate(c)
	if response.StatusCode != 0 {
		c.JSON(response.StatusCode, EmailResponse{JsonResponse: response})
	}

	if service.UserEmail != 1 {
		res.JsonResponse.StatusCode = http.StatusBadRequest
		res.JsonResponse.Message = "Error: Your service does not have GetEmail authority."
		return c.JSON(http.StatusBadRequest, res)
	}

	res.JsonResponse.StatusCode = http.StatusOK
	res.JsonResponse.Message = "ok"
	res.Email = user.Email
	return c.JSON(http.StatusOK, res)
}
