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

func (c *couponRepository) Get() ([]model.Coupon, error) {
	var coupon []model.Coupon
	err := c.db.Find(&coupon).Error
	return coupon, err
}
func (c *couponRepository) Create(coupon model.Coupon) (model.Coupon, error) {
	err := c.db.Save(&coupon).Error
	return coupon, err
}

func (c *couponRepository) Edit(coupon model.Coupon) (model.Coupon, error) {

	// stub
	return model.Coupon{}, nil
}
