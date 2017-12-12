package structs

type Register interface {
	RegisterNewUser(user User) error
	UpdateNewUser(user User) error
	RegisterService(user User) error
	UpdateService(user User) error
}

type MySql interface {
	DataBaseInit() error

	// Service
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
type ServiceNewPage struct {
	Service Service
	Message string
	Csrf    string
}
type UserMyPage struct {
	User    User
	Service []Service
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

type Service struct {
	Uid         string
	ServiceName string
	Email       string
	Url         string
	CallbackUrl string
	Token       string
	Secret      string

	UserName     bool
	UserEmail    bool
	UserImage    bool
	UserAge      bool
	UserBirthday bool
	UserPhone    bool
	UserAddress  bool
	CreatedAt    string
	UpdatedAt    string
}

type Confirmed struct {
	Id string
	UserUid string
	ServiceUid string
}