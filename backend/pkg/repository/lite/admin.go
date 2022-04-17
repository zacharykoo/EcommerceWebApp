package lite

import (
	"github.com/zacharykoo/EcommerceWebApp/backend/pkg/model"
	"gorm.io/gorm"
)

type adminRepository struct {
	db *gorm.DB
}

func GetAdminRepository(db *gorm.DB) adminRepository {
	return adminRepository{
		db: db,
	}
}

func (c *adminRepository) Get() ([]model.Admin, error) {
	var admin []model.Admin
	err := c.db.Find(&admin).Error
	return admin, err
}
func (c *adminRepository) Create(admin model.Admin) (model.Admin, error) {
	c.db.Save(&admin)
	admin.AdminID = admin.ID
	err := c.db.Updates(&admin).Error
	return admin, err
}

func (c *adminRepository) Edit(admin model.Admin) (model.Admin, error) {
	err := c.db.Where("adminID = ?", admin.AdminID).Updates(&admin).Error
	return admin, err
}
