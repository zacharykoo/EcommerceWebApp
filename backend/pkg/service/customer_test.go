package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/zacharykoo/EcommerceWebApp/backend/pkg/model"
)

func TestCreateCustomer(t *testing.T) {

	customer := model.Customer{
		FirstName:   "zachary ",
		LastName:    "koo",
		PhoneNumber: "123-456-7891",
	}
	b, err := json.Marshal(customer)

	if err != nil {
		t.Fatalf("unable to marshal customer: %v", err)
	}

	reader := bytes.NewReader(b)

	resp, err := http.Post("http://localhost:8081/api/customer", "application/json", reader)
	if err != nil {
		t.Fatalf("unable to create request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status code was not OK, we got: %v", resp.StatusCode)
	}
}
