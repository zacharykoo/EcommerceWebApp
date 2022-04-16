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
	err := c.db.Save(&rewards).Error
	return rewards, err
}

func (c *rewardsRepository) Edit(rewards model.Rewards) (model.Rewards, error) {

	// stub
	return model.Rewards{}, nil
}
