package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/zacharykoo/EcommerceWebApp/backend/pkg/model"
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
		someCustomer, err := c.repo.Get()
		if err != nil {
			fmt.Printf("unable to get customer: %v", err)
			w.Write([]byte("error getting customer"))
		}
		b, err := json.Marshal(someCustomer)
		if err != nil {
			fmt.Printf("unable to marshal customer: %v", err)
			w.Write([]byte("error getting customer"))
		}
		w.Write(b)
	}
}

func (c *customer) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("unable to read body: %v", err)
			return
		}
		var customer model.Customer
		err = json.Unmarshal(body, &customer)
		if err != nil {
			fmt.Printf("unable to unmarshal into customer: %v", err)
			return
		}

		customer, err = c.repo.Create(customer)
		w.Write([]byte(fmt.Sprintf("created customer with ID: %v", customer.ID)))
	}
}

func (c *customer) Edit() http.HandlerFunc {
	// stubs
	return nil
}
