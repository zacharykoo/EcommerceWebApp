package model

import "gorm.io/gorm"

type Coupon struct {
	gorm.Model
	CouponID   uint   `gorm:"column:couponID" json:"couponID"`
	CouponInfo string `gorm:"column:couponInfo" json:"couponInfo"`
}
