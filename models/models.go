package models

type EditingUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Credentials struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	Id       string `json:"id"`
	UserName string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type Profile struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Github   string `json:"github"`
	Linkedin string `json:"linkedin"`
	Whatsapp string `json:"whatsapp"`
}
