package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/zacharykoo/EcommerceWebApp/backend/pkg/model"
	"github.com/zacharykoo/EcommerceWebApp/backend/pkg/repository"
)

type shoppingCart struct {
	repo repository.ShoppingCartRepository
}

func GetShoppingCartService(repo repository.ShoppingCartRepository) ShoppingCartService {
	return &shoppingCart{
		repo: repo,
	}
}

func (c *shoppingCart) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		someShoppingCart, err := c.repo.Get()
		if err != nil {
			fmt.Printf("unable to get shopping cart: %v", err)
			w.Write([]byte("error getting shopping cart"))
		}
		b, err := json.Marshal(someShoppingCart)
		if err != nil {
			fmt.Printf("unable to marshal shopping cart: %v", err)
			w.Write([]byte("error marshalling shopping cart"))
		}

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(b)
		if err != nil {
			fmt.Printf("Unable to send marshal shopping cart: %v", err)
			w.Write([]byte("error cannot send shopping cart"))
		}
	}
}

func (c *shoppingCart) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("unable to read body: %v", err)
			return
		}
		var shoppingCart model.ShoppingCart
		err = json.Unmarshal(body, &shoppingCart)
		if err != nil {
			fmt.Printf("unable to unmarshal into shopping cart: %v", err)
			return
		}

		shoppingCart, _ = c.repo.Create(shoppingCart)
		shoppingCart.CartID = shoppingCart.ID
		w.Write([]byte(fmt.Sprintf("created rewards with No: %v", shoppingCart.ID)))
	}
}

func (c *shoppingCart) Edit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("unable to read body: %v", err)
			return
		}
		var shoppingCart model.ShoppingCart
		err = json.Unmarshal(body, &shoppingCart)
		if err != nil {
			fmt.Printf("unable to unmarshal into shopping cart: %v", err)
			return
		}

		shoppingCart, _ = c.repo.Edit(shoppingCart)
		w.Write([]byte(fmt.Sprintf("created shopping cart with ID: %v", shoppingCart.ID)))
	}
}
