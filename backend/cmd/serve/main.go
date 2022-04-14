package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello world")
	r := mux.NewRouter()
	r.HandleFunc("/api/customer", getCustomer("zachary"))
	fmt.Println("Serving on :8081...")
	http.ListenAndServe(":8081", r)
}

func getCustomer(zachary string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		test := r.URL.Query()[zachary]
		w.Write([]byte("Hello World " + test[0]))
	}
}
