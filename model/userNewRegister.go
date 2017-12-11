package model

import (
	"github.com/Azunyan1111/multilogin/structs"
	"github.com/Azunyan1111/multilogin/mysql"
)

func RegisterNewUser(user structs.User)error{
	uid, err := mysql.InsertUser(user)
	if err != nil{
		return err
	}
	if user.Image != ""{
		if err := mysql.UpdateUserImage(uid, user.Image); err != nil{
			return err
		}
	}
	if user.Age != ""{
		if err := mysql.UpdateUserAge(uid, user.Age); err != nil{
			return err
		}
	}
	if user.Birthday != ""{
		if err := mysql.UpdateUserBirthday(uid,user.Birthday); err != nil{
			return err
		}
	}
	if user.Phone != ""{
		if err := mysql.UpdateUserPhone(uid,user.Phone); err != nil{
			return err
		}
	}
	if user.Address != ""{
		if err := mysql.UpdateUserAddress(uid, user.Address); err != nil{
			return err
		}
	}
	return nil
}