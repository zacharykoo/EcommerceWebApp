package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/zacharykoo/EcommerceWebApp/backend/pkg/model"
)

func TestCreateAdmin(t *testing.T) {

	admin := model.Admin{
		Name:    "Mickey",
		Address: "111 Dalt Wisney",
	}
	b, err := json.Marshal(admin)

	if err != nil {
		t.Fatalf("unable to marshal admin: %v", err)
	}

	reader := bytes.NewReader(b)

	resp, err := http.Post("http://localhost:8081/api/admin", "application/json", reader)
	if err != nil {
		t.Fatalf("unable to create request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status code was not OK, we got: %v", resp.StatusCode)
	}
}

func TestUpdateAdmin(t *testing.T) {

	admin := model.Admin{
		Name:    "EDITED NAME!!!",
		Address: "Edited address!!!!",
		AdminID: 2,
	}
	b, err := json.Marshal(admin)

	if err != nil {
		t.Fatalf("unable to marshal admin: %v", err)
	}

	reader := bytes.NewReader(b)
	req, err := http.NewRequest(http.MethodPut, "http://localhost:8081/api/admin", reader)
	if err != nil {
		t.Fatalf("unable to request put method for admin: %v", err)
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
