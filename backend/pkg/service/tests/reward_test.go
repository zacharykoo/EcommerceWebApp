package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/zacharykoo/EcommerceWebApp/backend/pkg/model"
)

func TestCreateRewards(t *testing.T) {

	rewards := model.Rewards{
		PointAwarded:     100,
		AwardedDate:      "14/4/2022",
		ProductPurchased: "Nvidia RTX 3080",
	}
	b, err := json.Marshal(rewards)

	if err != nil {
		t.Fatalf("unable to marshal rewards: %v", err)
	}

	reader := bytes.NewReader(b)

	resp, err := http.Post("http://localhost:8081/api/rewards", "application/json", reader)
	if err != nil {
		t.Fatalf("unable to create request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status code was not OK, we got: %v", resp.StatusCode)
	}
}
