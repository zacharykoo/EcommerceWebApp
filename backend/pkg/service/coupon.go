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
		var couponID int
		var err error
		couponIDString := r.URL.Query().Get(KeyCouponID)
		if couponIDString == "" {
			couponID = 0
		} else {
			couponID, err = strconv.Atoi(couponIDString)
			if err != nil {
				fmt.Printf("unable to parse couponID: %v", err)
				w.Write([]byte("error getting coupon"))
				return
			}
		}
		someCoupon, err := c.repo.Get(couponID)
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
		w.Write([]byte(fmt.Sprintf("created rewards with couponID: %v", coupon.CouponID)))
	}
}

func (c *coupon) Edit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
		w.Write([]byte(fmt.Sprintf("Coupon with couponID %v is editted", coupon.CouponID)))
	}
}
