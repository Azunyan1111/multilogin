package model

type Register interface {
	RegisterNewUser(user User) error
	UpdateNewUser(user User) error
	RegisterService(user User) error
	UpdateService(user User) error
}

type MySql interface {
	DataBaseInit() error

	// User
	InsertNewUser(user User) error
	UpdateUser(user User) error

	// Service
	InsertService() error
	UpdateService() error
}

type UserNewPage struct {
	User    User
	Message string
	Csrf    string
}

type User struct {
	UserName string
	Email    string
	Image    string
	Age      string
	Birthday string
	Phone    string
	Address  string
}
