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

func (c *customerRepository) Get(ID int) ([]model.Customer, error) {
	var customers []model.Customer
	option := func(db *gorm.DB, ID int) *gorm.DB {
		if ID != 0 {
			return db.Where("membershipID = ?", ID)
		}
		return db
	}
	err := option(c.db, ID).Find(&customers).Error
	return customers, err
}

func (c *customerRepository) Create(customer model.Customer) (model.Customer, error) {
	c.db.Save(&customer)
	customer.MembershipID = customer.ID
	err := c.db.Updates(&customer).Error
	return customer, err
}

func (c *customerRepository) Edit(customer model.Customer) (model.Customer, error) {
	err := c.db.Where("membershipID = ?", customer.MembershipID).Updates(&customer).Error
	return customer, err
}
