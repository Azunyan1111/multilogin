package model

type UserNewPage struct {
	NewUser NewUser
	Message string
	Csrf string
}

type NewUser struct {
	UserName string
	Email string
	Image string
	Age string
	Birthday string
	Phone string
	Address string
}