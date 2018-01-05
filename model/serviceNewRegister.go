package model

import (
	"github.com/Azunyan1111/multilogin/mysql"
	"github.com/Azunyan1111/multilogin/structs"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func RegisterNewService(service structs.Service) (string, error) {
	service.Uid = uuid.New().String()
	service.Token = uuid.New().String()
	service.Secret = uuid.New().String()

	orm := mysql.GetOrm()
	orm.NewRecord(&service)
	if orm.Create(&service).RowsAffected != 1 {
		return "", errors.New("Insert Error:RegisterNewService")
	}
	return service.Uid, nil
}
