package repository

import "github.com/zacharykoo/EcommerceWebApp/backend/pkg/model"

type CustomerRepository interface {
	Get() model.Customer
	Set(model.Customer) model.Customer
	Edit(model.Customer) model.Customer
}
