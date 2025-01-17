package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string    `json:"name" gorm:"column:user_name;not null"`
	Email       string    `json:"email" gorm:"column:user_email;unique;not null"`
	Password    string    `json:"password" gorm:"column:user_password;not null"`
	PhoneNumber string    `json:"phone_number" gorm:"column:phone_number;not null"`
	RoleID      uint      `json:"role_id" gorm:"not null;default:4"`
	Companies   []Company `json:"Companies" gorm:"foreignKey:UserID"`
}
