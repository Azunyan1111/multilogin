package model

import (
	"github.com/Azunyan1111/multilogin/mysql"
	"github.com/Azunyan1111/multilogin/structs"
	"log"
)

func RegisterNewUser(user structs.Usered)(string,error) {
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

func UpdateNewUser(user structs.Usered)(error) {
	if user.UserName != "" {
		if err := mysql.UpdateUserName(user.Uid, user.UserName); err != nil {
			return  err
		}
	}
	if user.Email != "" {
		if err := mysql.UpdateUserEmail(user.Uid, user.Email); err != nil {
			return  err
		}
	}
	if user.Image != "" {
		if err := mysql.UpdateUserImage(user.Uid, user.Image); err != nil {
			return  err
		}
	}
	if user.Age != "" {
		log.Println(user.Uid, "aaa" , user.Age)
		if err := mysql.UpdateUserAge(user.Uid, user.Age); err != nil {
			log.Println(err)
			return  err
		}
	}
	if user.Birthday != "" {
		if err := mysql.UpdateUserBirthday(user.Uid, user.Birthday); err != nil {
			return  err
		}
	}
	if user.Phone != "" {
		if err := mysql.UpdateUserPhone(user.Uid, user.Phone); err != nil {
			return  err
		}
	}
	if user.Address != "" {
		if err := mysql.UpdateUserAddress(user.Uid, user.Address); err != nil {
			return  err
		}
	}
	return nil
}