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
		FirstName:   "Beff",
		LastName:    "Jezos",
		PhoneNumber: "101-202-3030",
		Address:     "123 Scamazon",
		Preference:  "Beauty",
		Birthday:    "2/12/1940",
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

func TestUpdateCustomer(t *testing.T) {

	customer := model.Customer{
		FirstName:    "JohnEdit",
		LastName:     "SmithEdit",
		MembershipID: 5,
	}
	b, err := json.Marshal(customer)

	if err != nil {
		t.Fatalf("unable to marshal customer: %v", err)
	}

	reader := bytes.NewReader(b)
	req, err := http.NewRequest(http.MethodPut, "http://localhost:8081/api/customer", reader)
	if err != nil {
		t.Fatalf("unable to request put method for customer: %v", err)
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

func TestGetCustomerWithID(t *testing.T) {
	membershipID := 4
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
