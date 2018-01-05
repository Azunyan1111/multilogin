package structs

type Register interface {
	RegisterNewUser(user Usered) error
	UpdateNewUser(user Usered) error
	RegisterService(user Usered) error
	UpdateService(user Usered) error
}

type MySql interface {
	DataBaseInit() error

	// Serviced
	InsertNewUser(user Usered) error
	UpdateUser(user Usered) error

	// Serviced
	InsertService() error
	UpdateService() error
}

type UserNewPage struct {
	User    Usered
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
type ServiceMyPage struct {
	Service      Service
	UserResponse []UserResponse
	Message      string
	Csrf         string
}
type ConfirmedPage struct {
	User    Usered
	Service Service
	Message string
	Csrf    string
}
type LoginPage struct {
	Email   string
	Message string
	Csrf    string
}
type LoginCodePage struct {
	Code    string
	Message string
	Csrf    string
}
type Usered struct {
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
	DeletedAt string
}

type Serviced struct {
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
	DeletedAt    string
}

type Confirmed struct {
	Id         string
	UserUid    string
	ServiceUid string
}

type Error struct {
	StatusCode int
	Message    string
}
type UserResponse struct {
	Uid      string `json:"Uuid"`
	UserName string `json:"UserName"`
	Email    string `json:"Email"`
	Image    string `json:"Image"`
	Age      string `json:"Age"`
	Birthday string `json:"Birthday"`
	Phone    string `json:"Phone"`
	Address  string `json:"Address"`
}
