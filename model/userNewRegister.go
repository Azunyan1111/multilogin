package model

import (
	"github.com/Azunyan1111/multilogin/structs"
	"github.com/Azunyan1111/multilogin/mysql"
)

func RegisterNewUser(user structs.User)error{
	//*/
	uid, err := mysql.InsertUser(user)
	if err != nil{
		return err
	}
	if err := mysql.UpdateUserName(user.UserName, uid); err != nil{
		return err
	}
	if err := mysql.UpdateUserImage(user.Image, uid); err != nil{
		return err
	}
	//*/
	return nil
}