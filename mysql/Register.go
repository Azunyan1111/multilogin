package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/Azunyan1111/multilogin/model"
	"os"
	"database/sql"
	"github.com/google/uuid"
)

var MyDB *sql.DB

func DataBaseInit() error{
	dataSource := os.Getenv("DOCKER_DATABASE_URL")
	var err error
	MyDB, err = sql.Open("mysql", dataSource)
	if err != nil {
		return err
	}
	return nil
}

// Select文を実行するとRowsAffectedに0が帰って来る
func ConnectionTest()(int64, error){
	row, err := MyDB.Exec("select * from users;")
	if err != nil {
		return -1 , err
	}
	ra, err := row.RowsAffected()
	if err != nil {
		return -1 , err
	}
	return ra, nil
}

func SelectUserByUuid(uuid string)(model.User, error){
	row := MyDB.QueryRow("select * from users where uuid = ?;", uuid)
	var dataBaseId int
	var user model.User
	if err := row.Scan(&dataBaseId, &user.Uid, &user.UserName, &user.Image, &user.Age, &user.Birthday, &user.Email,
		&user.EmailOK, &user.Phone, &user.PhoneOK, &user.Address, &user.CreatedAt, &user.UpdatedAt); err != nil{
		return model.User{}, err
	}
	return user, nil
}

func InsertUser(user model.User)(string, error){
	uid := uuid.New()
	_, err := MyDB.Exec("INSERT INTO `users` (`uuid`, `user`, `email`) VALUES (?, ?, ?);",uid,user.UserName, user.Email)
	if err != nil {
		return "", err
	}
	return uid.String(),nil
}

func DeleteUser(uid string)error{
	_, err := MyDB.Exec("DELETE FROM users WHERE uuid = ?;", uid)
	if err != nil {
		return err
	}
	return nil
}