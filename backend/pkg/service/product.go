package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/zacharykoo/EcommerceWebApp/backend/pkg/model"
	"github.com/zacharykoo/EcommerceWebApp/backend/pkg/repository"
)

type product struct {
	repo repository.ProductRepository
}

func GetProductService(repo repository.ProductRepository) ProductService {
	return &product{
		repo: repo,
	}
}

func (c *product) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		someProduct, err := c.repo.Get()
		if err != nil {
			fmt.Printf("unable to get product: %v", err)
			w.Write([]byte("error getting product"))
		}
		b, err := json.Marshal(someProduct)
		if err != nil {
			fmt.Printf("unable to marshal product: %v", err)
			w.Write([]byte("error marshalling product"))
		}

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(b)
		if err != nil {
			fmt.Printf("Unable to send marshal product: %v", err)
			w.Write([]byte("error cannot send product"))
		}
	}
}

func (c *product) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("unable to read body: %v", err)
			return
		}
		var product model.Product
		err = json.Unmarshal(body, &product)
		if err != nil {
			fmt.Printf("unable to unmarshal into rewards: %v", err)
			return
		}

		product, _ = c.repo.Create(product)
		product.ItemNumber = product.ID
		w.Write([]byte(fmt.Sprintf("created rewards with No: %v", product.ID)))
	}
}

func (c *product) Edit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("unable to read body: %v", err)
			return
		}
		var product model.Product
		err = json.Unmarshal(body, &product)
		if err != nil {
			fmt.Printf("unable to unmarshal into product: %v", err)
			return
		}

		product, _ = c.repo.Edit(product)
		w.Write([]byte(fmt.Sprintf("created product with ID: %v", product.ID)))
	}
}
