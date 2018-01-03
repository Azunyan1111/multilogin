package api

import (
	"net/http"
	"github.com/labstack/echo"
	"strconv"
)

type JsonResponse struct {
	StatusCode int	`json:"StatusCode"`
	Message string	`json:"Message"`
}

type sum struct {
	JsonResponse JsonResponse `json:"status"`
	Value        int          `json:"Value"`
}
// http://localhost:8040/api/sum?value1=1&value2=1
func Sum(c echo.Context) error {
	value1 := c.QueryParam("value1")
	value2 := c.QueryParam("value2")

	var res sum

	valueInt1, err := strconv.Atoi(value1)
	if err != nil{
		res.JsonResponse.StatusCode = http.StatusInternalServerError
		res.JsonResponse.Message = "Error Request Value is not number."
		c.JSON(http.StatusInternalServerError,res)
	}
	valueInt2, err := strconv.Atoi(value2)
	if err != nil{
		res.JsonResponse.StatusCode = http.StatusInternalServerError
		res.JsonResponse.Message = "Error Request Value is not number."
		c.JSON(http.StatusInternalServerError,res)
	}
	value := valueInt1 + valueInt2
	res.JsonResponse.StatusCode = http.StatusOK
	res.JsonResponse.Message = "ok"
	res.Value = value
	c.Response().Header().Set("Authorization","test")
	return c.JSON(http.StatusOK, res)
}