package lite

import (
	"github.com/zacharykoo/EcommerceWebApp/backend/pkg/model"
	"gorm.io/gorm"
)

type shoppingCartRepository struct {
	db *gorm.DB
}

func GetShoppingCartRepository(db *gorm.DB) shoppingCartRepository {
	return shoppingCartRepository{
		db: db,
	}
}

func (c *shoppingCartRepository) Get() ([]model.ShoppingCart, error) {
	var shoppingCart []model.ShoppingCart
	err := c.db.Find(&shoppingCart).Error
	return shoppingCart, err
}
func (c *shoppingCartRepository) Create(shoppingCart model.ShoppingCart) (model.ShoppingCart, error) {
	c.db.Save(&shoppingCart)
	shoppingCart.CartID = shoppingCart.ID
	err := c.db.Updates(&shoppingCart).Error
	return shoppingCart, err
}

func (c *shoppingCartRepository) Edit(shoppingCart model.ShoppingCart) (model.ShoppingCart, error) {
	err := c.db.Where("cartID = ?", shoppingCart.CartID).Updates(&shoppingCart).Error
	return shoppingCart, err
}
