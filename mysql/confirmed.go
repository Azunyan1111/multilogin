package mysql

import (
	"github.com/Azunyan1111/multilogin/structs"
	"github.com/pkg/errors"
)

func IsServiceByUuid(uuid string) bool {
	row := MyDB.QueryRow("select count(id) from service where uuid = ?;", uuid)
	var count int
	if err := row.Scan(&count); err != nil {
		return false
	}
	if count == 1 {
		return true
	}
	return false
}

func SelectServiceByuUid(uid string) (structs.Serviced, error) {
	row := MyDB.QueryRow("select `id`, `uuid`, `name`, `email`, `url`, `url_callback`, `token`, `secret`,"+
		" `p_name`, `p_image`, `p_age`, `p_birthday`, `p_email`, `p_phone`, `p_address`, `created_at`, `updated_at` from service where uuid = ?;", uid)
	var dataBaseId int
	var user structs.Serviced
	if err := row.Scan(&dataBaseId, &user.Uid, &user.ServiceName, &user.Email, &user.Url, &user.CallbackUrl, &user.Token,
		&user.Secret, &user.UserName, &user.UserImage, &user.UserAge, &user.UserBirthday, &user.Email, &user.UserPhone,
		&user.UserAddress, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return structs.Serviced{}, err
	}
	return user, nil
}

func InsertConfirmedByUidAndUid(userUid string, serviceUid string) error {
	result, err := MyDB.Exec(`INSERT INTO 'confirmed_service'' ('user_uuid'', 'service_uuid') VALUES (?, ?);`, userUid, serviceUid)
	if err != nil {
		return err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if count != 1 {
		return errors.New("Insert Error!")
	}
	return nil
}

// Test Only
func IsNotFoundConfirmedByuUserUidAndServiceUid(userUid string, serviceUid string) (bool, error) {
	row := MyDB.QueryRow(`select count(id) from confirmed_service where user_uuid = ? && service_uuid = ?;`, userUid, serviceUid)
	var count int
	if err := row.Scan(&count); err != nil {
		return false, err
	}
	if count == 0 {
		return true, nil
	} else {
		return false, nil
	}
}
