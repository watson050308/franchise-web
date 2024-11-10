package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"not null;default:null"`
	Email    string `json:"email" gorm:"unique;not null;default:null"`
	Password string `json:"password" gorm:"not null;default:null"`
}
