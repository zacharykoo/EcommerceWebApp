package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/zacharykoo/EcommerceWebApp/backend/pkg/model"
	"github.com/zacharykoo/EcommerceWebApp/backend/pkg/repository"
)

type rewards struct {
	repo repository.RewardsRepository
}

func GetRewardsService(repo repository.RewardsRepository) RewardsService {
	return &rewards{
		repo: repo,
	}
}

func (c *rewards) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		someRewards, err := c.repo.Get()
		if err != nil {
			fmt.Printf("unable to get rewards: %v", err)
			w.Write([]byte("error getting rewards"))
		}
		b, err := json.Marshal(someRewards)
		if err != nil {
			fmt.Printf("unable to marshal rewards: %v", err)
			w.Write([]byte("error marshalling rewards"))
		}

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(b)
		if err != nil {
			fmt.Printf("Unable to send marshal rewards: %v", err)
			w.Write([]byte("error cannot send rewards"))
		}
	}
}

func (c *rewards) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("unable to read body: %v", err)
			return
		}
		var rewards model.Rewards
		err = json.Unmarshal(body, &rewards)
		if err != nil {
			fmt.Printf("unable to unmarshal into rewards: %v", err)
			return
		}

		rewards, _ = c.repo.Create(rewards)
		rewards.Rewardpt_no = rewards.ID
		w.Write([]byte(fmt.Sprintf("created rewards with No: %v", rewards.ID)))
	}
}

func (c *rewards) Edit() http.HandlerFunc {
	// stubs
	return nil
}
