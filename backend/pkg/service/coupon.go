package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/zacharykoo/EcommerceWebApp/backend/pkg/model"
	"github.com/zacharykoo/EcommerceWebApp/backend/pkg/repository"
)

type coupon struct {
	repo repository.CouponRepository
}

func GetCouponService(repo repository.CouponRepository) CouponService {
	return &coupon{
		repo: repo,
	}
}

func (c *coupon) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		someCoupon, err := c.repo.Get()
		if err != nil {
			fmt.Printf("unable to get coupon: %v", err)
			w.Write([]byte("error getting coupon"))
		}
		b, err := json.Marshal(someCoupon)
		if err != nil {
			fmt.Printf("unable to marshal coupon: %v", err)
			w.Write([]byte("error marshalling coupon"))
		}

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(b)
		if err != nil {
			fmt.Printf("Unable to send marshal coupon: %v", err)
			w.Write([]byte("error cannot send coupon"))
		}
	}
}

func (c *coupon) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("unable to read body: %v", err)
			return
		}
		var coupon model.Coupon
		err = json.Unmarshal(body, &coupon)
		if err != nil {
			fmt.Printf("unable to unmarshal into rewards: %v", err)
			return
		}

		coupon, _ = c.repo.Create(coupon)
		coupon.CouponID = coupon.ID
		w.Write([]byte(fmt.Sprintf("created rewards with No: %v", coupon.ID)))
	}
}

func (c *coupon) Edit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("unable to read body: %v", err)
			return
		}
		var coupon model.Coupon
		err = json.Unmarshal(body, &coupon)
		if err != nil {
			fmt.Printf("unable to unmarshal into customer: %v", err)
			return
		}

		coupon, _ = c.repo.Edit(coupon)
		w.Write([]byte(fmt.Sprintf("created customer with ID: %v", coupon.ID)))
	}
}
