package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name  string `json:"role_name" gorm:"column:role_name;not null"`
	Users []User `json:"users" gorm:"foreignKey:RoleID"`
}
