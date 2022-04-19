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
	KeyRewardpt_no  = "rewardpt_no"
	KeyItem_no      = "item_no"
	KeyCartID       = "cartID"
	KeyShipmentID   = "shipmentID"
	KeyOrder_no     = "order_no"
	KeyCouponID     = "couponID"
	KeyAdminID      = "adminID"
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

		var membershipID int
		var err error
		membershipIDString := r.URL.Query().Get(KeyMembershipID)
		if membershipIDString == "" {
			membershipID = 0
		} else {
			membershipID, err = strconv.Atoi(membershipIDString)
			if err != nil {
				fmt.Printf("unable to parse membershipID: %v", err)
				w.Write([]byte("error getting customer"))
				return
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
		w.Write([]byte(fmt.Sprintf("created customer with membershipID: %v", customer.MembershipID)))
	}
}

func (c *customer) Edit() http.HandlerFunc {
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

		customer, _ = c.repo.Edit(customer)
		w.Write([]byte(fmt.Sprintf("Customer with membershipID %v is editted", customer.MembershipID)))
	}
}
