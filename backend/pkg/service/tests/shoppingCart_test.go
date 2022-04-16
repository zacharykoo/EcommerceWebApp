package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/zacharykoo/EcommerceWebApp/backend/pkg/model"
)

func TestCreateShoppingCart(t *testing.T) {

	shoppingCart := model.ShoppingCart{
		Quantity:      10,
		TotalPrice:    10000,
		ProductsAdded: "RTX 3080",
	}
	b, err := json.Marshal(shoppingCart)

	if err != nil {
		t.Fatalf("unable to marshal shopping cart: %v", err)
	}

	reader := bytes.NewReader(b)

	resp, err := http.Post("http://localhost:8081/api/shoppingCart", "application/json", reader)
	if err != nil {
		t.Fatalf("unable to create request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status code was not OK, we got: %v", resp.StatusCode)
	}
}
