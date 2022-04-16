package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	OrderNumber  uint   `gorm:"column:order_no" json:"order_no"`
	MembershipID uint   `gorm:"column:membershipID" json:"membershipID"`
	PhoneNumber  int    `gorm:"column:phone_no" json:"phone_no"`
	Email        string `gorm:"column:email" json:"email"`
	City         string `gorm:"column:city" json:"city"`
	State        string `gorm:"column:state" json:"state"`
	Street       string `gorm:"column:street" json:"street"`
	Zip          string `gorm:"column:zip" json:"zip"`
}
