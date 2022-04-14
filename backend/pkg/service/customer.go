package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zacharykoo/EcommerceWebApp/backend/pkg/repository"
)

type customer struct {
	repo repository.CustomerRepository
}

func GetCustomerService(repo repository.CustomerRepository) CustomerService {
	return &customer{
		repo: repo,
	}
}

func (c *customer) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		someCustomer := c.repo.Get()
		b, err := json.Marshal(someCustomer)
		if err != nil {
			fmt.Printf("unable to marshal customer: %v", err)
			// send error to frontend/api
		}
		w.Write(b)
	}
}

func (c *customer) Set() http.HandlerFunc {
	// stubs
	return nil
}

func (c *customer) Edit() http.HandlerFunc {
	// stubs
	return nil
}
