package lite

import (
	"github.com/zacharykoo/EcommerceWebApp/backend/pkg/model"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func GetProductRepository(db *gorm.DB) productRepository {
	return productRepository{
		db: db,
	}
}

func (c *productRepository) Get() ([]model.Product, error) {
	var product []model.Product
	err := c.db.Find(&product).Error
	return product, err
}
func (c *productRepository) Create(product model.Product) (model.Product, error) {
	c.db.Save(&product)
	product.ItemNumber = product.ID
	err := c.db.Updates(&product).Error
	return product, err
}

func (c *productRepository) Edit(product model.Product) (model.Product, error) {
	err := c.db.Where("item_no = ?", product.ItemNumber).Updates(&product).Error
	return product, err
}
