package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

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
		var rewardpt_no int
		var err error
		rewardpt_noString := r.URL.Query().Get(KeyRewardpt_no)
		if rewardpt_noString == "" {
			rewardpt_no = 0
		} else {
			rewardpt_no, err = strconv.Atoi(rewardpt_noString)
			if err != nil {
				fmt.Printf("unable to parse rewardpt_no: %v", err)
				w.Write([]byte("error getting rewards"))
				return
			}
		}
		someRewards, err := c.repo.Get(rewardpt_no)
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
		w.Write([]byte(fmt.Sprintf("created rewards with rewardpt_no: %v", rewards.Rewardpt_no)))
	}
}

func (c *rewards) Edit() http.HandlerFunc {
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

		rewards, _ = c.repo.Edit(rewards)
		w.Write([]byte(fmt.Sprintf("Rewards with rewardpt_no %v is editted", rewards.Rewardpt_no)))
	}
}
