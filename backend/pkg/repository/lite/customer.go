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

func (c *customerRepository) Get() model.Customer {
	// stub
	// Database guy works on getting the db stuff from c.db
	// https://gorm.io/docs/index.html
	return model.Customer{}
}
func (c *customerRepository) Set(model.Customer) model.Customer {
	// stub

	return model.Customer{}
}
func (c *customerRepository) Edit(model.Customer) model.Customer {

	// stub
	return model.Customer{}
}
