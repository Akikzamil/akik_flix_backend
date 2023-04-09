package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	IsAdmin  bool   `json:"is_admin"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
