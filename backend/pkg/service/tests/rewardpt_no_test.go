package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/zacharykoo/EcommerceWebApp/backend/pkg/model"
)

func TestCreateRewardpt_no(t *testing.T) {

	rewardpt_no := model.Rewardpt_no{
		Reward_no:    20,
		MembershipID: 5,
	}
	b, err := json.Marshal(rewardpt_no)

	if err != nil {
		t.Fatalf("unable to marshal rewardpt_no: %v", err)
	}

	reader := bytes.NewReader(b)

	resp, err := http.Post("http://localhost:8081/api/rewardpt_no", "application/json", reader)
	if err != nil {
		t.Fatalf("unable to create request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status code was not OK, we got: %v", resp.StatusCode)
	}
}

func TestUpdateRewardpt_no(t *testing.T) {

	rewardpt_no := model.Rewardpt_no{
		MembershipID: 2,
		Reward_no:    65,
	}
	b, err := json.Marshal(rewardpt_no)

	if err != nil {
		t.Fatalf("unable to marshal rewardpt_no: %v", err)
	}

	reader := bytes.NewReader(b)
	req, err := http.NewRequest(http.MethodPut, "http://localhost:8081/api/rewardpt_no", reader)
	if err != nil {
		t.Fatalf("unable to request put method for rewardpt_no: %v", err)
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
