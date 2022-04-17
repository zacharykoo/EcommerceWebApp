package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/zacharykoo/EcommerceWebApp/backend/pkg/model"
)

func TestCreateShipment(t *testing.T) {

	shipment := model.Shipment{
		TrackingInfo:    "Shipping from China",
		ExpectedArrival: "10/4/2022",
		TransportType:   "Ship",
	}
	b, err := json.Marshal(shipment)

	if err != nil {
		t.Fatalf("unable to marshal shipment: %v", err)
	}

	reader := bytes.NewReader(b)

	resp, err := http.Post("http://localhost:8081/api/shipment", "application/json", reader)
	if err != nil {
		t.Fatalf("unable to create request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status code was not OK, we got: %v", resp.StatusCode)
	}
}

func TestUpdateShipment(t *testing.T) {

	shipment := model.Shipment{
		ShipmentID:   2,
		TrackingInfo: "Edited info",
	}
	b, err := json.Marshal(shipment)

	if err != nil {
		t.Fatalf("unable to marshal shipment: %v", err)
	}

	reader := bytes.NewReader(b)
	req, err := http.NewRequest(http.MethodPut, "http://localhost:8081/api/shipment", reader)
	if err != nil {
		t.Fatalf("unable to request put method for shipment: %v", err)
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
