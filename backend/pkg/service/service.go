package service

import "net/http"

type CustomerService interface {
	Get() http.HandlerFunc
	Set() http.HandlerFunc
	Edit() http.HandlerFunc
}
