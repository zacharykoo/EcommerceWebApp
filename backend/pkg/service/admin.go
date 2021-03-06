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

type admin struct {
	repo repository.AdminRepository
}

func GetAdminService(repo repository.AdminRepository) AdminService {
	return &admin{
		repo: repo,
	}
}

func (c *admin) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var adminID int
		var err error
		adminIDString := r.URL.Query().Get(KeyAdminID)
		if adminIDString == "" {
			adminID = 0
		} else {
			adminID, err = strconv.Atoi(adminIDString)
			if err != nil {
				fmt.Printf("unable to parse adminID: %v", err)
				w.Write([]byte("error getting admin"))
				return
			}
		}
		someAdmin, err := c.repo.Get(adminID)
		if err != nil {
			fmt.Printf("unable to get admin: %v", err)
			w.Write([]byte("error getting admin"))
		}
		b, err := json.Marshal(someAdmin)
		if err != nil {
			fmt.Printf("unable to marshal admin: %v", err)
			w.Write([]byte("error marshalling admin"))
		}

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(b)
		if err != nil {
			fmt.Printf("Unable to send marshal admin: %v", err)
			w.Write([]byte("error cannot send admin"))
		}
	}
}

func (c *admin) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("unable to read body: %v", err)
			return
		}
		var admin model.Admin
		err = json.Unmarshal(body, &admin)
		if err != nil {
			fmt.Printf("unable to unmarshal into admin: %v", err)
			return
		}

		admin, _ = c.repo.Create(admin)
		admin.AdminID = admin.ID
		w.Write([]byte(fmt.Sprintf("created admin with adminID: %v", admin.AdminID)))
	}
}

func (c *admin) Edit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("unable to read body: %v", err)
			return
		}
		var admin model.Admin
		err = json.Unmarshal(body, &admin)
		if err != nil {
			fmt.Printf("unable to unmarshal into admin: %v", err)
			return
		}

		admin, _ = c.repo.Edit(admin)
		w.Write([]byte(fmt.Sprintf("Admin with adminID %v is editted", admin.AdminID)))
	}
}
