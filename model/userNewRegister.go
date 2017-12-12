package model

import (
	"github.com/Azunyan1111/multilogin/mysql"
	"github.com/Azunyan1111/multilogin/structs"
)

func RegisterNewUser(user structs.User)(string,error) {
	uid, err := mysql.InsertUser(user)
	if err != nil {
		return "",err
	}
	if user.Image != "" {
		if err := mysql.UpdateUserImage(uid, user.Image); err != nil {
			return uid, err
		}
	}
	if user.Age != "" {
		if err := mysql.UpdateUserAge(uid, user.Age); err != nil {
			return uid, err
		}
	}
	if user.Birthday != "" {
		if err := mysql.UpdateUserBirthday(uid, user.Birthday); err != nil {
			return uid, err
		}
	}
	if user.Phone != "" {
		if err := mysql.UpdateUserPhone(uid, user.Phone); err != nil {
			return uid, err
		}
	}
	if user.Address != "" {
		if err := mysql.UpdateUserAddress(uid, user.Address); err != nil {
			return uid, err
		}
	}
	return uid, err
}
