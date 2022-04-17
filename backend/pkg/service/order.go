package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/zacharykoo/EcommerceWebApp/backend/pkg/model"
	"github.com/zacharykoo/EcommerceWebApp/backend/pkg/repository"
)

type order struct {
	repo repository.OrderRepository
}

func GetOrderService(repo repository.OrderRepository) OrderService {
	return &order{
		repo: repo,
	}
}

func (c *order) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		someOrder, err := c.repo.Get()
		if err != nil {
			fmt.Printf("unable to get order: %v", err)
			w.Write([]byte("error getting order"))
		}
		b, err := json.Marshal(someOrder)
		if err != nil {
			fmt.Printf("unable to marshal order: %v", err)
			w.Write([]byte("error marshalling order"))
		}

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(b)
		if err != nil {
			fmt.Printf("Unable to send marshal order: %v", err)
			w.Write([]byte("error cannot send order"))
		}
	}
}

func (c *order) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("unable to read body: %v", err)
			return
		}
		var order model.Order
		err = json.Unmarshal(body, &order)
		if err != nil {
			fmt.Printf("unable to unmarshal into rewards: %v", err)
			return
		}

		order, _ = c.repo.Create(order)
		order.OrderNumber = order.ID
		w.Write([]byte(fmt.Sprintf("created rewards with No: %v", order.ID)))
	}
}

func (c *order) Edit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("unable to read body: %v", err)
			return
		}
		var order model.Order
		err = json.Unmarshal(body, &order)
		if err != nil {
			fmt.Printf("unable to unmarshal into order: %v", err)
			return
		}

		order, _ = c.repo.Edit(order)
		w.Write([]byte(fmt.Sprintf("created order with ID: %v", order.ID)))
	}
}
