package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Product struct {
	Id          int     `json:"id"`
	Sku         string  `json:"sku"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}

var products = []*Product{
	{
		Id:          1,
		Sku:         "sku1",
		Name:        "Test Product Name 1",
		Description: " Test Product Description 2",
		Price:       100.77,
		Stock:       10,
	},
	{
		Id:          2,
		Sku:         "sku2",
		Name:        "Test Product Name 2",
		Description: " Test Product Description 2",
		Price:       100.77,
		Stock:       20,
	},
}

func allProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("allProducts API hit NewServeMux.")
	json.NewEncoder(w).Encode(products)

}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page Endpoint NewServeMux.")
}

// func getProduct(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Get a single product by ID.")
// }
// func updateProduct(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Update product of given ID.")
// }
// func deleteProduct(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Delete a product")
// }
// func createProduct(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Create a product")
// }

func handleRequests() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homePage)
	mux.HandleFunc("/products", allProducts)
	// http.HandleFunc("/products/<id:int>", getProduct)
	// http.HandleFunc("/products/<id:int>", updateProduct)
	// http.HandleFunc("/products/<id:int>", deleteProduct)
	// http.HandleFunc("/products/<id:int>", createProduct)
	log.Fatal(http.ListenAndServe(":3000", mux))
}

func main() {
	fmt.Println("go apis with NewServeMux")
	handleRequests()
}
