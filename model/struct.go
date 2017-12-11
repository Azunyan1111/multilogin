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
	Uid       string
	UserName  string
	Email     string
	EmailOK   bool
	Image     string
	Age       string
	Birthday  string
	Phone     string
	PhoneOK   bool
	Address   string
	CreatedAt string
	UpdatedAt string
}
