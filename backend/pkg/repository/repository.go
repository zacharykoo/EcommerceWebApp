package repository

import "github.com/zacharykoo/EcommerceWebApp/backend/pkg/model"

type CustomerRepository interface {
	Get(ID int) ([]model.Customer, error)
	Create(model.Customer) (model.Customer, error)
	Edit(model.Customer) (model.Customer, error)
}

type Rewardpt_noRepository interface {
	Get() ([]model.Rewardpt_no, error)
	Create(model.Rewardpt_no) (model.Rewardpt_no, error)
	Edit(model.Rewardpt_no) (model.Rewardpt_no, error)
}

type RewardsRepository interface {
	Get() ([]model.Rewards, error)
	Create(model.Rewards) (model.Rewards, error)
	Edit(model.Rewards) (model.Rewards, error)
}

type ProductRepository interface {
	Get() ([]model.Product, error)
	Create(model.Product) (model.Product, error)
	Edit(model.Product) (model.Product, error)
}

type ShoppingCartRepository interface {
	Get() ([]model.ShoppingCart, error)
	Create(model.ShoppingCart) (model.ShoppingCart, error)
	Edit(model.ShoppingCart) (model.ShoppingCart, error)
}

type OrderRepository interface {
	Get() ([]model.Order, error)
	Create(model.Order) (model.Order, error)
	Edit(model.Order) (model.Order, error)
}

type CouponRepository interface {
	Get() ([]model.Coupon, error)
	Create(model.Coupon) (model.Coupon, error)
	Edit(model.Coupon) (model.Coupon, error)
}

type AdminRepository interface {
	Get() ([]model.Admin, error)
	Create(model.Admin) (model.Admin, error)
	Edit(model.Admin) (model.Admin, error)
}

type ShipmentRepository interface {
	Get() ([]model.Shipment, error)
	Create(model.Shipment) (model.Shipment, error)
	Edit(model.Shipment) (model.Shipment, error)
}
