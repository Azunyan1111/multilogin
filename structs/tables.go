package structs

import "github.com/jinzhu/gorm"

type ConfirmedService struct {
	GormModel gorm.Model
	UserUid string	`gorm:"column:user_uuid"`
	ServiceUid string	`gorm:"column:service_uuid"`
}

type Service struct {
	GormModel gorm.Model

	Uid         string `gorm:"column:uuid"`
	ServiceName string `gorm:"column:name"`
	Email       string `gorm:"column:email"`
	Url         string `gorm:"column:url"`
	CallbackUrl string `gorm:"column:url_callback"`
	Token       string `gorm:"column:token"`
	Secret      string `gorm:"column:secret"`

	UserName     bool `gorm:"column:p_name"`
	UserEmail    bool `gorm:"column:p_email"`
	UserImage    bool `gorm:"column:p_image"`
	UserAge      bool `gorm:"column:p_age"`
	UserBirthday bool `gorm:"column:p_birthday"`
	UserPhone    bool `gorm:"column:p_phone"`
	UserAddress  bool `gorm:"column:p_address"`
}

type User struct {
	GormModel gorm.Model

	Uid       string `gorm:"column:uuid"`
	UserName  string `gorm:"column:user"`
	Email     string `gorm:"column:email"`
	EmailOK   bool `gorm:"column:email_ok"`
	Image     string `gorm:"column:image"`
	Age       string `gorm:"column:age"`
	Birthday  string `gorm:"column:birthday"`
	Phone     string `gorm:"column:phone"`
	PhoneOK   bool `gorm:"column:phone_ok"`
	Address   string `gorm:"column:address"`
}

// set TableName
func (User) TableName() string {
	return "users"
}

func (Service) TableName() string {
	return "service"
}

func (ConfirmedService) TableName() string {
	return "confirmed_service"
}