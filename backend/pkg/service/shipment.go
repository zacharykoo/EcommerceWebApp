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

type shipment struct {
	repo repository.ShipmentRepository
}

func GetShipmentService(repo repository.ShipmentRepository) ShipmentService {
	return &shipment{
		repo: repo,
	}
}

func (c *shipment) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var shipmentID int
		var err error
		shipmentIDString := r.URL.Query().Get(KeyShipmentID)
		if shipmentIDString == "" {
			shipmentID = 0
		} else {
			shipmentID, err = strconv.Atoi(shipmentIDString)
			if err != nil {
				fmt.Printf("unable to parse shipmentID: %v", err)
				w.Write([]byte("error getting shipment"))
				return
			}
		}
		someShipment, err := c.repo.Get(shipmentID)
		if err != nil {
			fmt.Printf("unable to get shipment: %v", err)
			w.Write([]byte("error getting shipment"))
		}
		b, err := json.Marshal(someShipment)
		if err != nil {
			fmt.Printf("unable to marshal shipment: %v", err)
			w.Write([]byte("error marshalling shipment"))
		}

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(b)
		if err != nil {
			fmt.Printf("Unable to send marshal shipment: %v", err)
			w.Write([]byte("error cannot send shipment"))
		}
	}
}

func (c *shipment) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("unable to read body: %v", err)
			return
		}
		var shipment model.Shipment
		err = json.Unmarshal(body, &shipment)
		if err != nil {
			fmt.Printf("unable to unmarshal into shipment: %v", err)
			return
		}

		shipment, _ = c.repo.Create(shipment)
		shipment.ShipmentID = shipment.ID
		w.Write([]byte(fmt.Sprintf("created shipment with shipmentID: %v", shipment.ShipmentID)))
	}
}

func (c *shipment) Edit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("unable to read body: %v", err)
			return
		}
		var shipment model.Shipment
		err = json.Unmarshal(body, &shipment)
		if err != nil {
			fmt.Printf("unable to unmarshal into shipment: %v", err)
			return
		}

		shipment, _ = c.repo.Edit(shipment)
		w.Write([]byte(fmt.Sprintf("Shipment with shipmentID %v is editted", shipment.ShipmentID)))
	}
}
