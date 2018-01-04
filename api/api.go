package api

import (
	"net/http"
	"github.com/labstack/echo"
	"strconv"
)

// http://localhost:8040/api/sum?value1=1&value2=1
func Sum(c echo.Context) error {
	value1 := c.QueryParam("value1")
	value2 := c.QueryParam("value2")

	var res SumResponse

	valueInt1, err := strconv.Atoi(value1)
	if err != nil{
		res.JsonResponse.StatusCode = http.StatusInternalServerError
		res.JsonResponse.Message = "Error: Request Value1 is not number."
		c.JSON(http.StatusInternalServerError,res)
	}
	valueInt2, err := strconv.Atoi(value2)
	if err != nil{
		res.JsonResponse.StatusCode = http.StatusInternalServerError
		res.JsonResponse.Message = "Error: Request Value2 is not number."
		c.JSON(http.StatusInternalServerError,res)
	}
	value := valueInt1 + valueInt2
	res.JsonResponse.StatusCode = http.StatusOK
	res.JsonResponse.Message = "ok"
	res.Value = value
	c.Response().Header().Set("Authorization","test")
	return c.JSON(http.StatusOK, res)
}


// TODO:認証を作る。ヘッダーの中身全部打ち込まれるからスプリットしよう
func IsBadSignature(authorization string)(bool){
	return false
}
// TODO:authorizationにはアクセストークンが入っているので、アクセストークンからserviceUUIDを取得しよう
func GetServiceUid(authorization string)(string){
	return "124ah368-1eha-7h81-2345-365a24h6522y"
}