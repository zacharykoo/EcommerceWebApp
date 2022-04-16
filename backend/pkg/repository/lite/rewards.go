package lite

import (
	"github.com/zacharykoo/EcommerceWebApp/backend/pkg/model"
	"gorm.io/gorm"
)

type rewardsRepository struct {
	db *gorm.DB
}

func GetRewardsRepository(db *gorm.DB) rewardsRepository {
	return rewardsRepository{
		db: db,
	}
}

func (c *rewardsRepository) Get() ([]model.Rewards, error) {
	var rewards []model.Rewards
	err := c.db.Find(&rewards).Error
	return rewards, err
}
func (c *rewardsRepository) Create(rewards model.Rewards) (model.Rewards, error) {
	c.db.Save(&rewards)
	rewards.Rewardpt_no = rewards.ID
	err := c.db.Updates(&rewards).Error
	return rewards, err
}

func (c *rewardsRepository) Edit(rewards model.Rewards) (model.Rewards, error) {
	err := c.db.Where("rewardpt_no = ?", rewards.Rewardpt_no).Updates(&rewards).Error
	return rewards, err
}
