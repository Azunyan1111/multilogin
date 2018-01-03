package api

import (
	"net/http"
	"github.com/labstack/echo"
	"strconv"
	"github.com/Azunyan1111/multilogin/structs"
	"github.com/Azunyan1111/multilogin/mysql"
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
// http://localhost:8040/api/user/name?uuid=26d2983e-3d5a-421c-bf6f-d4608025e555
func GetName(c echo.Context) error {
	// response struct
	var res NameResponse
	// orm
	orm := mysql.GetOrm()

	userUid := c.QueryParam("uuid")
	if userUid == ""{
		res.JsonResponse.StatusCode = http.StatusBadRequest
		res.JsonResponse.Message = "Error: Request uuid not found. example url '/api/name?uuid=26d2983e-3d5a-421c-bf6f-d4608025e555'"
		return c.JSON(http.StatusBadRequest,res)
	}
	if IsBadSignature(c.Request().Header.Get("Authorization")){
		res.JsonResponse.StatusCode = http.StatusBadRequest
		res.JsonResponse.Message = "Error: Request Signature is bad."
		return c.JSON(http.StatusBadRequest,res)
	}
	// 権限を確認する
	var service structs.Service
	serviceUid := GetServiceUid(c.Request().Header.Get("Authorization"))
	if orm.Find(&service,"uuid = ?",serviceUid).RowsAffected != 1{
		res.JsonResponse.StatusCode = http.StatusBadRequest
		res.JsonResponse.Message = "Error: Your service is not registered."
		return c.JSON(http.StatusBadRequest,res)
	}
	if service.UserName != 1{
		res.JsonResponse.StatusCode = http.StatusBadRequest
		res.JsonResponse.Message = "Error: Your service does not have GetName authority."
		return c.JSON(http.StatusBadRequest,res)
	}
	// 権限があるのでユーザーの情報を返す
	var user structs.User
	if orm.Find(&user,"uuid = ?",userUid).RowsAffected != 1{
		res.JsonResponse.StatusCode = http.StatusBadRequest
		res.JsonResponse.Message = "Error: The specified user does not exist"
		return c.JSON(http.StatusBadRequest,res)
	}
	res.JsonResponse.StatusCode = http.StatusOK
	res.JsonResponse.Message = "ok"
	res.Name = user.UserName
	return c.JSON(http.StatusOK,res)
}

// TODO:認証を作る。ヘッダーの中身全部打ち込まれるからスプリットしよう
func IsBadSignature(authorization string)(bool){
	return false
}
// TODO:authorizationにはアクセストークンが入っているので、アクセストークンからserviceUUIDを取得しよう
func GetServiceUid(authorization string)(string){
	return "025ad602-7dba-4c08-8226-704b65f2873c"
}