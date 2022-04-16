package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func TestGetCustomerWithID(t *testing.T) {
	membershipID := 0
	url := fmt.Sprintf("http://localhost:8081/api/customer?membershipID=%d", membershipID)

	resp, err := http.Get(url)
	if err != nil {
		t.Fatalf("Unable to get customer from membershipID: %v", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
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
}
