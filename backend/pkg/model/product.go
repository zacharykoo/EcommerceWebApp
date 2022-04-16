package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ItemNumber   uint   `gorm:"column:item_no" json:"item_no"`
	Price        int    `gorm:"column:price" json:"price"`
	Category     string `gorm:"column:category" json:"category"`
	ProductImage string `gorm:"column:product_image" json:"product_image"`
}
