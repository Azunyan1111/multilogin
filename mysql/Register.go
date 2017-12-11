package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"os"
	"database/sql"
	"github.com/google/uuid"
	"github.com/Azunyan1111/multilogin/structs"
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

func SelectUserByUuid(uuid string)(structs.User, error){
	row := MyDB.QueryRow("select * from users where uuid = ?;", uuid)
	var dataBaseId int
	var user structs.User
	if err := row.Scan(&dataBaseId, &user.Uid, &user.UserName, &user.Image, &user.Age, &user.Birthday, &user.Email,
		&user.EmailOK, &user.Phone, &user.PhoneOK, &user.Address, &user.CreatedAt, &user.UpdatedAt); err != nil{
		return structs.User{}, err
	}
	return user, nil
}

func InsertUser(user structs.User)(string, error){
	uid := uuid.New()
	_, err := MyDB.Exec("INSERT INTO `users` (`uuid`, `user`, `email`) VALUES (?, ?, ?);",uid,user.UserName, user.Email)
	if err != nil {
		return "", err
	}
	return uid.String(),nil
}

func DeleteUserByUid(uid string)error{
	_, err := MyDB.Exec("DELETE FROM users WHERE uuid = ?;", uid)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUserByTestUser()error{
	_, err := MyDB.Exec("DELETE FROM users WHERE user = 'TestUser114514';")
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserName(uid string,userName string)error{
	_, err := MyDB.Exec("UPDATE users SET user = ? WHERE uuid = ?;",userName, uid)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserImage(uid string,image string)error{
	_, err := MyDB.Exec("UPDATE users SET image = ? WHERE uuid = ?;",image, uid)
	if err != nil {
		return err
	}
	return nil
}