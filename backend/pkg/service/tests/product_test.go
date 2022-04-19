package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/zacharykoo/EcommerceWebApp/backend/pkg/model"
)

func TestCreateProduct(t *testing.T) {

	product := model.Product{
		Price:        10,
		ItemName:     "The Ordinary",
		Description:  "Apply to face for shinny forehead",
		Category:     "Beauty",
		ProductImage: "https://images.unsplash.com/photo-1580870069867-74c57ee1bb07?crop=entropy&cs=tinysrgb&fit=crop&fm=jpg&h=900&ixid=MnwxfDB8MXxyYW5kb218MHx8cHJvZHVjdHx8fHx8fDE2NTAxNzAyNzI&ixlib=rb-1.2.1&q=80&utm_campaign=api-credit&utm_medium=referral&utm_source=unsplash_source&w=1600",
	}
	b, err := json.Marshal(product)

	if err != nil {
		t.Fatalf("unable to marshal product: %v", err)
	}

	reader := bytes.NewReader(b)

	resp, err := http.Post("http://localhost:8081/api/product", "application/json", reader)
	if err != nil {
		t.Fatalf("unable to create request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status code was not OK, we got: %v", resp.StatusCode)
	}
}

func TestUpdateProduct(t *testing.T) {

	product := model.Product{
		ItemNumber: 2,
		Price:      1200,
		Category:   "Edited",
	}

	b, err := json.Marshal(product)

	if err != nil {
		t.Fatalf("unable to marshal product: %v", err)
	}

	reader := bytes.NewReader(b)
	req, err := http.NewRequest(http.MethodPut, "http://localhost:8081/api/product", reader)
	if err != nil {
		t.Fatalf("unable to request put method for product: %v", err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("unable to create request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status code was not OK, we got: %v", resp.StatusCode)
	}
}
