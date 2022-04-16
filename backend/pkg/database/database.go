// package database will provide a database connection to sqlite for now
package database

import (
	"github.com/glebarez/sqlite"
	"github.com/zacharykoo/EcommerceWebApp/backend/pkg/model"
	"gorm.io/gorm"
)

func ConnectSQLite() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("EcommerceWebApp.sqlite"), &gorm.Config{})
	return db, err
}

func MigrateTables(db *gorm.DB) error {
	err := db.AutoMigrate(
		model.Customer{},
		model.Rewardpt_no{},
		model.Rewards{},
		model.Product{},
		model.ShoppingCart{},
		model.Order{},
		model.Coupon{},
		model.Admin{},
		model.Shipment{},
	)
	return err
}
