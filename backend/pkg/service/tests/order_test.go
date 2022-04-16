package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/zacharykoo/EcommerceWebApp/backend/pkg/model"
)

func TestCreateOrder(t *testing.T) {

	order := model.Order{
		MembershipID: 2,
		PhoneNumber:  4031234567,
		Email:        "example@gmail.com",
		City:         "Calgary",
		State:        "Alberta",
		Street:       "123 example",
		Zip:          "A2B 3C4",
	}
	b, err := json.Marshal(order)

	if err != nil {
		t.Fatalf("unable to marshal order: %v", err)
	}

	reader := bytes.NewReader(b)

	resp, err := http.Post("http://localhost:8081/api/order", "application/json", reader)
	if err != nil {
		t.Fatalf("unable to create request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status code was not OK, we got: %v", resp.StatusCode)
	}
}

func TestUpdateOrder(t *testing.T) {

	order := model.Order{
		PhoneNumber: 9999999999,
		Email:       "Edited Email",
		OrderNumber: 3,
	}
	b, err := json.Marshal(order)

	if err != nil {
		t.Fatalf("unable to marshal order: %v", err)
	}

	reader := bytes.NewReader(b)
	req, err := http.NewRequest(http.MethodPut, "http://localhost:8081/api/order", reader)
	if err != nil {
		t.Fatalf("unable to request put method for order: %v", err)
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
