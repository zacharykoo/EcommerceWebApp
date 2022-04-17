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
		TrackingInfo:    "All the way up Mars",
		ExpectedArrival: "16/4/2022",
		TransportType:   "Private Jet",
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
