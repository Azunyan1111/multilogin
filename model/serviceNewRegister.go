package model

import (
	"github.com/Azunyan1111/multilogin/mysql"
	"github.com/Azunyan1111/multilogin/structs"
	"github.com/google/uuid"
)

func RegisterNewService(service structs.Service) error {
	service.Token = uuid.New().String()
	service.Secret = uuid.New().String()
	uid, err := mysql.InsertService(service)
	if err != nil {
		return err
	}

	if service.Url != "" {
		if err := mysql.UpdateServiceUrlByUid(uid, service.Url); err != nil {
			return err
		}
	}
	if service.CallbackUrl != "" {
		if err := mysql.UpdateServiceCallbackUrlByUid(uid, service.CallbackUrl); err != nil {
			return err
		}
	}

	if service.UserName {
		if err := mysql.UpdateServiceNameByUid(uid, service.UserName); err != nil {
			return err
		}
	}
	if service.UserEmail {
		if err := mysql.UpdateServiceEmailByUid(uid, service.UserEmail); err != nil {
			return err
		}
	}
	if service.UserImage {
		if err := mysql.UpdateServiceImageByUid(uid, service.UserImage); err != nil {
			return err
		}
	}
	if service.UserAge {
		if err := mysql.UpdateServiceAgeByUid(uid, service.UserAge); err != nil {
			return err
		}
	}
	if service.UserBirthday {
		if err := mysql.UpdateServiceBirthdayByUid(uid, service.UserBirthday); err != nil {
			return err
		}
	}
	if service.UserPhone {
		if err := mysql.UpdateServicePhoneByUid(uid, service.UserPhone); err != nil {
			return err
		}
	}
	if service.UserAddress {
		if err := mysql.UpdateServiceAddressByUid(uid, service.UserAddress); err != nil {
			return err
		}
	}

	return nil
}
