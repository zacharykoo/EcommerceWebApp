package lite

import (
	"github.com/zacharykoo/EcommerceWebApp/backend/pkg/model"
	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func GetOrderRepository(db *gorm.DB) orderRepository {
	return orderRepository{
		db: db,
	}
}

func (c *orderRepository) Get() ([]model.Order, error) {
	var order []model.Order
	err := c.db.Find(&order).Error
	return order, err
}
func (c *orderRepository) Create(order model.Order) (model.Order, error) {
	c.db.Save(&order)
	order.OrderNumber = order.ID
	err := c.db.Updates(&order).Error
	return order, err
}

func (c *orderRepository) Edit(order model.Order) (model.Order, error) {
	err := c.db.Where("order_no = ?", order.OrderNumber).Updates(&order).Error
	return order, err
}
