package lite

import (
	"github.com/zacharykoo/EcommerceWebApp/backend/pkg/model"
	"gorm.io/gorm"
)

type couponRepository struct {
	db *gorm.DB
}

func GetCouponRepository(db *gorm.DB) couponRepository {
	return couponRepository{
		db: db,
	}
}

func (c *couponRepository) Get(ID int) ([]model.Coupon, error) {
	var coupon []model.Coupon
	option := func(db *gorm.DB, ID int) *gorm.DB {
		if ID != 0 {
			return db.Where("couponID = ?", ID)
		}
		return db
	}
	err := option(c.db, ID).Find(&coupon).Error
	return coupon, err
}
func (c *couponRepository) Create(coupon model.Coupon) (model.Coupon, error) {
	c.db.Save(&coupon)
	coupon.CouponID = coupon.ID
	err := c.db.Updates(&coupon).Error
	return coupon, err
}

func (c *couponRepository) Edit(coupon model.Coupon) (model.Coupon, error) {
	err := c.db.Where("couponID = ?", coupon.CouponID).Updates(&coupon).Error
	return coupon, err
}
