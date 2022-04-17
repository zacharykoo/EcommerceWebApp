package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	AdminID uint   `gorm:"column:adminID" json:"adminID"`
	Name    string `gorm:"column:name" json:"name"`
	Address string `gorm:"column:address" json:"address"`
}
