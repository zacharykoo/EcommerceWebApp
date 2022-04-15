// package database will provide a database connection to sqlite for now
package database

import (
	"github.com/zacharykoo/EcommerceWebApp/backend/pkg/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectSQLite() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("EcommerceWebApp.sqlite"), &gorm.Config{})
	return db, err
}

func MigrateTables(db *gorm.DB) error {
	err := db.AutoMigrate(
		model.Customer{},
	)
	return err
}
