package model

import "gorm.io/gorm"

type Shipment struct {
	gorm.Model
	ShipmentID      uint   `gorm:"column:shipmentID" json:"shipmentID"`
	TrackingInfo    string `gorm:"column:trackingInfo" json:"trackingInfo"`
	ExpectedArrival string `gorm:"column:expectedArrival" json:"expectedArrival"`
	TransportType   string `gorm:"column:transportType" json:"transportType"`
}
