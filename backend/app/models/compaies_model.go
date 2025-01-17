package model

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name                   string `json:"name" gorm:"column:company_name;not null"`
	TelNumber              string `json:"tel_number" gorm:"not null"`
	BusinessID             string `json:"business_id" gorm:"unique;not null"`
	ResponsiblePerson      string `json:"responsible_person" gorm:"not null"`
	ResponsiblePersonPhone string `json:"responsible_person_phone" gorm:"not null"`
	PrimaryContectPhone    string `json:"primary_contect_phone" gorm:"not null"`
	CompanyEmail           string `json:"company_email" gorm:"not null"`
	CompanyAddress         string `json:"company_address" gorm:"not null"`
	WebUrl                 string `json:"web_url"`
	BrandID                uint   `json:"brand_id"`
	UserID                 uint   `json:"user_id" gorm:"not null"`
}
