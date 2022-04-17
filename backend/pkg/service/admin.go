package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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
		enableCors(&w)
		someAdmin, err := c.repo.Get()
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
		enableCors(&w)
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
		w.Write([]byte(fmt.Sprintf("created admin with No: %v", admin.ID)))
	}
}

func (c *admin) Edit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
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
		w.Write([]byte(fmt.Sprintf("created admin with ID: %v", admin.ID)))
	}
}
