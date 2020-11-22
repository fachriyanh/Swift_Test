package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Order Struct (Model)
type Order struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Amount int    `json:"amount"`
	Code   string `json:"code"`
}

type Product struct {
	IDProduct   string `json:"idproduct"`
	DateExpired string `json:"dateexpired"`
}

var products []Product
var orders []Order

func getProduct() string {

	oldest := (products[0]).DateExpired
	idx := 0
	i := 0
	for _, item := range products {
		if item.DateExpired < oldest {
			oldest = item.DateExpired
			idx = i
		}
		i++
	}
	return products[idx].IDProduct
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	id := getProduct()

	w.Header().Set("Content-Type", "application/json")
	var order Order
	_ = json.NewDecoder(r.Body).Decode(&order)
	order.ID = "JKT2715"
	order.Code = id
	order.Status = "confirmed"
	orders = append(orders, order)
	json.NewEncoder(w).Encode(order)
}

func main() {
	//Init ROuter
	r := mux.NewRouter()

	//Mock Data
	products = append(products, Product{IDProduct: "NF6H", DateExpired: "20220202"})
	products = append(products, Product{IDProduct: "NF6H", DateExpired: "20220201"})

	// Route Handlerrs / Endpoint
	r.HandleFunc("/api/order", createOrder).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))
}
