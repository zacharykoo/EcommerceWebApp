package service

import "net/http"

type Customer interface {
	Get() http.Handler
	Set() http.Handler
	Edit() http.Handler
}
