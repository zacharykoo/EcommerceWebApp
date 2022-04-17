package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/zacharykoo/EcommerceWebApp/backend/pkg/model"
	"github.com/zacharykoo/EcommerceWebApp/backend/pkg/repository"
)

const (
	KeyMembershipID = "membershipID"
)

type customer struct {
	repo repository.CustomerRepository
}

func GetCustomerService(repo repository.CustomerRepository) CustomerService {
	return &customer{
		repo: repo,
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func (c *customer) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		var membershipID int
		var err error
		membershipIDString, ok := r.URL.Query()[KeyMembershipID]
		if !ok {
			membershipID = 0
		} else {
			membershipID, err = strconv.Atoi(membershipIDString[0])
			if err != nil {
				fmt.Printf("unable to parse membershipID: %v", err)
				w.Write([]byte("error getting customer"))
			}
		}
		someCustomer, err := c.repo.Get(membershipID)
		if err != nil {
			fmt.Printf("unable to get customer: %v", err)
			w.Write([]byte("error getting customer"))
		}
		b, err := json.Marshal(someCustomer)
		if err != nil {
			fmt.Printf("unable to marshal customer: %v", err)
			w.Write([]byte("error getting customer"))
		}

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(b)
		if err != nil {
			fmt.Printf("Unable to send marshal customer: %v", err)
			w.Write([]byte("error cannot send customer"))
		}
	}
}

func (c *customer) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
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

		customer, _ = c.repo.Create(customer)
		w.Write([]byte(fmt.Sprintf("created customer with ID: %v", customer.ID)))
	}
}

func (c *customer) Edit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
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

		customer, _ = c.repo.Edit(customer)
		w.Write([]byte(fmt.Sprintf("created customer with ID: %v", customer.ID)))
	}
}
