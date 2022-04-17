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
		Reward_no:    1,
		MembershipID: 1,
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
