package model

import "gorm.io/gorm"

type ShoppingCart struct {
	gorm.Model
	CartID      uint   `gorm:"column:cartID" json:"cartID"`
	ProductList string `gorm:"column:productList" json:"productList"`
}
