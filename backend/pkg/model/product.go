package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ItemNumber   uint   `gorm:"column:item_no" json:"item_no"`
	ItemName     string `gorm:"column:itemName" json:"itemName"`
	Description  string `gorm:"column:description" json:"description"`
	Price        int    `gorm:"column:price" json:"price"`
	Category     string `gorm:"column:category" json:"category"`
	ProductImage string `gorm:"column:product_image" json:"product_image"`
}
