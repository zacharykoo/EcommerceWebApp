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
		MembershipID: 1,
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
