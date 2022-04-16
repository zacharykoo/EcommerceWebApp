package model

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	FirstName    string `gorm:"column:fn" json:"fn"`
	LastName     string `gorm:"column:ln" json:"ln"`
	PhoneNumber  string `gorm:"column:phone_no" json:"phone_no"`
	Address      string `gorm:"column:address" json:"address"`
	Preference   string `gorm:"column:preference" json:"preference"`
	Birthday     string `gorm:"column:birthday" json:"birthday"`
	MembershipID string `gorm:"column:membershipID" json:"membershipID"`
}
