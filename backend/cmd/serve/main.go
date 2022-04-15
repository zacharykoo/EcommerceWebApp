package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zacharykoo/EcommerceWebApp/backend/pkg/database"
	"github.com/zacharykoo/EcommerceWebApp/backend/pkg/repository/lite"
	"github.com/zacharykoo/EcommerceWebApp/backend/pkg/service"
)

func main() {
	fmt.Println("Hello world")
	//service.GetCustomerService()

	db, err := database.ConnectSQLite()
	if err != nil {
		fmt.Printf("unable to connect to database: %v\n", err)
		return
	}

	err = database.MigrateTables(db)
	if err != nil {
		fmt.Printf("unable to migrate tables: %v", err)
		return
	}

	customerRepository := lite.GetCustomerRepository(db)
	customerService := service.GetCustomerService(&customerRepository)

	r := mux.NewRouter()

	r.Handle("/api/customer", customerService.Get()).Methods("GET")
	r.Handle("/api/customer", customerService.Create()).Methods("POST")

	fmt.Println("Serving on :8081...")
	http.ListenAndServe(":8081", r)
}
