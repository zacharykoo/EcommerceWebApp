package service

import "net/http"

type CustomerService interface {
	Get() http.HandlerFunc
	Create() http.HandlerFunc
	Edit() http.HandlerFunc
}

type Rewardpt_noService interface {
	Get() http.HandlerFunc
	Create() http.HandlerFunc
	Edit() http.HandlerFunc
}

type RewardsService interface {
	Get() http.HandlerFunc
	Create() http.HandlerFunc
	Edit() http.HandlerFunc
}

type ProductService interface {
	Get() http.HandlerFunc
	Create() http.HandlerFunc
	Edit() http.HandlerFunc
}

type ShoppingCartService interface {
	Get() http.HandlerFunc
	Create() http.HandlerFunc
	Edit() http.HandlerFunc
}

type OrderService interface {
	Get() http.HandlerFunc
	Create() http.HandlerFunc
	Edit() http.HandlerFunc
}
