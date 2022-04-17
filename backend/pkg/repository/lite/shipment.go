package lite

import (
	"github.com/zacharykoo/EcommerceWebApp/backend/pkg/model"
	"gorm.io/gorm"
)

type shipmentRepository struct {
	db *gorm.DB
}

func GetShipmentRepository(db *gorm.DB) shipmentRepository {
	return shipmentRepository{
		db: db,
	}
}

func (c *shipmentRepository) Get() ([]model.Shipment, error) {
	var shipment []model.Shipment
	err := c.db.Find(&shipment).Error
	return shipment, err
}
func (c *shipmentRepository) Create(shipment model.Shipment) (model.Shipment, error) {
	c.db.Save(&shipment)
	shipment.ShipmentID = shipment.ID
	err := c.db.Updates(&shipment).Error
	return shipment, err
}

func (c *shipmentRepository) Edit(shipment model.Shipment) (model.Shipment, error) {
	err := c.db.Where("shipmentID = ?", shipment.ShipmentID).Updates(&shipment).Error
	return shipment, err
}