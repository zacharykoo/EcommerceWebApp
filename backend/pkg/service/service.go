package service

import "net/http"

type CustomerService interface {
	Get() http.HandlerFunc
	Create() http.HandlerFunc
	Edit() http.HandlerFunc
}
