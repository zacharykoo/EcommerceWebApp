package lite

import (
	"github.com/zacharykoo/EcommerceWebApp/backend/pkg/model"
	"gorm.io/gorm"
)

type customerRepository struct {
	db *gorm.DB
}

func GetCustomerRepository(db *gorm.DB) customerRepository {
	return customerRepository{
		db: db,
	}
}

func (c *customerRepository) Get() (model.Customer, error) {
	var customers []model.Customer
	err := c.db.Find(&customers).Error
	return model.Customer{}, err
}
func (c *customerRepository) Create(customer model.Customer) (model.Customer, error) {
	err := c.db.Save(&customer).Error
	return customer, err
}

func (c *customerRepository) Edit(customer model.Customer) (model.Customer, error) {

	// stub
	return model.Customer{}, nil
}
