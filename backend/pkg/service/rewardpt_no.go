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

type rewardpt_no struct {
	repo repository.Rewardpt_noRepository
}

func GetRewardpt_noService(repo repository.Rewardpt_noRepository) Rewardpt_noService {
	return &rewardpt_no{
		repo: repo,
	}
}

func (c *rewardpt_no) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var membershipID int
		var err error
		membershipIDString := r.URL.Query().Get(KeyMembershipID)
		if membershipIDString == "" {
			membershipID = 0
		} else {
			membershipID, err = strconv.Atoi(membershipIDString)
			if err != nil {
				fmt.Printf("unable to parse membershipID: %v", err)
				w.Write([]byte("error getting customer"))
				return
			}
		}
		someRewardpt, err := c.repo.Get(membershipID)
		if err != nil {
			fmt.Printf("unable to get rewardpt_no: %v", err)
			w.Write([]byte("error getting rewardpt_no"))
		}
		b, err := json.Marshal(someRewardpt)
		if err != nil {
			fmt.Printf("unable to marshal rewardpt_no: %v", err)
			w.Write([]byte("error getting rewardpt_no"))
		}

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(b)
		if err != nil {
			fmt.Printf("Unable to send marshal rewardpt_no: %v", err)
			w.Write([]byte("error cannot send rewardpt_no"))
		}
	}
}

func (c *rewardpt_no) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("unable to read body: %v", err)
			return
		}
		var rewardpt_no model.Rewardpt_no
		err = json.Unmarshal(body, &rewardpt_no)
		if err != nil {
			fmt.Printf("unable to unmarshal into rewardpt_no: %v", err)
			return
		}

		rewardpt_no, _ = c.repo.Create(rewardpt_no)
		w.Write([]byte(fmt.Sprintf("created rewardpt_no with No: %v", rewardpt_no.ID)))
	}
}

func (c *rewardpt_no) Edit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("unable to read body: %v", err)
			return
		}
		var rewardpt_no model.Rewardpt_no
		err = json.Unmarshal(body, &rewardpt_no)
		if err != nil {
			fmt.Printf("unable to unmarshal into rewardpt_no: %v", err)
			return
		}

		rewardpt_no, _ = c.repo.Edit(rewardpt_no)
		w.Write([]byte(fmt.Sprintf("created rewardpt_no with ID: %v", rewardpt_no.ID)))
	}
}
