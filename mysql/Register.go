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

func InsertNewUser(user model.User) error{
	_, err := MyDB.Exec("INSERT INTO `users` (`uuid`, `user`, `email`) VALUES (?, ?, ?);",uuid.New(),user.UserName, user.Email)
	if err != nil {
		return err
	}
	return nil
}
