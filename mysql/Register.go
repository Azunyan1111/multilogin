package mysql

import (
	"database/sql"
	"github.com/Azunyan1111/multilogin/structs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"os"
)

var MyDB *sql.DB

func DataBaseInit() error {
	dataSource := os.Getenv("DOCKER_DATABASE_URL")
	var err error
	MyDB, err = sql.Open("mysql", dataSource)
	if err != nil {
		return err
	}
	return nil
}

// Select文を実行するとRowsAffectedに0が帰って来る
func ConnectionTest() (int64, error) {
	row, err := MyDB.Exec("select * from users;")
	if err != nil {
		return -1, err
	}
	ra, err := row.RowsAffected()
	if err != nil {
		return -1, err
	}
	return ra, nil
}

func SelectUserByUuid(uuid string) (structs.User, error) {
	row := MyDB.QueryRow("select * from users where uuid = ?;", uuid)
	var dataBaseId int
	var user structs.User
	if err := row.Scan(&dataBaseId, &user.Uid, &user.UserName, &user.Image, &user.Age, &user.Birthday, &user.Email,
		&user.EmailOK, &user.Phone, &user.PhoneOK, &user.Address, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return structs.User{}, err
	}
	return user, nil
}

func InsertUser(user structs.User) (string, error) {
	uid := uuid.New()
	_, err := MyDB.Exec("INSERT INTO `users` (`uuid`, `user`, `email`) VALUES (?, ?, ?);", uid, user.UserName, user.Email)
	if err != nil {
		return "", err
	}
	return uid.String(), nil
}

func InsertService(service structs.Service) (string, error) {
	uid := uuid.New()
	_, err := MyDB.Exec("INSERT INTO `service` (`uuid`, `name`, `email`, `token`, `secret`) VALUES (?, ?, ?,?,?);", uid, service.ServiceName, service.Email, service.Token, service.Secret)
	if err != nil {
		return "", err
	}
	return uid.String(), nil
}

func DeleteUserByUid(uid string) error {
	_, err := MyDB.Exec("DELETE FROM users WHERE uuid = ?;", uid)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUserByTestUser() error {
	_, err := MyDB.Exec("DELETE FROM users WHERE user = 'TestUser114514';")
	if err != nil {
		return err
	}
	return nil
}
func DeleteUserByTestService() error {
	_, err := MyDB.Exec("DELETE FROM service WHERE name = 'TestUser114514';")
	if err != nil {
		return err
	}
	return nil
}
func SelectUserByTestUser() (structs.User, error) {
	row := MyDB.QueryRow("select * from users where user = 'TestUser114514';")
	var dataBaseId int
	var user structs.User
	if err := row.Scan(&dataBaseId, &user.Uid, &user.UserName, &user.Image, &user.Age, &user.Birthday, &user.Email,
		&user.EmailOK, &user.Phone, &user.PhoneOK, &user.Address, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return structs.User{}, err
	}
	return user, nil
}

func SelectUserByTestService() (structs.Service, error) {
	row := MyDB.QueryRow("select * from service where name = 'TestUser114514';")
	var dataBaseId int
	var user structs.Service
	if err := row.Scan(&dataBaseId, &user.Uid, &user.ServiceName, &user.Email, &user.Url, &user.CallbackUrl, &user.Token,
		&user.Secret, &user.UserName, &user.UserImage, &user.UserAge, &user.UserBirthday, &user.UserEmail,
		&user.UserPhone, &user.UserAddress, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return structs.Service{}, err
	}
	return user, nil
}

func UpdateUserName(uid string, userName string) error {
	_, err := MyDB.Exec("UPDATE users SET user = ? WHERE uuid = ?;", userName, uid)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserImage(uid string, image string) error {
	_, err := MyDB.Exec("UPDATE users SET image = ? WHERE uuid = ?;", image, uid)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserAge(uid string, age string) error {
	_, err := MyDB.Exec("UPDATE users SET age = ? WHERE uuid = ?;", age, uid)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserBirthday(uid string, birthday string) error {
	_, err := MyDB.Exec("UPDATE users SET birthday = ? WHERE uuid = ?;", birthday, uid)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserEmail(uid string, email string) error {
	_, err := MyDB.Exec("UPDATE users SET email = ? WHERE uuid = ?;", email, uid)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserEmailOk(uid string, emailOk bool) error {
	_, err := MyDB.Exec("UPDATE users SET emailOk = ? WHERE uuid = ?;", emailOk, uid)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserPhone(uid string, phone string) error {
	_, err := MyDB.Exec("UPDATE users SET phone = ? WHERE uuid = ?;", phone, uid)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserPhoneOk(uid string, phoneOk bool) error {
	_, err := MyDB.Exec("UPDATE users SET phoneOk = ? WHERE uuid = ?;", phoneOk, uid)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserAddress(uid string, address string) error {
	_, err := MyDB.Exec("UPDATE users SET address = ? WHERE uuid = ?;", address, uid)
	if err != nil {
		return err
	}
	return nil
}

// Service
func UpdateServiceUrlByUid(uid string, url string) error {
	_, err := MyDB.Exec("UPDATE service SET url = ? WHERE uuid = ?;", url, uid)
	if err != nil {
		return err
	}
	return nil
}
func UpdateServiceCallbackUrlByUid(uid string, callbackUrl string) error {
	_, err := MyDB.Exec("UPDATE service SET url_callback = ? WHERE uuid = ?;", callbackUrl, uid)
	if err != nil {
		return err
	}
	return nil
}
func UpdateServiceNameByUid(uid string, name bool) error {
	_, err := MyDB.Exec("UPDATE service SET p_name = ? WHERE uuid = ?;", name, uid)
	if err != nil {
		return err
	}
	return nil
}
func UpdateServiceImageByUid(uid string, image bool) error {
	_, err := MyDB.Exec("UPDATE service SET p_image = ? WHERE uuid = ?;", image, uid)
	if err != nil {
		return err
	}
	return nil
}
func UpdateServiceAgeByUid(uid string, age bool) error {
	_, err := MyDB.Exec("UPDATE service SET p_age = ? WHERE uuid = ?;", age, uid)
	if err != nil {
		return err
	}
	return nil
}
func UpdateServiceBirthdayByUid(uid string, birthday bool) error {
	_, err := MyDB.Exec("UPDATE service SET p_birthday = ? WHERE uuid = ?;", birthday, uid)
	if err != nil {
		return err
	}
	return nil
}
func UpdateServiceEmailByUid(uid string, email bool) error {
	_, err := MyDB.Exec("UPDATE service SET p_email = ? WHERE uuid = ?;", email, uid)
	if err != nil {
		return err
	}
	return nil
}
func UpdateServicePhoneByUid(uid string, phone bool) error {
	_, err := MyDB.Exec("UPDATE service SET p_phone = ? WHERE uuid = ?;", phone, uid)
	if err != nil {
		return err
	}
	return nil
}
func UpdateServiceAddressByUid(uid string, address bool) error {
	_, err := MyDB.Exec("UPDATE service SET p_address = ? WHERE uuid = ?;", address, uid)
	if err != nil {
		return err
	}
	return nil
}
