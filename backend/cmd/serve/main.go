package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zacharykoo/EcommerceWebApp/backend/pkg/repository/lite"
	"github.com/zacharykoo/EcommerceWebApp/backend/pkg/service"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Hello world")

	db := ConnectToDatabase()
	customerRepository := lite.GetCustomerRepository(db)
	customerService := service.GetCustomerService(&customerRepository)

	r := mux.NewRouter()
	r.Handle("/api/customer", customerService.Get())
	fmt.Println("Serving on :8081...")
	http.ListenAndServe(":8081", r)

}

// make this connect to a sqlite database, or w/e db
// might need to make this return an error later too
func ConnectToDatabase() *gorm.DB {

	// STUB, NEED TO DO THIS
	return nil
}
