package api

import (
	"github.com/Azunyan1111/multilogin/structs"
	"github.com/labstack/echo"
	"github.com/Azunyan1111/multilogin/mysql"
	"net/http"
)

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
	// ユーザーが認証しているかを確認する
	var con structs.ConfirmedService
	if orm.Find(&con,"user_uuid = ? and service_uuid = ?",userUid,serviceUid).RowsAffected != 1{
		res.JsonResponse.StatusCode = http.StatusBadRequest
		res.JsonResponse.Message = "Error: This user is not working with your service."
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
