package repository

import "github.com/zacharykoo/EcommerceWebApp/backend/pkg/model"

type CustomerRepository interface {
	Get() (model.Customer, error)
	Create(model.Customer) (model.Customer, error)
	Edit(model.Customer) (model.Customer, error)
}
