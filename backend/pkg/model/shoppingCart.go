package model

import "gorm.io/gorm"

type ShoppingCart struct {
	gorm.Model
	CartID uint `gorm:"column:cartID" json:"cartID"`
	//This should be slice but "unsupported data type: &[]"
	Quantity   int `gorm:"column:quantity" json:"quantity"`
	TotalPrice int `gorm:"column:totalPrice" json:"totalPrice"`
	//This should be slice but "unsupported data type: &[]"
	ProductsAdded string `gorm:"column:productsAdded" json:"productsAdded"`
}

type Items struct {
	gorm.Model
	Quantity int `gorm:"column:quantity" json:"quantity"`
}
