package model

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	FirstName   string `gorm:"column:fn"`
	LastName    string `gorm:"column:ln"`
	PhoneNumber string `gorm:"column:phone_no"`
	Address     string `gorm:"column:address"`
	Preference  string `gorm:"column:preference"`
	Birthday    string `gorm:"column:birthday"`
}
