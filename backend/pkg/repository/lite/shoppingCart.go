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

func (c *shoppingCartRepository) Get(ID int) ([]model.ShoppingCart, error) {
	var shoppingCart []model.ShoppingCart
	option := func(db *gorm.DB, ID int) *gorm.DB {
		if ID != 0 {
			return db.Where("cartID = ?", ID)
		}
		return db
	}
	err := option(c.db, ID).Find(&shoppingCart).Error
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
