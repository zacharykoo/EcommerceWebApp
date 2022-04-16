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
		Price:        10000,
		Category:     "Hardware",
		ProductImage: "someImage1",
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
