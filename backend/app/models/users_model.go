package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string `json:"name" gorm:"column:user_name;not null;default:null"`
	Email       string `json:"email" gorm:"column:user_email;unique;not null;default:null"`
	Password    string `json:"password" gorm:"column:user_password;not null;default:null"`
	PhoneNumber string `json:"phone_number"`
}
