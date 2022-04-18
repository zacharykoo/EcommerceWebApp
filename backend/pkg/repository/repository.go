package repository

import "github.com/zacharykoo/EcommerceWebApp/backend/pkg/model"

type CustomerRepository interface {
	Get(ID int) ([]model.Customer, error)
	Create(model.Customer) (model.Customer, error)
	Edit(model.Customer) (model.Customer, error)
}

type Rewardpt_noRepository interface {
	Get(ID int) ([]model.Rewardpt_no, error)
	Create(model.Rewardpt_no) (model.Rewardpt_no, error)
	Edit(model.Rewardpt_no) (model.Rewardpt_no, error)
}

type RewardsRepository interface {
	Get(ID int) ([]model.Rewards, error)
	Create(model.Rewards) (model.Rewards, error)
	Edit(model.Rewards) (model.Rewards, error)
}

type ProductRepository interface {
	Get(ID int) ([]model.Product, error)
	Create(model.Product) (model.Product, error)
	Edit(model.Product) (model.Product, error)
}

type ShoppingCartRepository interface {
	Get(ID int) ([]model.ShoppingCart, error)
	Create(model.ShoppingCart) (model.ShoppingCart, error)
	Edit(model.ShoppingCart) (model.ShoppingCart, error)
}

type OrderRepository interface {
	Get(ID int) ([]model.Order, error)
	Create(model.Order) (model.Order, error)
	Edit(model.Order) (model.Order, error)
}

type CouponRepository interface {
	Get(ID int) ([]model.Coupon, error)
	Create(model.Coupon) (model.Coupon, error)
	Edit(model.Coupon) (model.Coupon, error)
}

type AdminRepository interface {
	Get(ID int) ([]model.Admin, error)
	Create(model.Admin) (model.Admin, error)
	Edit(model.Admin) (model.Admin, error)
}

type ShipmentRepository interface {
	Get(ID int) ([]model.Shipment, error)
	Create(model.Shipment) (model.Shipment, error)
	Edit(model.Shipment) (model.Shipment, error)
}
