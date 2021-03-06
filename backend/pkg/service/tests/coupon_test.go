package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/zacharykoo/EcommerceWebApp/backend/pkg/model"
)

func TestCreateCoupon(t *testing.T) {

	coupon := model.Coupon{
		CouponInfo: "Coupon for 10 percent off!!",
	}
	b, err := json.Marshal(coupon)

	if err != nil {
		t.Fatalf("unable to marshal coupon: %v", err)
	}

	reader := bytes.NewReader(b)

	resp, err := http.Post("http://localhost:8081/api/coupon", "application/json", reader)
	if err != nil {
		t.Fatalf("unable to create request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status code was not OK, we got: %v", resp.StatusCode)
	}
}

func TestUpdateCoupon(t *testing.T) {

	coupon := model.Coupon{
		CouponInfo: "Coupon for free shipment",
	}
	b, err := json.Marshal(coupon)

	if err != nil {
		t.Fatalf("unable to marshal coupon: %v", err)
	}

	reader := bytes.NewReader(b)
	req, err := http.NewRequest(http.MethodPut, "http://localhost:8081/api/coupon", reader)
	if err != nil {
		t.Fatalf("unable to request put method for coupon: %v", err)
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
