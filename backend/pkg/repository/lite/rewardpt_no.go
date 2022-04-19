package lite

import (
	"github.com/zacharykoo/EcommerceWebApp/backend/pkg/model"
	"gorm.io/gorm"
)

type rewardpt_noRepository struct {
	db *gorm.DB
}

func GetRewardpt_noRepository(db *gorm.DB) rewardpt_noRepository {
	return rewardpt_noRepository{
		db: db,
	}
}

func (c *rewardpt_noRepository) Get(ID int) ([]model.Rewardpt_no, error) {
	var rewardpt_no []model.Rewardpt_no
	option := func(db *gorm.DB, ID int) *gorm.DB {
		if ID != 0 {
			return db.Where("membershipID = ?", ID)
		}
		return db
	}
	err := option(c.db, ID).Find(&rewardpt_no).Error
	return rewardpt_no, err
}
func (c *rewardpt_noRepository) Create(rewardpt_no model.Rewardpt_no) (model.Rewardpt_no, error) {
	err := c.db.Save(&rewardpt_no).Error
	return rewardpt_no, err
}

func (c *rewardpt_noRepository) Edit(rewardpt_no model.Rewardpt_no) (model.Rewardpt_no, error) {
	err := c.db.Where("membershipID = ?", rewardpt_no.MembershipID).Updates(&rewardpt_no).Error
	return rewardpt_no, err
}
