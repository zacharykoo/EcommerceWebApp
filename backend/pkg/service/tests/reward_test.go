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
		PointAwarded:     800,
		AwardedDate:      "15/4/2022",
		ProductPurchased: "Nvidia RTX 3090",
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

func TestUpdateReward(t *testing.T) {

	reward := model.Rewards{
		ProductPurchased: "edited",
		Rewardpt_no:      1,
	}
	b, err := json.Marshal(reward)

	if err != nil {
		t.Fatalf("unable to marshal reward: %v", err)
	}

	reader := bytes.NewReader(b)
	req, err := http.NewRequest(http.MethodPut, "http://localhost:8081/api/rewards", reader)
	if err != nil {
		t.Fatalf("unable to request put method for reward: %v", err)
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
