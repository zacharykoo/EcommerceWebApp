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

func (c *productRepository) Get(ID int) ([]model.Product, error) {
	var product []model.Product
	option := func(db *gorm.DB, ID int) *gorm.DB {
		if ID != 0 {
			return db.Where("item_no = ?", ID)
		}
		return db
	}
	err := option(c.db, ID).Find(&product).Error
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
