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
	err := c.db.Save(&product).Error
	return product, err
}

func (c *productRepository) Edit(product model.Product) (model.Product, error) {

	// stub
	return model.Product{}, nil
}
