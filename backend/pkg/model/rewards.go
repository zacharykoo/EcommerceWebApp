package model

import "gorm.io/gorm"

type Rewards struct {
	gorm.Model
	Rewardpt_no      uint   `gorm:"column:rewardpt_no" json:"rewardpt_no"`
	PointAwarded     int    `gorm:"column:point_awarded" json:"point_awarded"`
	AwardedDate      string `gorm:"column:awarded_date" json:"awarded_date"`
	ProductPurchased string `gorm:"column:product_purchased" json:"product_purchased"`
}
