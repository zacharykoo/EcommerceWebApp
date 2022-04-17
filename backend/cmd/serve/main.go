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

	rewardpt_noRepository := lite.GetRewardpt_noRepository(db)
	rewardpt_noService := service.GetRewardpt_noService(&rewardpt_noRepository)

	rewardsRepository := lite.GetRewardsRepository(db)
	rewardsService := service.GetRewardsService(&rewardsRepository)

	productRepository := lite.GetProductRepository(db)
	productService := service.GetProductService(&productRepository)

	shoppingCartRepository := lite.GetShoppingCartRepository(db)
	shoppingCartService := service.GetShoppingCartService(&shoppingCartRepository)

	orderRepository := lite.GetOrderRepository(db)
	orderService := service.GetOrderService(&orderRepository)

	couponRepository := lite.GetCouponRepository(db)
	couponService := service.GetCouponService(&couponRepository)

	adminRepository := lite.GetAdminRepository(db)
	adminService := service.GetAdminService(&adminRepository)

	shipmentRepository := lite.GetShipmentRepository(db)
	shipmentService := service.GetShipmentService(&shipmentRepository)

	r := mux.NewRouter()

	r.Handle("/api/customer", customerService.Get()).Methods("GET")
	r.Handle("/api/customer", customerService.Create()).Methods("POST")
	r.Handle("/api/customer", customerService.Edit()).Methods("PUT")

	r.Handle("/api/rewardpt_no", rewardpt_noService.Get()).Methods("GET")
	r.Handle("/api/rewardpt_no", rewardpt_noService.Create()).Methods("POST")

	r.Handle("/api/rewards", rewardsService.Get()).Methods("GET")
	r.Handle("/api/rewards", rewardsService.Create()).Methods("POST")

	r.Handle("/api/product", productService.Get()).Methods("GET")
	r.Handle("/api/product", productService.Create()).Methods("POST")

	r.Handle("/api/shoppingCart", shoppingCartService.Get()).Methods("GET")
	r.Handle("/api/shoppingCart", shoppingCartService.Create()).Methods("POST")

	r.Handle("/api/order", orderService.Get()).Methods("GET")
	r.Handle("/api/order", orderService.Create()).Methods("POST")

	r.Handle("/api/coupon", couponService.Get()).Methods("GET")
	r.Handle("/api/coupon", couponService.Create()).Methods("POST")
	r.Handle("/api/coupon", couponService.Edit()).Methods("PUT")

	r.Handle("/api/admin", adminService.Get()).Methods("GET")
	r.Handle("/api/admin", adminService.Create()).Methods("POST")
	r.Handle("/api/admin", adminService.Edit()).Methods("PUT")

	r.Handle("/api/shipment", shipmentService.Get()).Methods("GET")
	r.Handle("/api/shipment", shipmentService.Create()).Methods("POST")

	fmt.Println("Serving on :8081...")
	http.ListenAndServe(":8081", r)
}
