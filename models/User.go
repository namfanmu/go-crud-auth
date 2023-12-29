package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string  `gorm:"unique" json:"username"`
	Password string  `json:"password"`
	Age      int     `json:"age"`
	Roles    []*Role `gorm:"many2many:user_roles;" json:"roles"`
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
